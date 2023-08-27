# \SSHKeysAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSSHKey**](SSHKeysAPI.md#CreateSSHKey) | **Post** /regions/:region/ssh-keys | 
[**DeleteSSHKey**](SSHKeysAPI.md#DeleteSSHKey) | **Post** /regions/:region/ssh-keys/:name | 
[**GetAllSSHKeys**](SSHKeysAPI.md#GetAllSSHKeys) | **Get** /regions/:region/ssh-keys | 
[**GetSSHKey**](SSHKeysAPI.md#GetSSHKey) | **Get** /regions/:region/ssh-keys/:name | 



## CreateSSHKey

> MessageResponse CreateSSHKey(ctx, region).CreateSSHKeyRequest(createSSHKeyRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    createSSHKeyRequest := *openapiclient.NewCreateSSHKeyRequest() // CreateSSHKeyRequest | info about the new SSH key to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysAPI.CreateSSHKey(context.Background(), region).CreateSSHKeyRequest(createSSHKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.CreateSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSSHKey`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.CreateSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createSSHKeyRequest** | [**CreateSSHKeyRequest**](CreateSSHKeyRequest.md) | info about the new SSH key to be created | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSSHKey

> MessageResponse DeleteSSHKey(ctx, region, name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    name := "name_example" // string | SSH key name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysAPI.DeleteSSHKey(context.Background(), region, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.DeleteSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSSHKey`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.DeleteSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**name** | **string** | SSH key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSSHKeys

> SSHKey GetAllSSHKeys(ctx, region).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysAPI.GetAllSSHKeys(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.GetAllSSHKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSSHKeys`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.GetAllSSHKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSSHKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSSHKey

> SSHKey GetSSHKey(ctx, region, name).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    name := "name_example" // string | SSH key name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SSHKeysAPI.GetSSHKey(context.Background(), region, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SSHKeysAPI.GetSSHKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSSHKey`: SSHKey
    fmt.Fprintf(os.Stdout, "Response from `SSHKeysAPI.GetSSHKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**name** | **string** | SSH key name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSSHKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SSHKey**](SSHKey.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

