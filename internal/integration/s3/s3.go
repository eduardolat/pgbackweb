package s3

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
)

type Client struct{}

func New() *Client {
	return &Client{}
}

// createS3Client creates a new S3 client
func createS3Client(
	accessKey, secretKey, region, endpoint string,
) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:           aws.String(region),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create aws session: %w", err)
	}

	return s3.New(sess), nil
}

// Ping tests the connection to S3
func (Client) Ping(
	accessKey, secretKey, region, endpoint, bucketName string,
) error {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return err
	}

	_, err = s3Client.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return fmt.Errorf("failed to ping S3 bucket: %w", err)
	}

	return nil
}

// Upload uploads a file to S3
func (Client) Upload(
	accessKey, secretKey, region, endpoint, bucketName, key string,
	fileContent []byte,
) (string, error) {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(fileContent)
	key = strutil.RemoveLeadingSlash(key)
	contentType := strutil.GetContentTypeFromFileName(key)

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(bucketName),
		Key:         aws.String(key),
		Body:        aws.ReadSeekCloser(reader),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %w", err)
	}

	return key, nil
}

// Delete deletes a file from S3
func (Client) Delete(
	accessKey, secretKey, region, endpoint, bucketName, key string,
) error {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return err
	}

	key = strutil.RemoveLeadingSlash(key)

	_, err = s3Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}

	return nil
}

// GetDownloadLink generates a presigned URL for downloading a file from S3
func (Client) GetDownloadLink(
	accessKey, secretKey, region, endpoint, bucketName, key string,
	expiration time.Duration,
) (string, error) {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return "", err
	}

	key = strutil.RemoveLeadingSlash(key)
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})

	url, err := req.Presign(expiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return url, nil
}
