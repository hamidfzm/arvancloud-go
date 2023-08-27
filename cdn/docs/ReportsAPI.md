# \ReportsAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkReportsTrafficsTotal**](ReportsAPI.md#BulkReportsTrafficsTotal) | **Post** /bulk/reports/traffics | Get traffic report for multiple domains
[**BulkReportsVisitorsTotal**](ReportsAPI.md#BulkReportsVisitorsTotal) | **Post** /bulk/reports/visitors | Get visitor report for multiple domains
[**ReportsAttacksAttackers**](ReportsAPI.md#ReportsAttacksAttackers) | **Get** /domains/{domain}/reports/attacks/attackers | Get list of attackers information
[**ReportsAttacksIndex**](ReportsAPI.md#ReportsAttacksIndex) | **Get** /domains/{domain}/reports/attacks/list | Get list of attacks details
[**ReportsAttacksMap**](ReportsAPI.md#ReportsAttacksMap) | **Get** /domains/{domain}/reports/attacks/map | Get geo-map of attacks
[**ReportsAttacksShow**](ReportsAPI.md#ReportsAttacksShow) | **Get** /domains/{domain}/reports/attacks | Get report of attacks
[**ReportsAttacksUri**](ReportsAPI.md#ReportsAttacksUri) | **Get** /domains/{domain}/reports/attacks/uri | Get list of URLs under attack
[**ReportsDnsGeo**](ReportsAPI.md#ReportsDnsGeo) | **Get** /domains/{domain}/reports/dns-geo | Get DNS requests as geo-map
[**ReportsDnsRequests**](ReportsAPI.md#ReportsDnsRequests) | **Get** /domains/{domain}/reports/dns-requests | Get response time report
[**ReportsErrorLogDetails**](ReportsAPI.md#ReportsErrorLogDetails) | **Get** /domains/{domain}/reports/error-log-details | Get detail of an error
[**ReportsErrorLogs**](ReportsAPI.md#ReportsErrorLogs) | **Get** /domains/{domain}/reports/error-logs | Get list of errors
[**ReportsErrorLogsChart**](ReportsAPI.md#ReportsErrorLogsChart) | **Get** /domains/{domain}/reports/error-logs/chart | Get chart view of errors
[**ReportsResponseTimeIndex**](ReportsAPI.md#ReportsResponseTimeIndex) | **Get** /domains/{domain}/reports/response-time | Get response time report
[**ReportsStatusIndex**](ReportsAPI.md#ReportsStatusIndex) | **Get** /domains/{domain}/reports/status | Get status codes pie chart
[**ReportsStatusSummary**](ReportsAPI.md#ReportsStatusSummary) | **Get** /domains/{domain}/reports/status/summary | Get an overview of status codes pie chart
[**ReportsTrafficsMap**](ReportsAPI.md#ReportsTrafficsMap) | **Get** /domains/{domain}/reports/traffics/map | Get traffic as geo-map
[**ReportsTrafficsSaved**](ReportsAPI.md#ReportsTrafficsSaved) | **Get** /domains/{domain}/reports/traffics/saved | Get traffic saved to total pie chart
[**ReportsTrafficsTotal**](ReportsAPI.md#ReportsTrafficsTotal) | **Get** /domains/{domain}/reports/traffics | Get traffic report for domain
[**ReportsTransportLayerProxiesTraffics**](ReportsAPI.md#ReportsTransportLayerProxiesTraffics) | **Get** /domains/{domain}/reports/transport-layer-proxies/{transportLayerProxyId}/traffics | Get traffic report for transport layer proxy
[**ReportsVisitorsHighRequestIps**](ReportsAPI.md#ReportsVisitorsHighRequestIps) | **Get** /domains/{domain}/reports/high-request-ips | Get report of IPs with highest number of requests
[**ReportsVisitorsIndex**](ReportsAPI.md#ReportsVisitorsIndex) | **Get** /domains/{domain}/reports/visitors | Get report of visitors for domain



## BulkReportsTrafficsTotal

> BulkTrafficReportData BulkReportsTrafficsTotal(ctx).Period(period).Since(since).Until(until).BulkReportsTrafficsTotalRequest(bulkReportsTrafficsTotalRequest).Execute()

Get traffic report for multiple domains

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    bulkReportsTrafficsTotalRequest := *openapiclient.NewBulkReportsTrafficsTotalRequest() // BulkReportsTrafficsTotalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.BulkReportsTrafficsTotal(context.Background()).Period(period).Since(since).Until(until).BulkReportsTrafficsTotalRequest(bulkReportsTrafficsTotalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.BulkReportsTrafficsTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkReportsTrafficsTotal`: BulkTrafficReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.BulkReportsTrafficsTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkReportsTrafficsTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **bulkReportsTrafficsTotalRequest** | [**BulkReportsTrafficsTotalRequest**](BulkReportsTrafficsTotalRequest.md) |  | 

### Return type

[**BulkTrafficReportData**](BulkTrafficReportData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkReportsVisitorsTotal

> BulkVisitorReportData BulkReportsVisitorsTotal(ctx).Period(period).Since(since).Until(until).BulkReportsVisitorsTotalRequest(bulkReportsVisitorsTotalRequest).Execute()

Get visitor report for multiple domains

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    bulkReportsVisitorsTotalRequest := *openapiclient.NewBulkReportsVisitorsTotalRequest() // BulkReportsVisitorsTotalRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.BulkReportsVisitorsTotal(context.Background()).Period(period).Since(since).Until(until).BulkReportsVisitorsTotalRequest(bulkReportsVisitorsTotalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.BulkReportsVisitorsTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkReportsVisitorsTotal`: BulkVisitorReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.BulkReportsVisitorsTotal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkReportsVisitorsTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **bulkReportsVisitorsTotalRequest** | [**BulkReportsVisitorsTotalRequest**](BulkReportsVisitorsTotalRequest.md) |  | 

### Return type

[**BulkVisitorReportData**](BulkVisitorReportData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsAttacksAttackers

> ReportsAttacksAttackers200Response ReportsAttacksAttackers(ctx, domain).Period(period).Execute()

Get list of attackers information

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
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsAttacksAttackers(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsAttacksAttackers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksAttackers`: ReportsAttacksAttackers200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsAttacksAttackers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksAttackersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**ReportsAttacksAttackers200Response**](ReportsAttacksAttackers200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsAttacksIndex

> ReportsAttacksIndex200Response ReportsAttacksIndex(ctx, domain).Period(period).PerPage(perPage).Page(page).Execute()

Get list of attacks details

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
    period := "period_example" // string | Select period -ending now- for report (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsAttacksIndex(context.Background(), domain).Period(period).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsAttacksIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksIndex`: ReportsAttacksIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsAttacksIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**ReportsAttacksIndex200Response**](ReportsAttacksIndex200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsAttacksMap

> AttackReportMapData ReportsAttacksMap(ctx, domain).Period(period).Execute()

Get geo-map of attacks

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
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsAttacksMap(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsAttacksMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksMap`: AttackReportMapData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsAttacksMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**AttackReportMapData**](AttackReportMapData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsAttacksShow

> ReportsAttacksShow200Response ReportsAttacksShow(ctx, domain).Period(period).Execute()

Get report of attacks

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
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsAttacksShow(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsAttacksShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksShow`: ReportsAttacksShow200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsAttacksShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**ReportsAttacksShow200Response**](ReportsAttacksShow200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsAttacksUri

> AttackReportUriData ReportsAttacksUri(ctx, domain).Period(period).Execute()

Get list of URLs under attack

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
    period := "period_example" // string | Select period -ending now- for report (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsAttacksUri(context.Background(), domain).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsAttacksUri``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsAttacksUri`: AttackReportUriData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsAttacksUri`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsAttacksUriRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 

### Return type

[**AttackReportUriData**](AttackReportUriData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsDnsGeo

> DnsGeoReportData ReportsDnsGeo(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get DNS requests as geo-map

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsDnsGeo(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsDnsGeo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsDnsGeo`: DnsGeoReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsDnsGeo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDnsGeoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**DnsGeoReportData**](DnsGeoReportData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsDnsRequests

> DnsRequestReportData ReportsDnsRequests(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get response time report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsDnsRequests(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsDnsRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsDnsRequests`: DnsRequestReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsDnsRequests`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsDnsRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**DnsRequestReportData**](DnsRequestReportData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsErrorLogDetails

> ReportsErrorLogDetails200Response ReportsErrorLogDetails(ctx, domain).Period(period).Since(since).Until(until).Error_(error_).Execute()

Get detail of an error

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    error_ := "error__example" // string | Error message to search for (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsErrorLogDetails(context.Background(), domain).Period(period).Since(since).Until(until).Error_(error_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsErrorLogDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsErrorLogDetails`: ReportsErrorLogDetails200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsErrorLogDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsErrorLogDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **error_** | **string** | Error message to search for | 

### Return type

[**ReportsErrorLogDetails200Response**](ReportsErrorLogDetails200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsErrorLogs

> ErrorLogsData ReportsErrorLogs(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get list of errors

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsErrorLogs(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsErrorLogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsErrorLogs`: ErrorLogsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsErrorLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsErrorLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**ErrorLogsData**](ErrorLogsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsErrorLogsChart

> ErrorLogChartData ReportsErrorLogsChart(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get chart view of errors

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsErrorLogsChart(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsErrorLogsChart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsErrorLogsChart`: ErrorLogChartData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsErrorLogsChart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsErrorLogsChartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**ErrorLogChartData**](ErrorLogChartData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsResponseTimeIndex

> ResponseTimeData ReportsResponseTimeIndex(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get response time report

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsResponseTimeIndex(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsResponseTimeIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsResponseTimeIndex`: ResponseTimeData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsResponseTimeIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsResponseTimeIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**ResponseTimeData**](ResponseTimeData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsStatusIndex

> StatusCodeReportData ReportsStatusIndex(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get status codes pie chart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsStatusIndex(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsStatusIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsStatusIndex`: StatusCodeReportData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsStatusIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsStatusIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**StatusCodeReportData**](StatusCodeReportData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsStatusSummary

> StatusCodeSummaryData ReportsStatusSummary(ctx, domain).Period(period).Since(since).Until(until).Execute()

Get an overview of status codes pie chart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsStatusSummary(context.Background(), domain).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsStatusSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsStatusSummary`: StatusCodeSummaryData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsStatusSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsStatusSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**StatusCodeSummaryData**](StatusCodeSummaryData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsTrafficsMap

> MapTrafficsData ReportsTrafficsMap(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get traffic as geo-map

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsTrafficsMap(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsTrafficsMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTrafficsMap`: MapTrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsTrafficsMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTrafficsMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**MapTrafficsData**](MapTrafficsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsTrafficsSaved

> SavedTrafficsData ReportsTrafficsSaved(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get traffic saved to total pie chart

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsTrafficsSaved(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsTrafficsSaved``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTrafficsSaved`: SavedTrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsTrafficsSaved`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTrafficsSavedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**SavedTrafficsData**](SavedTrafficsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsTrafficsTotal

> TrafficsData ReportsTrafficsTotal(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get traffic report for domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsTrafficsTotal(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsTrafficsTotal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTrafficsTotal`: TrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsTrafficsTotal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTrafficsTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**TrafficsData**](TrafficsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsTransportLayerProxiesTraffics

> TransportLayerProxyTrafficsData ReportsTransportLayerProxiesTraffics(ctx, domain, transportLayerProxyId).Period(period).Since(since).Until(until).Execute()

Get traffic report for transport layer proxy

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    transportLayerProxyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Transport layer proxy id
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsTransportLayerProxiesTraffics(context.Background(), domain, transportLayerProxyId).Period(period).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsTransportLayerProxiesTraffics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsTransportLayerProxiesTraffics`: TransportLayerProxyTrafficsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsTransportLayerProxiesTraffics`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 
**transportLayerProxyId** | **string** | Transport layer proxy id | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsTransportLayerProxiesTrafficsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 

### Return type

[**TransportLayerProxyTrafficsData**](TransportLayerProxyTrafficsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsVisitorsHighRequestIps

> ReportsVisitorsHighRequestIps200Response ReportsVisitorsHighRequestIps(ctx, domain).Period(period).Since(since).Until(until).PerPage(perPage).Page(page).Execute()

Get report of IPs with highest number of requests

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsVisitorsHighRequestIps(context.Background(), domain).Period(period).Since(since).Until(until).PerPage(perPage).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsVisitorsHighRequestIps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsVisitorsHighRequestIps`: ReportsVisitorsHighRequestIps200Response
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsVisitorsHighRequestIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsVisitorsHighRequestIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]

### Return type

[**ReportsVisitorsHighRequestIps200Response**](ReportsVisitorsHighRequestIps200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReportsVisitorsIndex

> VisitorsData ReportsVisitorsIndex(ctx, domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()

Get report of visitors for domain

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/cdn"
)

func main() {
    domain := "example.com" // string | Domain name
    period := "period_example" // string | Select period -ending now- for report (optional)
    since := time.Now() // time.Time |  (optional)
    until := time.Now() // time.Time |  (optional)
    filterSubdomain := "filterSubdomain_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReportsAPI.ReportsVisitorsIndex(context.Background(), domain).Period(period).Since(since).Until(until).FilterSubdomain(filterSubdomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportsAPI.ReportsVisitorsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReportsVisitorsIndex`: VisitorsData
    fmt.Fprintf(os.Stdout, "Response from `ReportsAPI.ReportsVisitorsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | Domain name | 

### Other Parameters

Other parameters are passed through a pointer to a apiReportsVisitorsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** | Select period -ending now- for report | 
 **since** | **time.Time** |  | 
 **until** | **time.Time** |  | 
 **filterSubdomain** | **string** |  | 

### Return type

[**VisitorsData**](VisitorsData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

