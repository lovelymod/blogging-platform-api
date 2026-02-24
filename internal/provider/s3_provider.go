package provider

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type s3Provider struct {
	client     *s3.Client
	domain     string
	bucketName string
}

type S3Provider interface {
	UploadImage(ctx context.Context, fileHeader *multipart.FileHeader, path string) (string, error)
}

func NewS3Provider(s3Client *s3.Client, domain string, bucketName string) S3Provider {
	return &s3Provider{
		client:     s3Client,
		domain:     domain,
		bucketName: bucketName,
	}
}

func (p *s3Provider) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader, path string) (string, error) {
	file, _ := fileHeader.Open()
	defer file.Close()

	objectKey := path + "/" + fileHeader.Filename

	_, err := p.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &p.bucketName,
		Key:    &objectKey,
		Body:   file,
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	publicUrl := fmt.Sprintf("%s/%s", p.domain, objectKey)

	return publicUrl, nil
}
