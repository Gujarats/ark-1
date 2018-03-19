package ark

import (
	"log"
	"strings"
	"testing"
)

func TestFindAndReplace(t *testing.T) {
	testObjects := []struct {
		config      map[string]string
		accessKeyId string
		secretKeyId string
		data        []byte
		expected    string
	}{
		// 0
		{
			config: map[string]string{
				AccessKey: "AccesKeyFromConfig",
				SecretKey: "SecretKeyFromConfig",
			},

			accessKeyId: "12345",
			secretKeyId: "1234",

			data: []byte(`
				AccesKeyFromConfig=0000
				SecretKeyFromConfig=0000
			`),

			expected: string(
				[]byte("\nAccesKeyFromConfig=12345\nSecretKeyFromConfig=1234"),
			),
		},

		// 1 empty data []byte
		{
			config: map[string]string{
				AccessKey: "AccesKeyFromConfig",
				SecretKey: "SecretKeyFromConfig",
			},
			accessKeyId: "12345",
			secretKeyId: "1234",
			data:        []byte(``),
			expected: string(
				[]byte("\nAccesKeyFromConfig=12345\nSecretKeyFromConfig=1234"),
			),
		},

		// 2 already has some data
		{
			config: map[string]string{
				AccessKey: "AccesKeyFromConfig",
				SecretKey: "SecretKeyFromConfig",
			},
			accessKeyId: "12345",
			secretKeyId: "1234",
			data: []byte(`#some data here
				AccesKeyFromConfig=0000
				SecretKeyFromConfig=0000`),

			expected: string(
				[]byte("#some data here\nAccesKeyFromConfig=12345\nSecretKeyFromConfig=1234"),
			),
		},
	}

	for _, testObject := range testObjects {
		result := findAndReplace(testObject.config, testObject.accessKeyId, testObject.secretKeyId, testObject.data)
		if !strings.Contains(result, testObject.expected) {
			log.Fatalf("expected = %+v, result = %+v\n", testObject.expected, result)
		}
	}
}
