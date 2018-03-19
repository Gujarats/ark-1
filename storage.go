package ark

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/ssm"
)

// Store the keys to parameter store
// the key is /bei/developers/s3read/access for access key
// the key is /bei/developers/s3read/secret for secret key
// to make things easier in the future, the key should be access like a path
func StoreKeys(sess *session.Session, accessKey *iam.AccessKey) error {
	svc := ssm.New(sess)

	err := putParameterStoreKey(svc, AccessReaderPathKey, accessKey.AccessKeyId)
	err = putParameterStoreKey(svc, SecretReaderPathKey, accessKey.SecretAccessKey)

	if err != nil {
		return err
	}

	return nil
}

func putParameterStoreKey(svc *ssm.SSM, key string, value *string) error {
	_, err := svc.PutParameter(&ssm.PutParameterInput{
		Name:      aws.String(key),
		Value:     value,
		Type:      aws.String(TypeSecure),
		Overwrite: aws.Bool(true),
	})

	if err != nil {
		return err
	}

	return nil
}
