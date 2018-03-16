package ark

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func RetriveKeys(sess *session.Session) {
	ssm.New(sess)
}

// get the keys from parameter store
func GetValueFromParameterStore(svc *ssm.SSM, keyName string, decrypt bool) (*ssm.GetParameterOutput, error) {
	result, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(keyName),
		WithDecryption: aws.Bool(decrypt),
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// NOTE :  be aware if gradle.properties file size too large it will return an error
func UpdateGradleProperties(configKey map[string]string, accessKeyId string, secretKey string) error {
	gradlePropertiesPath := gradlePropertiesPath()

	// open the file
	// create if not exist
	file, err := os.OpenFile(gradlePropertiesPath, O_RDWR|O_CREATE|O_TRUNC, 0644)
	if err != nil {
		return err
	}

	// read the content using buffer
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	// find and replace data
	result := findAndReplace(configKey, accessKeyId, secretKey, data)

	_, err = file.Write([]byte(result))
	if err != nil {
		return err
	}

	return nil
}

func findAndReplace(config map[string]string, accessKeyId string, secretKeyId string, data []byte) string {
	lines := strings.Split(string(data), "\n")

	for i, line := range lines {
		if strings.Contains(line, config[AccessKey]) {
			lines[i] = config[AccessKey] + "=" + accessKeyId
		} else if strings.Contains(line, config[SecretKey]) {
			lines[i] = config[SecretKey] + "=" + secretKeyId
		}
	}

	result := strings.Join(lines, "\n")

	return result
}
