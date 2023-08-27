/*
Arvan LIVE

Testing DomainAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package live

import (
	"context"
	openapiclient "github.com/hamidfzm/arvancloud-go/live"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_live_DomainAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DomainAPIService DomainGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DomainAPI.DomainGet(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DomainAPIService DomainPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DomainAPI.DomainPost(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}