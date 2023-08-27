# \RedirectAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RedirectShow**](RedirectAPI.md#RedirectShow) | **Get** /domains/{domain}/settings/www-redirect | Get redirect settings
[**RedirectUpdate**](RedirectAPI.md#RedirectUpdate) | **Put** /domains/{domain}/settings/www-redirect | Update redirect settings



## RedirectShow

> RedirectData RedirectShow(ctx, domain).Execute()

Get redirect settings

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
    resp, r, err := apiClient.RedirectAPI.RedirectShow(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectAPI.RedirectShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedirectShow`: RedirectData
    fmt.Fprintf(os.Stdout, "Response from `RedirectAPI.RedirectShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedirectShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RedirectData**](RedirectData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedirectUpdate

> MessageResponse RedirectUpdate(ctx, domain).Redirect(redirect).Execute()

Update redirect settings

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
    redirect := *openapiclient.NewRedirect() // Redirect |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RedirectAPI.RedirectUpdate(context.Background(), domain).Redirect(redirect).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RedirectAPI.RedirectUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RedirectUpdate`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `RedirectAPI.RedirectUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedirectUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redirect** | [**Redirect**](Redirect.md) |  | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

