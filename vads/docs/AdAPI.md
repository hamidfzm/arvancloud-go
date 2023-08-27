# \AdAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdsAdDelete**](AdAPI.md#AdsAdDelete) | **Delete** /ads/{ad} | Remove the specified ad.
[**AdsAdGet**](AdAPI.md#AdsAdGet) | **Get** /ads/{ad} | Return the specified ad.
[**AdsAdPut**](AdAPI.md#AdsAdPut) | **Put** /ads/{ad} | Update the specified ad.
[**ChannelsChannelAdsGet**](AdAPI.md#ChannelsChannelAdsGet) | **Get** /channels/{channel}/ads | Return all channel&#39;s ads.
[**ChannelsChannelAdsPost**](AdAPI.md#ChannelsChannelAdsPost) | **Post** /channels/{channel}/ads | Store a newly ad for specific channel.



## AdsAdDelete

> AdsAdDelete(ctx, ad).Execute()

Remove the specified ad.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdAPI.AdsAdDelete(context.Background(), ad).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.AdsAdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ad** | **string** | The Id of ad | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdsAdDeleteRequest struct via the builder pattern


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


## AdsAdGet

> AdsAdGet(ctx, ad).Execute()

Return the specified ad.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdAPI.AdsAdGet(context.Background(), ad).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.AdsAdGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ad** | **string** | The Id of ad | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdsAdGetRequest struct via the builder pattern


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


## AdsAdPut

> AdsAdPut(ctx, ad).Body(body).Execute()

Update the specified ad.

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
    body := *openapiclient.NewAdsAdPutRequest() // AdsAdPutRequest | Ad's details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdAPI.AdsAdPut(context.Background(), ad).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.AdsAdPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ad** | **string** | The Id of ad | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdsAdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**AdsAdPutRequest**](AdsAdPutRequest.md) | Ad&#39;s details | 

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


## ChannelsChannelAdsGet

> ChannelsChannelAdsGet(ctx, channel).Filter(filter).Page(page).PerPage(perPage).Execute()

Return all channel's ads.

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
    filter := "filter_example" // string | Filter result (optional)
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit for query (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdAPI.ChannelsChannelAdsGet(context.Background(), channel).Filter(filter).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.ChannelsChannelAdsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelAdsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 
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


## ChannelsChannelAdsPost

> ChannelsChannelAdsPost(ctx, channel).Title(title).AdType(adType).SkipType(skipType).ClickThrough(clickThrough).MediaFile(mediaFile).Description(description).SkipOffset(skipOffset).Execute()

Store a newly ad for specific channel.

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
    title := "title_example" // string | Title of ad
    adType := "adType_example" // string | Ad type
    skipType := "skipType_example" // string | Skip type
    clickThrough := "clickThrough_example" // string | Click URL when user click on ad
    mediaFile := os.NewFile(1234, "some_file") // *os.File | Media file for ad (Accept video/mp4 for pre_roll ads)
    description := "description_example" // string | Description of ad (optional)
    skipOffset := int32(56) // int32 | Skip offset in seconds (required if skip type is allow) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdAPI.ChannelsChannelAdsPost(context.Background(), channel).Title(title).AdType(adType).SkipType(skipType).ClickThrough(clickThrough).MediaFile(mediaFile).Description(description).SkipOffset(skipOffset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdAPI.ChannelsChannelAdsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelAdsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **title** | **string** | Title of ad | 
 **adType** | **string** | Ad type | 
 **skipType** | **string** | Skip type | 
 **clickThrough** | **string** | Click URL when user click on ad | 
 **mediaFile** | ***os.File** | Media file for ad (Accept video/mp4 for pre_roll ads) | 
 **description** | **string** | Description of ad | 
 **skipOffset** | **int32** | Skip offset in seconds (required if skip type is allow) | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

