package bootstrap

import (
	"blogging-platform-api/internal/entity"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	s3Config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func SetUpS3(config *entity.Config) *s3.Client {
	s3Cfg, err := s3Config.LoadDefaultConfig(context.TODO(),
		s3Config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.R2_ACCESSKEY_ID, config.R2_ACCESSKEY_SECRET, "")),
		s3Config.WithRegion("auto"), // Required by SDK but not used by R2
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(s3Cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", config.R2_ACCOUNT_ID))
	})

	return client
}
