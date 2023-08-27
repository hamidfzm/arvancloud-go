/*
Arvancloud PaaS REST API

Testing RouteAPIService

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

func Test_paas_RouteAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RouteAPIService CreateNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.RouteAPI.CreateNamespacedRoute(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService DeleteNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.DeleteNamespacedRoute(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService DeletecollectionNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.RouteAPI.DeletecollectionNamespacedRoute(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService ListNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.RouteAPI.ListNamespacedRoute(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService PatchNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.PatchNamespacedRoute(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService PatchNamespacedRouteStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.PatchNamespacedRouteStatus(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService ReadNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.ReadNamespacedRoute(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService ReadNamespacedRouteStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.ReadNamespacedRouteStatus(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService ReplaceNamespacedRoute", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.ReplaceNamespacedRoute(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RouteAPIService ReplaceNamespacedRouteStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.RouteAPI.ReplaceNamespacedRouteStatus(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}