# \BucketLifecycleConfigurationAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBucketLifecycle**](BucketLifecycleConfigurationAPI.md#DeleteBucketLifecycle) | **Delete** /{Bucket}#lifecycle | 
[**GetBucketLifecycleConfiguration**](BucketLifecycleConfigurationAPI.md#GetBucketLifecycleConfiguration) | **Get** /{Bucket}#lifecycle | 
[**PutBucketLifecycleConfiguration**](BucketLifecycleConfigurationAPI.md#PutBucketLifecycleConfiguration) | **Put** /{Bucket}#lifecycle | 



## DeleteBucketLifecycle

> DeleteBucketLifecycle(ctx, bucket).Lifecycle(lifecycle).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func main() {
    bucket := "bucket_example" // string | The bucket name of the lifecycle to delete.
    lifecycle := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketLifecycleConfigurationAPI.DeleteBucketLifecycle(context.Background(), bucket).Lifecycle(lifecycle).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketLifecycleConfigurationAPI.DeleteBucketLifecycle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket name of the lifecycle to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketLifecycleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycle** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

 (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBucketLifecycleConfiguration

> GetBucketLifecycleConfigurationOutput GetBucketLifecycleConfiguration(ctx, bucket).Lifecycle(lifecycle).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func main() {
    bucket := "bucket_example" // string | The name of the bucket for which to get the lifecycle information.
    lifecycle := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketLifecycleConfigurationAPI.GetBucketLifecycleConfiguration(context.Background(), bucket).Lifecycle(lifecycle).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketLifecycleConfigurationAPI.GetBucketLifecycleConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketLifecycleConfiguration`: GetBucketLifecycleConfigurationOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketLifecycleConfigurationAPI.GetBucketLifecycleConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The name of the bucket for which to get the lifecycle information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketLifecycleConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycle** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetBucketLifecycleConfigurationOutput**](GetBucketLifecycleConfigurationOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketLifecycleConfiguration

> PutBucketLifecycleConfiguration(ctx, bucket).Lifecycle(lifecycle).PutBucketLifecycleConfigurationRequest(putBucketLifecycleConfigurationRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func main() {
    bucket := "bucket_example" // string | The name of the bucket for which to set the configuration.
    lifecycle := true // bool | 
    putBucketLifecycleConfigurationRequest := *openapiclient.NewPutBucketLifecycleConfigurationRequest() // PutBucketLifecycleConfigurationRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketLifecycleConfigurationAPI.PutBucketLifecycleConfiguration(context.Background(), bucket).Lifecycle(lifecycle).PutBucketLifecycleConfigurationRequest(putBucketLifecycleConfigurationRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketLifecycleConfigurationAPI.PutBucketLifecycleConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The name of the bucket for which to set the configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketLifecycleConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lifecycle** | **bool** |  | 
 **putBucketLifecycleConfigurationRequest** | [**PutBucketLifecycleConfigurationRequest**](PutBucketLifecycleConfigurationRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

 (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

