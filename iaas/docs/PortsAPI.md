# \PortsAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisablePort**](PortsAPI.md#DisablePort) | **Patch** /regions/:region/ports/:id/disable | 
[**DisablePortSecurity**](PortsAPI.md#DisablePortSecurity) | **Patch** /regions/:region/ports/:id/disablePortSecurity | 
[**EnablePort**](PortsAPI.md#EnablePort) | **Patch** /regions/:region/ports/:id/enable | 
[**EnablePortSecurity**](PortsAPI.md#EnablePortSecurity) | **Patch** /regions/:region/ports/:id/enablePortSecurity | 



## DisablePort

> MessageResponse DisablePort(ctx, region, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | port id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsAPI.DisablePort(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.DisablePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisablePort`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PortsAPI.DisablePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisablePortSecurity

> MessageResponse DisablePortSecurity(ctx, region, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | port id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsAPI.DisablePortSecurity(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.DisablePortSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisablePortSecurity`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PortsAPI.DisablePortSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisablePortSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePort

> MessageResponse EnablePort(ctx, region, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | port id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsAPI.EnablePort(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.EnablePort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnablePort`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PortsAPI.EnablePort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnablePortSecurity

> MessageResponse EnablePortSecurity(ctx, region, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | port id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PortsAPI.EnablePortSecurity(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PortsAPI.EnablePortSecurity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnablePortSecurity`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PortsAPI.EnablePortSecurity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | port id | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnablePortSecurityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

