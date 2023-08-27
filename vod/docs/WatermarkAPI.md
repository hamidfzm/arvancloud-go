# \WatermarkAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelWatermarksGet**](WatermarkAPI.md#ChannelsChannelWatermarksGet) | **Get** /channels/{channel}/watermarks | Return all channel&#39;s watermarks.
[**ChannelsChannelWatermarksPost**](WatermarkAPI.md#ChannelsChannelWatermarksPost) | **Post** /channels/{channel}/watermarks | Store a newly created Watermark.
[**WatermarksWatermarkDelete**](WatermarkAPI.md#WatermarksWatermarkDelete) | **Delete** /watermarks/{watermark} | Remove the specified watermark.
[**WatermarksWatermarkGet**](WatermarkAPI.md#WatermarksWatermarkGet) | **Get** /watermarks/{watermark} | Return the specified watermark.
[**WatermarksWatermarkPatch**](WatermarkAPI.md#WatermarksWatermarkPatch) | **Patch** /watermarks/{watermark} | Update the specified watermark.



## ChannelsChannelWatermarksGet

> ChannelsChannelWatermarksGet(ctx, channel).Filter(filter).Page(page).PerPage(perPage).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Return all channel's watermarks.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    filter := "filter_example" // string | Filter result (optional)
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit for query (optional)
    secureIp := "secureIp_example" // string | The IP address for generate secure links for. If channel is secure default is request IP (optional)
    secureExpireTime := int32(56) // int32 | The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.ChannelsChannelWatermarksGet(context.Background(), channel).Filter(filter).Page(page).PerPage(perPage).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.ChannelsChannelWatermarksGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelWatermarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit for query | 
 **secureIp** | **string** | The IP address for generate secure links for. If channel is secure default is request IP | 
 **secureExpireTime** | **int32** | The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now | 

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


## ChannelsChannelWatermarksPost

> ChannelsChannelWatermarksPost(ctx, channel).Title(title).Watermark(watermark).Description(description).Execute()

Store a newly created Watermark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    title := "title_example" // string | Title of watermark
    watermark := os.NewFile(1234, "some_file") // *os.File | Watermark file
    description := "description_example" // string | Description of watermark (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.ChannelsChannelWatermarksPost(context.Background(), channel).Title(title).Watermark(watermark).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.ChannelsChannelWatermarksPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelWatermarksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | Title of watermark | 
 **watermark** | ***os.File** | Watermark file | 
 **description** | **string** | Description of watermark | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WatermarksWatermarkDelete

> WatermarksWatermarkDelete(ctx, watermark).Execute()

Remove the specified watermark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    watermark := "watermark_example" // string | The Id of watermark

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.WatermarksWatermarkDelete(context.Background(), watermark).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.WatermarksWatermarkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watermark** | **string** | The Id of watermark | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatermarksWatermarkDeleteRequest struct via the builder pattern


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


## WatermarksWatermarkGet

> WatermarksWatermarkGet(ctx, watermark).Execute()

Return the specified watermark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    watermark := "watermark_example" // string | The Id of watermark

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.WatermarksWatermarkGet(context.Background(), watermark).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.WatermarksWatermarkGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watermark** | **string** | The Id of watermark | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatermarksWatermarkGetRequest struct via the builder pattern


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


## WatermarksWatermarkPatch

> WatermarksWatermarkPatch(ctx, watermark).Body(body).Execute()

Update the specified watermark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    watermark := "watermark_example" // string | The Id of watermark
    body := *openapiclient.NewWatermarksWatermarkPatchRequest() // WatermarksWatermarkPatchRequest | Watermark details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.WatermarksWatermarkPatch(context.Background(), watermark).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.WatermarksWatermarkPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**watermark** | **string** | The Id of watermark | 

### Other Parameters

Other parameters are passed through a pointer to a apiWatermarksWatermarkPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**WatermarksWatermarkPatchRequest**](WatermarksWatermarkPatchRequest.md) | Watermark details | 

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

