package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	fmt.Println("Get data from SSM: ")

	cfg, _ := config.LoadDefaultConfig(context.TODO())
	api := ssm.NewFromConfig(cfg)

	output, _ := api.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String("sample"),
		WithDecryption: aws.Bool(true),
	})

	fmt.Println(*output.Parameter.Value)
}
