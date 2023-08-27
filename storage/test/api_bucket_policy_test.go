/*
ArvanCloud S3 Services

Testing BucketPolicyAPIService

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

func Test_storage_BucketPolicyAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BucketPolicyAPIService DeleteBucketPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.BucketPolicyAPI.DeleteBucketPolicy(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketPolicyAPIService GetBucketPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		resp, httpRes, err := apiClient.BucketPolicyAPI.GetBucketPolicy(context.Background(), bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketPolicyAPIService GetBucketPolicyStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		resp, httpRes, err := apiClient.BucketPolicyAPI.GetBucketPolicyStatus(context.Background(), bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BucketPolicyAPIService PutBucketPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var bucket string

		httpRes, err := apiClient.BucketPolicyAPI.PutBucketPolicy(context.Background(), bucket).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
