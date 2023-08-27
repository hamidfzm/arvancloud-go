# \ListAPI

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListsDestroy**](ListAPI.md#ListsDestroy) | **Delete** /dynamic-fields/{id} | Delete List
[**ListsIndex**](ListAPI.md#ListsIndex) | **Get** /dynamic-fields | Get the list of Lists
[**ListsShow**](ListAPI.md#ListsShow) | **Get** /dynamic-fields/{id} | Get an existing List
[**ListsStore**](ListAPI.md#ListsStore) | **Post** /dynamic-fields | Store new List
[**ListsUpdate**](ListAPI.md#ListsUpdate) | **Patch** /dynamic-fields/{id} | Update an existing List



## ListsDestroy

> MessageResponse ListsDestroy(ctx, id).Execute()

Delete List

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListAPI.ListsDestroy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsDestroy`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListsDestroy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsIndex

> ListsIndex200Response ListsIndex(ctx).PerPage(perPage).Page(page).Scope(scope).Type_(type_).Name(name).Execute()

Get the list of Lists

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
    perPage := int32(56) // int32 | Set how many items returned per page (optional)
    page := int32(56) // int32 | Set the desired page number (optional) (default to 1)
    scope := "scope_example" // string |  (optional)
    type_ := "type__example" // string |  (optional)
    name := "name_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListAPI.ListsIndex(context.Background()).PerPage(perPage).Page(page).Scope(scope).Type_(type_).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsIndex`: ListsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListsIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **perPage** | **int32** | Set how many items returned per page | 
 **page** | **int32** | Set the desired page number | [default to 1]
 **scope** | **string** |  | 
 **type_** | **string** |  | 
 **name** | **string** |  | 

### Return type

[**ListsIndex200Response**](ListsIndex200Response.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsShow

> DynamicFieldData ListsShow(ctx, id).Execute()

Get an existing List

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListAPI.ListsShow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsShow`: DynamicFieldData
    fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DynamicFieldData**](DynamicFieldData.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsStore

> DynamicFieldResponse ListsStore(ctx).DynamicField(dynamicField).Execute()

Store new List

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
    dynamicField := *openapiclient.NewDynamicField("Name_example", "Type_example", []openapiclient.DynamicFieldValue{*openapiclient.NewDynamicFieldValue()}) // DynamicField |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListAPI.ListsStore(context.Background()).DynamicField(dynamicField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListsStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsStore`: DynamicFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListsStore`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListsStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dynamicField** | [**DynamicField**](DynamicField.md) |  | 

### Return type

[**DynamicFieldResponse**](DynamicFieldResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListsUpdate

> DynamicFieldResponse ListsUpdate(ctx, id).DynamicField(dynamicField).Execute()

Update an existing List

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    dynamicField := *openapiclient.NewDynamicField("Name_example", "Type_example", []openapiclient.DynamicFieldValue{*openapiclient.NewDynamicFieldValue()}) // DynamicField |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListAPI.ListsUpdate(context.Background(), id).DynamicField(dynamicField).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListAPI.ListsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListsUpdate`: DynamicFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `ListAPI.ListsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicField** | [**DynamicField**](DynamicField.md) |  | 

### Return type

[**DynamicFieldResponse**](DynamicFieldResponse.md)

### Authorization

[ApiKey](../README.md#ApiKey), [UserToken](../README.md#UserToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

