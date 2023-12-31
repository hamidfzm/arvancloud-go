/*
Arvan VOD

Testing WatermarkAPIService

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

func Test_vod_WatermarkAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WatermarkAPIService ChannelsChannelWatermarksGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel string

		httpRes, err := apiClient.WatermarkAPI.ChannelsChannelWatermarksGet(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WatermarkAPIService ChannelsChannelWatermarksPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var channel string

		httpRes, err := apiClient.WatermarkAPI.ChannelsChannelWatermarksPost(context.Background(), channel).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WatermarkAPIService WatermarksWatermarkDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var watermark string

		httpRes, err := apiClient.WatermarkAPI.WatermarksWatermarkDelete(context.Background(), watermark).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WatermarkAPIService WatermarksWatermarkGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var watermark string

		httpRes, err := apiClient.WatermarkAPI.WatermarksWatermarkGet(context.Background(), watermark).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WatermarkAPIService WatermarksWatermarkPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var watermark string

		httpRes, err := apiClient.WatermarkAPI.WatermarksWatermarkPatch(context.Background(), watermark).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
