# \ThumbnailAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VideosVideoThumbnailGet**](ThumbnailAPI.md#VideosVideoThumbnailGet) | **Get** /videos/{video}/thumbnail | Display video thumbnail.
[**VideosVideoThumbnailPost**](ThumbnailAPI.md#VideosVideoThumbnailPost) | **Post** /videos/{video}/thumbnail | Store a newly created thumbnail.



## VideosVideoThumbnailGet

> VideosVideoThumbnailGet(ctx, video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Display video thumbnail.

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
    r, err := apiClient.ThumbnailAPI.VideosVideoThumbnailGet(context.Background(), video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailAPI.VideosVideoThumbnailGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoThumbnailGetRequest struct via the builder pattern


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


## VideosVideoThumbnailPost

> VideosVideoThumbnailPost(ctx, video).Thumbnail(thumbnail).ThumbnailTime(thumbnailTime).Execute()

Store a newly created thumbnail.

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
    thumbnail := os.NewFile(1234, "some_file") // *os.File | The png file. (optional)
    thumbnailTime := int32(56) // int32 | Specific video time (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ThumbnailAPI.VideosVideoThumbnailPost(context.Background(), video).Thumbnail(thumbnail).ThumbnailTime(thumbnailTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThumbnailAPI.VideosVideoThumbnailPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoThumbnailPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **thumbnail** | ***os.File** | The png file. | 
 **thumbnailTime** | **int32** | Specific video time | 

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

