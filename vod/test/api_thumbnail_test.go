/*
Arvan VOD

Testing ThumbnailAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package vod

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func Test_vod_ThumbnailAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ThumbnailAPIService VideosVideoThumbnailGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var video string

		httpRes, err := apiClient.ThumbnailAPI.VideosVideoThumbnailGet(context.Background(), video).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThumbnailAPIService VideosVideoThumbnailPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var video string

		httpRes, err := apiClient.ThumbnailAPI.VideosVideoThumbnailPost(context.Background(), video).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
