# \MultipartAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AbortMultipartUpload**](MultipartAPI.md#AbortMultipartUpload) | **Delete** /{Bucket}/{Key}#uploadId | 
[**CompleteMultipartUpload**](MultipartAPI.md#CompleteMultipartUpload) | **Post** /{Bucket}/{Key}#uploadId | 
[**CreateMultipartUpload**](MultipartAPI.md#CreateMultipartUpload) | **Post** /{Bucket}/{Key}#uploads | 
[**ListMultipartUploads**](MultipartAPI.md#ListMultipartUploads) | **Get** /{Bucket}#uploads | 
[**ListParts**](MultipartAPI.md#ListParts) | **Get** /{Bucket}/{Key}#uploadId | 
[**UploadPart**](MultipartAPI.md#UploadPart) | **Put** /{Bucket}/{Key}#partNumber&amp;uploadId | 



## AbortMultipartUpload

> map[string]interface{} AbortMultipartUpload(ctx, bucket, key).UploadId(uploadId).XAmzSecurityToken(xAmzSecurityToken).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name to which the upload was taking place. </p>
    key := "key_example" // string | Key of the object for which the multipart upload was initiated.
    uploadId := "uploadId_example" // string | Upload ID that identifies the multipart upload.
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipartAPI.AbortMultipartUpload(context.Background(), bucket, key).UploadId(uploadId).XAmzSecurityToken(xAmzSecurityToken).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipartAPI.AbortMultipartUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AbortMultipartUpload`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultipartAPI.AbortMultipartUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name to which the upload was taking place. &lt;/p&gt; | 
**key** | **string** | Key of the object for which the multipart upload was initiated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAbortMultipartUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadId** | **string** | Upload ID that identifies the multipart upload. | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzRequestPayer** | **string** |  | 
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


## CompleteMultipartUpload

> CompleteMultipartUploadOutput CompleteMultipartUpload(ctx, bucket, key).UploadId(uploadId).CompleteMultipartUploadRequest(completeMultipartUploadRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>Name of the bucket to which the multipart upload was initiated.</p> </p>
    key := "key_example" // string | Object key for which the multipart upload was initiated.
    uploadId := "uploadId_example" // string | ID for the initiated multipart upload.
    completeMultipartUploadRequest := *openapiclient.NewCompleteMultipartUploadRequest() // CompleteMultipartUploadRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipartAPI.CompleteMultipartUpload(context.Background(), bucket, key).UploadId(uploadId).CompleteMultipartUploadRequest(completeMultipartUploadRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipartAPI.CompleteMultipartUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CompleteMultipartUpload`: CompleteMultipartUploadOutput
    fmt.Fprintf(os.Stdout, "Response from `MultipartAPI.CompleteMultipartUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;Name of the bucket to which the multipart upload was initiated.&lt;/p&gt; &lt;/p&gt; | 
**key** | **string** | Object key for which the multipart upload was initiated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCompleteMultipartUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadId** | **string** | ID for the initiated multipart upload. | 
 **completeMultipartUploadRequest** | [**CompleteMultipartUploadRequest**](CompleteMultipartUploadRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**CompleteMultipartUploadOutput**](CompleteMultipartUploadOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMultipartUpload

> CreateMultipartUploadOutput CreateMultipartUpload(ctx, bucket, key).Uploads(uploads).CreateMultipartUploadRequest(createMultipartUploadRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).CacheControl(cacheControl).ContentDisposition(contentDisposition).ContentEncoding(contentEncoding).ContentLanguage(contentLanguage).ContentType(contentType).Expires(expires).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzServerSideEncryption(xAmzServerSideEncryption).XAmzStorageClass(xAmzStorageClass).XAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzServerSideEncryptionAwsKmsKeyId(xAmzServerSideEncryptionAwsKmsKeyId).XAmzServerSideEncryptionContext(xAmzServerSideEncryptionContext).XAmzServerSideEncryptionBucketKeyEnabled(xAmzServerSideEncryptionBucketKeyEnabled).XAmzRequestPayer(xAmzRequestPayer).XAmzTagging(xAmzTagging).XAmzObjectLockMode(xAmzObjectLockMode).XAmzObjectLockRetainUntilDate(xAmzObjectLockRetainUntilDate).XAmzObjectLockLegalHold(xAmzObjectLockLegalHold).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The name of the bucket to which to initiate the upload</p> <p>
    key := "key_example" // string | Object key for which the multipart upload is to be initiated.
    uploads := true // bool | 
    createMultipartUploadRequest := *openapiclient.NewCreateMultipartUploadRequest() // CreateMultipartUploadRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzAcl := "xAmzAcl_example" // string | <p>The canned ACL to apply to the object.</p> </p> (optional)
    cacheControl := "cacheControl_example" // string | Specifies caching behavior along the request/reply chain. (optional)
    contentDisposition := "contentDisposition_example" // string | Specifies presentational information for the object. (optional)
    contentEncoding := "contentEncoding_example" // string | Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. (optional)
    contentLanguage := "contentLanguage_example" // string | The language the content is in. (optional)
    contentType := "contentType_example" // string | A standard MIME type describing the format of the object data. (optional)
    expires := time.Now() // time.Time | The date and time at which the object is no longer cacheable. (optional)
    xAmzGrantFullControl := "xAmzGrantFullControl_example" // string | <p>Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.</p> </p> (optional)
    xAmzGrantRead := "xAmzGrantRead_example" // string | <p>Allows grantee to read the object data and its metadata.</p> </p> (optional)
    xAmzGrantReadAcp := "xAmzGrantReadAcp_example" // string | <p>Allows grantee to read the object ACL.</p> </p> (optional)
    xAmzGrantWriteAcp := "xAmzGrantWriteAcp_example" // string | <p>Allows grantee to write the ACL for the applicable object.</p> </p> (optional)
    xAmzServerSideEncryption := "xAmzServerSideEncryption_example" // string | The server-side encryption algorithm used when storing this object in ArvanCloud S3 (for example, AES256, aws:kms). (optional)
    xAmzStorageClass := "xAmzStorageClass_example" // string | By default, ArvanCloud S3 uses the STANDARD Storage Class to store newly created objects. The STANDARD storage class provides high durability and high availability. Depending on performance needs, you can specify a different Storage Class. (optional)
    xAmzWebsiteRedirectLocation := "xAmzWebsiteRedirectLocation_example" // string | If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. ArvanCloud S3 stores the value of this header in the object metadata. (optional)
    xAmzServerSideEncryptionCustomerAlgorithm := "xAmzServerSideEncryptionCustomerAlgorithm_example" // string | Specifies the algorithm to use to when encrypting the object (for example, AES256). (optional)
    xAmzServerSideEncryptionCustomerKey := "xAmzServerSideEncryptionCustomerKey_example" // string | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the <code>x-amz-server-side-encryption-customer-algorithm</code> header. (optional)
    xAmzServerSideEncryptionCustomerKeyMD5 := "xAmzServerSideEncryptionCustomerKeyMD5_example" // string | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. (optional)
    xAmzServerSideEncryptionAwsKmsKeyId := "xAmzServerSideEncryptionAwsKmsKeyId_example" // string | Specifies the ID of the symmetric customer managed key to use for object encryption. All GET and PUT requests for an object protected by Amazon Web Services KMS will fail if not made via SSL or using SigV4. For information about configuring using any of the officially supported Amazon Web Services SDKs and Amazon Web Services CLI, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version\">Specifying the Signature Version in Request Authentication</a> in the <i>S3 User Guide</i>. (optional)
    xAmzServerSideEncryptionContext := "xAmzServerSideEncryptionContext_example" // string | Specifies the Amazon Web Services KMS Encryption Context to use for object encryption. The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs. (optional)
    xAmzServerSideEncryptionBucketKeyEnabled := true // bool | <p>Specifies whether ArvanCloud S3 should use an S3 Bucket Key for object encryption with server-side encryption using AWS KMS (SSE-KMS). Setting this header to <code>true</code> causes ArvanCloud S3 to use an S3 Bucket Key for object encryption with SSE-KMS.</p> <p>Specifying this header with an object action doesn’t affect bucket-level settings for S3 Bucket Key.</p> (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzTagging := "xAmzTagging_example" // string | The tag-set for the object. The tag-set must be encoded as URL Query parameters. (optional)
    xAmzObjectLockMode := "xAmzObjectLockMode_example" // string | Specifies the Object Lock mode that you want to apply to the uploaded object. (optional)
    xAmzObjectLockRetainUntilDate := time.Now() // time.Time | Specifies the date and time when you want the Object Lock to expire. (optional)
    xAmzObjectLockLegalHold := "xAmzObjectLockLegalHold_example" // string | Specifies whether you want to apply a Legal Hold to the uploaded object. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipartAPI.CreateMultipartUpload(context.Background(), bucket, key).Uploads(uploads).CreateMultipartUploadRequest(createMultipartUploadRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).CacheControl(cacheControl).ContentDisposition(contentDisposition).ContentEncoding(contentEncoding).ContentLanguage(contentLanguage).ContentType(contentType).Expires(expires).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzServerSideEncryption(xAmzServerSideEncryption).XAmzStorageClass(xAmzStorageClass).XAmzWebsiteRedirectLocation(xAmzWebsiteRedirectLocation).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzServerSideEncryptionAwsKmsKeyId(xAmzServerSideEncryptionAwsKmsKeyId).XAmzServerSideEncryptionContext(xAmzServerSideEncryptionContext).XAmzServerSideEncryptionBucketKeyEnabled(xAmzServerSideEncryptionBucketKeyEnabled).XAmzRequestPayer(xAmzRequestPayer).XAmzTagging(xAmzTagging).XAmzObjectLockMode(xAmzObjectLockMode).XAmzObjectLockRetainUntilDate(xAmzObjectLockRetainUntilDate).XAmzObjectLockLegalHold(xAmzObjectLockLegalHold).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipartAPI.CreateMultipartUpload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMultipartUpload`: CreateMultipartUploadOutput
    fmt.Fprintf(os.Stdout, "Response from `MultipartAPI.CreateMultipartUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The name of the bucket to which to initiate the upload&lt;/p&gt; &lt;p&gt; | 
**key** | **string** | Object key for which the multipart upload is to be initiated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMultipartUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploads** | **bool** |  | 
 **createMultipartUploadRequest** | [**CreateMultipartUploadRequest**](CreateMultipartUploadRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzAcl** | **string** | &lt;p&gt;The canned ACL to apply to the object.&lt;/p&gt; &lt;/p&gt; | 
 **cacheControl** | **string** | Specifies caching behavior along the request/reply chain. | 
 **contentDisposition** | **string** | Specifies presentational information for the object. | 
 **contentEncoding** | **string** | Specifies what content encodings have been applied to the object and thus what decoding mechanisms must be applied to obtain the media-type referenced by the Content-Type header field. | 
 **contentLanguage** | **string** | The language the content is in. | 
 **contentType** | **string** | A standard MIME type describing the format of the object data. | 
 **expires** | **time.Time** | The date and time at which the object is no longer cacheable. | 
 **xAmzGrantFullControl** | **string** | &lt;p&gt;Gives the grantee READ, READ_ACP, and WRITE_ACP permissions on the object.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantRead** | **string** | &lt;p&gt;Allows grantee to read the object data and its metadata.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantReadAcp** | **string** | &lt;p&gt;Allows grantee to read the object ACL.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzGrantWriteAcp** | **string** | &lt;p&gt;Allows grantee to write the ACL for the applicable object.&lt;/p&gt; &lt;/p&gt; | 
 **xAmzServerSideEncryption** | **string** | The server-side encryption algorithm used when storing this object in ArvanCloud S3 (for example, AES256, aws:kms). | 
 **xAmzStorageClass** | **string** | By default, ArvanCloud S3 uses the STANDARD Storage Class to store newly created objects. The STANDARD storage class provides high durability and high availability. Depending on performance needs, you can specify a different Storage Class. | 
 **xAmzWebsiteRedirectLocation** | **string** | If the bucket is configured as a website, redirects requests for this object to another object in the same bucket or to an external URL. ArvanCloud S3 stores the value of this header in the object metadata. | 
 **xAmzServerSideEncryptionCustomerAlgorithm** | **string** | Specifies the algorithm to use to when encrypting the object (for example, AES256). | 
 **xAmzServerSideEncryptionCustomerKey** | **string** | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the &lt;code&gt;x-amz-server-side-encryption-customer-algorithm&lt;/code&gt; header. | 
 **xAmzServerSideEncryptionCustomerKeyMD5** | **string** | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. | 
 **xAmzServerSideEncryptionAwsKmsKeyId** | **string** | Specifies the ID of the symmetric customer managed key to use for object encryption. All GET and PUT requests for an object protected by Amazon Web Services KMS will fail if not made via SSL or using SigV4. For information about configuring using any of the officially supported Amazon Web Services SDKs and Amazon Web Services CLI, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version\&quot;&gt;Specifying the Signature Version in Request Authentication&lt;/a&gt; in the &lt;i&gt;S3 User Guide&lt;/i&gt;. | 
 **xAmzServerSideEncryptionContext** | **string** | Specifies the Amazon Web Services KMS Encryption Context to use for object encryption. The value of this header is a base64-encoded UTF-8 string holding JSON with the encryption context key-value pairs. | 
 **xAmzServerSideEncryptionBucketKeyEnabled** | **bool** | &lt;p&gt;Specifies whether ArvanCloud S3 should use an S3 Bucket Key for object encryption with server-side encryption using AWS KMS (SSE-KMS). Setting this header to &lt;code&gt;true&lt;/code&gt; causes ArvanCloud S3 to use an S3 Bucket Key for object encryption with SSE-KMS.&lt;/p&gt; &lt;p&gt;Specifying this header with an object action doesn’t affect bucket-level settings for S3 Bucket Key.&lt;/p&gt; | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzTagging** | **string** | The tag-set for the object. The tag-set must be encoded as URL Query parameters. | 
 **xAmzObjectLockMode** | **string** | Specifies the Object Lock mode that you want to apply to the uploaded object. | 
 **xAmzObjectLockRetainUntilDate** | **time.Time** | Specifies the date and time when you want the Object Lock to expire. | 
 **xAmzObjectLockLegalHold** | **string** | Specifies whether you want to apply a Legal Hold to the uploaded object. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**CreateMultipartUploadOutput**](CreateMultipartUploadOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMultipartUploads

> ListMultipartUploadsOutput ListMultipartUploads(ctx, bucket).Uploads(uploads).XAmzSecurityToken(xAmzSecurityToken).Delimiter(delimiter).EncodingType(encodingType).KeyMarker(keyMarker).MaxUploads(maxUploads).Prefix(prefix).UploadIdMarker(uploadIdMarker).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxUploads2(maxUploads2).KeyMarker2(keyMarker2).UploadIdMarker2(uploadIdMarker2).Execute()





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
    bucket := "bucket_example" // string | <p>The name of the bucket to which the multipart upload was initiated. </p>
    uploads := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    delimiter := "delimiter_example" // string | <p>Character you use to group keys.</p> <p>All keys that contain the same string between the prefix, if specified, and the first occurrence of the delimiter after the prefix are grouped under a single result element, <code>CommonPrefixes</code>. If you don't specify the prefix parameter, then the substring starts at the beginning of the key. The keys that are grouped under <code>CommonPrefixes</code> result element are not returned elsewhere in the response.</p> (optional)
    encodingType := "encodingType_example" // string |  (optional)
    keyMarker := "keyMarker_example" // string | <p>Together with upload-id-marker, this parameter specifies the multipart upload after which listing should begin.</p> <p>If <code>upload-id-marker</code> is not specified, only the keys lexicographically greater than the specified <code>key-marker</code> will be included in the list.</p> <p>If <code>upload-id-marker</code> is specified, any multipart uploads for a key equal to the <code>key-marker</code> might also be included, provided those multipart uploads have upload IDs lexicographically greater than the specified <code>upload-id-marker</code>.</p> (optional)
    maxUploads := int32(56) // int32 | Sets the maximum number of multipart uploads, from 1 to 1,000, to return in the response body. 1,000 is the maximum number of uploads that can be returned in a response. (optional)
    prefix := "prefix_example" // string | Lists in-progress uploads only for those keys that begin with the specified prefix. You can use prefixes to separate a bucket into different grouping of keys. (You can think of using prefix to make groups in the same way you'd use a folder in a file system.) (optional)
    uploadIdMarker := "uploadIdMarker_example" // string | Together with key-marker, specifies the multipart upload after which listing should begin. If key-marker is not specified, the upload-id-marker parameter is ignored. Otherwise, any multipart uploads for a key equal to the key-marker might be included in the list only if they have an upload ID lexicographically greater than the specified <code>upload-id-marker</code>. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    maxUploads2 := "maxUploads_example" // string | Pagination limit (optional)
    keyMarker2 := "keyMarker_example" // string | Pagination token (optional)
    uploadIdMarker2 := "uploadIdMarker_example" // string | Pagination token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipartAPI.ListMultipartUploads(context.Background(), bucket).Uploads(uploads).XAmzSecurityToken(xAmzSecurityToken).Delimiter(delimiter).EncodingType(encodingType).KeyMarker(keyMarker).MaxUploads(maxUploads).Prefix(prefix).UploadIdMarker(uploadIdMarker).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxUploads2(maxUploads2).KeyMarker2(keyMarker2).UploadIdMarker2(uploadIdMarker2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipartAPI.ListMultipartUploads``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMultipartUploads`: ListMultipartUploadsOutput
    fmt.Fprintf(os.Stdout, "Response from `MultipartAPI.ListMultipartUploads`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The name of the bucket to which the multipart upload was initiated. &lt;/p&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMultipartUploadsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **uploads** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **delimiter** | **string** | &lt;p&gt;Character you use to group keys.&lt;/p&gt; &lt;p&gt;All keys that contain the same string between the prefix, if specified, and the first occurrence of the delimiter after the prefix are grouped under a single result element, &lt;code&gt;CommonPrefixes&lt;/code&gt;. If you don&#39;t specify the prefix parameter, then the substring starts at the beginning of the key. The keys that are grouped under &lt;code&gt;CommonPrefixes&lt;/code&gt; result element are not returned elsewhere in the response.&lt;/p&gt; | 
 **encodingType** | **string** |  | 
 **keyMarker** | **string** | &lt;p&gt;Together with upload-id-marker, this parameter specifies the multipart upload after which listing should begin.&lt;/p&gt; &lt;p&gt;If &lt;code&gt;upload-id-marker&lt;/code&gt; is not specified, only the keys lexicographically greater than the specified &lt;code&gt;key-marker&lt;/code&gt; will be included in the list.&lt;/p&gt; &lt;p&gt;If &lt;code&gt;upload-id-marker&lt;/code&gt; is specified, any multipart uploads for a key equal to the &lt;code&gt;key-marker&lt;/code&gt; might also be included, provided those multipart uploads have upload IDs lexicographically greater than the specified &lt;code&gt;upload-id-marker&lt;/code&gt;.&lt;/p&gt; | 
 **maxUploads** | **int32** | Sets the maximum number of multipart uploads, from 1 to 1,000, to return in the response body. 1,000 is the maximum number of uploads that can be returned in a response. | 
 **prefix** | **string** | Lists in-progress uploads only for those keys that begin with the specified prefix. You can use prefixes to separate a bucket into different grouping of keys. (You can think of using prefix to make groups in the same way you&#39;d use a folder in a file system.) | 
 **uploadIdMarker** | **string** | Together with key-marker, specifies the multipart upload after which listing should begin. If key-marker is not specified, the upload-id-marker parameter is ignored. Otherwise, any multipart uploads for a key equal to the key-marker might be included in the list only if they have an upload ID lexicographically greater than the specified &lt;code&gt;upload-id-marker&lt;/code&gt;. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **maxUploads2** | **string** | Pagination limit | 
 **keyMarker2** | **string** | Pagination token | 
 **uploadIdMarker2** | **string** | Pagination token | 

### Return type

[**ListMultipartUploadsOutput**](ListMultipartUploadsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListParts

> ListPartsOutput ListParts(ctx, bucket, key).UploadId(uploadId).XAmzSecurityToken(xAmzSecurityToken).MaxParts(maxParts).PartNumberMarker(partNumberMarker).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxParts2(maxParts2).PartNumberMarker2(partNumberMarker2).Execute()





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
    bucket := "bucket_example" // string | <p>The name of the bucket to which the parts are being uploaded. </p> </p>
    key := "key_example" // string | Object key for which the multipart upload was initiated.
    uploadId := "uploadId_example" // string | Upload ID identifying the multipart upload whose parts are being listed.
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    maxParts := int32(56) // int32 | Sets the maximum number of parts to return. (optional)
    partNumberMarker := int32(56) // int32 | Specifies the part after which listing should begin. Only parts with higher part numbers will be listed. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    maxParts2 := "maxParts_example" // string | Pagination limit (optional)
    partNumberMarker2 := "partNumberMarker_example" // string | Pagination token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipartAPI.ListParts(context.Background(), bucket, key).UploadId(uploadId).XAmzSecurityToken(xAmzSecurityToken).MaxParts(maxParts).PartNumberMarker(partNumberMarker).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxParts2(maxParts2).PartNumberMarker2(partNumberMarker2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipartAPI.ListParts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListParts`: ListPartsOutput
    fmt.Fprintf(os.Stdout, "Response from `MultipartAPI.ListParts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The name of the bucket to which the parts are being uploaded. &lt;/p&gt; &lt;/p&gt; | 
**key** | **string** | Object key for which the multipart upload was initiated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPartsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uploadId** | **string** | Upload ID identifying the multipart upload whose parts are being listed. | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **maxParts** | **int32** | Sets the maximum number of parts to return. | 
 **partNumberMarker** | **int32** | Specifies the part after which listing should begin. Only parts with higher part numbers will be listed. | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **maxParts2** | **string** | Pagination limit | 
 **partNumberMarker2** | **string** | Pagination token | 

### Return type

[**ListPartsOutput**](ListPartsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadPart

> map[string]interface{} UploadPart(ctx, bucket, key).PartNumber(partNumber).UploadId(uploadId).UploadPartRequest(uploadPartRequest).XAmzSecurityToken(xAmzSecurityToken).ContentLength(contentLength).ContentMD5(contentMD5).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The name of the bucket to which the multipart upload was initiated.</p> </p>
    key := "key_example" // string | Object key for which the multipart upload was initiated.
    partNumber := int32(56) // int32 | Part number of part being uploaded. This is a positive integer between 1 and 10,000.
    uploadId := "uploadId_example" // string | Upload ID identifying the multipart upload whose part is being uploaded.
    uploadPartRequest := *openapiclient.NewUploadPartRequest() // UploadPartRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    contentLength := int32(56) // int32 | Size of the body in bytes. This parameter is useful when the size of the body cannot be determined automatically. (optional)
    contentMD5 := "contentMD5_example" // string | The base64-encoded 128-bit MD5 digest of the part data. This parameter is auto-populated when using the command from the CLI. This parameter is required if object lock parameters are specified. (optional)
    xAmzServerSideEncryptionCustomerAlgorithm := "xAmzServerSideEncryptionCustomerAlgorithm_example" // string | Specifies the algorithm to use to when encrypting the object (for example, AES256). (optional)
    xAmzServerSideEncryptionCustomerKey := "xAmzServerSideEncryptionCustomerKey_example" // string | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the <code>x-amz-server-side-encryption-customer-algorithm header</code>. This must be the same encryption key specified in the initiate multipart upload request. (optional)
    xAmzServerSideEncryptionCustomerKeyMD5 := "xAmzServerSideEncryptionCustomerKeyMD5_example" // string | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MultipartAPI.UploadPart(context.Background(), bucket, key).PartNumber(partNumber).UploadId(uploadId).UploadPartRequest(uploadPartRequest).XAmzSecurityToken(xAmzSecurityToken).ContentLength(contentLength).ContentMD5(contentMD5).XAmzServerSideEncryptionCustomerAlgorithm(xAmzServerSideEncryptionCustomerAlgorithm).XAmzServerSideEncryptionCustomerKey(xAmzServerSideEncryptionCustomerKey).XAmzServerSideEncryptionCustomerKeyMD5(xAmzServerSideEncryptionCustomerKeyMD5).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MultipartAPI.UploadPart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadPart`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `MultipartAPI.UploadPart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The name of the bucket to which the multipart upload was initiated.&lt;/p&gt; &lt;/p&gt; | 
**key** | **string** | Object key for which the multipart upload was initiated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadPartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **partNumber** | **int32** | Part number of part being uploaded. This is a positive integer between 1 and 10,000. | 
 **uploadId** | **string** | Upload ID identifying the multipart upload whose part is being uploaded. | 
 **uploadPartRequest** | [**UploadPartRequest**](UploadPartRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **contentLength** | **int32** | Size of the body in bytes. This parameter is useful when the size of the body cannot be determined automatically. | 
 **contentMD5** | **string** | The base64-encoded 128-bit MD5 digest of the part data. This parameter is auto-populated when using the command from the CLI. This parameter is required if object lock parameters are specified. | 
 **xAmzServerSideEncryptionCustomerAlgorithm** | **string** | Specifies the algorithm to use to when encrypting the object (for example, AES256). | 
 **xAmzServerSideEncryptionCustomerKey** | **string** | Specifies the customer-provided encryption key for ArvanCloud S3 to use in encrypting data. This value is used to store the object and then it is discarded; ArvanCloud S3 does not store the encryption key. The key must be appropriate for use with the algorithm specified in the &lt;code&gt;x-amz-server-side-encryption-customer-algorithm header&lt;/code&gt;. This must be the same encryption key specified in the initiate multipart upload request. | 
 **xAmzServerSideEncryptionCustomerKeyMD5** | **string** | Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321. ArvanCloud S3 uses this header for a message integrity check to ensure that the encryption key was transmitted without error. | 
 **xAmzRequestPayer** | **string** |  | 
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

