# \BucketWebsiteConfigurationAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBucketWebsite**](BucketWebsiteConfigurationAPI.md#DeleteBucketWebsite) | **Delete** /{Bucket}#website | 
[**GetBucketWebsite**](BucketWebsiteConfigurationAPI.md#GetBucketWebsite) | **Get** /{Bucket}#website | 
[**PutBucketWebsite**](BucketWebsiteConfigurationAPI.md#PutBucketWebsite) | **Put** /{Bucket}#website | 



## DeleteBucketWebsite

> DeleteBucketWebsite(ctx, bucket).Website(website).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The bucket name for which you want to remove the website configuration. 
    website := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketWebsiteConfigurationAPI.DeleteBucketWebsite(context.Background(), bucket).Website(website).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketWebsiteConfigurationAPI.DeleteBucketWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket name for which you want to remove the website configuration.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **website** | **bool** |  | 
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


## GetBucketWebsite

> GetBucketWebsiteOutput GetBucketWebsite(ctx, bucket).Website(website).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The bucket name for which to get the website configuration.
    website := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketWebsiteConfigurationAPI.GetBucketWebsite(context.Background(), bucket).Website(website).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketWebsiteConfigurationAPI.GetBucketWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketWebsite`: GetBucketWebsiteOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketWebsiteConfigurationAPI.GetBucketWebsite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket name for which to get the website configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **website** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetBucketWebsiteOutput**](GetBucketWebsiteOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketWebsite

> PutBucketWebsite(ctx, bucket).Website(website).PutBucketWebsiteRequest(putBucketWebsiteRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The bucket name.
    website := true // bool | 
    putBucketWebsiteRequest := *openapiclient.NewPutBucketWebsiteRequest(*openapiclient.NewPutBucketWebsiteRequestWebsiteConfiguration()) // PutBucketWebsiteRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    contentMD5 := "contentMD5_example" // string | <p>The base64-encoded 128-bit MD5 digest of the data. You must use this header as a message integrity check to verify that the request body was not corrupted in transit. For more information, see <a href=\"http://www.ietf.org/rfc/rfc1864.txt\">RFC 1864</a>.</p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketWebsiteConfigurationAPI.PutBucketWebsite(context.Background(), bucket).Website(website).PutBucketWebsiteRequest(putBucketWebsiteRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketWebsiteConfigurationAPI.PutBucketWebsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketWebsiteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **website** | **bool** |  | 
 **putBucketWebsiteRequest** | [**PutBucketWebsiteRequest**](PutBucketWebsiteRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **contentMD5** | **string** | &lt;p&gt;The base64-encoded 128-bit MD5 digest of the data. You must use this header as a message integrity check to verify that the request body was not corrupted in transit. For more information, see &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864&lt;/a&gt;.&lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
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

