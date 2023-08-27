# \FloatingIPsAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachFloatIPToServer**](FloatingIPsAPI.md#AttachFloatIPToServer) | **Patch** /regions/:region/float-ips/:id/attach | 
[**CreateFloatingIP**](FloatingIPsAPI.md#CreateFloatingIP) | **Post** /regions/:region/float-ips | 
[**DeleteFloatingIP**](FloatingIPsAPI.md#DeleteFloatingIP) | **Delete** /regions/:region/float-ips/:id | 
[**DetachFloatIPFromServer**](FloatingIPsAPI.md#DetachFloatIPFromServer) | **Patch** /regions/:region/float-ips/detach | 
[**GetAllFloatingIPs**](FloatingIPsAPI.md#GetAllFloatingIPs) | **Get** /regions/:region/float-ips | 
[**GetFloatingIPInternalIPs**](FloatingIPsAPI.md#GetFloatingIPInternalIPs) | **Get** /regions/:region/float-ips/ips | 



## AttachFloatIPToServer

> MessageResponse AttachFloatIPToServer(ctx, region, id).FloatIPAttachRequest(floatIPAttachRequest).Execute()





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
    id := "id_example" // string | float-ip id
    floatIPAttachRequest := *openapiclient.NewFloatIPAttachRequest() // FloatIPAttachRequest | Floating-ip attachment info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPsAPI.AttachFloatIPToServer(context.Background(), region, id).FloatIPAttachRequest(floatIPAttachRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPsAPI.AttachFloatIPToServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachFloatIPToServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPsAPI.AttachFloatIPToServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | float-ip id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachFloatIPToServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **floatIPAttachRequest** | [**FloatIPAttachRequest**](FloatIPAttachRequest.md) | Floating-ip attachment info | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFloatingIP

> FloatIP CreateFloatingIP(ctx, region).CreateFloatIPRequest(createFloatIPRequest).Execute()





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
    createFloatIPRequest := *openapiclient.NewCreateFloatIPRequest() // CreateFloatIPRequest | Info about the new float-ip to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPsAPI.CreateFloatingIP(context.Background(), region).CreateFloatIPRequest(createFloatIPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPsAPI.CreateFloatingIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFloatingIP`: FloatIP
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPsAPI.CreateFloatingIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFloatingIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createFloatIPRequest** | [**CreateFloatIPRequest**](CreateFloatIPRequest.md) | Info about the new float-ip to be created | 

### Return type

[**FloatIP**](FloatIP.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFloatingIP

> MessageResponse DeleteFloatingIP(ctx, region).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPsAPI.DeleteFloatingIP(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPsAPI.DeleteFloatingIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFloatingIP`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPsAPI.DeleteFloatingIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFloatingIPRequest struct via the builder pattern


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


## DetachFloatIPFromServer

> MessageResponse DetachFloatIPFromServer(ctx, region).DetachFloatIPRequest(detachFloatIPRequest).Execute()





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
    detachFloatIPRequest := *openapiclient.NewDetachFloatIPRequest() // DetachFloatIPRequest | Floating-ip attachment info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPsAPI.DetachFloatIPFromServer(context.Background(), region).DetachFloatIPRequest(detachFloatIPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPsAPI.DetachFloatIPFromServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachFloatIPFromServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPsAPI.DetachFloatIPFromServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachFloatIPFromServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detachFloatIPRequest** | [**DetachFloatIPRequest**](DetachFloatIPRequest.md) | Floating-ip attachment info | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFloatingIPs

> FloatIP GetAllFloatingIPs(ctx, region).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPsAPI.GetAllFloatingIPs(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPsAPI.GetAllFloatingIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllFloatingIPs`: FloatIP
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPsAPI.GetAllFloatingIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFloatingIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FloatIP**](FloatIP.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFloatingIPInternalIPs

> ServerIPInfo GetFloatingIPInternalIPs(ctx, region).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FloatingIPsAPI.GetFloatingIPInternalIPs(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FloatingIPsAPI.GetFloatingIPInternalIPs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFloatingIPInternalIPs`: ServerIPInfo
    fmt.Fprintf(os.Stdout, "Response from `FloatingIPsAPI.GetFloatingIPInternalIPs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFloatingIPInternalIPsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerIPInfo**](ServerIPInfo.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

