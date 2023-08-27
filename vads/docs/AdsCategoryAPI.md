# \AdsCategoryAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdsAdCategoriesCategoryDelete**](AdsCategoryAPI.md#AdsAdCategoriesCategoryDelete) | **Delete** /ads/{ad}/categories/{category} | Detach category from ad.
[**AdsAdCategoriesCategoryPut**](AdsCategoryAPI.md#AdsAdCategoriesCategoryPut) | **Put** /ads/{ad}/categories/{category} | Attach category to ad
[**AdsAdCategoriesGet**](AdsCategoryAPI.md#AdsAdCategoriesGet) | **Get** /ads/{ad}/categories | Return all ad&#39;s categories.



## AdsAdCategoriesCategoryDelete

> AdsAdCategoriesCategoryDelete(ctx, ad, category).Execute()

Detach category from ad.

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
    category := "category_example" // string | The Id of category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdsCategoryAPI.AdsAdCategoriesCategoryDelete(context.Background(), ad, category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsCategoryAPI.AdsAdCategoriesCategoryDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ad** | **string** | The Id of ad | 
**category** | **string** | The Id of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdsAdCategoriesCategoryDeleteRequest struct via the builder pattern


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


## AdsAdCategoriesCategoryPut

> AdsAdCategoriesCategoryPut(ctx, ad, category).Execute()

Attach category to ad

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
    category := "category_example" // string | The Id of category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdsCategoryAPI.AdsAdCategoriesCategoryPut(context.Background(), ad, category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsCategoryAPI.AdsAdCategoriesCategoryPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ad** | **string** | The Id of ad | 
**category** | **string** | The Id of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdsAdCategoriesCategoryPutRequest struct via the builder pattern


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


## AdsAdCategoriesGet

> AdsAdCategoriesGet(ctx, ad).Execute()

Return all ad's categories.

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
    r, err := apiClient.AdsCategoryAPI.AdsAdCategoriesGet(context.Background(), ad).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsCategoryAPI.AdsAdCategoriesGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAdsAdCategoriesGetRequest struct via the builder pattern


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

