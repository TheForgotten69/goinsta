package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFeedTagLike(t *testing.T) {
	insta, err := getRandomAccount()
	require.NoError(t, err)

	feedTag, err := insta.Feed.Tags("golang")
	require.NoError(t, err)

	for _, item := range feedTag.RankedItems {
		// media, err := insta.GetMedia(item.ID)
		// if err != nil {
		// 	t.Fatal(err)
		// 	return
		// }
		// err = media.Items[0].Like()

		require.NoError(t, item.Like())

		t.Logf("media %s liked by goinsta", item.ID)
	}
}

func TestFeedTagNext(t *testing.T) {
	insta, err := getRandomAccount()
	require.NoError(t, err)

	feedTag, err := insta.Feed.Tags("golang")
	require.NoError(t, err)

	initNextID := feedTag.NextID
	require.True(t, feedTag.Next(), "Failed to fetch next page")

	assert.Equal(t, "ok", feedTag.Status)

	gotNextID := feedTag.NextID
	assert.NotEqual(t, initNextID, gotNextID, "NextID must differ after FeedTag.Next() call")
}
