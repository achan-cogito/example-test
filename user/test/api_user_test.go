/*
Example.com

Testing UserAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package userapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/achan-cogito/example-test"
)

func Test_userapi_UserAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test UserAPIService GetUserByName", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		resp, httpRes, err := apiClient.UserAPI.GetUserByName(context.Background(), username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserAPIService UpdateUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var username string

		httpRes, err := apiClient.UserAPI.UpdateUser(context.Background(), username).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
