package executions

import "context"

func (s *Service) SoftDeleteExpiredExecutions(ctx context.Context) error {
	expiredExecutions, err := s.dbgen.ExecutionsServiceGetExpiredExecutions(ctx)
	if err != nil {
		return err
	}

	for _, execution := range expiredExecutions {
		if err := s.SoftDeleteExecution(ctx, execution.ID); err != nil {
			return err
		}
	}

	return nil
}
