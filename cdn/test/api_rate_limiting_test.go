/*
ArvanCloud CDN Services

Testing RateLimitingAPIService

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

func Test_cdn_RateLimitingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RateLimitingAPIService RateLimitingActionsReprioritize", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingActionsReprioritize(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingRulesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingRulesDestroy(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingRulesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingRulesIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingRulesShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingRulesShow(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingRulesStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingRulesStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingRulesUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingRulesUpdate(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingSettingsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingSettingsIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RateLimitingAPIService RateLimitingSettingsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RateLimitingAPI.RateLimitingSettingsUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
