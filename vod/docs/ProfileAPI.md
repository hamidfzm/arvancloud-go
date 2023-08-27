# \ProfileAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelProfilesGet**](ProfileAPI.md#ChannelsChannelProfilesGet) | **Get** /channels/{channel}/profiles | Return all channel&#39;s profiles.
[**ChannelsChannelProfilesPost**](ProfileAPI.md#ChannelsChannelProfilesPost) | **Post** /channels/{channel}/profiles | Store a newly created profile.
[**ProfilesProfileDelete**](ProfileAPI.md#ProfilesProfileDelete) | **Delete** /profiles/{profile} | Remove the specified profile.
[**ProfilesProfileGet**](ProfileAPI.md#ProfilesProfileGet) | **Get** /profiles/{profile} | Return the specified profile.
[**ProfilesProfilePatch**](ProfileAPI.md#ProfilesProfilePatch) | **Patch** /profiles/{profile} | Update the specified profile.



## ChannelsChannelProfilesGet

> ChannelsChannelProfilesGet(ctx, channel).Filter(filter).Page(page).PerPage(perPage).Execute()

Return all channel's profiles.

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
    channel := "channel_example" // string | The Id of channel
    filter := "filter_example" // string | Filter result (optional)
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfileAPI.ChannelsChannelProfilesGet(context.Background(), channel).Filter(filter).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.ChannelsChannelProfilesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelProfilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 
 **page** | **int32** | Page number | 
 **perPage** | **int32** | Page limit | 

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


## ChannelsChannelProfilesPost

> ChannelsChannelProfilesPost(ctx, channel).Body(body).Execute()

Store a newly created profile.

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
    channel := "channel_example" // string | The Id of channel
    body := *openapiclient.NewChannelsChannelProfilesPostRequest("Title_example", "ConvertMode_example") // ChannelsChannelProfilesPostRequest | Profile details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfileAPI.ChannelsChannelProfilesPost(context.Background(), channel).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.ChannelsChannelProfilesPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiChannelsChannelProfilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ChannelsChannelProfilesPostRequest**](ChannelsChannelProfilesPostRequest.md) | Profile details | 

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


## ProfilesProfileDelete

> ProfilesProfileDelete(ctx, profile).Execute()

Remove the specified profile.

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
    profile := "profile_example" // string | The Id of profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfileAPI.ProfilesProfileDelete(context.Background(), profile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.ProfilesProfileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profile** | **string** | The Id of profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesProfileDeleteRequest struct via the builder pattern


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


## ProfilesProfileGet

> ProfilesProfileGet(ctx, profile).Execute()

Return the specified profile.

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
    profile := "profile_example" // string | The Id of profile

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfileAPI.ProfilesProfileGet(context.Background(), profile).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.ProfilesProfileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profile** | **string** | The Id of profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesProfileGetRequest struct via the builder pattern


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


## ProfilesProfilePatch

> ProfilesProfilePatch(ctx, profile).Body(body).Execute()

Update the specified profile.

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
    profile := "profile_example" // string | The Id of profile
    body := *openapiclient.NewProfilesProfilePatchRequest() // ProfilesProfilePatchRequest | Profile details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ProfileAPI.ProfilesProfilePatch(context.Background(), profile).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProfileAPI.ProfilesProfilePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**profile** | **string** | The Id of profile | 

### Other Parameters

Other parameters are passed through a pointer to a apiProfilesProfilePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ProfilesProfilePatchRequest**](ProfilesProfilePatchRequest.md) | Profile details | 

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

