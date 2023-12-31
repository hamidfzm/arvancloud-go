/*
ArvanCloud S3 Services

Testing BucketCORSAPIService

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

func Test_storage_BucketCORSAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BucketCORSAPIService DeleteBucketCors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.BucketCORSAPI.DeleteBucketCors(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketCORSAPIService GetBucketCors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		resp, httpRes, err := apiClient.BucketCORSAPI.GetBucketCors(context.Background(), bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketCORSAPIService PutBucketCors", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.BucketCORSAPI.PutBucketCors(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
