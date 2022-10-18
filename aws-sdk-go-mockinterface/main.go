package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3GetObjectAPI interfaceを定義し、プロダクションコードかテストかに応じて実装をDIできるようにする
type S3GetObjectAPI interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

// GetObjectFromS3 は S3GetObjectAPI interfaceを利用する。 api.GetObjectを呼び出せばよいことだけを知っている
func GetObjectFromS3(ctx context.Context, api S3GetObjectAPI, bucket, key string) ([]byte, error) {
	object, err := api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer object.Body.Close()

	return ioutil.ReadAll(object.Body)
}


func main() {
	var bucket = "sample-bucket-thaim"
	var key = "sample-object-key"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg)
	params := &s3.GetObjectInput{
		Bucket: &bucket,
		Key: &key,
	}

	object, err := client.GetObject(context.Background(), params)
	if err != nil {
		log.Fatal(err)
	}
	defer object.Body.Close()

	objectString, _ := ioutil.ReadAll(object.Body)
	fmt.Println(string(objectString))
}
