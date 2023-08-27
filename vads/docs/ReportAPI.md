# \ReportAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdsAdReportsTrackEventPeriodGet**](ReportAPI.md#AdsAdReportsTrackEventPeriodGet) | **Get** /ads/{ad}/reports/track/{event}/{period} | Ad track report per event.
[**CampaignsCampaignAdsAdReportsTrackEventPeriodGet**](ReportAPI.md#CampaignsCampaignAdsAdReportsTrackEventPeriodGet) | **Get** /campaigns/{campaign}/ads/{ad}/reports/track/{event}/{period} | Ad in campaign track report per event.
[**CampaignsCampaignReportsTrackEventPeriodGet**](ReportAPI.md#CampaignsCampaignReportsTrackEventPeriodGet) | **Get** /campaigns/{campaign}/reports/track/{event}/{period} | Campaign track report per event.



## AdsAdReportsTrackEventPeriodGet

> AdsAdReportsTrackEventPeriodGet(ctx, ad, event, period).From(from).To(to).Execute()

Ad track report per event.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    ad := "ad_example" // string | The Id of ad
    event := "event_example" // string | Event type
    period := "period_example" // string | Report period
    from := "from_example" // string | Starting datetime of report. Format: 'YYYY-MM-DD H:i:s' like: '2018-01-01 00:00:00'
    to := "to_example" // string | Ending datetime of report. Format: 'YYYY-MM-DD H:i:s' like: '2018-01-01 00:00:00'

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportAPI.AdsAdReportsTrackEventPeriodGet(context.Background(), ad, event, period).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.AdsAdReportsTrackEventPeriodGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ad** | **string** | The Id of ad | 
**event** | **string** | Event type | 
**period** | **string** | Report period | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdsAdReportsTrackEventPeriodGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **string** | Starting datetime of report. Format: &#39;YYYY-MM-DD H:i:s&#39; like: &#39;2018-01-01 00:00:00&#39; | 
 **to** | **string** | Ending datetime of report. Format: &#39;YYYY-MM-DD H:i:s&#39; like: &#39;2018-01-01 00:00:00&#39; | 

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


## CampaignsCampaignAdsAdReportsTrackEventPeriodGet

> CampaignsCampaignAdsAdReportsTrackEventPeriodGet(ctx, campaign, ad, event, period).From(from).To(to).Execute()

Ad in campaign track report per event.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    campaign := "campaign_example" // string | The Id of campaign
    ad := "ad_example" // string | The Id of ad
    event := "event_example" // string | Event type
    period := "period_example" // string | Report period
    from := "from_example" // string | Starting datetime of report. Format: 'YYYY-MM-DD H:i:s' like: '2018-01-01 00:00:00'
    to := "to_example" // string | Ending datetime of report. Format: 'YYYY-MM-DD H:i:s' like: '2018-01-01 00:00:00'

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportAPI.CampaignsCampaignAdsAdReportsTrackEventPeriodGet(context.Background(), campaign, ad, event, period).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.CampaignsCampaignAdsAdReportsTrackEventPeriodGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaign** | **string** | The Id of campaign | 
**ad** | **string** | The Id of ad | 
**event** | **string** | Event type | 
**period** | **string** | Report period | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignAdsAdReportsTrackEventPeriodGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **from** | **string** | Starting datetime of report. Format: &#39;YYYY-MM-DD H:i:s&#39; like: &#39;2018-01-01 00:00:00&#39; | 
 **to** | **string** | Ending datetime of report. Format: &#39;YYYY-MM-DD H:i:s&#39; like: &#39;2018-01-01 00:00:00&#39; | 

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


## CampaignsCampaignReportsTrackEventPeriodGet

> CampaignsCampaignReportsTrackEventPeriodGet(ctx, campaign, event, period).From(from).To(to).Execute()

Campaign track report per event.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vads"
)

func main() {
    campaign := "campaign_example" // string | The Id of campaign
    event := "event_example" // string | Event type
    period := "period_example" // string | Report period
    from := "from_example" // string | Starting datetime of report. Format: 'YYYY-MM-DD H:i:s' like: '2018-01-01 00:00:00'
    to := "to_example" // string | Ending datetime of report. Format: 'YYYY-MM-DD H:i:s' like: '2018-01-01 00:00:00'

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReportAPI.CampaignsCampaignReportsTrackEventPeriodGet(context.Background(), campaign, event, period).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReportAPI.CampaignsCampaignReportsTrackEventPeriodGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaign** | **string** | The Id of campaign | 
**event** | **string** | Event type | 
**period** | **string** | Report period | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignReportsTrackEventPeriodGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **string** | Starting datetime of report. Format: &#39;YYYY-MM-DD H:i:s&#39; like: &#39;2018-01-01 00:00:00&#39; | 
 **to** | **string** | Ending datetime of report. Format: &#39;YYYY-MM-DD H:i:s&#39; like: &#39;2018-01-01 00:00:00&#39; | 

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

