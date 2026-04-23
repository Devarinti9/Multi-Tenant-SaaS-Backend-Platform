package storage

import (
	"bytes"
	"context"
	"fmt"

	"multi-tenant-saas-backend-platform/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	credentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(cfg config.Config) (*s3.Client, error) {
	var awsCfg aws.Config
	var err error

	if cfg.S3AccessKey != "" && cfg.S3SecretKey != "" {
		awsCfg, err = awsConfig.LoadDefaultConfig(
			context.Background(),
			awsConfig.WithRegion(cfg.AWSRegion),
			awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.S3AccessKey, cfg.S3SecretKey, "")),
		)
	} else {
		awsCfg, err = awsConfig.LoadDefaultConfig(context.Background(), awsConfig.WithRegion(cfg.AWSRegion))
	}
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		if cfg.S3Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.S3Endpoint)
			o.UsePathStyle = cfg.UsePathStyleS3
		}
	})

	return client, nil
}

func UploadOrganizationManifest(ctx context.Context, client *s3.Client, bucket string, organization string, payload []byte) (string, error) {
	key := fmt.Sprintf("organizations/%s/manifest.json", organization)
	_, err := client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   bytes.NewReader(payload),
	})
	if err != nil {
		return "", err
	}
	return key, nil
}
