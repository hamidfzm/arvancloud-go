/*
ArvanCloud CDN Services

Testing ActiveHealthCheckAPIService

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

func Test_cdn_ActiveHealthCheckAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var healthcheck string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckDestroy(context.Background(), domain, healthcheck).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckReportsDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckReportsDetails(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckReportsSummary", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckReportsSummary(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var healthcheck string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckShow(context.Background(), domain, healthcheck).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService ActiveHealthCheckUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var healthcheck string

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.ActiveHealthCheckUpdate(context.Background(), domain, healthcheck).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActiveHealthCheckAPIService HealthChecksZonesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ActiveHealthCheckAPI.HealthChecksZonesIndex(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
