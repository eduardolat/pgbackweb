package users

import "context"

func (s *Service) GetUsersQty(ctx context.Context) (int64, error) {
	return s.dbgen.UsersServiceGetUsersQty(ctx)
}
