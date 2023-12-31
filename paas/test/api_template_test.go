/*
Arvancloud PaaS REST API

Testing TemplateAPIService

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

func Test_paas_TemplateAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TemplateAPIService DeleteNamespacedTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.TemplateAPI.DeleteNamespacedTemplate(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService ListNamespacedTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string

		resp, httpRes, err := apiClient.TemplateAPI.ListNamespacedTemplate(context.Background(), namespace).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService ListTemplateForAllNamespaces", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TemplateAPI.ListTemplateForAllNamespaces(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService PatchNamespacedTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.TemplateAPI.PatchNamespacedTemplate(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService ReadNamespacedTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.TemplateAPI.ReadNamespacedTemplate(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplateAPIService ReplaceNamespacedTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var namespace string
		var name string

		resp, httpRes, err := apiClient.TemplateAPI.ReplaceNamespacedTemplate(context.Background(), namespace, name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
