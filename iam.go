package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
)

func CreateAccessKey() (string, error) {

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(env.Region),
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	// Create a IAM service client.
	svc := iam.New(sess)

	result, err := svc.CreateAccessKey(&iam.CreateAccessKeyInput{
		UserName: aws.String(env.User),
	})

	if err != nil {
		log.Fatal(err)
	}

	return result.AccessKey.String(), nil

}
