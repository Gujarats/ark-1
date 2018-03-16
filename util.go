package ark

import (
	"log"
	"os/user"
	"syscall"
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
)

const (
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	O_TRUNC  int = syscall.O_TRUNC
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
