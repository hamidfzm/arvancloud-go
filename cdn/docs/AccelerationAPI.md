# \AccelerationAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccelerationShow**](AccelerationAPI.md#AccelerationShow) | **Get** /domains/{domain}/acceleration | Get the content of acceleration settings
[**AccelerationUpdate**](AccelerationAPI.md#AccelerationUpdate) | **Patch** /domains/{domain}/acceleration | Update the content of acceleration settings
[**ImageResizeShow**](AccelerationAPI.md#ImageResizeShow) | **Get** /domains/{domain}/image-resize | Get the content of image resize settings
[**ImageResizeUpdate**](AccelerationAPI.md#ImageResizeUpdate) | **Patch** /domains/{domain}/image-resize | Update the content of image resize settings



## AccelerationShow

> AccelerationResponse AccelerationShow(ctx, domain).Execute()

Get the content of acceleration settings

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationAPI.AccelerationShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationAPI.AccelerationShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccelerationShow`: AccelerationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationAPI.AccelerationShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccelerationShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccelerationResponse**](AccelerationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccelerationUpdate

> AccelerationResponse AccelerationUpdate(ctx, domain).AccelerationUpdate(accelerationUpdate).Execute()

Update the content of acceleration settings

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
    accelerationUpdate := *openapiclient.NewAccelerationUpdate() // AccelerationUpdate |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationAPI.AccelerationUpdate(context.Background(), domain).AccelerationUpdate(accelerationUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationAPI.AccelerationUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccelerationUpdate`: AccelerationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationAPI.AccelerationUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiAccelerationUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accelerationUpdate** | [**AccelerationUpdate**](AccelerationUpdate.md) |  | 

### Return type

[**AccelerationResponse**](AccelerationResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageResizeShow

> ImageResizeResponse ImageResizeShow(ctx, domain).Execute()

Get the content of image resize settings

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationAPI.ImageResizeShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationAPI.ImageResizeShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageResizeShow`: ImageResizeResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationAPI.ImageResizeShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageResizeShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ImageResizeResponse**](ImageResizeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ImageResizeUpdate

> ImageResizeResponse ImageResizeUpdate(ctx, domain).ImageResize(imageResize).Execute()

Update the content of image resize settings

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
    imageResize := *openapiclient.NewImageResize() // ImageResize |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccelerationAPI.ImageResizeUpdate(context.Background(), domain).ImageResize(imageResize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccelerationAPI.ImageResizeUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ImageResizeUpdate`: ImageResizeResponse
    fmt.Fprintf(os.Stdout, "Response from `AccelerationAPI.ImageResizeUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiImageResizeUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **imageResize** | [**ImageResize**](ImageResize.md) |  | 

### Return type

[**ImageResizeResponse**](ImageResizeResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

