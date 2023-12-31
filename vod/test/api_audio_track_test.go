/*
Arvan VOD

Testing AudioTrackAPIService

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

func Test_vod_AudioTrackAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AudioTrackAPIService AudioTracksAudioTrackDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var audioTrack string

		httpRes, err := apiClient.AudioTrackAPI.AudioTracksAudioTrackDelete(context.Background(), audioTrack).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioTrackAPIService AudioTracksAudioTrackGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var audioTrack string

		httpRes, err := apiClient.AudioTrackAPI.AudioTracksAudioTrackGet(context.Background(), audioTrack).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioTrackAPIService VideosVideoAudioTracksGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var video string

		httpRes, err := apiClient.AudioTrackAPI.VideosVideoAudioTracksGet(context.Background(), video).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AudioTrackAPIService VideosVideoAudioTracksPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var video string

		httpRes, err := apiClient.AudioTrackAPI.VideosVideoAudioTracksPost(context.Background(), video).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
