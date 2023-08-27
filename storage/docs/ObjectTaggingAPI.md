# \ObjectTaggingAPI

All URIs are relative to *https://s3.ir-thr-at1.arvanstorage.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectTagging**](ObjectTaggingAPI.md#DeleteObjectTagging) | **Delete** /{Bucket}/{Key}#tagging | 
[**GetObjectTagging**](ObjectTaggingAPI.md#GetObjectTagging) | **Get** /{Bucket}/{Key}#tagging | 
[**PutObjectTagging**](ObjectTaggingAPI.md#PutObjectTagging) | **Put** /{Bucket}/{Key}#tagging | 



## DeleteObjectTagging

> map[string]interface{} DeleteObjectTagging(ctx, bucket, key).Tagging(tagging).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name containing the objects from which to remove the tags. </p>
    key := "key_example" // string | The key that identifies the object in the bucket from which to remove all tags.
    tagging := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    versionId := "versionId_example" // string | The versionId of the object that the tag-set will be removed from. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectTaggingAPI.DeleteObjectTagging(context.Background(), bucket, key).Tagging(tagging).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectTaggingAPI.DeleteObjectTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteObjectTagging`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectTaggingAPI.DeleteObjectTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name containing the objects from which to remove the tags. &lt;/p&gt; | 
**key** | **string** | The key that identifies the object in the bucket from which to remove all tags. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagging** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **versionId** | **string** | The versionId of the object that the tag-set will be removed from. | 
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


## GetObjectTagging

> GetObjectTaggingOutput GetObjectTagging(ctx, bucket, key).Tagging(tagging).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).XAmzRequestPayer(xAmzRequestPayer).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name containing the object for which to get the tagging information. </p> </p>
    key := "key_example" // string | Object key for which to get the tagging information.
    tagging := true // bool | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    versionId := "versionId_example" // string | The versionId of the object for which to get the tagging information. (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectTaggingAPI.GetObjectTagging(context.Background(), bucket, key).Tagging(tagging).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).XAmzRequestPayer(xAmzRequestPayer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectTaggingAPI.GetObjectTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetObjectTagging`: GetObjectTaggingOutput
    fmt.Fprintf(os.Stdout, "Response from `ObjectTaggingAPI.GetObjectTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name containing the object for which to get the tagging information. &lt;/p&gt; &lt;/p&gt; | 
**key** | **string** | Object key for which to get the tagging information. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagging** | **bool** |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **versionId** | **string** | The versionId of the object for which to get the tagging information. | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **xAmzRequestPayer** | **string** |  | 

### Return type

[**GetObjectTaggingOutput**](GetObjectTaggingOutput.md)

### Authorization

[hmac](../README.md#hmac)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObjectTagging

> map[string]interface{} PutObjectTagging(ctx, bucket, key).Tagging(tagging).PutBucketTaggingRequest(putBucketTaggingRequest).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).XAmzRequestPayer(xAmzRequestPayer).Execute()





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
    bucket := "bucket_example" // string | <p>The bucket name containing the object. </p> </p>
    key := "key_example" // string | Name of the object key.
    tagging := true // bool | 
    putBucketTaggingRequest := *openapiclient.NewPutBucketTaggingRequest(*openapiclient.NewPutBucketTaggingRequestTagging("TODO")) // PutBucketTaggingRequest | 
    xAmzSecurityToken := "xAmzSecurityToken_example" // string | This parameter is currently not supported and is not required. (optional)
    versionId := "versionId_example" // string | The versionId of the object that the tag-set will be added to. (optional)
    contentMD5 := "contentMD5_example" // string | <p>The MD5 hash for the request body.</p> <p>For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.</p> (optional)
    xAmzExpectedBucketOwner := "xAmzExpectedBucketOwner_example" // string | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP <code>403 (Access Denied)</code> error. (optional)
    xAmzRequestPayer := "xAmzRequestPayer_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ObjectTaggingAPI.PutObjectTagging(context.Background(), bucket, key).Tagging(tagging).PutBucketTaggingRequest(putBucketTaggingRequest).XAmzSecurityToken(xAmzSecurityToken).VersionId(versionId).ContentMD5(contentMD5).XAmzExpectedBucketOwner(xAmzExpectedBucketOwner).XAmzRequestPayer(xAmzRequestPayer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ObjectTaggingAPI.PutObjectTagging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutObjectTagging`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ObjectTaggingAPI.PutObjectTagging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bucket** | **string** | &lt;p&gt;The bucket name containing the object. &lt;/p&gt; &lt;/p&gt; | 
**key** | **string** | Name of the object key. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectTaggingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tagging** | **bool** |  | 
 **putBucketTaggingRequest** | [**PutBucketTaggingRequest**](PutBucketTaggingRequest.md) |  | 
 **xAmzSecurityToken** | **string** | This parameter is currently not supported and is not required. | 
 **versionId** | **string** | The versionId of the object that the tag-set will be added to. | 
 **contentMD5** | **string** | &lt;p&gt;The MD5 hash for the request body.&lt;/p&gt; &lt;p&gt;For requests made using the Amazon Web Services Command Line Interface (CLI) or Amazon Web Services SDKs, this field is calculated automatically.&lt;/p&gt; | 
 **xAmzExpectedBucketOwner** | **string** | The account ID of the expected bucket owner. If the bucket is owned by a different account, the request will fail with an HTTP &lt;code&gt;403 (Access Denied)&lt;/code&gt; error. | 
 **xAmzRequestPayer** | **string** |  | 

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

