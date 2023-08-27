# \NetworksAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachNetworkToServer**](NetworksAPI.md#AttachNetworkToServer) | **Patch** /regions/:region/networks/:id/attach | 
[**DetachNetworkFromServer**](NetworksAPI.md#DetachNetworkFromServer) | **Patch** /regions/:region/networks/:id/detach | 
[**GetAllNetworks**](NetworksAPI.md#GetAllNetworks) | **Get** /regions/:region/networks | 



## AttachNetworkToServer

> MessageResponse AttachNetworkToServer(ctx, region, id).AttachServerToNetworkRequest(attachServerToNetworkRequest).Execute()





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
    id := "id_example" // string | network id
    attachServerToNetworkRequest := *openapiclient.NewAttachServerToNetworkRequest() // AttachServerToNetworkRequest | id of the server which needs to be detached

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksAPI.AttachNetworkToServer(context.Background(), region, id).AttachServerToNetworkRequest(attachServerToNetworkRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.AttachNetworkToServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachNetworkToServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.AttachNetworkToServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | network id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachNetworkToServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachServerToNetworkRequest** | [**AttachServerToNetworkRequest**](AttachServerToNetworkRequest.md) | id of the server which needs to be detached | 

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


## DetachNetworkFromServer

> MessageResponse DetachNetworkFromServer(ctx, region, id).DetachNetworkFromServerRequest(detachNetworkFromServerRequest).Execute()





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
    id := "id_example" // string | network id
    detachNetworkFromServerRequest := *openapiclient.NewDetachNetworkFromServerRequest() // DetachNetworkFromServerRequest | id of the server which needs to be detached

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworksAPI.DetachNetworkFromServer(context.Background(), region, id).DetachNetworkFromServerRequest(detachNetworkFromServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DetachNetworkFromServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachNetworkFromServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DetachNetworkFromServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | network id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachNetworkFromServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **detachNetworkFromServerRequest** | [**DetachNetworkFromServerRequest**](DetachNetworkFromServerRequest.md) | id of the server which needs to be detached | 

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


## GetAllNetworks

> Network GetAllNetworks(ctx, region).Execute()





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
    resp, r, err := apiClient.NetworksAPI.GetAllNetworks(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetAllNetworks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllNetworks`: Network
    fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetAllNetworks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllNetworksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Network**](Network.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

