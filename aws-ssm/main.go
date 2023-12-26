package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	params := "sample"
	if len(os.Args) == 2 {
		params = os.Args[1]
	}

	fmt.Println("Get data from SSM: " + params)

	cfg, _ := config.LoadDefaultConfig(context.TODO())
	api := ssm.NewFromConfig(cfg)

	output, _ := api.GetParameter(context.TODO(), &ssm.GetParameterInput{
		Name:           aws.String(params),
		WithDecryption: aws.Bool(true),
	})

	fmt.Println(*output.Parameter.Value)
}
