# \TransportLayerProxyAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransportLayerProxiesDestroy**](TransportLayerProxyAPI.md#TransportLayerProxiesDestroy) | **Delete** /domains/{domain}/transport-layer-proxies/{transportLayerProxyId} | delete a transport layer proxy
[**TransportLayerProxiesIndex**](TransportLayerProxyAPI.md#TransportLayerProxiesIndex) | **Get** /domains/{domain}/transport-layer-proxies | Show list of transport layer proxies for given domain
[**TransportLayerProxiesShow**](TransportLayerProxyAPI.md#TransportLayerProxiesShow) | **Get** /domains/{domain}/transport-layer-proxies/{transportLayerProxyId} | Show a transport layer proxy&#39;s details based on given id
[**TransportLayerProxiesStore**](TransportLayerProxyAPI.md#TransportLayerProxiesStore) | **Post** /domains/{domain}/transport-layer-proxies | Create new transport layer proxy
[**TransportLayerProxiesUpdate**](TransportLayerProxyAPI.md#TransportLayerProxiesUpdate) | **Put** /domains/{domain}/transport-layer-proxies/{transportLayerProxyId} | Update a transport layer proxy



## TransportLayerProxiesDestroy

> MessageResponse TransportLayerProxiesDestroy(ctx, domain, transportLayerProxyId).Execute()

delete a transport layer proxy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    transportLayerProxyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport layer proxy id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportLayerProxyAPI.TransportLayerProxiesDestroy(context.Background(), domain, transportLayerProxyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportLayerProxyAPI.TransportLayerProxiesDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransportLayerProxiesDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportLayerProxyAPI.TransportLayerProxiesDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**transportLayerProxyId** | **string** | Transport layer proxy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransportLayerProxiesDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransportLayerProxiesIndex

> TransportLayerProxiesIndex200Response TransportLayerProxiesIndex(ctx, domain).PerPage(perPage).Page(page).Execute()

Show list of transport layer proxies for given domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportLayerProxyAPI.TransportLayerProxiesIndex(context.Background(), domain).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportLayerProxyAPI.TransportLayerProxiesIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransportLayerProxiesIndex`: TransportLayerProxiesIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `TransportLayerProxyAPI.TransportLayerProxiesIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransportLayerProxiesIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**TransportLayerProxiesIndex200Response**](TransportLayerProxiesIndex200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransportLayerProxiesShow

> TransportLayerProxyResponse TransportLayerProxiesShow(ctx, domain, transportLayerProxyId).Execute()

Show a transport layer proxy's details based on given id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    transportLayerProxyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport layer proxy id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportLayerProxyAPI.TransportLayerProxiesShow(context.Background(), domain, transportLayerProxyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportLayerProxyAPI.TransportLayerProxiesShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransportLayerProxiesShow`: TransportLayerProxyResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportLayerProxyAPI.TransportLayerProxiesShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**transportLayerProxyId** | **string** | Transport layer proxy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransportLayerProxiesShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TransportLayerProxyResponse**](TransportLayerProxyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransportLayerProxiesStore

> TransportLayerProxyResponse TransportLayerProxiesStore(ctx, domain).TransportLayerProxyStore(transportLayerProxyStore).Execute()

Create new transport layer proxy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    transportLayerProxyStore := *openapiclient.NewTransportLayerProxyStore("AppName_example", "sub.example.com", int32(123), "ProxyProtocol_example", "BalanceAlgorithm_example", "FirewallDefaultAction_example") // TransportLayerProxyStore |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportLayerProxyAPI.TransportLayerProxiesStore(context.Background(), domain).TransportLayerProxyStore(transportLayerProxyStore).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportLayerProxyAPI.TransportLayerProxiesStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransportLayerProxiesStore`: TransportLayerProxyResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportLayerProxyAPI.TransportLayerProxiesStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransportLayerProxiesStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transportLayerProxyStore** | [**TransportLayerProxyStore**](TransportLayerProxyStore.md) |  | 

### Return type

[**TransportLayerProxyResponse**](TransportLayerProxyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransportLayerProxiesUpdate

> TransportLayerProxyResponse TransportLayerProxiesUpdate(ctx, domain, transportLayerProxyId).TransportLayerProxyUpdate(transportLayerProxyUpdate).Execute()

Update a transport layer proxy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    transportLayerProxyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport layer proxy id
    transportLayerProxyUpdate := *openapiclient.NewTransportLayerProxyUpdate("AppName_example", "ProxyProtocol_example", "BalanceAlgorithm_example") // TransportLayerProxyUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransportLayerProxyAPI.TransportLayerProxiesUpdate(context.Background(), domain, transportLayerProxyId).TransportLayerProxyUpdate(transportLayerProxyUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransportLayerProxyAPI.TransportLayerProxiesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransportLayerProxiesUpdate`: TransportLayerProxyResponse
    fmt.Fprintf(os.Stdout, "Response from `TransportLayerProxyAPI.TransportLayerProxiesUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**transportLayerProxyId** | **string** | Transport layer proxy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransportLayerProxiesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **transportLayerProxyUpdate** | [**TransportLayerProxyUpdate**](TransportLayerProxyUpdate.md) |  | 

### Return type

[**TransportLayerProxyResponse**](TransportLayerProxyResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

