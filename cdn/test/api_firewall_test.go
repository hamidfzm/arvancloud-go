/*
ArvanCloud CDN Services

Testing FirewallAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cdn

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func Test_cdn_FirewallAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FirewallAPIService FirewallReprioritize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallReprioritize(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallRulesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallRulesDestroy(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallRulesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallRulesIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallRulesShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallRulesShow(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallRulesStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallRulesStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallRulesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallRulesUpdate(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallSettingsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallSettingsIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FirewallAPIService FirewallSettingsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.FirewallAPI.FirewallSettingsUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
