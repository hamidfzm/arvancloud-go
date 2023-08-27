# \BucketCORSAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteBucketCors**](BucketCORSAPI.md#DeleteBucketCors) | **Delete** /{Bucket}#cors | 
[**GetBucketCors**](BucketCORSAPI.md#GetBucketCors) | **Get** /{Bucket}#cors | 
[**PutBucketCors**](BucketCORSAPI.md#PutBucketCors) | **Put** /{Bucket}#cors | 



## DeleteBucketCors

> DeleteBucketCors(ctx, bucket).Cors(cors).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | Specifies the bucket whose <code>cors</code> configuration is being deleted.
    cors := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketCORSAPI.DeleteBucketCors(context.Background(), bucket).Cors(cors).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketCORSAPI.DeleteBucketCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Specifies the bucket whose &lt;code&gt;cors&lt;/code&gt; configuration is being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cors** | **bool** |  | 
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


## GetBucketCors

> GetBucketCorsOutput GetBucketCors(ctx, bucket).Cors(cors).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The bucket name for which to get the cors configuration.
    cors := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketCORSAPI.GetBucketCors(context.Background(), bucket).Cors(cors).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketCORSAPI.GetBucketCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketCors`: GetBucketCorsOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketCORSAPI.GetBucketCors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket name for which to get the cors configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cors** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetBucketCorsOutput**](GetBucketCorsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketCors

> PutBucketCors(ctx, bucket).Cors(cors).PutBucketCorsRequest(putBucketCorsRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | Specifies the bucket impacted by the <code>cors</code>configuration.
    cors := true // bool | 
    putBucketCorsRequest := *openapiclient.NewPutBucketCorsRequest(*openapiclient.NewPutBucketCorsRequestCORSConfiguration("TODO")) // PutBucketCorsRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    contentMD5 := "contentMD5_example" // string | <p>The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to <a href=\"http://www.ietf.org/rfc/rfc1864.txt\">RFC 1864.</a> </p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketCORSAPI.PutBucketCors(context.Background(), bucket).Cors(cors).PutBucketCorsRequest(putBucketCorsRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketCORSAPI.PutBucketCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Specifies the bucket impacted by the &lt;code&gt;cors&lt;/code&gt;configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cors** | **bool** |  | 
 **putBucketCorsRequest** | [**PutBucketCorsRequest**](PutBucketCorsRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **contentMD5** | **string** | &lt;p&gt;The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864.&lt;/a&gt; &lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
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

