package main

import (
	"log"

	"github.com/Gujarats/ark"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	config := getConfig()

	sess, err := ark.CreateSessionWithProfile(config.Region, config.Profile)
	if err != nil {
		log.Fatal(err)
	}

	svc := ssm.New(sess)

	secretKey, err := ark.GetValueFromParameterStore(svc, config.SecretKey, true)
	if err != nil {
		log.Fatal(err)
	}

	accesKey, err := ark.GetValueFromParameterStore(svc, config.SecretKey, true)
	if err != nil {
		log.Fatal(err)
	}

	configGradle := make(map[string]string)
	configGradle[ark.AccessKey] = config.GradleAccessKey
	configGradle[ark.SecretKey] = config.GradleSecretKey

	err = ark.UpdateGradleProperties(configGradle, *accesKey.Parameter.Value, *secretKey.Parameter.Value)
	if err != nil {
		log.Fatal(err)
	}
}
