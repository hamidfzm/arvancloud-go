# \WatermarkAPI

All URIs are relative to *https://napi.arvancloud.ir/live/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatermarksGet**](WatermarkAPI.md#WatermarksGet) | **Get** /watermarks | Return all watermarks.
[**WatermarksPost**](WatermarkAPI.md#WatermarksPost) | **Post** /watermarks | Store a newly created Watermark.
[**WatermarksWatermarkDelete**](WatermarkAPI.md#WatermarksWatermarkDelete) | **Delete** /watermarks/{watermark} | Remove the specified watermark.
[**WatermarksWatermarkGet**](WatermarkAPI.md#WatermarksWatermarkGet) | **Get** /watermarks/{watermark} | Return the specified watermark.
[**WatermarksWatermarkPatch**](WatermarkAPI.md#WatermarksWatermarkPatch) | **Patch** /watermarks/{watermark} | Update the specified watermark.



## WatermarksGet

> WatermarksGet(ctx).Filter(filter).Page(page).PerPage(perPage).Execute()

Return all watermarks.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/live"
)

func main() {
    filter := "filter_example" // string | Filter result (optional)
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit for query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.WatermarksGet(context.Background()).Filter(filter).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.WatermarksGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatermarksGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter result | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit for query | 

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


## WatermarksPost

> WatermarksPost(ctx).Title(title).Watermark(watermark).Description(description).Execute()

Store a newly created Watermark.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/live"
)

func main() {
    title := "title_example" // string | Title of watermark
    watermark := os.NewFile(1234, "some_file") // *os.File | Watermark file
    description := "description_example" // string | Description of watermark (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.WatermarkAPI.WatermarksPost(context.Background()).Title(title).Watermark(watermark).Description(description).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WatermarkAPI.WatermarksPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWatermarksPostRequest struct via the builder pattern


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
    openapiclient "github.com/hamidfzm/arvancloud-go/live"
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
    openapiclient "github.com/hamidfzm/arvancloud-go/live"
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
    openapiclient "github.com/hamidfzm/arvancloud-go/live"
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

