# \CampaignsAdAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignsCampaignAdsAdDelete**](CampaignsAdAPI.md#CampaignsCampaignAdsAdDelete) | **Delete** /campaigns/{campaign}/ads/{ad} | Detach ad from campaign.
[**CampaignsCampaignAdsAdGet**](CampaignsAdAPI.md#CampaignsCampaignAdsAdGet) | **Get** /campaigns/{campaign}/ads/{ad} | Show attach detail of specific campaign&#39;s ad.
[**CampaignsCampaignAdsAdPut**](CampaignsAdAPI.md#CampaignsCampaignAdsAdPut) | **Put** /campaigns/{campaign}/ads/{ad} | Update the specified campaign&#39;s ad.
[**CampaignsCampaignAdsGet**](CampaignsAdAPI.md#CampaignsCampaignAdsGet) | **Get** /campaigns/{campaign}/ads | Return all campaign&#39;s ads.
[**CampaignsCampaignAdsPost**](CampaignsAdAPI.md#CampaignsCampaignAdsPost) | **Post** /campaigns/{campaign}/ads | Attach ad to campaign



## CampaignsCampaignAdsAdDelete

> CampaignsCampaignAdsAdDelete(ctx, campaign, ad).Execute()

Detach ad from campaign.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsAdAPI.CampaignsCampaignAdsAdDelete(context.Background(), campaign, ad).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAdAPI.CampaignsCampaignAdsAdDelete``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignAdsAdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CampaignsCampaignAdsAdGet

> CampaignsCampaignAdsAdGet(ctx, campaign, ad).Execute()

Show attach detail of specific campaign's ad.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsAdAPI.CampaignsCampaignAdsAdGet(context.Background(), campaign, ad).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAdAPI.CampaignsCampaignAdsAdGet``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignAdsAdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## CampaignsCampaignAdsAdPut

> CampaignsCampaignAdsAdPut(ctx, campaign, ad).Body(body).Execute()

Update the specified campaign's ad.

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
    body := *openapiclient.NewCampaignsCampaignAdsAdPutRequest() // CampaignsCampaignAdsAdPutRequest | Campaign's ad details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsAdAPI.CampaignsCampaignAdsAdPut(context.Background(), campaign, ad).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAdAPI.CampaignsCampaignAdsAdPut``: %v\n", err)
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

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignAdsAdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**CampaignsCampaignAdsAdPutRequest**](CampaignsCampaignAdsAdPutRequest.md) | Campaign&#39;s ad details | 

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


## CampaignsCampaignAdsGet

> CampaignsCampaignAdsGet(ctx, campaign).Execute()

Return all campaign's ads.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsAdAPI.CampaignsCampaignAdsGet(context.Background(), campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAdAPI.CampaignsCampaignAdsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaign** | **string** | The Id of campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignAdsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## CampaignsCampaignAdsPost

> CampaignsCampaignAdsPost(ctx, campaign).Body(body).Execute()

Attach ad to campaign

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
    body := *openapiclient.NewCampaignsCampaignAdsPostRequest("AdId_example", int32(123)) // CampaignsCampaignAdsPostRequest | Campaign's ad details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignsAdAPI.CampaignsCampaignAdsPost(context.Background(), campaign).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAdAPI.CampaignsCampaignAdsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaign** | **string** | The Id of campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiCampaignsCampaignAdsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignsCampaignAdsPostRequest**](CampaignsCampaignAdsPostRequest.md) | Campaign&#39;s ad details | 

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

