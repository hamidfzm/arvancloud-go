/*

Testing PlansAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package iaas

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func Test_iaas_PlansAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PlansAPIService GetAllRegionPlans", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string

		resp, httpRes, err := apiClient.PlansAPI.GetAllRegionPlans(context.Background(), region).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PlansAPIService GetPlanByID", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var region string
		var id string

		resp, httpRes, err := apiClient.PlansAPI.GetPlanByID(context.Background(), region, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}