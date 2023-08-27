/*
ArvanCloud CDN Services

Testing EmailForwardingAPIService

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

func Test_cdn_EmailForwardingAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EmailForwardingAPIService EmailForwardingsActivate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsActivate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsAliasesDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingAliasId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesDestroy(context.Background(), domain, emailForwardingId, emailForwardingAliasId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsAliasesIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesIndex(context.Background(), domain, emailForwardingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsAliasesStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesStore(context.Background(), domain, emailForwardingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsAliasesToggleActivation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingAliasId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesToggleActivation(context.Background(), domain, emailForwardingId, emailForwardingAliasId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsAliasesUpdateRecipients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingAliasId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsAliasesUpdateRecipients(context.Background(), domain, emailForwardingId, emailForwardingAliasId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsCatchAllActivate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsCatchAllActivate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsCatchAllDeactivate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsCatchAllDeactivate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsDeactivate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsDeactivate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsRecipientsDestroy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingRecipientId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsDestroy(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsRecipientsIndex", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsIndex(context.Background(), domain, emailForwardingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsRecipientsResendVerification", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingRecipientId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsResendVerification(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsRecipientsSetDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingRecipientId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsSetDefault(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsRecipientsStore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsStore(context.Background(), domain, emailForwardingId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsRecipientsVerify", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string
		var emailForwardingId string
		var emailForwardingRecipientId string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsRecipientsVerify(context.Background(), domain, emailForwardingId, emailForwardingRecipientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EmailForwardingAPIService EmailForwardingsStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.EmailForwardingAPI.EmailForwardingsStats(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}