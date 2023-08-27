# \AudioTrackAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AudioTracksAudioTrackDelete**](AudioTrackAPI.md#AudioTracksAudioTrackDelete) | **Delete** /audio-tracks/{audio_track} | Remove the specified audio track.
[**AudioTracksAudioTrackGet**](AudioTrackAPI.md#AudioTracksAudioTrackGet) | **Get** /audio-tracks/{audio_track} | Return the specified audio track.
[**VideosVideoAudioTracksGet**](AudioTrackAPI.md#VideosVideoAudioTracksGet) | **Get** /videos/{video}/audio-tracks | Display a listing of the Audio Tracks.
[**VideosVideoAudioTracksPost**](AudioTrackAPI.md#VideosVideoAudioTracksPost) | **Post** /videos/{video}/audio-tracks | Store a newly created audio track.



## AudioTracksAudioTrackDelete

> AudioTracksAudioTrackDelete(ctx, audioTrack).Execute()

Remove the specified audio track.

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
    audioTrack := "audioTrack_example" // string | The Id of audio track

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioTrackAPI.AudioTracksAudioTrackDelete(context.Background(), audioTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioTrackAPI.AudioTracksAudioTrackDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audioTrack** | **string** | The Id of audio track | 

### Other Parameters

Other parameters are passed through a pointer to a apiAudioTracksAudioTrackDeleteRequest struct via the builder pattern


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


## AudioTracksAudioTrackGet

> AudioTracksAudioTrackGet(ctx, audioTrack).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Return the specified audio track.

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
    audioTrack := "audioTrack_example" // string | The Id of audio track
    secureIp := "secureIp_example" // string | The IP address for generate secure links for. If channel is secure default is request IP (optional)
    secureExpireTime := int32(56) // int32 | The Unix Timestamp for expire secure links.      *          If channel is secure default is 24 hours later from now (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioTrackAPI.AudioTracksAudioTrackGet(context.Background(), audioTrack).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioTrackAPI.AudioTracksAudioTrackGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**audioTrack** | **string** | The Id of audio track | 

### Other Parameters

Other parameters are passed through a pointer to a apiAudioTracksAudioTrackGetRequest struct via the builder pattern


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


## VideosVideoAudioTracksGet

> VideosVideoAudioTracksGet(ctx, video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Display a listing of the Audio Tracks.

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
    r, err := apiClient.AudioTrackAPI.VideosVideoAudioTracksGet(context.Background(), video).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioTrackAPI.VideosVideoAudioTracksGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoAudioTracksGetRequest struct via the builder pattern


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


## VideosVideoAudioTracksPost

> VideosVideoAudioTracksPost(ctx, video).Lang(lang).AudioTrack(audioTrack).Execute()

Store a newly created audio track.

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
    lang := "lang_example" // string | Track language
    audioTrack := os.NewFile(1234, "some_file") // *os.File | The Mp3 or ACC Audio file.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AudioTrackAPI.VideosVideoAudioTracksPost(context.Background(), video).Lang(lang).AudioTrack(audioTrack).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AudioTrackAPI.VideosVideoAudioTracksPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoAudioTracksPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | Track language | 
 **audioTrack** | ***os.File** | The Mp3 or ACC Audio file. | 

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

