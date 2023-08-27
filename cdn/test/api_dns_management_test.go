/*
ArvanCloud CDN Services

Testing DNSManagementAPIService

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

func Test_cdn_DNSManagementAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DNSManagementAPIService DnsRecordsCloud", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsCloud(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsDestroy(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsDnsSecShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsDnsSecShow(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsDnsSecUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsDnsSecUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsImport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsImport(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsIndex(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsShow(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsStore(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DNSManagementAPIService DnsRecordsUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var id string

		resp, httpRes, err := apiClient.DNSManagementAPI.DnsRecordsUpdate(context.Background(), domain, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}