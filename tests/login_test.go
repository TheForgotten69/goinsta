package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestImportAccount(t *testing.T) {
	insta, err := getRandomAccount()
	require.NoError(t, err)

	t.Logf("logged into Instagram as user '%s'", insta.Account.Username)
}
