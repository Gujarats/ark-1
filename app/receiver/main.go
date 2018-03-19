package main

import (
	"log"

	"github.com/Gujarats/ark"
	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func main() {
	config := getConfig()

	sess, err := ark.CreateSessionWithProfile(config.Region, config.Profile)
	if err != nil {
		logger.Debug("error :: ", err.Error())
		log.Fatal(err)
	}

	creds := stscreds.NewCredentials(sess, config.RoleName)

	svc := ssm.New(sess, &aws.Config{Credentials: creds})

	secretKey, err := ark.GetValueFromParameterStore(svc, config.SecretKey, true)
	if err != nil {
		logger.Debug("error :: ", err.Error())
		log.Fatal(err)
	}

	accessKey, err := ark.GetValueFromParameterStore(svc, config.AccessKey, true)
	if err != nil {
		logger.Debug("error :: ", err.Error())
		log.Fatal(err)
	}

	configGradle := make(map[string]string)
	configGradle[ark.AccessKey] = config.GradleAccessKey
	configGradle[ark.SecretKey] = config.GradleSecretKey

	logger.Debugf("configGradle :: ", "%+v\n", configGradle)
	logger.Debugf("accessKey value :: ", "%+v\n", *accessKey.Parameter.Value)
	logger.Debugf("accessKey value :: ", "%+v\n", *secretKey.Parameter.Value)

	err = ark.UpdateGradleProperties(configGradle, *accessKey.Parameter.Value, *secretKey.Parameter.Value)
	if err != nil {
		logger.Debug("error :: ", err.Error())
		log.Fatal(err)
	}
}
