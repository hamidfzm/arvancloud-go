/*
Arvancloud PaaS REST API

Testing ServiceAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package paas

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func Test_paas_ServiceAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceAPIService CreateNamespacedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ServiceAPI.CreateNamespacedService(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService DeleteNamespacedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.DeleteNamespacedService(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ListNamespacedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ServiceAPI.ListNamespacedService(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService PatchNamespacedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.PatchNamespacedService(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ReadNamespacedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.ReadNamespacedService(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceAPIService ReplaceNamespacedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ServiceAPI.ReplaceNamespacedService(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
