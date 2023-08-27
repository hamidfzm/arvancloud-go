# \TagAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TagsGet**](TagAPI.md#TagsGet) | **Get** /tags | Return all user&#39;s tags.
[**TagsPost**](TagAPI.md#TagsPost) | **Post** /tags | Store a newly created Tag.
[**VideosVideoTagsGet**](TagAPI.md#VideosVideoTagsGet) | **Get** /videos/{video}/tags | Return all vidoe&#39;s tags.
[**VideosVideoTagsPut**](TagAPI.md#VideosVideoTagsPut) | **Put** /videos/{video}/tags | Update the video tags.



## TagsGet

> TagsGet(ctx, filter2).Filter(filter).Sort(sort).Execute()

Return all user's tags.

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
    filter2 := "filter_example" // string | filter tags
    filter := "filter_example" // string | Filter result (optional)
    sort := "sort_example" // string | Sort ex: asc ,desc  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagAPI.TagsGet(context.Background(), filter2).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.TagsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filter2** | **string** | filter tags | 

### Other Parameters

Other parameters are passed through a pointer to a apiTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 
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


## TagsPost

> TagsPost(ctx).Title(title).Execute()

Store a newly created Tag.

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
    title := "title_example" // string | Title of tag

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagAPI.TagsPost(context.Background()).Title(title).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.TagsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTagsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **title** | **string** | Title of tag | 

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


## VideosVideoTagsGet

> VideosVideoTagsGet(ctx, video).Filter(filter).Execute()

Return all vidoe's tags.

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
    filter := "filter_example" // string | Filter result (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagAPI.VideosVideoTagsGet(context.Background(), video).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.VideosVideoTagsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoTagsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 

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


## VideosVideoTagsPut

> VideosVideoTagsPut(ctx, video).Body(body).Execute()

Update the video tags.

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
    body := []string{"Property_example"} // []string | tags (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TagAPI.VideosVideoTagsPut(context.Background(), video).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TagAPI.VideosVideoTagsPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiVideosVideoTagsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | tags | 

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

