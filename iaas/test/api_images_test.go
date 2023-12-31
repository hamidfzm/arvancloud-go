/*

Testing ImagesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package iaas

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func Test_iaas_ImagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesAPIService AppendDataToImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string

		resp, httpRes, err := apiClient.ImagesAPI.AppendDataToImage(context.Background(), region).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService DeleteImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string
		var id string

		resp, httpRes, err := apiClient.ImagesAPI.DeleteImage(context.Background(), region, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService GetImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string

		resp, httpRes, err := apiClient.ImagesAPI.GetImages(context.Background(), region).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService ImportImageFromURL", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string

		resp, httpRes, err := apiClient.ImagesAPI.ImportImageFromURL(context.Background(), region).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService UploadImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string

		resp, httpRes, err := apiClient.ImagesAPI.UploadImage(context.Background(), region).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
