package main

import (
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kelseyhightower/envconfig"
)

// NOTE : all the tag use lower case but
// in the environtment variable must use Upper case
// eg : `envconfig:region` in environment variable use ARTIFACT_REGION
// we USE ARTIFACT for the prefix
type Environment struct {
	// speficy region eg : ap-southeast-1
	Region string `envconfig:"region"`

	// speficy single IAM user
	User string `envconfig:"user"`

	// speficy multiple IAM user
	Users []string `envconfig:"users"`

	// speficy multiple GROUP of aws using ARN
	Group string `envconfig:"group"`
}

var env Environment

func init() {
	err := envconfig.Process("MYAPP", &env)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	lambda.Start(CreateAccessKey)
}
