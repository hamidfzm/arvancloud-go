# \CampaignAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CampaignsCampaignDelete**](CampaignAPI.md#CampaignsCampaignDelete) | **Delete** /campaigns/{campaign} | Remove the specified campaign.
[**CampaignsCampaignGet**](CampaignAPI.md#CampaignsCampaignGet) | **Get** /campaigns/{campaign} | Return the specified campaign.
[**CampaignsCampaignPut**](CampaignAPI.md#CampaignsCampaignPut) | **Put** /campaigns/{campaign} | Update the specified campaign.
[**ChannelsChannelCampaignsGet**](CampaignAPI.md#ChannelsChannelCampaignsGet) | **Get** /channels/{channel}/campaigns | Return all channel&#39;s campaigns.
[**ChannelsChannelCampaignsPost**](CampaignAPI.md#ChannelsChannelCampaignsPost) | **Post** /channels/{channel}/campaigns | Store a newly campaign for specific channel.



## CampaignsCampaignDelete

> CampaignsCampaignDelete(ctx, campaign).Execute()

Remove the specified campaign.

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
    r, err := apiClient.CampaignAPI.CampaignsCampaignDelete(context.Background(), campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CampaignsCampaignDelete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCampaignsCampaignDeleteRequest struct via the builder pattern


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


## CampaignsCampaignGet

> CampaignsCampaignGet(ctx, campaign).Execute()

Return the specified campaign.

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
    r, err := apiClient.CampaignAPI.CampaignsCampaignGet(context.Background(), campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CampaignsCampaignGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCampaignsCampaignGetRequest struct via the builder pattern


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


## CampaignsCampaignPut

> CampaignsCampaignPut(ctx, campaign).Body(body).Execute()

Update the specified campaign.

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
    body := *openapiclient.NewCampaignsCampaignPutRequest() // CampaignsCampaignPutRequest | Ad's ad details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignAPI.CampaignsCampaignPut(context.Background(), campaign).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.CampaignsCampaignPut``: %v\n", err)
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

Other parameters are passed through a pointer to a apiCampaignsCampaignPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CampaignsCampaignPutRequest**](CampaignsCampaignPutRequest.md) | Ad&#39;s ad details | 

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


## ChannelsChannelCampaignsGet

> ChannelsChannelCampaignsGet(ctx, channel).Page(page).PerPage(perPage).Execute()

Return all channel's campaigns.

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
    channel := "channel_example" // string | The Id of channel
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit for query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignAPI.ChannelsChannelCampaignsGet(context.Background(), channel).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.ChannelsChannelCampaignsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelCampaignsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit for query | 

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


## ChannelsChannelCampaignsPost

> ChannelsChannelCampaignsPost(ctx, channel).Body(body).Execute()

Store a newly campaign for specific channel.

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
    channel := "channel_example" // string | The Id of channel
    body := *openapiclient.NewChannelsChannelCampaignsPostRequest("Title_example", "SkipType_example", false) // ChannelsChannelCampaignsPostRequest | Ad's ad details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CampaignAPI.ChannelsChannelCampaignsPost(context.Background(), channel).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.ChannelsChannelCampaignsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelCampaignsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChannelsChannelCampaignsPostRequest**](ChannelsChannelCampaignsPostRequest.md) | Ad&#39;s ad details | 

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

