package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)


// S3GetObject is a structure wraps s3 client
type S3GetObject struct {
	Client S3GetObjectAPI
}

// NewS3GetObject returns a new S3GetObject object
func NewS3GetObject() (*S3GetObject, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	return &S3GetObject{
		Client: client,
	}, nil
}


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

// CallGetObjectFromS3Interface call GetObjectFromS3 w/o depends on aws
func CallGetObjectFromS3Interface() {
	var bucket = "sample-bucket-thaim"
	var key = "sample-object-key"

	api, err := NewS3GetObject()
	if err != nil {
		log.Fatal(err)
	}

	objectString, err := GetObjectFromS3(context.TODO(), api.Client, bucket, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(objectString))
}

// CallGetObjectFromS3Direct call GetObjectFromS3 using aws sdk
func CallGetObjectFromS3Direct() {
	var bucket = "sample-bucket-thaim"
	var key = "sample-object-key"

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	var api S3GetObjectAPI = s3.NewFromConfig(cfg)

	objectString, err := GetObjectFromS3(context.TODO(), api, bucket, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(objectString))
}

func main() {
	CallGetObjectFromS3Interface()
	CallGetObjectFromS3Direct()
}
