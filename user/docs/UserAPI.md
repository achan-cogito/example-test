# \UserAPI

All URIs are relative to *http://example.com/api/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserByName**](UserAPI.md#GetUserByName) | **Get** /users/{username} | Get user by user name
[**UpdateUser**](UserAPI.md#UpdateUser) | **Put** /users/{username} | Updated user



## GetUserByName

> User GetUserByName(ctx, username).PrettyPrint(prettyPrint).WithEmail(withEmail).Execute()

Get user by user name



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/achan-cogito/user-api-go"
)

func main() {
	username := "username_example" // string | The name that needs to be fetched
	prettyPrint := true // bool | Pretty print response (optional)
	withEmail := true // bool | Filter users without email (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAPI.GetUserByName(context.Background(), username).PrettyPrint(prettyPrint).WithEmail(withEmail).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.GetUserByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserByName`: User
	fmt.Fprintf(os.Stdout, "Response from `UserAPI.GetUserByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The name that needs to be fetched | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **prettyPrint** | **bool** | Pretty print response | 
 **withEmail** | **bool** | Filter users without email | 

### Return type

[**User**](User.md)

### Authorization

[api_key](../README.md#api_key), [main_auth](../README.md#main_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUser

> UpdateUser(ctx, username).User(user).PrettyPrint(prettyPrint).Execute()

Updated user



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/achan-cogito/user-api-go"
)

func main() {
	username := "username_example" // string | The name that needs to be updated
	user := *openapiclient.NewUser() // User | Updated user object
	prettyPrint := true // bool | Pretty print response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.UserAPI.UpdateUser(context.Background(), username).User(user).PrettyPrint(prettyPrint).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAPI.UpdateUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The name that needs to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | Updated user object | 
 **prettyPrint** | **bool** | Pretty print response | 

### Return type

 (empty response body)

### Authorization

[main_auth](../README.md#main_auth)

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

