/*
ArvanCloud S3 Services

Testing ObjectAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package storage

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func Test_storage_ObjectAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectAPIService DeleteObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string
		var key string

		resp, httpRes, err := apiClient.ObjectAPI.DeleteObject(context.Background(), bucket, key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectAPIService GetObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string
		var key string

		resp, httpRes, err := apiClient.ObjectAPI.GetObject(context.Background(), bucket, key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectAPIService HeadObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string
		var key string

		resp, httpRes, err := apiClient.ObjectAPI.HeadObject(context.Background(), bucket, key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectAPIService PutObject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string
		var key string

		resp, httpRes, err := apiClient.ObjectAPI.PutObject(context.Background(), bucket, key).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
