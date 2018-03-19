package ark

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/Gujarats/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

func RetriveKeys(sess *session.Session) {
	ssm.New(sess)
}

type Receiver struct {
	svc *ssm.SSM
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
	file, err := os.OpenFile(gradlePropertiesPath, os.O_RDWR|os.O_CREATE, 0644)
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
	if result == "" {
		return errors.New("data must not be null for writer to build.properties")
	}

	logger.Debug("result :: ", result)

	_, err = file.WriteAt([]byte(result), 0)
	if err != nil {
		return err
	}

	return nil
}

// find specific line and replace with new one
// this function replace the line with the same gradle.properties key with new udpated value access & secret keys
// config here use map to specify :
// key on the gradle eg :
// --- gradle.properties
// secretKey=ExampleSecretKey
// accessKey=ExampleAccessKey
// -- END gralde.properties
// from the value above the key is `secretKey` and use as value for this `config` in this function parameter
func findAndReplace(config map[string]string, accessKeyId string, secretKeyId string, data []byte) string {
	var result string
	var lines []string
	// check if data is empty
	dataSpaceRemoved := removeSpace(string(data))
	// check if data already has keys
	isContainKeys := strings.Contains(string(data), config[AccessKey])
	if len(data) == 0 || dataSpaceRemoved == "" || !isContainKeys {
		accessKeyLine := "\n" + config[AccessKey] + "=" + accessKeyId
		secretKeyLine := config[SecretKey] + "=" + secretKeyId

		// because using "or" condition above there is chance if data has some content on it
		// to avoid truncating file content
		lines = append(lines, string(data))

		lines = append(lines, accessKeyLine)
		lines = append(lines, secretKeyLine)
	} else {
		lines = strings.Split(string(data), "\n")

		for i, line := range lines {
			if strings.Contains(line, config[AccessKey]) {
				lines[i] = config[AccessKey] + "=" + accessKeyId
			} else if strings.Contains(line, config[SecretKey]) {
				lines[i] = config[SecretKey] + "=" + secretKeyId
			}
		}
	}

	result = strings.Join(lines, "\n")

	return result
}

func removeSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
