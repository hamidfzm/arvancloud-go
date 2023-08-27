# \ObjectACLAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetObjectAcl**](ObjectACLAPI.md#GetObjectAcl) | **Get** /{Bucket}/{Key}#acl | 
[**PutObjectAcl**](ObjectACLAPI.md#PutObjectAcl) | **Put** /{Bucket}/{Key}#acl | 



## GetObjectAcl

> GetObjectAclOutput GetObjectAcl(ctx, bucket, key).Acl(acl).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name that contains the object for which to get the ACL information. </p> <p>
    key := "key_example" // string | The key of the object for which to get the ACL information.
    acl := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    versionId := "versionId_example" // string | VersionId used to reference a specific version of the object. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectACLAPI.GetObjectAcl(context.Background(), bucket, key).Acl(acl).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectACLAPI.GetObjectAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectAcl`: GetObjectAclOutput
    fmt.Fprintf(os.Stdout, "Response from `ObjectACLAPI.GetObjectAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name that contains the object for which to get the ACL information. &lt;/p&gt; &lt;p&gt; | 
**key** | **string** | The key of the object for which to get the ACL information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acl** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **versionId** | **string** | VersionId used to reference a specific version of the object. | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetObjectAclOutput**](GetObjectAclOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObjectAcl

> map[string]interface{} PutObjectAcl(ctx, bucket, key).Acl(acl).PutBucketAclRequest(putBucketAclRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).ContentMD5(contentMD5).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWrite(xAmzGrantWrite).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzRequestPayer(xAmzRequestPayer).VersionId(versionId).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name that contains the object to which you want to attach the ACL. </p> <p>
    key := "key_example" // string | <p>Key for which the PUT action was initiated.</p> <p>W
    acl := true // bool | 
    putBucketAclRequest := *openapiclient.NewPutBucketAclRequest() // PutBucketAclRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzAcl := "xAmzAcl_example" // string | The canned ACL to apply to the object. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL\">Canned ACL</a>. (optional)
    contentMD5 := "contentMD5_example" // string | <p>The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to <a href=\"http://www.ietf.org/rfc/rfc1864.txt\">RFC 1864.&gt;</a> </p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzGrantFullControl := "xAmzGrantFullControl_example" // string | <p>Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.</p> </p> (optional)
    xAmzGrantRead := "xAmzGrantRead_example" // string | <p>Allows grantee to list the objects in the bucket.</p> </p> (optional)
    xAmzGrantReadAcp := "xAmzGrantReadAcp_example" // string | <p>Allows grantee to read the bucket ACL.</p> </p> (optional)
    xAmzGrantWrite := "xAmzGrantWrite_example" // string | <p>Allows grantee to create new objects in the bucket.</p> <p>For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.</p> (optional)
    xAmzGrantWriteAcp := "xAmzGrantWriteAcp_example" // string | <p>Allows grantee to write the ACL for the applicable bucket.</p> </p> (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    versionId := "versionId_example" // string | VersionId used to reference a specific version of the object. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectACLAPI.PutObjectAcl(context.Background(), bucket, key).Acl(acl).PutBucketAclRequest(putBucketAclRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).ContentMD5(contentMD5).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWrite(xAmzGrantWrite).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzRequestPayer(xAmzRequestPayer).VersionId(versionId).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectACLAPI.PutObjectAcl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutObjectAcl`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectACLAPI.PutObjectAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name that contains the object to which you want to attach the ACL. &lt;/p&gt; &lt;p&gt; | 
**key** | **string** | &lt;p&gt;Key for which the PUT action was initiated.&lt;/p&gt; &lt;p&gt;W | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **acl** | **bool** |  | 
 **putBucketAclRequest** | [**PutBucketAclRequest**](PutBucketAclRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzAcl** | **string** | The canned ACL to apply to the object. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL\&quot;&gt;Canned ACL&lt;/a&gt;. | 
 **contentMD5** | **string** | &lt;p&gt;The base64-encoded 128-bit MD5 digest of the data. This header must be used as a message integrity check to verify that the request body was not corrupted in transit. For more information, go to &lt;a href&#x3D;\&quot;http://www.ietf.org/rfc/rfc1864.txt\&quot;&gt;RFC 1864.&amp;gt;&lt;/a&gt; &lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
 **xAmzGrantFullControl** | **string** | &lt;p&gt;Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantRead** | **string** | &lt;p&gt;Allows grantee to list the objects in the bucket.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantReadAcp** | **string** | &lt;p&gt;Allows grantee to read the bucket ACL.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantWrite** | **string** | &lt;p&gt;Allows grantee to create new objects in the bucket.&lt;/p&gt; &lt;p&gt;For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.&lt;/p&gt; | 
 **xAmzGrantWriteAcp** | **string** | &lt;p&gt;Allows grantee to write the ACL for the applicable bucket.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzRequestPayer** | **string** |  | 
 **versionId** | **string** | VersionId used to reference a specific version of the object. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

**map[string]interface{}**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

