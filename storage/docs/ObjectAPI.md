# \ObjectAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObject**](ObjectAPI.md#DeleteObject) | **Delete** /{Bucket}/{Key} | 
[**GetObject**](ObjectAPI.md#GetObject) | **Get** /{Bucket}/{Key} | 
[**HeadObject**](ObjectAPI.md#HeadObject) | **Head** /{Bucket}/{Key} | 
[**PutObject**](ObjectAPI.md#PutObject) | **Put** /{Bucket}/{Key} | 



## DeleteObject

> map[string]interface{} DeleteObject(ctx, bucket, key).XAmzSecurityToken(xAmzSecurityToken).XAmzMfa(xAmzMfa).VersionId(versionId).XAmzRequestPayer(xAmzRequestPayer).XAmzBypassGovernanceRetention(xAmzBypassGovernanceRetention).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name of the bucket containing the object. </p>
    key := "key_example" // string | Key name of the object to delete.
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzMfa := "xAmzMfa_example" // string | The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA delete enabled. (optional)
    versionId := "versionId_example" // string | VersionId used to reference a specific version of the object. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzBypassGovernanceRetention := true // bool | Indicates whether S3 Object Lock should bypass Governance-mode restrictions to process this operation. To use this header, you must have the <code>s3:PutBucketPublicAccessBlock</code> permission. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectAPI.DeleteObject(context.Background(), bucket, key).XAmzSecurityToken(xAmzSecurityToken).XAmzMfa(xAmzMfa).VersionId(versionId).XAmzRequestPayer(xAmzRequestPayer).XAmzBypassGovernanceRetention(xAmzBypassGovernanceRetention).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.DeleteObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.DeleteObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name of the bucket containing the object. &lt;/p&gt; | 
**key** | **string** | Key name of the object to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzMfa** | **string** | The concatenation of the authentication device&#39;s serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA delete enabled. | 
 **versionId** | **string** | VersionId used to reference a specific version of the object. | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzBypassGovernanceRetention** | **bool** | Indicates whether S3 Object Lock should bypass Governance-mode restrictions to process this operation. To use this header, you must have the &lt;code&gt;s3:PutBucketPublicAccessBlock&lt;/code&gt; permission. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

**map[string]interface{}**

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObject

> GetObjectOutput GetObject(ctx, bucket, key).XAmzSecurityToken(xAmzSecurityToken).IfMatch(ifMatch).IfModifiedSince(ifModifiedSince).IfNoneMatch(ifNoneMatch).IfUnmodifiedSince(ifUnmodifiedSince).Range_(range_).ResponseCacheControl(responseCacheControl).ResponseContentDisposition(responseContentDisposition).ResponseContentEncoding(responseContentEncoding).ResponseContentLanguage(responseContentLanguage).ResponseContentType(responseContentType).ResponseExpires(responseExpires).VersionId(versionId).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzRequestPayer(xAmzRequestPayer).PartNumber(partNumber).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func main() {
    bucket := "bucket_example" // string | <p>The bucket name containing the object. </p>
    key := "key_example" // string | Key of the object to get.
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    ifMatch := "ifMatch_example" // string | Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412 (precondition failed). (optional)
    ifModifiedSince := time.Now() // time.Time | Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified). (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not modified). (optional)
    ifUnmodifiedSince := time.Now() // time.Time | Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition failed). (optional)
    range_ := "range__example" // string | <p>Downloads the specified range bytes of an object. For more information about the HTTP Range header, see <a href=\"https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35\">https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35</a>.</p> <note> <p>ArvanCloud S3 doesn't support retrieving multiple ranges of data per <code>GET</code> request.</p> </note> (optional)
    responseCacheControl := "responseCacheControl_example" // string | Sets the <code>Cache-Control</code> header of the response. (optional)
    responseContentDisposition := "responseContentDisposition_example" // string | Sets the <code>Content-Disposition</code> header of the response (optional)
    responseContentEncoding := "responseContentEncoding_example" // string | Sets the <code>Content-Encoding</code> header of the response. (optional)
    responseContentLanguage := "responseContentLanguage_example" // string | Sets the <code>Content-Language</code> header of the response. (optional)
    responseContentType := "responseContentType_example" // string | Sets the <code>Content-Type</code> header of the response. (optional)
    responseExpires := time.Now() // time.Time | Sets the <code>Expires</code> header of the response. (optional)
    versionId := "versionId_example" // string | VersionId used to reference a specific version of the object. (optional)
    xAmzServerSideEncryptionCustomerAlgorithm := "xAmzServerSideEncryptionCustomerAlgorithm_example" // string | Specifies the algorithm to use to when decrypting the object (for example, AES256). (optional)
    xAmzServerSideEncryptionCustomerKey := "xAmzServerSideEncryptionCustomerKey_example" // string | Specifies the customer-provided encryption key for ArvanCloud S3 used to encrypt the data. This value is used to decrypt the object when recovering it and must match the one used when storing the data. The key must be appropriate for use with the algorithm specified in the <code>x-amz-server-side-encryption-customer-algorithm</code> header. (optional)
    xAmzServerSideEncryptionCustomerKeyMD5 := "xAmzServerSideEncryptionCustomerKeyMD5_example" // string | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    partNumber := int32(56) // int32 | Part number of the object being read. This is a positive integer between 1 and 10,000. Effectively performs a 'ranged' GET request for the part specified. Useful for downloading just a part of an object. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectAPI.GetObject(context.Background(), bucket, key).XAmzSecurityToken(xAmzSecurityToken).IfMatch(ifMatch).IfModifiedSince(ifModifiedSince).IfNoneMatch(ifNoneMatch).IfUnmodifiedSince(ifUnmodifiedSince).Range_(range_).ResponseCacheControl(responseCacheControl).ResponseContentDisposition(responseContentDisposition).ResponseContentEncoding(responseContentEncoding).ResponseContentLanguage(responseContentLanguage).ResponseContentType(responseContentType).ResponseExpires(responseExpires).VersionId(versionId).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzRequestPayer(xAmzRequestPayer).PartNumber(partNumber).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.GetObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObject`: GetObjectOutput
    fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.GetObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name containing the object. &lt;/p&gt; | 
**key** | **string** | Key of the object to get. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **ifMatch** | **string** | Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412 (precondition failed). | 
 **ifModifiedSince** | **time.Time** | Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified). | 
 **ifNoneMatch** | **string** | Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not modified). | 
 **ifUnmodifiedSince** | **time.Time** | Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition failed). | 
 **range_** | **string** | &lt;p&gt;Downloads the specified range bytes of an object. For more information about the HTTP Range header, see &lt;a href&#x3D;\&quot;https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35\&quot;&gt;https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35&lt;/a&gt;.&lt;/p&gt; &lt;note&gt; &lt;p&gt;ArvanCloud S3 doesn&#39;t support retrieving multiple ranges of data per &lt;code&gt;GET&lt;/code&gt; request.&lt;/p&gt; &lt;/note&gt; | 
 **responseCacheControl** | **string** | Sets the &lt;code&gt;Cache-Control&lt;/code&gt; header of the response. | 
 **responseContentDisposition** | **string** | Sets the &lt;code&gt;Content-Disposition&lt;/code&gt; header of the response | 
 **responseContentEncoding** | **string** | Sets the &lt;code&gt;Content-Encoding&lt;/code&gt; header of the response. | 
 **responseContentLanguage** | **string** | Sets the &lt;code&gt;Content-Language&lt;/code&gt; header of the response. | 
 **responseContentType** | **string** | Sets the &lt;code&gt;Content-Type&lt;/code&gt; header of the response. | 
 **responseExpires** | **time.Time** | Sets the &lt;code&gt;Expires&lt;/code&gt; header of the response. | 
 **versionId** | **string** | VersionId used to reference a specific version of the object. | 
 **xAmzServerSideEncryptionCustomerAlgorithm** | **string** | Specifies the algorithm to use to when decrypting the object (for example, AES256). | 
 **xAmzServerSideEncryptionCustomerKey** | **string** | Specifies the customer-provided encryption key for ArvanCloud S3 used to encrypt the data. This value is used to decrypt the object when recovering it and must match the one used when storing the data. The key must be appropriate for use with the algorithm specified in the &lt;code&gt;x-amz-server-side-encryption-customer-algorithm&lt;/code&gt; header. | 
 **xAmzServerSideEncryptionCustomerKeyMD5** | **string** | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. | 
 **xAmzRequestPayer** | **string** |  | 
 **partNumber** | **int32** | Part number of the object being read. This is a positive integer between 1 and 10,000. Effectively performs a &#39;ranged&#39; GET request for the part specified. Useful for downloading just a part of an object. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetObjectOutput**](GetObjectOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadObject

> HeadObjectOutput HeadObject(ctx, bucket, key).XAmzSecurityToken(xAmzSecurityToken).IfMatch(ifMatch).IfModifiedSince(ifModifiedSince).IfNoneMatch(ifNoneMatch).IfUnmodifiedSince(ifUnmodifiedSince).Range_(range_).VersionId(versionId).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzRequestPayer(xAmzRequestPayer).PartNumber(partNumber).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func main() {
    bucket := "bucket_example" // string | <p>The name of the bucket containing the object.</p>
    key := "key_example" // string | The object key.
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    ifMatch := "ifMatch_example" // string | Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412 (precondition failed). (optional)
    ifModifiedSince := time.Now() // time.Time | Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified). (optional)
    ifNoneMatch := "ifNoneMatch_example" // string | Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not modified). (optional)
    ifUnmodifiedSince := time.Now() // time.Time | Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition failed). (optional)
    range_ := "range__example" // string | <p>Downloads the specified range bytes of an object. For more information about the HTTP Range header, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35</a>.</p> <note> <p>ArvanCloud S3 doesn't support retrieving multiple ranges of data per <code>GET</code> request.</p> </note> (optional)
    versionId := "versionId_example" // string | VersionId used to reference a specific version of the object. (optional)
    xAmzServerSideEncryptionCustomerAlgorithm := "xAmzServerSideEncryptionCustomerAlgorithm_example" // string | Specifies the algorithm to use to when encrypting the object (for example, AES256). (optional)
    xAmzServerSideEncryptionCustomerKey := "xAmzServerSideEncryptionCustomerKey_example" // string | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the <code>x-amz-server-side-encryption-customer-algorithm</code> header. (optional)
    xAmzServerSideEncryptionCustomerKeyMD5 := "xAmzServerSideEncryptionCustomerKeyMD5_example" // string | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    partNumber := int32(56) // int32 | Part number of the object being read. This is a positive integer between 1 and 10,000. Effectively performs a 'ranged' HEAD request for the part specified. Useful querying about the size of the part and the number of parts in this object. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectAPI.HeadObject(context.Background(), bucket, key).XAmzSecurityToken(xAmzSecurityToken).IfMatch(ifMatch).IfModifiedSince(ifModifiedSince).IfNoneMatch(ifNoneMatch).IfUnmodifiedSince(ifUnmodifiedSince).Range_(range_).VersionId(versionId).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzRequestPayer(xAmzRequestPayer).PartNumber(partNumber).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.HeadObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HeadObject`: HeadObjectOutput
    fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.HeadObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The name of the bucket containing the object.&lt;/p&gt; | 
**key** | **string** | The object key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **ifMatch** | **string** | Return the object only if its entity tag (ETag) is the same as the one specified, otherwise return a 412 (precondition failed). | 
 **ifModifiedSince** | **time.Time** | Return the object only if it has been modified since the specified time, otherwise return a 304 (not modified). | 
 **ifNoneMatch** | **string** | Return the object only if its entity tag (ETag) is different from the one specified, otherwise return a 304 (not modified). | 
 **ifUnmodifiedSince** | **time.Time** | Return the object only if it has not been modified since the specified time, otherwise return a 412 (precondition failed). | 
 **range_** | **string** | &lt;p&gt;Downloads the specified range bytes of an object. For more information about the HTTP Range header, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35&lt;/a&gt;.&lt;/p&gt; &lt;note&gt; &lt;p&gt;ArvanCloud S3 doesn&#39;t support retrieving multiple ranges of data per &lt;code&gt;GET&lt;/code&gt; request.&lt;/p&gt; &lt;/note&gt; | 
 **versionId** | **string** | VersionId used to reference a specific version of the object. | 
 **xAmzServerSideEncryptionCustomerAlgorithm** | **string** | Specifies the algorithm to use to when encrypting the object (for example, AES256). | 
 **xAmzServerSideEncryptionCustomerKey** | **string** | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the &lt;code&gt;x-amz-server-side-encryption-customer-algorithm&lt;/code&gt; header. | 
 **xAmzServerSideEncryptionCustomerKeyMD5** | **string** | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. | 
 **xAmzRequestPayer** | **string** |  | 
 **partNumber** | **int32** | Part number of the object being read. This is a positive integer between 1 and 10,000. Effectively performs a &#39;ranged&#39; HEAD request for the part specified. Useful querying about the size of the part and the number of parts in this object. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**HeadObjectOutput**](HeadObjectOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObject

> map[string]interface{} PutObject(ctx, bucket, key).PutObjectRequest(putObjectRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).CacheControl(cacheControl).ContentDisposition(contentDisposition).ContentEncoding(contentEncoding).ContentLanguage(contentLanguage).ContentLength(contentLength).ContentMD5(contentMD5).ContentType(contentType).Expires(expires).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzServerSideEncryption(xAmzServerSideEncryption).XAmzStorageClass(xAmzStorageClass).XAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzServerSideEncryptionAwsKmsKeyId(xAmzServerSideEncryptionAwsKmsKeyId).XAmzServerSideEncryptionContext(xAmzServerSideEncryptionContext).XAmzServerSideEncryptionBucketKeyEnabled(xAmzServerSideEncryptionBucketKeyEnabled).XAmzRequestPayer(xAmzRequestPayer).XAmzTagging(xAmzTagging).XAmzObjectLockMode(xAmzObjectLockMode).XAmzObjectLockRetainUntilDate(xAmzObjectLockRetainUntilDate).XAmzObjectLockLegalHold(xAmzObjectLockLegalHold).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/hamidfzm/arvancloud-go/storage"
)

func main() {
    bucket := "bucket_example" // string | <p>The bucket name to which the PUT action was initiated. </p> <p>
    key := "key_example" // string | Object key for which the PUT action was initiated.
    putObjectRequest := *openapiclient.NewPutObjectRequest() // PutObjectRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzAcl := "xAmzAcl_example" // string | <p>The canned ACL to apply to the object. For more information, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL\">Canned ACL</a>.</p> </p> (optional)
    cacheControl := "cacheControl_example" // string |  Can be used to specify caching behavior along the request/reply chain. For more information, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9</a>. (optional)
    contentDisposition := "contentDisposition_example" // string | Specifies presentational information for the object. For more information, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1</a>. (optional)
    contentEncoding := "contentEncoding_example" // string | Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. For more information, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11</a>. (optional)
    contentLanguage := "contentLanguage_example" // string | The language the content is in. (optional)
    contentLength := int32(56) // int32 | Size of the body in bytes. This parameter is useful when the size of the body cannot be determined automatically. For more information, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13</a>. (optional)
    contentMD5 := "contentMD5_example" // string | The base64-encoded 128-bit MD5 digest of the message (without the headers) according to RFC 1864. This header can be used as a message integrity check to verify that the data is the same data that was originally sent. Although it is optional, we recommend using the Content-MD5 mechanism as an end-to-end integrity check. For more information about REST request authentication, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html\">REST Authentication</a>. (optional)
    contentType := "contentType_example" // string | A standard MIME type describing the format of the contents. For more information, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17</a>. (optional)
    expires := time.Now() // time.Time | The date and time at which the object is no longer cacheable. For more information, see <a href=\"http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21\">http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21</a>. (optional)
    xAmzGrantFullControl := "xAmzGrantFullControl_example" // string | <p>Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.</p> </p> (optional)
    xAmzGrantRead := "xAmzGrantRead_example" // string | <p>Allows grantee to read the object data and its metadata.</p> </p> (optional)
    xAmzGrantReadAcp := "xAmzGrantReadAcp_example" // string | <p>Allows grantee to read the object ACL.</p> </p> (optional)
    xAmzGrantWriteAcp := "xAmzGrantWriteAcp_example" // string | <p>Allows grantee to write the ACL for the applicable object.</p> </p> (optional)
    xAmzServerSideEncryption := "xAmzServerSideEncryption_example" // string | The server-side encryption algorithm used when storing this object in ArvanCloud S3 (for example, AES256, aws:kms). (optional)
    xAmzStorageClass := "xAmzStorageClass_example" // string | By default, ArvanCloud S3 uses the STANDARD Storage Class to store newly created objects. The STANDARD storage class provides high durability and high availability. Depending on performance needs, you can specify a different Storage Class. (optional)
    xAmzWebsiteRedirectLocation := "xAmzWebsiteRedirectLocation_example" // string | <p>If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. ArvanCloud S3 stores the value of this header in the object metadata. For information about object metadata, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html\">Object Key and Metadata</a>.</p> <p>In the following example, the request header sets the redirect to an object (anotherPage.html) in the same bucket:</p> <p> <code>x-amz-website-redirect-location: /anotherPage.html</code> </p> <p>In the following example, the request header sets the object redirect to another website:</p> <p> <code>x-amz-website-redirect-location: http://www.example.com/</code> </p> </p> (optional)
    xAmzServerSideEncryptionCustomerAlgorithm := "xAmzServerSideEncryptionCustomerAlgorithm_example" // string | Specifies the algorithm to use to when encrypting the object (for example, AES256). (optional)
    xAmzServerSideEncryptionCustomerKey := "xAmzServerSideEncryptionCustomerKey_example" // string | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the <code>x-amz-server-side-encryption-customer-algorithm</code> header. (optional)
    xAmzServerSideEncryptionCustomerKeyMD5 := "xAmzServerSideEncryptionCustomerKeyMD5_example" // string | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. (optional)
    xAmzServerSideEncryptionAwsKmsKeyId := "xAmzServerSideEncryptionAwsKmsKeyId_example" // string | If <code>x-amz-server-side-encryption</code> is present and has the value of <code>aws:kms</code>, this header specifies the ID of the Amazon Web Services Key Management Service (Amazon Web Services KMS) symmetrical customer managed key that was used for the object. If you specify <code>x-amz-server-side-encryption:aws:kms</code>, but do not provide<code> x-amz-server-side-encryption-aws-kms-key-id</code>, ArvanCloud S3 uses the Amazon Web Services managed key to protect the data. If the KMS key does not exist in the same account issuing the command, you must use the full ARN and not just the ID.  (optional)
    xAmzServerSideEncryptionContext := "xAmzServerSideEncryptionContext_example" // string | Specifies the Amazon Web Services KMS Encryption Context to use for object encryption. The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs. (optional)
    xAmzServerSideEncryptionBucketKeyEnabled := true // bool | <p>Specifies whether ArvanCloud S3 should use an S3 Bucket Key for object encryption with server-side encryption using AWS KMS (SSE-KMS). Setting this header to <code>true</code> causes ArvanCloud S3 to use an S3 Bucket Key for object encryption with SSE-KMS.</p> <p>Specifying this header with a PUT action doesn’t affect bucket-level settings for S3 Bucket Key.</p> (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzTagging := "xAmzTagging_example" // string | The tag-set for the object. The tag-set must be encoded as URL Query parameters. (For example, \"Key1=Value1\") (optional)
    xAmzObjectLockMode := "xAmzObjectLockMode_example" // string | The Object Lock mode that you want to apply to this object. (optional)
    xAmzObjectLockRetainUntilDate := time.Now() // time.Time | The date and time when you want this object's Object Lock to expire. Must be formatted as a timestamp parameter. (optional)
    xAmzObjectLockLegalHold := "xAmzObjectLockLegalHold_example" // string | Specifies whether a legal hold will be applied to this object. For more information about S3 Object Lock, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html\">Object Lock</a>. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectAPI.PutObject(context.Background(), bucket, key).PutObjectRequest(putObjectRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).CacheControl(cacheControl).ContentDisposition(contentDisposition).ContentEncoding(contentEncoding).ContentLanguage(contentLanguage).ContentLength(contentLength).ContentMD5(contentMD5).ContentType(contentType).Expires(expires).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzServerSideEncryption(xAmzServerSideEncryption).XAmzStorageClass(xAmzStorageClass).XAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzServerSideEncryptionAwsKmsKeyId(xAmzServerSideEncryptionAwsKmsKeyId).XAmzServerSideEncryptionContext(xAmzServerSideEncryptionContext).XAmzServerSideEncryptionBucketKeyEnabled(xAmzServerSideEncryptionBucketKeyEnabled).XAmzRequestPayer(xAmzRequestPayer).XAmzTagging(xAmzTagging).XAmzObjectLockMode(xAmzObjectLockMode).XAmzObjectLockRetainUntilDate(xAmzObjectLockRetainUntilDate).XAmzObjectLockLegalHold(xAmzObjectLockLegalHold).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectAPI.PutObject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutObject`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectAPI.PutObject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name to which the PUT action was initiated. &lt;/p&gt; &lt;p&gt; | 
**key** | **string** | Object key for which the PUT action was initiated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putObjectRequest** | [**PutObjectRequest**](PutObjectRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzAcl** | **string** | &lt;p&gt;The canned ACL to apply to the object. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL\&quot;&gt;Canned ACL&lt;/a&gt;.&lt;/p&gt; &lt;/p&gt; | 
 **cacheControl** | **string** |  Can be used to specify caching behavior along the request/reply chain. For more information, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.9&lt;/a&gt;. | 
 **contentDisposition** | **string** | Specifies presentational information for the object. For more information, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec19.html#sec19.5.1&lt;/a&gt;. | 
 **contentEncoding** | **string** | Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. For more information, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.11&lt;/a&gt;. | 
 **contentLanguage** | **string** | The language the content is in. | 
 **contentLength** | **int32** | Size of the body in bytes. This parameter is useful when the size of the body cannot be determined automatically. For more information, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.13&lt;/a&gt;. | 
 **contentMD5** | **string** | The base64-encoded 128-bit MD5 digest of the message (without the headers) according to RFC 1864. This header can be used as a message integrity check to verify that the data is the same data that was originally sent. Although it is optional, we recommend using the Content-MD5 mechanism as an end-to-end integrity check. For more information about REST request authentication, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html\&quot;&gt;REST Authentication&lt;/a&gt;. | 
 **contentType** | **string** | A standard MIME type describing the format of the contents. For more information, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.17&lt;/a&gt;. | 
 **expires** | **time.Time** | The date and time at which the object is no longer cacheable. For more information, see &lt;a href&#x3D;\&quot;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21\&quot;&gt;http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.21&lt;/a&gt;. | 
 **xAmzGrantFullControl** | **string** | &lt;p&gt;Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantRead** | **string** | &lt;p&gt;Allows grantee to read the object data and its metadata.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantReadAcp** | **string** | &lt;p&gt;Allows grantee to read the object ACL.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantWriteAcp** | **string** | &lt;p&gt;Allows grantee to write the ACL for the applicable object.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzServerSideEncryption** | **string** | The server-side encryption algorithm used when storing this object in ArvanCloud S3 (for example, AES256, aws:kms). | 
 **xAmzStorageClass** | **string** | By default, ArvanCloud S3 uses the STANDARD Storage Class to store newly created objects. The STANDARD storage class provides high durability and high availability. Depending on performance needs, you can specify a different Storage Class. | 
 **xAmzWebsiteRedirectLocation** | **string** | &lt;p&gt;If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. ArvanCloud S3 stores the value of this header in the object metadata. For information about object metadata, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMetadata.html\&quot;&gt;Object Key and Metadata&lt;/a&gt;.&lt;/p&gt; &lt;p&gt;In the following example, the request header sets the redirect to an object (anotherPage.html) in the same bucket:&lt;/p&gt; &lt;p&gt; &lt;code&gt;x-amz-website-redirect-location: /anotherPage.html&lt;/code&gt; &lt;/p&gt; &lt;p&gt;In the following example, the request header sets the object redirect to another website:&lt;/p&gt; &lt;p&gt; &lt;code&gt;x-amz-website-redirect-location: http://www.example.com/&lt;/code&gt; &lt;/p&gt; &lt;/p&gt; | 
 **xAmzServerSideEncryptionCustomerAlgorithm** | **string** | Specifies the algorithm to use to when encrypting the object (for example, AES256). | 
 **xAmzServerSideEncryptionCustomerKey** | **string** | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the &lt;code&gt;x-amz-server-side-encryption-customer-algorithm&lt;/code&gt; header. | 
 **xAmzServerSideEncryptionCustomerKeyMD5** | **string** | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. | 
 **xAmzServerSideEncryptionAwsKmsKeyId** | **string** | If &lt;code&gt;x-amz-server-side-encryption&lt;/code&gt; is present and has the value of &lt;code&gt;aws:kms&lt;/code&gt;, this header specifies the ID of the Amazon Web Services Key Management Service (Amazon Web Services KMS) symmetrical customer managed key that was used for the object. If you specify &lt;code&gt;x-amz-server-side-encryption:aws:kms&lt;/code&gt;, but do not provide&lt;code&gt; x-amz-server-side-encryption-aws-kms-key-id&lt;/code&gt;, ArvanCloud S3 uses the Amazon Web Services managed key to protect the data. If the KMS key does not exist in the same account issuing the command, you must use the full ARN and not just the ID.  | 
 **xAmzServerSideEncryptionContext** | **string** | Specifies the Amazon Web Services KMS Encryption Context to use for object encryption. The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs. | 
 **xAmzServerSideEncryptionBucketKeyEnabled** | **bool** | &lt;p&gt;Specifies whether ArvanCloud S3 should use an S3 Bucket Key for object encryption with server-side encryption using AWS KMS (SSE-KMS). Setting this header to &lt;code&gt;true&lt;/code&gt; causes ArvanCloud S3 to use an S3 Bucket Key for object encryption with SSE-KMS.&lt;/p&gt; &lt;p&gt;Specifying this header with a PUT action doesn’t affect bucket-level settings for S3 Bucket Key.&lt;/p&gt; | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzTagging** | **string** | The tag-set for the object. The tag-set must be encoded as URL Query parameters. (For example, \&quot;Key1&#x3D;Value1\&quot;) | 
 **xAmzObjectLockMode** | **string** | The Object Lock mode that you want to apply to this object. | 
 **xAmzObjectLockRetainUntilDate** | **time.Time** | The date and time when you want this object&#39;s Object Lock to expire. Must be formatted as a timestamp parameter. | 
 **xAmzObjectLockLegalHold** | **string** | Specifies whether a legal hold will be applied to this object. For more information about S3 Object Lock, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html\&quot;&gt;Object Lock&lt;/a&gt;. | 
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

