# \StreamAPI

All URIs are relative to *https://napi.arvancloud.ir/live/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StreamsGet**](StreamAPI.md#StreamsGet) | **Get** /streams | Return all streams.
[**StreamsPost**](StreamAPI.md#StreamsPost) | **Post** /streams | Store a newly created stream.
[**StreamsStreamDelete**](StreamAPI.md#StreamsStreamDelete) | **Delete** /streams/{stream} | Remove the specified stream.
[**StreamsStreamGet**](StreamAPI.md#StreamsStreamGet) | **Get** /streams/{stream} | Return the specified stream.
[**StreamsStreamPatch**](StreamAPI.md#StreamsStreamPatch) | **Patch** /streams/{stream} | Update the specified stream.
[**StreamsStreamStartRecordGet**](StreamAPI.md#StreamsStreamStartRecordGet) | **Get** /streams/{stream}/start-record | Start record live stream.
[**StreamsStreamStopRecordGet**](StreamAPI.md#StreamsStreamStopRecordGet) | **Get** /streams/{stream}/stop-record | Stop record live stream.



## StreamsGet

> StreamsGet(ctx).Filter(filter).Page(page).PerPage(perPage).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Return all streams.

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
    secureIp := "secureIp_example" // string | The IP address for generate secure links for. If stream is secure default is request IP (optional)
    secureExpireTime := int32(56) // int32 | The Unix Timestamp for expire secure links.      *          If stream is secure default is 24 hours later from now (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsGet(context.Background()).Filter(filter).Page(page).PerPage(perPage).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStreamsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter result | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit for query | 
 **secureIp** | **string** | The IP address for generate secure links for. If stream is secure default is request IP | 
 **secureExpireTime** | **int32** | The Unix Timestamp for expire secure links.      *          If stream is secure default is 24 hours later from now | 

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


## StreamsPost

> StreamsPost(ctx).Body(body).Execute()

Store a newly created stream.

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
    body := *openapiclient.NewStreamsPostRequest("Title_example", "Type_example", "Mode_example", "Slug_example", int32(123), []openapiclient.StreamsPostRequestConvertInfoInner{*openapiclient.NewStreamsPostRequestConvertInfoInner()}) // StreamsPostRequest | Stream details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStreamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**StreamsPostRequest**](StreamsPostRequest.md) | Stream details | 

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


## StreamsStreamDelete

> StreamsStreamDelete(ctx, stream).Execute()

Remove the specified stream.

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
    stream := "stream_example" // string | The Id of stream

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsStreamDelete(context.Background(), stream).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsStreamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | The Id of stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamsStreamDeleteRequest struct via the builder pattern


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


## StreamsStreamGet

> StreamsStreamGet(ctx, stream).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()

Return the specified stream.

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
    stream := "stream_example" // string | The Id of stream
    secureIp := "secureIp_example" // string | The IP address for generate secure links for. If stream is secure default is request IP (optional)
    secureExpireTime := int32(56) // int32 | The Unix Timestamp for expire secure links.      *          If stream is secure default is 24 hours later from now (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsStreamGet(context.Background(), stream).SecureIp(secureIp).SecureExpireTime(secureExpireTime).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsStreamGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | The Id of stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamsStreamGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **secureIp** | **string** | The IP address for generate secure links for. If stream is secure default is request IP | 
 **secureExpireTime** | **int32** | The Unix Timestamp for expire secure links.      *          If stream is secure default is 24 hours later from now | 

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


## StreamsStreamPatch

> StreamsStreamPatch(ctx, stream).Body(body).Execute()

Update the specified stream.

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
    stream := "stream_example" // string | The Id of stream
    body := *openapiclient.NewStreamsStreamPatchRequest() // StreamsStreamPatchRequest | Stream details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsStreamPatch(context.Background(), stream).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsStreamPatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stream** | **string** | The Id of stream | 

### Other Parameters

Other parameters are passed through a pointer to a apiStreamsStreamPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**StreamsStreamPatchRequest**](StreamsStreamPatchRequest.md) | Stream details | 

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


## StreamsStreamStartRecordGet

> StreamsStreamStartRecordGet(ctx).Execute()

Start record live stream.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsStreamStartRecordGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsStreamStartRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamsStreamStartRecordGetRequest struct via the builder pattern


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


## StreamsStreamStopRecordGet

> StreamsStreamStopRecordGet(ctx).Execute()

Stop record live stream.



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.StreamAPI.StreamsStreamStopRecordGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StreamAPI.StreamsStreamStopRecordGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStreamsStreamStopRecordGetRequest struct via the builder pattern


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

