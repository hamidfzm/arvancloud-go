# \ReportsAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReportMetric**](ReportsAPI.md#GetReportMetric) | **Get** /regions/:region/reports/:id/:metric | 
[**GetServerReports**](ReportsAPI.md#GetServerReports) | **Get** /regions/:region/reports/:id | 



## GetReportMetric

> ServerReports GetReportMetric(ctx, region, id, metric).Execute()





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
    id := "id_example" // string | server id
    metric := "metric_example" // string | the metric name to be returned

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetReportMetric(context.Background(), region, id, metric).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetReportMetric``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReportMetric`: ServerReports
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetReportMetric`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 
**metric** | **string** | the metric name to be returned | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportMetricRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerReports**](ServerReports.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerReports

> ServerReports GetServerReports(ctx, region, id, period).Execute()





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
    id := "id_example" // string | server id
    period := "period_example" // string | period of the reports, ex. `1m`, `1h`, `1d`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.GetServerReports(context.Background(), region, id, period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.GetServerReports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerReports`: ServerReports
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.GetServerReports`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 
**period** | **string** | period of the reports, ex. &#x60;1m&#x60;, &#x60;1h&#x60;, &#x60;1d&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerReportsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**ServerReports**](ServerReports.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

