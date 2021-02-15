package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchUser(t *testing.T) {
	count := 20

	insta, err := getRandomAccount()
	require.NoError(t, err)

	result, err := insta.Search.User("a", count)
	require.NoError(t, err)

	require.Equal(t, "ok", result.Status)

	t.Logf("result length is %d", len(result.Users))

	for _, user := range result.Users {
		t.Logf("user %s with id %d\n", user.Username, user.ID)
	}
}
