package tests

import (
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

const picsumurl = "https://picsum.photos/800/800"

func TestUploadPhoto(t *testing.T) {
	insta, err := getRandomAccount()
	require.NoError(t, err)

	log.Println("Download random photo")

	var client http.Client
	request, err := http.NewRequest(http.MethodGet, picsumurl, nil)
	require.NoError(t, err)

	resp, err := client.Do(request)
	require.NoError(t, err)

	defer resp.Body.Close()

	postedPhoto, err := insta.UploadPhoto(resp.Body, "awesome photo test! :)", 1, 1)
	require.NoError(t, err)

	log.Printf("Success upload photo %s", postedPhoto.ID)
}

func TestUploadAlbum(t *testing.T) {
	insta, err := getRandomAccount()
	require.NoError(t, err)

	log.Println("Download random photo")

	var (
		responses []*http.Response
		photos    []io.Reader
	)

	var client http.Client
	for i := 0; i < 3; i++ {
		request, err := http.NewRequest(http.MethodGet, picsumurl, nil)
		require.NoError(t, err)

		resp, err := client.Do(request)
		require.NoError(t, err)

		responses = append(responses, resp)
		photos = append(photos, resp.Body)
	}

	t.Cleanup(func() {
		for i := range responses {
			require.NoError(t, responses[i].Body.Close())
		}
	})

	postedPhoto, err := insta.UploadAlbum(photos, "awesome photo album test! :)", 1, 1)
	require.NoError(t, err)

	log.Printf("Success upload album %s", postedPhoto.ID)
}
