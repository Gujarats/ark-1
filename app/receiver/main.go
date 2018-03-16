package main

import "github.com/Gujarats/ark"

// TODO :
// using viper for gettting the setup
func main() {
	config := getConfig()
	sess, err := ark.CreateSessionWithProfile(config.Region, config.Profile)
	if err != nil {
		return err
	}
}
