/*
ArvanCloud S3 Services

Testing BucketLifecycleConfigurationAPIService

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

func Test_storage_BucketLifecycleConfigurationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BucketLifecycleConfigurationAPIService DeleteBucketLifecycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.BucketLifecycleConfigurationAPI.DeleteBucketLifecycle(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketLifecycleConfigurationAPIService GetBucketLifecycleConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		resp, httpRes, err := apiClient.BucketLifecycleConfigurationAPI.GetBucketLifecycleConfiguration(context.Background(), bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketLifecycleConfigurationAPIService PutBucketLifecycleConfiguration", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.BucketLifecycleConfigurationAPI.PutBucketLifecycleConfiguration(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}