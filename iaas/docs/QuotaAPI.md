# \QuotaAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserLimits**](QuotaAPI.md#GetUserLimits) | **Get** /regions/:region/quota | 



## GetUserLimits

> Quota GetUserLimits(ctx, region).Execute()





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
    resp, r, err := apiClient.QuotaAPI.GetUserLimits(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `QuotaAPI.GetUserLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserLimits`: Quota
    fmt.Fprintf(os.Stdout, "Response from `QuotaAPI.GetUserLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Quota**](Quota.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

