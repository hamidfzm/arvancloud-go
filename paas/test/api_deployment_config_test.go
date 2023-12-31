/*
Arvancloud PaaS REST API

Testing DeploymentConfigAPIService

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

func Test_paas_DeploymentConfigAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeploymentConfigAPIService CreateNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.CreateNamespacedDeploymentConfig(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService DeleteNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.DeleteNamespacedDeploymentConfig(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService DeletecollectionNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.DeletecollectionNamespacedDeploymentConfig(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService ListNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.ListNamespacedDeploymentConfig(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService PatchNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.PatchNamespacedDeploymentConfig(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService PatchNamespacedDeploymentConfigScale", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.PatchNamespacedDeploymentConfigScale(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService ReadNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.ReadNamespacedDeploymentConfig(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService ReadNamespacedDeploymentConfigScale", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.ReadNamespacedDeploymentConfigScale(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService ReplaceNamespacedDeploymentConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.ReplaceNamespacedDeploymentConfig(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeploymentConfigAPIService ReplaceNamespacedDeploymentConfigScale", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.DeploymentConfigAPI.ReplaceNamespacedDeploymentConfigScale(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
