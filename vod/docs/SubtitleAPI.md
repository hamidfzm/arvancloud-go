# \SubtitleAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubtitlesSubtitleDelete**](SubtitleAPI.md#SubtitlesSubtitleDelete) | **Delete** /subtitles/{subtitle} | Remove the specified subtitle.
[**SubtitlesSubtitleGet**](SubtitleAPI.md#SubtitlesSubtitleGet) | **Get** /subtitles/{subtitle} | Return the specified subtitle.
[**VideosVideoSubtitlesGet**](SubtitleAPI.md#VideosVideoSubtitlesGet) | **Get** /videos/{video}/subtitles | Display a listing of the subtitle.
[**VideosVideoSubtitlesPost**](SubtitleAPI.md#VideosVideoSubtitlesPost) | **Post** /videos/{video}/subtitles | Store a newly created subtitle.



## SubtitlesSubtitleDelete

> SubtitlesSubtitleDelete(ctx, subtitle).Execute()

Remove the specified subtitle.

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
    subtitle := "subtitle_example" // string | The Id of subtitle

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubtitleAPI.SubtitlesSubtitleDelete(context.Background(), subtitle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.SubtitlesSubtitleDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtitle** | **string** | The Id of subtitle | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubtitlesSubtitleDeleteRequest struct via the builder pattern


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


## SubtitlesSubtitleGet

> SubtitlesSubtitleGet(ctx, subtitle).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Return the specified subtitle.

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
    subtitle := "subtitle_example" // string | The Id of subtitle
    secureIp := "secureIp_example" // string | The IP address for generate secure links for. If channel is secure default is request IP (optional)
    secureExpireTime := int32(56) // int32 | The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubtitleAPI.SubtitlesSubtitleGet(context.Background(), subtitle).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.SubtitlesSubtitleGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subtitle** | **string** | The Id of subtitle | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubtitlesSubtitleGetRequest struct via the builder pattern


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


## VideosVideoSubtitlesGet

> VideosVideoSubtitlesGet(ctx, video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Display a listing of the subtitle.

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
    r, err := apiClient.SubtitleAPI.VideosVideoSubtitlesGet(context.Background(), video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.VideosVideoSubtitlesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoSubtitlesGetRequest struct via the builder pattern


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


## VideosVideoSubtitlesPost

> VideosVideoSubtitlesPost(ctx, video).Lang(lang).Subtitle(subtitle).Execute()

Store a newly created subtitle.

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
    lang := "lang_example" // string | Subtitle language
    subtitle := os.NewFile(1234, "some_file") // *os.File | The SRT or VTT subtitle file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SubtitleAPI.VideosVideoSubtitlesPost(context.Background(), video).Lang(lang).Subtitle(subtitle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubtitleAPI.VideosVideoSubtitlesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoSubtitlesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | Subtitle language | 
 **subtitle** | ***os.File** | The SRT or VTT subtitle file. | 

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

