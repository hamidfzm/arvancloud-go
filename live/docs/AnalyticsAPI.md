# \AnalyticsAPI

All URIs are relative to *https://napi.arvancloud.ir/live/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyticsPlayCountGet**](AnalyticsAPI.md#AnalyticsPlayCountGet) | **Get** /analytics/play-count | Return appropriate play count
[**AnalyticsTrafficGet**](AnalyticsAPI.md#AnalyticsTrafficGet) | **Get** /analytics/traffic | Return appropriate traffic
[**AnalyticsWatchTimeGet**](AnalyticsAPI.md#AnalyticsWatchTimeGet) | **Get** /analytics/watch-time | Return appropriate watch time



## AnalyticsPlayCountGet

> AnalyticsPlayCountGet(ctx).Period(period).Since(since).Until(until).Interval(interval).OrderByName(orderByName).OrderByOrder(orderByOrder).GroupBy(groupBy).Limit(limit).Offset(offset).Aggregate(aggregate).Timezone(timezone).FiltersStream(filtersStream).FiltersCountry(filtersCountry).FiltersAsn(filtersAsn).FiltersClientType(filtersClientType).FiltersClientFamily(filtersClientFamily).FiltersOsFamily(filtersOsFamily).FiltersDeviceType(filtersDeviceType).FiltersDeviceBrand(filtersDeviceBrand).FiltersDeviceModel(filtersDeviceModel).FiltersResolution(filtersResolution).Execute()

Return appropriate play count



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
    period := "period_example" // string | Enum: '1h' '2h' '3h' '6h' '12h' '24h' '3d' '7d' '14d' '1m'  (optional)
    since := "since_example" // string | string <date-time:Y-m-d H:i>|<date: Y-m-d> (optional)
    until := "until_example" // string | string <date-time:Y-m-d H:i>|<date: Y-m-d> (optional)
    interval := "interval_example" // string | Enum: 'minutely' 'hourly' 'daily' (optional)
    orderByName := "orderByName_example" // string | Enum: watch_time (optional)
    orderByOrder := "orderByOrder_example" // string | Enum: 'ASC' 'DESC' (optional)
    groupBy := "groupBy_example" // string | Enum: 'channel' 'video' 'country' 'asn' (optional)
    limit := int32(56) // int32 | Limit the number of entities (optional)
    offset := int32(56) // int32 | Define offset of entities (optional)
    aggregate := "aggregate_example" // string | Enum: 'SUM' 'AVG' 'MIN' 'MAX' 'COUNT' (optional)
    timezone := "timezone_example" // string | Timezone ex: Asia/Tehran (optional)
    filtersStream := []string{"Inner_example"} // []string | Filter by stream IDs (optional)
    filtersCountry := []string{"Inner_example"} // []string | Filter by country IDs (optional)
    filtersAsn := []string{"Inner_example"} // []string | Filter by asn numbers (optional)
    filtersClientType := []string{"Inner_example"} // []string | Filter by client type (optional)
    filtersClientFamily := []string{"Inner_example"} // []string | Filter by client family (optional)
    filtersOsFamily := []string{"Inner_example"} // []string | Filter by os family (optional)
    filtersDeviceType := []string{"Inner_example"} // []string | Filter by device type (optional)
    filtersDeviceBrand := []string{"Inner_example"} // []string | Filter by device brand (optional)
    filtersDeviceModel := []string{"Inner_example"} // []string | Filter by device model (optional)
    filtersResolution := []string{"Inner_example"} // []string | Filter by resolution (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnalyticsAPI.AnalyticsPlayCountGet(context.Background()).Period(period).Since(since).Until(until).Interval(interval).OrderByName(orderByName).OrderByOrder(orderByOrder).GroupBy(groupBy).Limit(limit).Offset(offset).Aggregate(aggregate).Timezone(timezone).FiltersStream(filtersStream).FiltersCountry(filtersCountry).FiltersAsn(filtersAsn).FiltersClientType(filtersClientType).FiltersClientFamily(filtersClientFamily).FiltersOsFamily(filtersOsFamily).FiltersDeviceType(filtersDeviceType).FiltersDeviceBrand(filtersDeviceBrand).FiltersDeviceModel(filtersDeviceModel).FiltersResolution(filtersResolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsPlayCountGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsPlayCountGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Enum: &#39;1h&#39; &#39;2h&#39; &#39;3h&#39; &#39;6h&#39; &#39;12h&#39; &#39;24h&#39; &#39;3d&#39; &#39;7d&#39; &#39;14d&#39; &#39;1m&#39;  | 
 **since** | **string** | string &lt;date-time:Y-m-d H:i&gt;|&lt;date: Y-m-d&gt; | 
 **until** | **string** | string &lt;date-time:Y-m-d H:i&gt;|&lt;date: Y-m-d&gt; | 
 **interval** | **string** | Enum: &#39;minutely&#39; &#39;hourly&#39; &#39;daily&#39; | 
 **orderByName** | **string** | Enum: watch_time | 
 **orderByOrder** | **string** | Enum: &#39;ASC&#39; &#39;DESC&#39; | 
 **groupBy** | **string** | Enum: &#39;channel&#39; &#39;video&#39; &#39;country&#39; &#39;asn&#39; | 
 **limit** | **int32** | Limit the number of entities | 
 **offset** | **int32** | Define offset of entities | 
 **aggregate** | **string** | Enum: &#39;SUM&#39; &#39;AVG&#39; &#39;MIN&#39; &#39;MAX&#39; &#39;COUNT&#39; | 
 **timezone** | **string** | Timezone ex: Asia/Tehran | 
 **filtersStream** | **[]string** | Filter by stream IDs | 
 **filtersCountry** | **[]string** | Filter by country IDs | 
 **filtersAsn** | **[]string** | Filter by asn numbers | 
 **filtersClientType** | **[]string** | Filter by client type | 
 **filtersClientFamily** | **[]string** | Filter by client family | 
 **filtersOsFamily** | **[]string** | Filter by os family | 
 **filtersDeviceType** | **[]string** | Filter by device type | 
 **filtersDeviceBrand** | **[]string** | Filter by device brand | 
 **filtersDeviceModel** | **[]string** | Filter by device model | 
 **filtersResolution** | **[]string** | Filter by resolution | 

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


## AnalyticsTrafficGet

> AnalyticsTrafficGet(ctx).Period(period).Since(since).Until(until).Interval(interval).OrderByName(orderByName).OrderByOrder(orderByOrder).GroupBy(groupBy).Limit(limit).Offset(offset).Aggregate(aggregate).Timezone(timezone).FiltersStream(filtersStream).FiltersCountry(filtersCountry).FiltersAsn(filtersAsn).FiltersClientType(filtersClientType).FiltersClientFamily(filtersClientFamily).FiltersOsFamily(filtersOsFamily).FiltersDeviceType(filtersDeviceType).FiltersDeviceBrand(filtersDeviceBrand).FiltersDeviceModel(filtersDeviceModel).FiltersResolution(filtersResolution).Execute()

Return appropriate traffic



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
    period := "period_example" // string | Enum: '1h' '2h' '3h' '6h' '12h' '24h' '3d' '7d' '14d' '1m' (optional)
    since := "since_example" // string | string <date-time:Y-m-d H:i>|<date: Y-m-d> (optional)
    until := "until_example" // string | string <date-time:Y-m-d H:i>|<date: Y-m-d> (optional)
    interval := "interval_example" // string | Enum: 'minutely' 'hourly' 'daily' (optional)
    orderByName := "orderByName_example" // string | Enum: traffic (optional)
    orderByOrder := "orderByOrder_example" // string | Enum: 'ASC' 'DESC' (optional)
    groupBy := "groupBy_example" // string | Enum: 'channel' 'video' 'country' 'asn' 'referer' (optional)
    limit := int32(56) // int32 | Limit the number of entities (optional)
    offset := int32(56) // int32 | Define offset of entities (optional)
    aggregate := "aggregate_example" // string | Enum: 'SUM' 'AVG' 'MIN' 'MAX' 'COUNT' (optional)
    timezone := "timezone_example" // string | Timezone ex: Asia/Tehran (optional)
    filtersStream := []string{"Inner_example"} // []string | Filter by stream IDs (optional)
    filtersCountry := []string{"Inner_example"} // []string | Filter by country IDs (optional)
    filtersAsn := []string{"Inner_example"} // []string | Filter by asn numbers (optional)
    filtersClientType := []string{"Inner_example"} // []string | Filter by client type (optional)
    filtersClientFamily := []string{"Inner_example"} // []string | Filter by client family (optional)
    filtersOsFamily := []string{"Inner_example"} // []string | Filter by os family (optional)
    filtersDeviceType := []string{"Inner_example"} // []string | Filter by device type (optional)
    filtersDeviceBrand := []string{"Inner_example"} // []string | Filter by device brand (optional)
    filtersDeviceModel := []string{"Inner_example"} // []string | Filter by device model (optional)
    filtersResolution := []string{"Inner_example"} // []string | Filter by resolution (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnalyticsAPI.AnalyticsTrafficGet(context.Background()).Period(period).Since(since).Until(until).Interval(interval).OrderByName(orderByName).OrderByOrder(orderByOrder).GroupBy(groupBy).Limit(limit).Offset(offset).Aggregate(aggregate).Timezone(timezone).FiltersStream(filtersStream).FiltersCountry(filtersCountry).FiltersAsn(filtersAsn).FiltersClientType(filtersClientType).FiltersClientFamily(filtersClientFamily).FiltersOsFamily(filtersOsFamily).FiltersDeviceType(filtersDeviceType).FiltersDeviceBrand(filtersDeviceBrand).FiltersDeviceModel(filtersDeviceModel).FiltersResolution(filtersResolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsTrafficGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsTrafficGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Enum: &#39;1h&#39; &#39;2h&#39; &#39;3h&#39; &#39;6h&#39; &#39;12h&#39; &#39;24h&#39; &#39;3d&#39; &#39;7d&#39; &#39;14d&#39; &#39;1m&#39; | 
 **since** | **string** | string &lt;date-time:Y-m-d H:i&gt;|&lt;date: Y-m-d&gt; | 
 **until** | **string** | string &lt;date-time:Y-m-d H:i&gt;|&lt;date: Y-m-d&gt; | 
 **interval** | **string** | Enum: &#39;minutely&#39; &#39;hourly&#39; &#39;daily&#39; | 
 **orderByName** | **string** | Enum: traffic | 
 **orderByOrder** | **string** | Enum: &#39;ASC&#39; &#39;DESC&#39; | 
 **groupBy** | **string** | Enum: &#39;channel&#39; &#39;video&#39; &#39;country&#39; &#39;asn&#39; &#39;referer&#39; | 
 **limit** | **int32** | Limit the number of entities | 
 **offset** | **int32** | Define offset of entities | 
 **aggregate** | **string** | Enum: &#39;SUM&#39; &#39;AVG&#39; &#39;MIN&#39; &#39;MAX&#39; &#39;COUNT&#39; | 
 **timezone** | **string** | Timezone ex: Asia/Tehran | 
 **filtersStream** | **[]string** | Filter by stream IDs | 
 **filtersCountry** | **[]string** | Filter by country IDs | 
 **filtersAsn** | **[]string** | Filter by asn numbers | 
 **filtersClientType** | **[]string** | Filter by client type | 
 **filtersClientFamily** | **[]string** | Filter by client family | 
 **filtersOsFamily** | **[]string** | Filter by os family | 
 **filtersDeviceType** | **[]string** | Filter by device type | 
 **filtersDeviceBrand** | **[]string** | Filter by device brand | 
 **filtersDeviceModel** | **[]string** | Filter by device model | 
 **filtersResolution** | **[]string** | Filter by resolution | 

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


## AnalyticsWatchTimeGet

> AnalyticsWatchTimeGet(ctx).Period(period).Since(since).Until(until).Interval(interval).OrderByName(orderByName).OrderByOrder(orderByOrder).GroupBy(groupBy).Limit(limit).Offset(offset).Aggregate(aggregate).Timezone(timezone).FiltersStream(filtersStream).FiltersCountry(filtersCountry).FiltersAsn(filtersAsn).FiltersClientType(filtersClientType).FiltersClientFamily(filtersClientFamily).FiltersOsFamily(filtersOsFamily).FiltersDeviceType(filtersDeviceType).FiltersDeviceBrand(filtersDeviceBrand).FiltersDeviceModel(filtersDeviceModel).FiltersResolution(filtersResolution).Execute()

Return appropriate watch time



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
    period := "period_example" // string | Enum: '1h' '2h' '3h' '6h' '12h' '24h' '3d' '7d' '14d' '1m' (optional)
    since := "since_example" // string | string <date-time:Y-m-d H:i>|<date: Y-m-d> (optional)
    until := "until_example" // string | string <date-time:Y-m-d H:i>|<date: Y-m-d> (optional)
    interval := "interval_example" // string | Enum: 'minutely' 'hourly' 'daily' (optional)
    orderByName := "orderByName_example" // string | Enum: watch_time (optional)
    orderByOrder := "orderByOrder_example" // string | Enum: 'ASC' 'DESC' (optional)
    groupBy := "groupBy_example" // string | Enum: 'channel' 'video' 'country' 'asn' (optional)
    limit := int32(56) // int32 | Limit the number of entities (optional)
    offset := int32(56) // int32 | Define offset of entities (optional)
    aggregate := "aggregate_example" // string | Enum: 'SUM' 'AVG' 'MIN' 'MAX' 'COUNT' (optional)
    timezone := "timezone_example" // string | Timezone ex: Asia/Tehran (optional)
    filtersStream := []string{"Inner_example"} // []string | Filter by stream IDs (optional)
    filtersCountry := []string{"Inner_example"} // []string | Filter by country IDs (optional)
    filtersAsn := []string{"Inner_example"} // []string | Filter by asn numbers (optional)
    filtersClientType := []string{"Inner_example"} // []string | Filter by client type (optional)
    filtersClientFamily := []string{"Inner_example"} // []string | Filter by client family (optional)
    filtersOsFamily := []string{"Inner_example"} // []string | Filter by os family (optional)
    filtersDeviceType := []string{"Inner_example"} // []string | Filter by device type (optional)
    filtersDeviceBrand := []string{"Inner_example"} // []string | Filter by device brand (optional)
    filtersDeviceModel := []string{"Inner_example"} // []string | Filter by device model (optional)
    filtersResolution := []string{"Inner_example"} // []string | Filter by resolution (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AnalyticsAPI.AnalyticsWatchTimeGet(context.Background()).Period(period).Since(since).Until(until).Interval(interval).OrderByName(orderByName).OrderByOrder(orderByOrder).GroupBy(groupBy).Limit(limit).Offset(offset).Aggregate(aggregate).Timezone(timezone).FiltersStream(filtersStream).FiltersCountry(filtersCountry).FiltersAsn(filtersAsn).FiltersClientType(filtersClientType).FiltersClientFamily(filtersClientFamily).FiltersOsFamily(filtersOsFamily).FiltersDeviceType(filtersDeviceType).FiltersDeviceBrand(filtersDeviceBrand).FiltersDeviceModel(filtersDeviceModel).FiltersResolution(filtersResolution).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsAPI.AnalyticsWatchTimeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyticsWatchTimeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **period** | **string** | Enum: &#39;1h&#39; &#39;2h&#39; &#39;3h&#39; &#39;6h&#39; &#39;12h&#39; &#39;24h&#39; &#39;3d&#39; &#39;7d&#39; &#39;14d&#39; &#39;1m&#39; | 
 **since** | **string** | string &lt;date-time:Y-m-d H:i&gt;|&lt;date: Y-m-d&gt; | 
 **until** | **string** | string &lt;date-time:Y-m-d H:i&gt;|&lt;date: Y-m-d&gt; | 
 **interval** | **string** | Enum: &#39;minutely&#39; &#39;hourly&#39; &#39;daily&#39; | 
 **orderByName** | **string** | Enum: watch_time | 
 **orderByOrder** | **string** | Enum: &#39;ASC&#39; &#39;DESC&#39; | 
 **groupBy** | **string** | Enum: &#39;channel&#39; &#39;video&#39; &#39;country&#39; &#39;asn&#39; | 
 **limit** | **int32** | Limit the number of entities | 
 **offset** | **int32** | Define offset of entities | 
 **aggregate** | **string** | Enum: &#39;SUM&#39; &#39;AVG&#39; &#39;MIN&#39; &#39;MAX&#39; &#39;COUNT&#39; | 
 **timezone** | **string** | Timezone ex: Asia/Tehran | 
 **filtersStream** | **[]string** | Filter by stream IDs | 
 **filtersCountry** | **[]string** | Filter by country IDs | 
 **filtersAsn** | **[]string** | Filter by asn numbers | 
 **filtersClientType** | **[]string** | Filter by client type | 
 **filtersClientFamily** | **[]string** | Filter by client family | 
 **filtersOsFamily** | **[]string** | Filter by os family | 
 **filtersDeviceType** | **[]string** | Filter by device type | 
 **filtersDeviceBrand** | **[]string** | Filter by device brand | 
 **filtersDeviceModel** | **[]string** | Filter by device model | 
 **filtersResolution** | **[]string** | Filter by resolution | 

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

