package tests

import (
	"bytes"
	"encoding/base64"
	"errors"
	"math/rand"
	"os"
	"strings"
	"testing"

	"github.com/TheForgotten69/goinsta/v2"
	"github.com/TheForgotten69/goinsta/v2/utilities"
	"github.com/stretchr/testify/require"
)

func readFromBase64(base64EncodedString string) (*goinsta.Instagram, error) {
	base64Bytes, err := base64.StdEncoding.DecodeString(base64EncodedString)
	if err != nil {
		return nil, err
	}
	return goinsta.ImportReader(bytes.NewReader(base64Bytes))
}

func availableEncodedAccounts() []string {
	output := make([]string, 0)

	environ := os.Environ()
	for _, env := range environ {
		if strings.HasPrefix(env, "INSTAGRAM_BASE64_") {
			index := strings.Index(env, "=")
			encodedString := env[index+1:]
			output = append(output, encodedString)
		}
	}

	return output
}

func getRandomAccount() (*goinsta.Instagram, error) {
	accounts := availableEncodedAccounts()
	if len(accounts) == 0 {
		return nil, errors.New("there is no encoded account in environ")
	}

	encodedAccount := accounts[rand.Intn(len(accounts))]
	return readFromBase64(encodedAccount)
}

func Test_getbase64(t *testing.T) {
	t.Skip()

	// run this test locally to get base64 encoded session that will be exported to env variables and used in tests.

	inst := goinsta.New("username", "password")

	require.NoError(t, inst.Login())

	base64, err := utilities.ExportAsBase64String(inst)
	require.NoError(t, err)

	t.Log(base64)
}
