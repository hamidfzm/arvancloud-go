/*
Arvancloud PaaS REST API

Testing ConfigMapAPIService

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

func Test_paas_ConfigMapAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ConfigMapAPIService CreateNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ConfigMapAPI.CreateNamespacedConfigMap(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMapAPIService DeleteNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ConfigMapAPI.DeleteNamespacedConfigMap(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMapAPIService DeletecollectionNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ConfigMapAPI.DeletecollectionNamespacedConfigMap(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMapAPIService ListNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.ConfigMapAPI.ListNamespacedConfigMap(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMapAPIService PatchNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ConfigMapAPI.PatchNamespacedConfigMap(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMapAPIService ReadNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ConfigMapAPI.ReadNamespacedConfigMap(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConfigMapAPIService ReplaceNamespacedConfigMap", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.ConfigMapAPI.ReplaceNamespacedConfigMap(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
