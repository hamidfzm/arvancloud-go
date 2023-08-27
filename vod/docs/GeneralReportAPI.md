# \GeneralReportAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReportGeoGet**](GeneralReportAPI.md#ReportGeoGet) | **Get** /report/geo | Return Domain Geo Report.
[**ReportStatisticsGet**](GeneralReportAPI.md#ReportStatisticsGet) | **Get** /report/statistics | Return Domain statistics report.
[**ReportTrafficsGet**](GeneralReportAPI.md#ReportTrafficsGet) | **Get** /report/traffics | Return Domain Traffic.
[**ReportUserAgentGet**](GeneralReportAPI.md#ReportUserAgentGet) | **Get** /report/user-agent | Return User Agent.
[**ReportVisitorsGet**](GeneralReportAPI.md#ReportVisitorsGet) | **Get** /report/visitors | Return Domain Visitors.



## ReportGeoGet

> ReportGeoGet(ctx).Period(period).Execute()

Return Domain Geo Report.

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
    period := "period_example" // string | Values: 1h,3h,6h,12h,24h,7d,30d

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeneralReportAPI.ReportGeoGet(context.Background()).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralReportAPI.ReportGeoGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportGeoGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Values: 1h,3h,6h,12h,24h,7d,30d | 

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


## ReportStatisticsGet

> ReportStatisticsGet(ctx).Execute()

Return Domain statistics report.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeneralReportAPI.ReportStatisticsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralReportAPI.ReportStatisticsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiReportStatisticsGetRequest struct via the builder pattern


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


## ReportTrafficsGet

> ReportTrafficsGet(ctx).Period(period).Execute()

Return Domain Traffic.

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
    period := "period_example" // string | Values: 1h,3h,6h,12h,24h,7d,30d

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeneralReportAPI.ReportTrafficsGet(context.Background()).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralReportAPI.ReportTrafficsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportTrafficsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Values: 1h,3h,6h,12h,24h,7d,30d | 

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


## ReportUserAgentGet

> ReportUserAgentGet(ctx).Period(period).Execute()

Return User Agent.

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
    period := "period_example" // string | Values: 1h,3h,6h,12h,24h,7d,30d

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeneralReportAPI.ReportUserAgentGet(context.Background()).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralReportAPI.ReportUserAgentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportUserAgentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Values: 1h,3h,6h,12h,24h,7d,30d | 

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


## ReportVisitorsGet

> ReportVisitorsGet(ctx).Period(period).Execute()

Return Domain Visitors.

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
    period := "period_example" // string | Values: 1h,3h,6h,12h,24h,7d,30d

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GeneralReportAPI.ReportVisitorsGet(context.Background()).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GeneralReportAPI.ReportVisitorsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReportVisitorsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Values: 1h,3h,6h,12h,24h,7d,30d | 

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

