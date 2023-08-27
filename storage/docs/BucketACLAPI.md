# \BucketACLAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBucketAcl**](BucketACLAPI.md#GetBucketAcl) | **Get** /{Bucket}#acl | 
[**PutBucketAcl**](BucketACLAPI.md#PutBucketAcl) | **Put** /{Bucket}#acl | 



## GetBucketAcl

> GetBucketAclOutput GetBucketAcl(ctx, bucket).Acl(acl).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | Specifies the S3 bucket whose ACL is being requested.
    acl := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketACLAPI.GetBucketAcl(context.Background(), bucket).Acl(acl).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketACLAPI.GetBucketAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBucketAcl`: GetBucketAclOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketACLAPI.GetBucketAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Specifies the S3 bucket whose ACL is being requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBucketAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acl** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetBucketAclOutput**](GetBucketAclOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutBucketAcl

> PutBucketAcl(ctx, bucket).Acl(acl).PutBucketAclRequest(putBucketAclRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).ContentMD5(contentMD5).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWrite(xAmzGrantWrite).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The bucket to which to apply the ACL.
    acl := true // bool | 
    putBucketAclRequest := *openapiclient.NewPutBucketAclRequest() // PutBucketAclRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzAcl := "xAmzAcl_example" // string | The canned ACL to apply to the bucket. (optional)
    contentMD5 := "contentMD5_example" // string | <p>The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to <a href=\"http://www.ietf.org/rfc/rfc1864.txt\">RFC 1864.</a> </p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzGrantFullControl := "xAmzGrantFullControl_example" // string | Allows grantee the read, write, read ACP, and write ACP permissions on the bucket. (optional)
    xAmzGrantRead := "xAmzGrantRead_example" // string | Allows grantee to list the objects in the bucket. (optional)
    xAmzGrantReadAcp := "xAmzGrantReadAcp_example" // string | Allows grantee to read the bucket ACL. (optional)
    xAmzGrantWrite := "xAmzGrantWrite_example" // string | <p>Allows grantee to create new objects in the bucket.</p> <p>For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.</p> (optional)
    xAmzGrantWriteAcp := "xAmzGrantWriteAcp_example" // string | Allows grantee to write the ACL for the applicable bucket. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketACLAPI.PutBucketAcl(context.Background(), bucket).Acl(acl).PutBucketAclRequest(putBucketAclRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).ContentMD5(contentMD5).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWrite(xAmzGrantWrite).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketACLAPI.PutBucketAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket to which to apply the ACL. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutBucketAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acl** | **bool** |  | 
 **putBucketAclRequest** | [**PutBucketAclRequest**](PutBucketAclRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzAcl** | **string** | The canned ACL to apply to the bucket. | 
 **contentMD5** | **string** | &lt;p&gt;The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864.&lt;/a&gt; &lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
 **xAmzGrantFullControl** | **string** | Allows grantee the read, write, read ACP, and write ACP permissions on the bucket. | 
 **xAmzGrantRead** | **string** | Allows grantee to list the objects in the bucket. | 
 **xAmzGrantReadAcp** | **string** | Allows grantee to read the bucket ACL. | 
 **xAmzGrantWrite** | **string** | &lt;p&gt;Allows grantee to create new objects in the bucket.&lt;/p&gt; &lt;p&gt;For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.&lt;/p&gt; | 
 **xAmzGrantWriteAcp** | **string** | Allows grantee to write the ACL for the applicable bucket. | 
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

