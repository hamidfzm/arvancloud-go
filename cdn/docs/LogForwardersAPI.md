# \LogForwardersAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LogForwardersDestroy**](LogForwardersAPI.md#LogForwardersDestroy) | **Delete** /domains/{domain}/log-forwarders/{logForwarderId} | delete a log forwarder
[**LogForwardersIndex**](LogForwardersAPI.md#LogForwardersIndex) | **Get** /domains/{domain}/log-forwarders | Show list of log forwarders for given domain
[**LogForwardersShow**](LogForwardersAPI.md#LogForwardersShow) | **Get** /domains/{domain}/log-forwarders/{logForwarderId} | Show a log forwarder&#39;s details based on given id
[**LogForwardersStore**](LogForwardersAPI.md#LogForwardersStore) | **Post** /domains/{domain}/log-forwarders | Create new log forwarder
[**LogForwardersUpdate**](LogForwardersAPI.md#LogForwardersUpdate) | **Put** /domains/{domain}/log-forwarders/{logForwarderId} | Update a log forwarder
[**LogForwardersUpdateStatus**](LogForwardersAPI.md#LogForwardersUpdateStatus) | **Patch** /domains/{domain}/log-forwarders/{logForwarderId}/status | Update a log forwarder&#39;s status



## LogForwardersDestroy

> MessageResponse LogForwardersDestroy(ctx, domain, logForwarderId).Execute()

delete a log forwarder

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
    logForwarderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Log Forwarder Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogForwardersAPI.LogForwardersDestroy(context.Background(), domain, logForwarderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogForwardersAPI.LogForwardersDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogForwardersDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `LogForwardersAPI.LogForwardersDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**logForwarderId** | **string** | Log Forwarder Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogForwardersDestroyRequest struct via the builder pattern


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


## LogForwardersIndex

> LogForwardersIndex200Response LogForwardersIndex(ctx, domain).PerPage(perPage).Page(page).Execute()

Show list of log forwarders for given domain

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
    resp, r, err := apiClient.LogForwardersAPI.LogForwardersIndex(context.Background(), domain).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogForwardersAPI.LogForwardersIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogForwardersIndex`: LogForwardersIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `LogForwardersAPI.LogForwardersIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogForwardersIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**LogForwardersIndex200Response**](LogForwardersIndex200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogForwardersShow

> LogForwarderResponse LogForwardersShow(ctx, domain, logForwarderId).Execute()

Show a log forwarder's details based on given id

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
    logForwarderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Log Forwarder Id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogForwardersAPI.LogForwardersShow(context.Background(), domain, logForwarderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogForwardersAPI.LogForwardersShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogForwardersShow`: LogForwarderResponse
    fmt.Fprintf(os.Stdout, "Response from `LogForwardersAPI.LogForwardersShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**logForwarderId** | **string** | Log Forwarder Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogForwardersShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**LogForwarderResponse**](LogForwarderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogForwardersStore

> LogForwarderResponse LogForwardersStore(ctx, domain).LogForwarder(logForwarder).Execute()

Create new log forwarder

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
    logForwarder := *openapiclient.NewLogForwarder("Name_example", "Description_example", "Type_example", "ConnectionType_example", openapiclient.LogForwarderDataFormat{LogForwarderAccessLogType: openapiclient.NewLogForwarderAccessLogType()}, openapiclient.LogForwarderSetting{LogForwarderDatadogConnectionType: openapiclient.NewLogForwarderDatadogConnectionType()}, false) // LogForwarder |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogForwardersAPI.LogForwardersStore(context.Background(), domain).LogForwarder(logForwarder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogForwardersAPI.LogForwardersStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogForwardersStore`: LogForwarderResponse
    fmt.Fprintf(os.Stdout, "Response from `LogForwardersAPI.LogForwardersStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogForwardersStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **logForwarder** | [**LogForwarder**](LogForwarder.md) |  | 

### Return type

[**LogForwarderResponse**](LogForwarderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogForwardersUpdate

> LogForwarderResponse LogForwardersUpdate(ctx, domain, logForwarderId).LogForwarder(logForwarder).Execute()

Update a log forwarder

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
    logForwarderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Log Forwarder Id
    logForwarder := *openapiclient.NewLogForwarder("Name_example", "Description_example", "Type_example", "ConnectionType_example", openapiclient.LogForwarderDataFormat{LogForwarderAccessLogType: openapiclient.NewLogForwarderAccessLogType()}, openapiclient.LogForwarderSetting{LogForwarderDatadogConnectionType: openapiclient.NewLogForwarderDatadogConnectionType()}, false) // LogForwarder |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogForwardersAPI.LogForwardersUpdate(context.Background(), domain, logForwarderId).LogForwarder(logForwarder).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogForwardersAPI.LogForwardersUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogForwardersUpdate`: LogForwarderResponse
    fmt.Fprintf(os.Stdout, "Response from `LogForwardersAPI.LogForwardersUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**logForwarderId** | **string** | Log Forwarder Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogForwardersUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **logForwarder** | [**LogForwarder**](LogForwarder.md) |  | 

### Return type

[**LogForwarderResponse**](LogForwarderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogForwardersUpdateStatus

> LogForwarderResponse LogForwardersUpdateStatus(ctx, domain, logForwarderId).UpdateBooleanStatus(updateBooleanStatus).Execute()

Update a log forwarder's status

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
    logForwarderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Log Forwarder Id
    updateBooleanStatus := *openapiclient.NewUpdateBooleanStatus() // UpdateBooleanStatus |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.LogForwardersAPI.LogForwardersUpdateStatus(context.Background(), domain, logForwarderId).UpdateBooleanStatus(updateBooleanStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LogForwardersAPI.LogForwardersUpdateStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogForwardersUpdateStatus`: LogForwarderResponse
    fmt.Fprintf(os.Stdout, "Response from `LogForwardersAPI.LogForwardersUpdateStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**logForwarderId** | **string** | Log Forwarder Id | 

### Other Parameters

Other parameters are passed through a pointer to a apiLogForwardersUpdateStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateBooleanStatus** | [**UpdateBooleanStatus**](UpdateBooleanStatus.md) |  | 

### Return type

[**LogForwarderResponse**](LogForwarderResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

