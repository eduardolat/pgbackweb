# OpenID Connect (OIDC) Integration

This document describes the OIDC integration added to pgbackweb to support authentication via external providers like Authentik, Keycloak, and other OIDC-compliant identity providers.

## Overview

The OIDC integration allows users to authenticate using external identity providers instead of local username/password combinations. This supports enterprise SSO workflows and centralized user management.

## Features

- Support for any OIDC-compliant identity provider
- Automatic user creation on first login
- User information synchronization on each login
- Seamless integration with existing authentication system
- Configurable user attribute mapping

## Configuration

Add the following environment variables to enable OIDC:

```bash
# Enable OIDC authentication
PBW_OIDC_ENABLED=true

# OIDC Provider Configuration
PBW_OIDC_ISSUER_URL=https://your-provider.com/auth/realms/your-realm
PBW_OIDC_CLIENT_ID=pgbackweb
PBW_OIDC_CLIENT_SECRET=your-client-secret
PBW_OIDC_REDIRECT_URL=https://your-domain.com/auth/oidc/callback

# Optional: Customize OIDC scopes (default: "openid profile email")
PBW_OIDC_SCOPES="openid profile email"

# Optional: Customize claim mappings
PBW_OIDC_USERNAME_CLAIM=preferred_username  # default: preferred_username
PBW_OIDC_EMAIL_CLAIM=email                  # default: email
PBW_OIDC_NAME_CLAIM=name                    # default: name
```

## Provider-Specific Setup

### Authentik

1. Create a new **OAuth2/OpenID Provider** in Authentik
2. Set the **Redirect URI** to: `https://your-domain.com/auth/oidc/callback`
3. Configure the **Client Type** as **Confidential**
4. Note the **Client ID** and **Client Secret**
5. Create a new **Application** and link it to the provider
6. Configure the **Issuer URL**: `https://your-authentik.com/application/o/your-app/`

### Keycloak

1. Create a new **Client** in your Keycloak realm
2. Set **Client Protocol** to `openid-connect`
3. Set **Access Type** to `confidential`
4. Add `https://your-domain.com/auth/oidc/callback` to **Valid Redirect URIs**
5. Note the **Client ID** and get the **Client Secret** from the Credentials tab
6. Configure the **Issuer URL**: `https://your-keycloak.com/auth/realms/your-realm`

### Generic OIDC Provider

For any OIDC-compliant provider:

1. Create a new OIDC client/application
2. Set the redirect URI to: `https://your-domain.com/auth/oidc/callback`
3. Ensure the client can access `openid`, `profile`, and `email` scopes
4. Note the issuer URL (usually ends with `/.well-known/openid_configuration`)

## Database Schema Changes

The OIDC integration adds the following columns to the `users` table:

```sql
ALTER TABLE users 
ADD COLUMN oidc_provider TEXT,
ADD COLUMN oidc_subject TEXT;

-- Make password nullable for OIDC users
ALTER TABLE users ALTER COLUMN password DROP NOT NULL;

-- Create unique index for OIDC users
CREATE UNIQUE INDEX users_oidc_provider_subject_idx 
ON users (oidc_provider, oidc_subject) 
WHERE oidc_provider IS NOT NULL AND oidc_subject IS NOT NULL;

-- Ensure users have either password or OIDC authentication
ALTER TABLE users ADD CONSTRAINT users_auth_method_check 
CHECK (
    (password IS NOT NULL AND oidc_provider IS NULL AND oidc_subject IS NULL) OR
    (password IS NULL AND oidc_provider IS NOT NULL AND oidc_subject IS NOT NULL)
);
```

## User Flow

1. **First-time users**: When an OIDC user logs in for the first time, a new user account is automatically created with information from the OIDC provider.

2. **Returning users**: Existing OIDC users are matched by their provider and subject ID. User information (name, email) is updated from the OIDC provider on each login.

3. **Mixed authentication**: The system supports both local users (with passwords) and OIDC users in the same instance.

## Security Considerations

- **State parameter**: CSRF protection using a random state parameter
- **Token validation**: ID tokens are cryptographically verified
- **Secure cookies**: State is stored in secure, HTTP-only cookies
- **Provider validation**: Only configured OIDC providers are accepted

## User Interface

When OIDC is enabled, the login page displays:
- A "Login with SSO" button at the top
- A divider separating SSO from traditional login
- The existing email/password form below

## Implementation Details

### Services

- **`internal/service/oidc/`**: Core OIDC authentication logic
- **`internal/view/web/oidc/`**: Web routes for OIDC authentication flow
- **`internal/config/`**: Environment variable configuration and validation

### Routes

- `GET /auth/oidc/login`: Initiates OIDC authentication flow
- `GET /auth/oidc/callback`: Handles OIDC provider callback

### Database Queries

- `OIDCServiceCreateUser`: Creates a new OIDC user
- `OIDCServiceGetUserByOIDC`: Retrieves user by provider and subject
- `OIDCServiceUpdateUser`: Updates existing OIDC user information

## Troubleshooting

### Common Issues

1. **Invalid redirect URI**: Ensure the redirect URI in your OIDC provider matches exactly: `https://your-domain.com/auth/oidc/callback`

2. **Certificate errors**: If using self-signed certificates, ensure your Go application trusts the certificates

3. **Claim mapping**: Verify that your OIDC provider returns the expected claims (`email`, `name`, `preferred_username`)

4. **Scopes**: Ensure your OIDC client has access to the required scopes (`openid`, `profile`, `email`)

### Debug Logging

The application logs OIDC authentication events. Check logs for:
- OIDC provider initialization errors
- Token exchange failures
- User creation/update events

## Migration from Local Authentication

Existing local users are unaffected by OIDC integration. To migrate users to OIDC:

1. Enable OIDC authentication
2. Users can continue using local authentication or switch to OIDC
3. No automatic migration is performed - users choose their preferred method

## Development

To run the application with OIDC in development:

```bash
# Set environment variables in .env file
echo "PBW_OIDC_ENABLED=true" >> .env
echo "PBW_OIDC_ISSUER_URL=https://your-dev-provider.com" >> .env
# ... other OIDC variables

# Run database migrations
task migrate up

# Generate database code
task gen-db

# Build and run
task dev
```
