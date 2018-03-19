package ark

import (
	"os"
	"syscall"

	"github.com/Gujarats/logger"
)

// this file currently used for set the environment variable

func SetEnvVariableAWS(accessKey, secretKey, tokenKey string) error {
	envAws := make(map[string]string)
	envAws[AWS_ACCESS_KEY_ID] = accessKey
	envAws[AWS_SECRET_ACCESS_KEY] = secretKey
	envAws[AWS_SESSION_TOKEN] = tokenKey

	for key, value := range envAws {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	logger.Debug("AWS_ACCESS_KEY_ID :: ", os.Getenv("AWS_ACCESS_KEY_ID"))
	logger.Debug("AWS_SECRET_ACCESS_KEY :: ", os.Getenv("AWS_SECRET_ACCESS_KEY"))

	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

	return nil
}
