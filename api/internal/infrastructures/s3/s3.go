package s3

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/boke0ya/beathub-api/internal/adapters"
	"github.com/boke0ya/beathub-api/internal/errors"
)

type S3Adapter struct {
	publicHost   string
	privateHost  string
	bucket       string
	accessKeyId  string
	accessSecret string
}

func NewS3Adapter(privateHost string, publicHost string, bucket string, accessKeyId string, accessSecret string) BucketAdapter {
	return S3Adapter{
		publicHost:   publicHost,
		privateHost:  privateHost,
		bucket:       bucket,
		accessKeyId:  accessKeyId,
		accessSecret: accessSecret,
	}
}

func (s S3Adapter) GetS3Client(host string) (*s3.Client, error) {
	creds := credentials.NewStaticCredentialsProvider(
		s.accessKeyId,
		s.accessSecret,
		"",
	)
	var cli *s3.Client
	if os.Getenv("ENV") == "development" {
		resolver := aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL: host + "/",
			}, nil
		})
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithEndpointResolver(resolver),
			config.WithCredentialsProvider(creds),
		)
		if err != nil {
			return nil, errors.New(errors.FailedToPersistBucketObject, err)
		}
		cli = s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		})
	} else {
		cfg, err := config.LoadDefaultConfig(
			context.TODO(),
			config.WithRegion("ap-northeast-1"),
			config.WithCredentialsProvider(creds),
		)
		if err != nil {
			return nil, errors.New(errors.FailedToPersistBucketObject, err)
		}
		cli = s3.NewFromConfig(cfg)
	}
	return cli, nil
}

func (a S3Adapter) GetS3PresignedClient() (*s3.PresignClient, error) {
	cli, err := a.GetS3Client(a.publicHost)
	if err != nil {
		return nil, errors.New(errors.FailedToPersistBucketObject, err)
	}
	psCli := s3.NewPresignClient(cli)
	return psCli, nil
}

func (s S3Adapter) CreatePutObjectUrl(key string) (string, error) {
	cli, err := s.GetS3PresignedClient()
	if err != nil {
		return "", errors.New(errors.FailedToPersistBucketObject, err)
	}
	req, err := cli.PresignPutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	})
	if err != nil {
		return "", errors.New(errors.FailedToPersistBucketObject, err)
	} else {
		return req.URL, nil
	}
}

func (s S3Adapter) GetObjectUrl(key string) string {
	if os.Getenv("ENV") == "development" {
		return s.publicHost + "/" + s.bucket + "/" + key
	} else {
		return fmt.Sprintf("https://%s.s3.ap-northeast-1.amazonaws.com/%s", s.bucket, key)
	}
}

func (s S3Adapter) DeleteObject(key string) error {
	cli, err := s.GetS3Client(s.privateHost)
	if err != nil {
		return errors.New(errors.FailedToPersistBucketObject, err)
	}
	input := &s3.DeleteObjectInput{
		Bucket: &s.bucket,
		Key:    &key,
	}
	if _, err = cli.DeleteObject(context.TODO(), input); err != nil {
		return errors.New(errors.FailedToPersistBucketObject, err)
	}
	return nil
}
