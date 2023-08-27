/*
ArvanCloud CDN Services

Testing RedirectAPIService

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

func Test_cdn_RedirectAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RedirectAPIService RedirectShow", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RedirectAPI.RedirectShow(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RedirectAPIService RedirectUpdate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var domain string

		resp, httpRes, err := apiClient.RedirectAPI.RedirectUpdate(context.Background(), domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
