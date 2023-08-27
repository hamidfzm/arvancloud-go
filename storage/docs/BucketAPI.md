# \BucketAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBucket**](BucketAPI.md#CreateBucket) | **Put** /{Bucket} | 
[**DeleteBucket**](BucketAPI.md#DeleteBucket) | **Delete** /{Bucket} | 
[**DeleteObjects**](BucketAPI.md#DeleteObjects) | **Post** /{Bucket}#delete | 
[**DeletePublicAccessBlock**](BucketAPI.md#DeletePublicAccessBlock) | **Delete** /{Bucket}#publicAccessBlock | 
[**GetPublicAccessBlock**](BucketAPI.md#GetPublicAccessBlock) | **Get** /{Bucket}#publicAccessBlock | 
[**HeadBucket**](BucketAPI.md#HeadBucket) | **Head** /{Bucket} | 
[**ListBuckets**](BucketAPI.md#ListBuckets) | **Get** /{Bucket} | 
[**ListObjectVersions**](BucketAPI.md#ListObjectVersions) | **Get** /{Bucket}#versions | 
[**ListObjects**](BucketAPI.md#ListObjects) | **Get** /{Bucket}#Listobjects | 
[**ListObjectsV2**](BucketAPI.md#ListObjectsV2) | **Get** /{Bucket}#list-type&#x3D;2 | 
[**PutPublicAccessBlock**](BucketAPI.md#PutPublicAccessBlock) | **Put** /{Bucket}#publicAccessBlock | 



## CreateBucket

> map[string]interface{} CreateBucket(ctx, bucket).CreateBucketRequest(createBucketRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWrite(xAmzGrantWrite).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzBucketObjectLockEnabled(xAmzBucketObjectLockEnabled).Execute()





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
    bucket := "bucket_example" // string | The name of the bucket to create.
    createBucketRequest := *openapiclient.NewCreateBucketRequest() // CreateBucketRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzAcl := "xAmzAcl_example" // string | The canned ACL to apply to the bucket. (optional)
    xAmzGrantFullControl := "xAmzGrantFullControl_example" // string | Allows grantee the read, write, read ACP, and write ACP permissions on the bucket. (optional)
    xAmzGrantRead := "xAmzGrantRead_example" // string | Allows grantee to list the objects in the bucket. (optional)
    xAmzGrantReadAcp := "xAmzGrantReadAcp_example" // string | Allows grantee to read the bucket ACL. (optional)
    xAmzGrantWrite := "xAmzGrantWrite_example" // string | <p>Allows grantee to create new objects in the bucket.</p> <p>For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.</p> (optional)
    xAmzGrantWriteAcp := "xAmzGrantWriteAcp_example" // string | Allows grantee to write the ACL for the applicable bucket. (optional)
    xAmzBucketObjectLockEnabled := true // bool | Specifies whether you want S3 Object Lock to be enabled for the new bucket. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.CreateBucket(context.Background(), bucket).CreateBucketRequest(createBucketRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzAcl(xAmzAcl).XAmzGrantFullControl(xAmzGrantFullControl).XAmzGrantRead(xAmzGrantRead).XAmzGrantReadAcp(xAmzGrantReadAcp).XAmzGrantWrite(xAmzGrantWrite).XAmzGrantWriteAcp(xAmzGrantWriteAcp).XAmzBucketObjectLockEnabled(xAmzBucketObjectLockEnabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.CreateBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBucket`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.CreateBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The name of the bucket to create. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createBucketRequest** | [**CreateBucketRequest**](CreateBucketRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzAcl** | **string** | The canned ACL to apply to the bucket. | 
 **xAmzGrantFullControl** | **string** | Allows grantee the read, write, read ACP, and write ACP permissions on the bucket. | 
 **xAmzGrantRead** | **string** | Allows grantee to list the objects in the bucket. | 
 **xAmzGrantReadAcp** | **string** | Allows grantee to read the bucket ACL. | 
 **xAmzGrantWrite** | **string** | &lt;p&gt;Allows grantee to create new objects in the bucket.&lt;/p&gt; &lt;p&gt;For the bucket and object owners of existing objects, also allows deletions and overwrites of those objects.&lt;/p&gt; | 
 **xAmzGrantWriteAcp** | **string** | Allows grantee to write the ACL for the applicable bucket. | 
 **xAmzBucketObjectLockEnabled** | **bool** | Specifies whether you want S3 Object Lock to be enabled for the new bucket. | 

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


## DeleteBucket

> DeleteBucket(ctx, bucket).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | Specifies the bucket being deleted.
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketAPI.DeleteBucket(context.Background(), bucket).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DeleteBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | Specifies the bucket being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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


## DeleteObjects

> DeleteObjectsOutput DeleteObjects(ctx, bucket).Delete(delete).DeleteObjectsRequest(deleteObjectsRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzMfa(xAmzMfa).XAmzRequestPayer(xAmzRequestPayer).XAmzBypassGovernanceRetention(xAmzBypassGovernanceRetention).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name containing the objects to delete. </p>
    delete := true // bool | 
    deleteObjectsRequest := *openapiclient.NewDeleteObjectsRequest(*openapiclient.NewDeleteObjectsRequestDelete("TODO")) // DeleteObjectsRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzMfa := "xAmzMfa_example" // string | The concatenation of the authentication device's serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA delete enabled. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)
    xAmzBypassGovernanceRetention := true // bool | Specifies whether you want to delete this object even if it has a Governance-type Object Lock in place. To use this header, you must have the <code>s3:PutBucketPublicAccessBlock</code> permission. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.DeleteObjects(context.Background(), bucket).Delete(delete).DeleteObjectsRequest(deleteObjectsRequest).XAmzSecurityToken(xAmzSecurityToken).XAmzMfa(xAmzMfa).XAmzRequestPayer(xAmzRequestPayer).XAmzBypassGovernanceRetention(xAmzBypassGovernanceRetention).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DeleteObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObjects`: DeleteObjectsOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.DeleteObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name containing the objects to delete. &lt;/p&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delete** | **bool** |  | 
 **deleteObjectsRequest** | [**DeleteObjectsRequest**](DeleteObjectsRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzMfa** | **string** | The concatenation of the authentication device&#39;s serial number, a space, and the value that is displayed on your authentication device. Required to permanently delete a versioned object if versioning is configured with MFA delete enabled. | 
 **xAmzRequestPayer** | **string** |  | 
 **xAmzBypassGovernanceRetention** | **bool** | Specifies whether you want to delete this object even if it has a Governance-type Object Lock in place. To use this header, you must have the &lt;code&gt;s3:PutBucketPublicAccessBlock&lt;/code&gt; permission. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**DeleteObjectsOutput**](DeleteObjectsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: text/xml
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePublicAccessBlock

> DeletePublicAccessBlock(ctx, bucket).PublicAccessBlock(publicAccessBlock).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | the ArvanCloud S3 bucket whose <code>PublicAccessBlock</code> configuration you want to delete. 
    publicAccessBlock := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketAPI.DeletePublicAccessBlock(context.Background(), bucket).PublicAccessBlock(publicAccessBlock).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.DeletePublicAccessBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | the ArvanCloud S3 bucket whose &lt;code&gt;PublicAccessBlock&lt;/code&gt; configuration you want to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePublicAccessBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicAccessBlock** | **bool** |  | 
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


## GetPublicAccessBlock

> GetPublicAccessBlockOutput GetPublicAccessBlock(ctx, bucket).PublicAccessBlock(publicAccessBlock).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The name of the ArvanCloud S3 bucket whose <code>PublicAccessBlock</code> configuration you want to retrieve. 
    publicAccessBlock := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.GetPublicAccessBlock(context.Background(), bucket).PublicAccessBlock(publicAccessBlock).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.GetPublicAccessBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicAccessBlock`: GetPublicAccessBlockOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.GetPublicAccessBlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The name of the ArvanCloud S3 bucket whose &lt;code&gt;PublicAccessBlock&lt;/code&gt; configuration you want to retrieve.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicAccessBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicAccessBlock** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

[**GetPublicAccessBlockOutput**](GetPublicAccessBlockOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HeadBucket

> HeadBucket(ctx, bucket).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name.</p>
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketAPI.HeadBucket(context.Background(), bucket).XAmzSecurityToken(xAmzSecurityToken).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.HeadBucket``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name.&lt;/p&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiHeadBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 

### Return type

 (empty response body)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuckets

> ListBucketsOutput ListBuckets(ctx).XAmzSecurityToken(xAmzSecurityToken).Execute()





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
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.ListBuckets(context.Background()).XAmzSecurityToken(xAmzSecurityToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.ListBuckets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuckets`: ListBucketsOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.ListBuckets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 

### Return type

[**ListBucketsOutput**](ListBucketsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectVersions

> ListObjectVersionsOutput ListObjectVersions(ctx, bucket).Versions(versions).XAmzSecurityToken(xAmzSecurityToken).Delimiter(delimiter).EncodingType(encodingType).KeyMarker(keyMarker).MaxKeys(maxKeys).Prefix(prefix).VersionIdMarker(versionIdMarker).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxKeys2(maxKeys2).KeyMarker2(keyMarker2).VersionIdMarker2(versionIdMarker2).Execute()





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
    bucket := "bucket_example" // string | The bucket name that contains the objects. 
    versions := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    delimiter := "delimiter_example" // string | A delimiter is a character that you specify to group keys. All keys that contain the same string between the <code>prefix</code> and the first occurrence of the delimiter are grouped under a single result element in CommonPrefixes. These groups are counted as one result against the max-keys limitation. These keys are not returned elsewhere in the response. (optional)
    encodingType := "encodingType_example" // string |  (optional)
    keyMarker := "keyMarker_example" // string | Specifies the key to start with when listing objects in a bucket. (optional)
    maxKeys := int32(56) // int32 | Sets the maximum number of keys returned in the response. By default the action returns up to 1,000 key names. The response might contain fewer keys but will never contain more. If additional keys satisfy the search criteria, but were not returned because max-keys was exceeded, the response contains &lt;isTruncated&gt;true&lt;/isTruncated&gt;. To return the additional keys, see key-marker and version-id-marker. (optional)
    prefix := "prefix_example" // string | Use this parameter to select only those keys that begin with the specified prefix. You can use prefixes to separate a bucket into different groupings of keys. (You can think of using prefix to make groups in the same way you'd use a folder in a file system.) You can use prefix with delimiter to roll up numerous objects into a single result under CommonPrefixes.  (optional)
    versionIdMarker := "versionIdMarker_example" // string | Specifies the object version you want to start listing from. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    maxKeys2 := "maxKeys_example" // string | Pagination limit (optional)
    keyMarker2 := "keyMarker_example" // string | Pagination token (optional)
    versionIdMarker2 := "versionIdMarker_example" // string | Pagination token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.ListObjectVersions(context.Background(), bucket).Versions(versions).XAmzSecurityToken(xAmzSecurityToken).Delimiter(delimiter).EncodingType(encodingType).KeyMarker(keyMarker).MaxKeys(maxKeys).Prefix(prefix).VersionIdMarker(versionIdMarker).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxKeys2(maxKeys2).KeyMarker2(keyMarker2).VersionIdMarker2(versionIdMarker2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.ListObjectVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjectVersions`: ListObjectVersionsOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.ListObjectVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The bucket name that contains the objects.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versions** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **delimiter** | **string** | A delimiter is a character that you specify to group keys. All keys that contain the same string between the &lt;code&gt;prefix&lt;/code&gt; and the first occurrence of the delimiter are grouped under a single result element in CommonPrefixes. These groups are counted as one result against the max-keys limitation. These keys are not returned elsewhere in the response. | 
 **encodingType** | **string** |  | 
 **keyMarker** | **string** | Specifies the key to start with when listing objects in a bucket. | 
 **maxKeys** | **int32** | Sets the maximum number of keys returned in the response. By default the action returns up to 1,000 key names. The response might contain fewer keys but will never contain more. If additional keys satisfy the search criteria, but were not returned because max-keys was exceeded, the response contains &amp;lt;isTruncated&amp;gt;true&amp;lt;/isTruncated&amp;gt;. To return the additional keys, see key-marker and version-id-marker. | 
 **prefix** | **string** | Use this parameter to select only those keys that begin with the specified prefix. You can use prefixes to separate a bucket into different groupings of keys. (You can think of using prefix to make groups in the same way you&#39;d use a folder in a file system.) You can use prefix with delimiter to roll up numerous objects into a single result under CommonPrefixes.  | 
 **versionIdMarker** | **string** | Specifies the object version you want to start listing from. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **maxKeys2** | **string** | Pagination limit | 
 **keyMarker2** | **string** | Pagination token | 
 **versionIdMarker2** | **string** | Pagination token | 

### Return type

[**ListObjectVersionsOutput**](ListObjectVersionsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjects

> ListObjectsOutput ListObjects(ctx, bucket).Delimiter(delimiter).EncodingType(encodingType).Marker(marker).MaxKeys(maxKeys).Prefix(prefix).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxKeys2(maxKeys2).Marker2(marker2).Execute()





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
    bucket := "bucket_example" // string | <p>The name of the bucket containing the objects.</p>
    delimiter := "delimiter_example" // string | A delimiter is a character you use to group keys. (optional)
    encodingType := "encodingType_example" // string |  (optional)
    marker := "marker_example" // string | Marker is where you want ArvanCloud S3 to start listing from. ArvanCloud S3 starts listing after this specified key. Marker can be any key in the bucket. (optional)
    maxKeys := int32(56) // int32 | Sets the maximum number of keys returned in the response. By default the action returns up to 1,000 key names. The response might contain fewer keys but will never contain more.  (optional)
    prefix := "prefix_example" // string | Limits the response to keys that begin with the specified prefix. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string | Confirms that the requester knows that she or he will be charged for the list objects request. Bucket owners need not specify this parameter in their requests. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    maxKeys2 := "maxKeys_example" // string | Pagination limit (optional)
    marker2 := "marker_example" // string | Pagination token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.ListObjects(context.Background(), bucket).Delimiter(delimiter).EncodingType(encodingType).Marker(marker).MaxKeys(maxKeys).Prefix(prefix).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxKeys2(maxKeys2).Marker2(marker2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.ListObjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjects`: ListObjectsOutput
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.ListObjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The name of the bucket containing the objects.&lt;/p&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **delimiter** | **string** | A delimiter is a character you use to group keys. | 
 **encodingType** | **string** |  | 
 **marker** | **string** | Marker is where you want ArvanCloud S3 to start listing from. ArvanCloud S3 starts listing after this specified key. Marker can be any key in the bucket. | 
 **maxKeys** | **int32** | Sets the maximum number of keys returned in the response. By default the action returns up to 1,000 key names. The response might contain fewer keys but will never contain more.  | 
 **prefix** | **string** | Limits the response to keys that begin with the specified prefix. | 
 **xAmzRequestPayer** | **string** | Confirms that the requester knows that she or he will be charged for the list objects request. Bucket owners need not specify this parameter in their requests. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **maxKeys2** | **string** | Pagination limit | 
 **marker2** | **string** | Pagination token | 

### Return type

[**ListObjectsOutput**](ListObjectsOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListObjectsV2

> ListObjectsV2Output ListObjectsV2(ctx, bucket).ListType(listType).XAmzSecurityToken(xAmzSecurityToken).Delimiter(delimiter).EncodingType(encodingType).MaxKeys(maxKeys).Prefix(prefix).ContinuationToken(continuationToken).FetchOwner(fetchOwner).StartAfter(startAfter).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxKeys2(maxKeys2).ContinuationToken2(continuationToken2).Execute()





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
    bucket := "bucket_example" // string | <p>Bucket name to list. </p> </p>
    listType := "listType_example" // string | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    delimiter := "delimiter_example" // string | A delimiter is a character you use to group keys. (optional)
    encodingType := "encodingType_example" // string | Encoding type used by ArvanCloud S3 to encode object keys in the response. (optional)
    maxKeys := int32(56) // int32 | Sets the maximum number of keys returned in the response. By default the action returns up to 1,000 key names. The response might contain fewer keys but will never contain more. (optional)
    prefix := "prefix_example" // string | Limits the response to keys that begin with the specified prefix. (optional)
    continuationToken := "continuationToken_example" // string | ContinuationToken indicates ArvanCloud S3 that the list is being continued on this bucket with a token. ContinuationToken is obfuscated and is not a real key. (optional)
    fetchOwner := true // bool | The owner field is not present in listV2 by default, if you want to return owner field with each key in the result then set the fetch owner field to true. (optional)
    startAfter := "startAfter_example" // string | StartAfter is where you want ArvanCloud S3 to start listing from. ArvanCloud S3 starts listing after this specified key. StartAfter can be any key in the bucket. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string | Confirms that the requester knows that she or he will be charged for the list objects request in V2 style. Bucket owners need not specify this parameter in their requests. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    maxKeys2 := "maxKeys_example" // string | Pagination limit (optional)
    continuationToken2 := "continuationToken_example" // string | Pagination token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BucketAPI.ListObjectsV2(context.Background(), bucket).ListType(listType).XAmzSecurityToken(xAmzSecurityToken).Delimiter(delimiter).EncodingType(encodingType).MaxKeys(maxKeys).Prefix(prefix).ContinuationToken(continuationToken).FetchOwner(fetchOwner).StartAfter(startAfter).XAmzRequestPayer(xAmzRequestPayer).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).MaxKeys2(maxKeys2).ContinuationToken2(continuationToken2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.ListObjectsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListObjectsV2`: ListObjectsV2Output
    fmt.Fprintf(os.Stdout, "Response from `BucketAPI.ListObjectsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;Bucket name to list. &lt;/p&gt; &lt;/p&gt; | 

### Other Parameters

Other parameters are passed through a pointer to a apiListObjectsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **listType** | **string** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **delimiter** | **string** | A delimiter is a character you use to group keys. | 
 **encodingType** | **string** | Encoding type used by ArvanCloud S3 to encode object keys in the response. | 
 **maxKeys** | **int32** | Sets the maximum number of keys returned in the response. By default the action returns up to 1,000 key names. The response might contain fewer keys but will never contain more. | 
 **prefix** | **string** | Limits the response to keys that begin with the specified prefix. | 
 **continuationToken** | **string** | ContinuationToken indicates ArvanCloud S3 that the list is being continued on this bucket with a token. ContinuationToken is obfuscated and is not a real key. | 
 **fetchOwner** | **bool** | The owner field is not present in listV2 by default, if you want to return owner field with each key in the result then set the fetch owner field to true. | 
 **startAfter** | **string** | StartAfter is where you want ArvanCloud S3 to start listing from. ArvanCloud S3 starts listing after this specified key. StartAfter can be any key in the bucket. | 
 **xAmzRequestPayer** | **string** | Confirms that the requester knows that she or he will be charged for the list objects request in V2 style. Bucket owners need not specify this parameter in their requests. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **maxKeys2** | **string** | Pagination limit | 
 **continuationToken2** | **string** | Pagination token | 

### Return type

[**ListObjectsV2Output**](ListObjectsV2Output.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPublicAccessBlock

> PutPublicAccessBlock(ctx, bucket).PublicAccessBlock(publicAccessBlock).PutPublicAccessBlockRequest(putPublicAccessBlockRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | The name of the ArvanCloud S3 bucket whose <code>PublicAccessBlock</code> configuration you want to set.
    publicAccessBlock := true // bool | 
    putPublicAccessBlockRequest := *openapiclient.NewPutPublicAccessBlockRequest(*openapiclient.NewPutPublicAccessBlockRequestPublicAccessBlockConfiguration()) // PutPublicAccessBlockRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    contentMD5 := "contentMD5_example" // string | <p>The MD5 hash of the <code>PutPublicAccessBlock</code> request body. </p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.BucketAPI.PutPublicAccessBlock(context.Background(), bucket).PublicAccessBlock(publicAccessBlock).PutPublicAccessBlockRequest(putPublicAccessBlockRequest).XAmzSecurityToken(xAmzSecurityToken).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BucketAPI.PutPublicAccessBlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | The name of the ArvanCloud S3 bucket whose &lt;code&gt;PublicAccessBlock&lt;/code&gt; configuration you want to set. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPublicAccessBlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publicAccessBlock** | **bool** |  | 
 **putPublicAccessBlockRequest** | [**PutPublicAccessBlockRequest**](PutPublicAccessBlockRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **contentMD5** | **string** | &lt;p&gt;The MD5 hash of the &lt;code&gt;PutPublicAccessBlock&lt;/code&gt; request body. &lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
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

