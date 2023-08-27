# \PTRAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePTR**](PTRAPI.md#CreatePTR) | **Post** /regions/:region/ptr/ | 
[**DeletePTR**](PTRAPI.md#DeletePTR) | **Delete** /regions/:region/ptr/:ip | 



## CreatePTR

> MessageResponse CreatePTR(ctx, region).CreatePTRRequest(createPTRRequest).Execute()





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
    createPTRRequest := *openapiclient.NewCreatePTRRequest() // CreatePTRRequest | IP and domain info of the PTR record

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PTRAPI.CreatePTR(context.Background(), region).CreatePTRRequest(createPTRRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PTRAPI.CreatePTR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePTR`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PTRAPI.CreatePTR`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePTRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPTRRequest** | [**CreatePTRRequest**](CreatePTRRequest.md) | IP and domain info of the PTR record | 

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


## DeletePTR

> MessageResponse DeletePTR(ctx, region, ip).Execute()





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
    ip := "ip_example" // string | the IP which PTR record points to

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PTRAPI.DeletePTR(context.Background(), region, ip).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PTRAPI.DeletePTR``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePTR`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `PTRAPI.DeletePTR`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**ip** | **string** | the IP which PTR record points to | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePTRRequest struct via the builder pattern


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

