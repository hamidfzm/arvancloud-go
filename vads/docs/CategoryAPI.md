# \CategoryAPI

All URIs are relative to *https://napi.arvancloud.ir/vads/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CategoriesCategoryDelete**](CategoryAPI.md#CategoriesCategoryDelete) | **Delete** /categories/{category} | Remove the specified category.
[**CategoriesCategoryGet**](CategoryAPI.md#CategoriesCategoryGet) | **Get** /categories/{category} | Return the specified category.
[**CategoriesCategoryPut**](CategoryAPI.md#CategoriesCategoryPut) | **Put** /categories/{category} | Update the specified category.
[**CategoriesGet**](CategoryAPI.md#CategoriesGet) | **Get** /categories | Return all user categories.
[**CategoriesPost**](CategoryAPI.md#CategoriesPost) | **Post** /categories | Store a newly category.



## CategoriesCategoryDelete

> CategoriesCategoryDelete(ctx, category).Execute()

Remove the specified category.

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
    category := "category_example" // string | The Id of category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CategoryAPI.CategoriesCategoryDelete(context.Background(), category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoriesCategoryDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | The Id of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryDeleteRequest struct via the builder pattern


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


## CategoriesCategoryGet

> CategoriesCategoryGet(ctx, category).Execute()

Return the specified category.

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
    category := "category_example" // string | The Id of category

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CategoryAPI.CategoriesCategoryGet(context.Background(), category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoriesCategoryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | The Id of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryGetRequest struct via the builder pattern


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


## CategoriesCategoryPut

> CategoriesCategoryPut(ctx, category).Body(body).Execute()

Update the specified category.

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
    category := "category_example" // string | The Id of category
    body := *openapiclient.NewCategoriesCategoryPutRequest() // CategoriesCategoryPutRequest | Category details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CategoryAPI.CategoriesCategoryPut(context.Background(), category).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoriesCategoryPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**category** | **string** | The Id of category | 

### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesCategoryPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CategoriesCategoryPutRequest**](CategoriesCategoryPutRequest.md) | Category details | 

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


## CategoriesGet

> CategoriesGet(ctx).Page(page).PerPage(perPage).Execute()

Return all user categories.

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
    page := int32(56) // int32 | Page number (optional)
    perPage := int32(56) // int32 | Page limit (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CategoryAPI.CategoriesGet(context.Background()).Page(page).PerPage(perPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## CategoriesPost

> CategoriesPost(ctx).Body(body).Execute()

Store a newly category.

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
    body := *openapiclient.NewCategoriesPostRequest("Title_example") // CategoriesPostRequest | Category details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CategoryAPI.CategoriesPost(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoryAPI.CategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CategoriesPostRequest**](CategoriesPostRequest.md) | Category details | 

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

