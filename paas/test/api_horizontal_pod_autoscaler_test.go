/*
Arvancloud PaaS REST API

Testing HorizontalPodAutoscalerAPIService

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

func Test_paas_HorizontalPodAutoscalerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HorizontalPodAutoscalerAPIService CreateNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.CreateNamespacedHorizontalPodAutoscaler(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService DeleteNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.DeleteNamespacedHorizontalPodAutoscaler(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService DeletecollectionNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.DeletecollectionNamespacedHorizontalPodAutoscaler(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService ListNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.ListNamespacedHorizontalPodAutoscaler(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService PatchNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.PatchNamespacedHorizontalPodAutoscaler(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService PatchNamespacedHorizontalPodAutoscalerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.PatchNamespacedHorizontalPodAutoscalerStatus(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService ReadNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.ReadNamespacedHorizontalPodAutoscaler(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService ReadNamespacedHorizontalPodAutoscalerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.ReadNamespacedHorizontalPodAutoscalerStatus(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService ReplaceNamespacedHorizontalPodAutoscaler", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.ReplaceNamespacedHorizontalPodAutoscaler(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test HorizontalPodAutoscalerAPIService ReplaceNamespacedHorizontalPodAutoscalerStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.HorizontalPodAutoscalerAPI.ReplaceNamespacedHorizontalPodAutoscalerStatus(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
