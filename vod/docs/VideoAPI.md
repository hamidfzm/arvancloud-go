# \VideoAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelVideosGet**](VideoAPI.md#ChannelsChannelVideosGet) | **Get** /channels/{channel}/videos | Return all channel&#39;s videos.
[**ChannelsChannelVideosPost**](VideoAPI.md#ChannelsChannelVideosPost) | **Post** /channels/{channel}/videos | Store a newly created video.
[**VideosBulkDelete**](VideoAPI.md#VideosBulkDelete) | **Delete** /videos/bulk | Remove the multiple video.
[**VideosVideoDelete**](VideoAPI.md#VideosVideoDelete) | **Delete** /videos/{video} | Remove the specified video.
[**VideosVideoGet**](VideoAPI.md#VideosVideoGet) | **Get** /videos/{video} | Return the specified video.
[**VideosVideoPatch**](VideoAPI.md#VideosVideoPatch) | **Patch** /videos/{video} | Update the specified video.



## ChannelsChannelVideosGet

> ChannelsChannelVideosGet(ctx, channel).Filter(filter).Page(page).PerPage(perPage).SecureIp(secureIp).SecureExpireTime(secureExpireTime).OrderBy(orderBy).Sort(sort).Execute()

Return all channel's videos.

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
    orderBy := "orderBy_example" // string | Order by ex: title, time (optional)
    sort := "sort_example" // string | Sort ex: asc ,desc  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VideoAPI.ChannelsChannelVideosGet(context.Background(), channel).Filter(filter).Page(page).PerPage(perPage).SecureIp(secureIp).SecureExpireTime(secureExpireTime).OrderBy(orderBy).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.ChannelsChannelVideosGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelVideosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit for query | 
 **secureIp** | **string** | The IP address for generate secure links for. If channel is secure default is request IP | 
 **secureExpireTime** | **int32** | The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now | 
 **orderBy** | **string** | Order by ex: title, time | 
 **sort** | **string** | Sort ex: asc ,desc  | 

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


## ChannelsChannelVideosPost

> ChannelsChannelVideosPost(ctx, channel).Body(body).Execute()

Store a newly created video.

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
    body := *openapiclient.NewChannelsChannelVideosPostRequest("Title_example", "ConvertMode_example") // ChannelsChannelVideosPostRequest | Video details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VideoAPI.ChannelsChannelVideosPost(context.Background(), channel).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.ChannelsChannelVideosPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelVideosPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChannelsChannelVideosPostRequest**](ChannelsChannelVideosPostRequest.md) | Video details | 

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


## VideosBulkDelete

> VideosBulkDelete(ctx, ids).Execute()

Remove the multiple video.

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
    ids := []string{"Inner_example"} // []string | The Id's of video's

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VideoAPI.VideosBulkDelete(context.Background(), ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.VideosBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]string**](string.md) | The Id&#39;s of video&#39;s | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideosBulkDeleteRequest struct via the builder pattern


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


## VideosVideoDelete

> VideosVideoDelete(ctx, video).Execute()

Remove the specified video.

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
    video := "video_example" // string | The Id of video

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VideoAPI.VideosVideoDelete(context.Background(), video).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.VideosVideoDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**video** | **string** | The Id of video | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideosVideoDeleteRequest struct via the builder pattern


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


## VideosVideoGet

> VideosVideoGet(ctx, video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Return the specified video.

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
    video := "video_example" // string | The Id of video
    secureIp := "secureIp_example" // string | The IP address for generate secure links for. If channel is secure default is request IP (optional)
    secureExpireTime := int32(56) // int32 | The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VideoAPI.VideosVideoGet(context.Background(), video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.VideosVideoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**video** | **string** | The Id of video | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideosVideoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## VideosVideoPatch

> VideosVideoPatch(ctx, video).Body(body).Execute()

Update the specified video.

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
    video := "video_example" // string | The Id of video
    body := *openapiclient.NewVideosVideoPatchRequest() // VideosVideoPatchRequest | Video details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VideoAPI.VideosVideoPatch(context.Background(), video).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VideoAPI.VideosVideoPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**video** | **string** | The Id of video | 

### Other Parameters

Other parameters are passed through a pointer to a apiVideosVideoPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VideosVideoPatchRequest**](VideosVideoPatchRequest.md) | Video details | 

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

