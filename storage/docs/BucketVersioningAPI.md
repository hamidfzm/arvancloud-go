# \BucketVersioningAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBucketVersioning**](BucketVersioningAPI.md#GetBucketVersioning) | **Get** /{Bucket}#versioning | 
[**PutBucketVersioning**](BucketVersioningAPI.md#PutBucketVersioning) | **Put** /{Bucket}#versioning | 



## GetBucketVersioning

> GetBucketVersioningOutput GetBucketVersioning(ctx, bucket).Versioning(versioning).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The name of the bucket for which to get the versioning information.
    versioning := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketVersioningAPI.GetBucketVersioning(context.Background(), bucket).Versioning(versioning).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketVersioningAPI.GetBucketVersioning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketVersioning`: GetBucketVersioningOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketVersioningAPI.GetBucketVersioning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The name of the bucket for which to get the versioning information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketVersioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versioning** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetBucketVersioningOutput**](GetBucketVersioningOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketVersioning

> PutBucketVersioning(ctx, bucket).Versioning(versioning).PutBucketVersioningRequest(putBucketVersioningRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzMfa(xAmzMfa).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    versioning := true // bool | 
    putBucketVersioningRequest := *openapiclient.NewPutBucketVersioningRequest(*openapiclient.NewPutBucketVersioningRequestVersioningConfiguration()) // PutBucketVersioningRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    contentMD5 := "contentMD5_example" // string | <p>&gt;The base64-encoded 128-bit MD5 digest of the data. You must use this header as a message integrity check to verify that the request body was not corrupted in transit. For more information, see <a href=\"http://www.ietf.org/rfc/rfc1864.txt\">RFC 1864</a>.</p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzMfa := "xAmzMfa_example" // string | The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketVersioningAPI.PutBucketVersioning(context.Background(), bucket).Versioning(versioning).PutBucketVersioningRequest(putBucketVersioningRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzMfa(xAmzMfa).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketVersioningAPI.PutBucketVersioning``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPutBucketVersioningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versioning** | **bool** |  | 
 **putBucketVersioningRequest** | [**PutBucketVersioningRequest**](PutBucketVersioningRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **contentMD5** | **string** | &lt;p&gt;&amp;gt;The base64-encoded 128-bit MD5 digest of the data. You must use this header as a message integrity check to verify that the request body was not corrupted in transit. For more information, see &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864&lt;/a&gt;.&lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
 **xAmzMfa** | **string** | The concatenation of the authentication device&#39;s serial number, a space, and the value that is displayed on your authentication device. | 
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

