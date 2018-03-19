package ark

import (
	"log"
	"os/user"
)

const (
	MaxIAMUser = 2
	Active     = "Active"
	Inactive   = "Inactive"
	TypeSecure = "SecureString"

	//key for parameter store
	AccessReaderPathKey = "/bei/developers/s3read/access"
	SecretReaderPathKey = "/bei/developers/s3read/secret"
	GradleDir           = "/.gradle/"

	// accessKey and secretKey for map
	AccessKey = "accessKey"
	SecretKey = "secretKey"

	// for storing to env variable
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN     = "AWS_SESSION_TOKEN"
)

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func gradleDir() string {
	return homeDir() + GradleDir
}

func gradlePropertiesPath() string {
	return gradleDir() + "gradle.properties"
}
