# \ChannelAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelDelete**](ChannelAPI.md#ChannelsChannelDelete) | **Delete** /channels/{channel} | Remove the specified channel.
[**ChannelsChannelGet**](ChannelAPI.md#ChannelsChannelGet) | **Get** /channels/{channel} | Return the specified channel.
[**ChannelsChannelPut**](ChannelAPI.md#ChannelsChannelPut) | **Put** /channels/{channel} | Update the specified channel.
[**ChannelsGet**](ChannelAPI.md#ChannelsGet) | **Get** /channels | Return all user channels.
[**ChannelsPost**](ChannelAPI.md#ChannelsPost) | **Post** /channels | Store a newly channel.



## ChannelsChannelDelete

> ChannelsChannelDelete(ctx, channel).Execute()

Remove the specified channel.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    channel := "channel_example" // string | The Id of channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelAPI.ChannelsChannelDelete(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelsChannelDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelGet

> ChannelsChannelGet(ctx, channel).Execute()

Return the specified channel.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    channel := "channel_example" // string | The Id of channel

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelAPI.ChannelsChannelGet(context.Background(), channel).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelsChannelGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelPut

> ChannelsChannelPut(ctx, channel).Body(body).Execute()

Update the specified channel.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    body := *openapiclient.NewChannelsChannelPutRequest() // ChannelsChannelPutRequest | Channel details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelAPI.ChannelsChannelPut(context.Background(), channel).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelsChannelPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChannelsChannelPutRequest**](ChannelsChannelPutRequest.md) | Channel details | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsGet

> ChannelsGet(ctx).Filter(filter).Page(page).PerPage(perPage).Execute()

Return all user channels.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    filter := "filter_example" // string | Filter result (optional)
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelAPI.ChannelsGet(context.Background()).Filter(filter).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter result | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsPost

> ChannelsPost(ctx).Body(body).Execute()

Store a newly channel.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    body := *openapiclient.NewChannelsPostRequest("Title_example") // ChannelsPostRequest | Channel details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ChannelAPI.ChannelsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelAPI.ChannelsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChannelsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ChannelsPostRequest**](ChannelsPostRequest.md) | Channel details | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

