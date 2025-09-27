package oidc

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/eduardolat/pgbackweb/internal/config"
	"github.com/eduardolat/pgbackweb/internal/database/dbgen"
	"golang.org/x/oauth2"
)

// Custom error types for better error handling
var (
	ErrEmailAlreadyExists = errors.New("email already exists with different authentication method")
	ErrOIDCNotEnabled     = errors.New("OIDC is not enabled")
	ErrInvalidToken       = errors.New("invalid or expired token")
	ErrMissingClaims      = errors.New("required user information missing from OIDC claims")
)

type Service struct {
	env      config.Env
	dbgen    *dbgen.Queries
	provider *oidc.Provider
	config   oauth2.Config
}

type UserInfo struct {
	Email    string
	Name     string
	Username string
	Subject  string
}

func New(env config.Env, dbgen *dbgen.Queries) (*Service, error) {
	if !env.PBW_OIDC_ENABLED {
		return &Service{env: env, dbgen: dbgen}, nil
	}

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, env.PBW_OIDC_ISSUER_URL)
	if err != nil {
		return nil, fmt.Errorf("failed to create OIDC provider: %w", err)
	}

	scopes := strings.Split(env.PBW_OIDC_SCOPES, " ")
	config := oauth2.Config{
		ClientID:     env.PBW_OIDC_CLIENT_ID,
		ClientSecret: env.PBW_OIDC_CLIENT_SECRET,
		RedirectURL:  env.PBW_OIDC_REDIRECT_URL,
		Endpoint:     provider.Endpoint(),
		Scopes:       scopes,
	}

	return &Service{
		env:      env,
		dbgen:    dbgen,
		provider: provider,
		config:   config,
	}, nil
}

func (s *Service) IsEnabled() bool {
	return s.env.PBW_OIDC_ENABLED
}

func (s *Service) GetAuthURL(state string) string {
	if !s.IsEnabled() {
		return ""
	}
	return s.config.AuthCodeURL(state)
}

func (s *Service) GenerateState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func (s *Service) ExchangeCode(ctx context.Context, code string) (*UserInfo, error) {
	if !s.IsEnabled() {
		return nil, ErrOIDCNotEnabled
	}

	token, err := s.config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, fmt.Errorf("no id_token field in oauth2 token")
	}

	verifier := s.provider.Verifier(&oidc.Config{ClientID: s.env.PBW_OIDC_CLIENT_ID})
	idToken, err := verifier.Verify(ctx, rawIDToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %w", err)
	}

	claims := make(map[string]interface{})
	if err := idToken.Claims(&claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	userInfo := &UserInfo{
		Subject: idToken.Subject,
	}

	// Extract email
	if email, ok := claims[s.env.PBW_OIDC_EMAIL_CLAIM].(string); ok {
		userInfo.Email = strings.ToLower(email)
	}

	// Extract name
	if name, ok := claims[s.env.PBW_OIDC_NAME_CLAIM].(string); ok {
		userInfo.Name = name
	}

	// Extract username
	if username, ok := claims[s.env.PBW_OIDC_USERNAME_CLAIM].(string); ok {
		userInfo.Username = username
	}

	// Fallback to email as username if username not provided
	if userInfo.Username == "" && userInfo.Email != "" {
		userInfo.Username = strings.Split(userInfo.Email, "@")[0]
	}

	// Fallback to username as name if name not provided
	if userInfo.Name == "" && userInfo.Username != "" {
		userInfo.Name = userInfo.Username
	}

	if userInfo.Email == "" || userInfo.Name == "" || userInfo.Subject == "" {
		return nil, ErrMissingClaims
	}

	return userInfo, nil
}

func (s *Service) CreateOrUpdateUser(ctx context.Context, userInfo *UserInfo) (*dbgen.User, error) {
	// Try to get existing OIDC user
	_, err := s.dbgen.OIDCServiceGetUserByOIDC(ctx, dbgen.OIDCServiceGetUserByOIDCParams{
		OidcProvider: sql.NullString{String: "oidc", Valid: true},
		OidcSubject:  sql.NullString{String: userInfo.Subject, Valid: true},
	})

	if err == nil {
		// OIDC user exists, update their information
		user, err := s.dbgen.OIDCServiceUpdateUser(ctx, dbgen.OIDCServiceUpdateUserParams{
			Name:         userInfo.Name,
			Email:        userInfo.Email,
			OidcProvider: sql.NullString{String: "oidc", Valid: true},
			OidcSubject:  sql.NullString{String: userInfo.Subject, Valid: true},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to update user: %w", err)
		}
		return &user, nil
	}

	// OIDC user doesn't exist, check if regular user with same email exists
	_, err = s.dbgen.AuthServiceLoginGetUserByEmail(ctx, strings.ToLower(userInfo.Email))
	if err == nil {
		// Regular user with same email exists - we cannot create OIDC user
		// This prevents account takeover and maintains data integrity
		return nil, ErrEmailAlreadyExists
	}

	// No existing user, create new OIDC user
	user, err := s.dbgen.OIDCServiceCreateUser(ctx, dbgen.OIDCServiceCreateUserParams{
		Name:         userInfo.Name,
		Email:        userInfo.Email,
		OidcProvider: sql.NullString{String: "oidc", Valid: true},
		OidcSubject:  sql.NullString{String: userInfo.Subject, Valid: true},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}
