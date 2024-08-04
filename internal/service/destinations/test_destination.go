package destinations

import "fmt"

func (s *Service) TestDestination(
	accessKey, secretKey, region, endpoint, bucketName string,
) error {
	err := s.ints.StorageClient.S3Ping(accessKey, secretKey, region, endpoint, bucketName)
	if err != nil {
		return fmt.Errorf("error pinging destination: %w", err)
	}

	return nil
}
