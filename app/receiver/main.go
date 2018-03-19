package main

import (
	"log"

	"github.com/Gujarats/ark"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// all config values if not set then the default would be empty "" or false
func main() {
	config := getConfig()

	sess, err := ark.CreateSessionWithProfile(config.Region, config.Profile)
	if err != nil {
		log.Fatal(err)
	}

	creds := stscreds.NewCredentials(sess, config.RoleName)

	svc := ssm.New(sess, &aws.Config{Credentials: creds})

	secretKey, err := ark.GetValueFromParameterStore(svc, config.SecretKey, true)
	if err != nil {
		log.Fatal(err)
	}

	accessKey, err := ark.GetValueFromParameterStore(svc, config.AccessKey, true)
	if err != nil {
		log.Fatal(err)
	}

	if config.UseEnvVariable {
		err := ark.SetEnvVariableAWS(*accessKey.Parameter.Value, *secretKey.Parameter.Value, "")
		if err != nil {
			log.Fatal(err)
		}
	}

	if config.UseGradleProperties {
		configGradle := make(map[string]string)
		configGradle[ark.AccessKey] = config.GradleAccessKey
		configGradle[ark.SecretKey] = config.GradleSecretKey

		err = ark.UpdateGradleProperties(configGradle, *accessKey.Parameter.Value, *secretKey.Parameter.Value)
		if err != nil {
			log.Fatal(err)
		}

	}
}
