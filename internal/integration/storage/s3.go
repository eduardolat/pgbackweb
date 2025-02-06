package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/eduardolat/pgbackweb/internal/util/strutil"
)

// createS3Client creates a new S3 client
func createS3Client(
	accessKey, secretKey, region, endpoint string,
) (*s3.Client, error) {
	credentialsProvider := credentials.NewStaticCredentialsProvider(
		accessKey, secretKey, "",
	)

	//nolint:all
	endpointResolver := aws.EndpointResolverFunc(func(
		_ string, _ string,
	) (aws.Endpoint, error) {
		return aws.Endpoint{
			HostnameImmutable: true,
			URL:               endpoint,
		}, nil
	})

	//nolint:all
	conf, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolver(endpointResolver),
		config.WithCredentialsProvider(credentialsProvider),
	)
	if err != nil {
		return nil, fmt.Errorf("error initializing storage config: %w", err)
	}

	s3Client := s3.NewFromConfig(conf)
	return s3Client, nil
}

// S3Test tests the connection to S3
func (Client) S3Test(
	accessKey, secretKey, region, endpoint, bucketName string,
) error {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return err
	}

	_, err = s3Client.HeadBucket(
		context.TODO(),
		&s3.HeadBucketInput{
			Bucket: aws.String(bucketName),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to test S3 bucket: %w", err)
	}

	return nil
}

// S3Upload uploads a file to S3 from a reader.
//
// Returns the file size, in bytes.
func (Client) S3Upload(
	accessKey, secretKey, region, endpoint, bucketName, key string,
	fileReader io.Reader,
) (int64, error) {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return 0, err
	}

	key = strutil.RemoveLeadingSlash(key)
	contentType := strutil.GetContentTypeFromFileName(key)

	uploader := manager.NewUploader(s3Client)
	_, err = uploader.Upload(
		context.TODO(),
		&s3.PutObjectInput{
			Bucket:      aws.String(bucketName),
			Key:         aws.String(key),
			Body:        fileReader,
			ContentType: aws.String(contentType),
		},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to upload file to S3: %w", err)
	}

	fileHead, err := s3Client.HeadObject(
		context.TODO(),
		&s3.HeadObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return 0, fmt.Errorf("failed to get uploaded file info from S3: %w", err)
	}

	var fileSize int64
	if fileHead.ContentLength != nil {
		fileSize = *fileHead.ContentLength
	}

	return fileSize, nil
}

// S3Delete deletes a file from S3
func (Client) S3Delete(
	accessKey, secretKey, region, endpoint, bucketName, key string,
) error {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return err
	}

	key = strutil.RemoveLeadingSlash(key)

	_, err = s3Client.DeleteObject(
		context.TODO(),
		&s3.DeleteObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		},
	)
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %w", err)
	}

	return nil
}

// S3GetDownloadLink generates a presigned URL for downloading a file from S3
func (Client) S3GetDownloadLink(
	accessKey, secretKey, region, endpoint, bucketName, key string,
	expiration time.Duration,
) (string, error) {
	s3Client, err := createS3Client(
		accessKey, secretKey, region, endpoint,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create S3 client: %w", err)
	}

	presigned, err := s3.NewPresignClient(s3Client).PresignGetObject(
		context.TODO(),
		&s3.GetObjectInput{
			Bucket: aws.String(bucketName),
			Key:    aws.String(key),
		},
		s3.WithPresignExpires(expiration),
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %w", err)
	}

	return presigned.URL, nil
}
