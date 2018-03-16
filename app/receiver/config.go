package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Region            string `viper:"region"`
	Profile           string `viper:"profile"`
	AwsConfigPath     string `viper:aws_config`
	AwsCredentialPath string `viper:aws_credential`
}

const (
	pathConfig = ".ark"
)

func getConfig() *Config {
	viper.AddConfigPath("$HOME/" + pathConfig)
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file must exist in ~/"+pathConfig+"./config.yaml: %s \n", err))
	}

	// read the config file to struct
	config := &Config{
		GradleCacheDir: viper.GetString("gradle"),
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarhsal config struct : %s \n", err))
	}

	return config
}
