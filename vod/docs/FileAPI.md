# \FileAPI

All URIs are relative to *https://napi.arvancloud.ir/vod/2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChannelsChannelFilesFileHead**](FileAPI.md#ChannelsChannelFilesFileHead) | **Head** /channels/{channel}/files/{file} | Get upload offset. See https://tus.io/ for more detail.
[**ChannelsChannelFilesFilePatch**](FileAPI.md#ChannelsChannelFilesFilePatch) | **Patch** /channels/{channel}/files/{file} | Upload and apply bytes to a file. See https://tus.io/ for more detail.
[**ChannelsChannelFilesGet**](FileAPI.md#ChannelsChannelFilesGet) | **Get** /channels/{channel}/files | Return all draft files of channel.
[**ChannelsChannelFilesPost**](FileAPI.md#ChannelsChannelFilesPost) | **Post** /channels/{channel}/files | Request a new upload file. See https://tus.io/ for more detail.
[**FilesFileDelete**](FileAPI.md#FilesFileDelete) | **Delete** /files/{file} | Remove the specified file.
[**FilesFileGet**](FileAPI.md#FilesFileGet) | **Get** /files/{file} | Return the specified file.



## ChannelsChannelFilesFileHead

> ChannelsChannelFilesFileHead(ctx, channel, file).Execute()

Get upload offset. See https://tus.io/ for more detail.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    file := "file_example" // string | The Id of file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileAPI.ChannelsChannelFilesFileHead(context.Background(), channel, file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.ChannelsChannelFilesFileHead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 
**file** | **string** | The Id of file | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelFilesFileHeadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelFilesFilePatch

> ChannelsChannelFilesFilePatch(ctx, channel, file).TusResumable(tusResumable).UploadOffset(uploadOffset).ContentType(contentType).Execute()

Upload and apply bytes to a file. See https://tus.io/ for more detail.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    file := "file_example" // string | The Id of file
    tusResumable := "tusResumable_example" // string | version of tus.io (default to "1.0.0")
    uploadOffset := int32(56) // int32 | request and response header indicates a byte offset within a resource.      * For uploading entire file in one request, set this to '0' (default to 0)
    contentType := "contentType_example" // string | Request content type (default to "application/offset+octet-stream")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileAPI.ChannelsChannelFilesFilePatch(context.Background(), channel, file).TusResumable(tusResumable).UploadOffset(uploadOffset).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.ChannelsChannelFilesFilePatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 
**file** | **string** | The Id of file | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelFilesFilePatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tusResumable** | **string** | version of tus.io | [default to &quot;1.0.0&quot;]
 **uploadOffset** | **int32** | request and response header indicates a byte offset within a resource.      * For uploading entire file in one request, set this to &#39;0&#39; | [default to 0]
 **contentType** | **string** | Request content type | [default to &quot;application/offset+octet-stream&quot;]

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelFilesGet

> ChannelsChannelFilesGet(ctx, channel).Filter(filter).Execute()

Return all draft files of channel.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    filter := "filter_example" // string | Filter result (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileAPI.ChannelsChannelFilesGet(context.Background(), channel).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.ChannelsChannelFilesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Filter result | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ChannelsChannelFilesPost

> ChannelsChannelFilesPost(ctx, channel).TusResumable(tusResumable).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()

Request a new upload file. See https://tus.io/ for more detail.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    channel := "channel_example" // string | The Id of channel
    tusResumable := "tusResumable_example" // string | version of tus.io (default to "1.0.0")
    uploadLength := int32(56) // int32 | To indicate the size of entire upload in bytes
    uploadMetadata := "uploadMetadata_example" // string | To add additional metadata to the upload creation request.      * MUST contain 'filename' and 'filetype'. From all available fields only these two fields will be used.      * MUST consist of one or more comma-separated key-value pairs.      * The key and value MUST be separated by a space.      * The key MUST NOT contain spaces and commas and MUST NOT be empty.      * The key SHOULD be ASCII encoded and the value MUST be Base64 encoded.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileAPI.ChannelsChannelFilesPost(context.Background(), channel).TusResumable(tusResumable).UploadLength(uploadLength).UploadMetadata(uploadMetadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.ChannelsChannelFilesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channel** | **string** | The Id of channel | 

### Other Parameters

Other parameters are passed through a pointer to a apiChannelsChannelFilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tusResumable** | **string** | version of tus.io | [default to &quot;1.0.0&quot;]
 **uploadLength** | **int32** | To indicate the size of entire upload in bytes | 
 **uploadMetadata** | **string** | To add additional metadata to the upload creation request.      * MUST contain &#39;filename&#39; and &#39;filetype&#39;. From all available fields only these two fields will be used.      * MUST consist of one or more comma-separated key-value pairs.      * The key and value MUST be separated by a space.      * The key MUST NOT contain spaces and commas and MUST NOT be empty.      * The key SHOULD be ASCII encoded and the value MUST be Base64 encoded. | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesFileDelete

> FilesFileDelete(ctx, file).Execute()

Remove the specified file.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    file := "file_example" // string | The Id of file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileAPI.FilesFileDelete(context.Background(), file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FilesFileDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | **string** | The Id of file | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesFileDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FilesFileGet

> FilesFileGet(ctx, file).Execute()

Return the specified file.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/vod"
)

func main() {
    file := "file_example" // string | The Id of file

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileAPI.FilesFileGet(context.Background(), file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileAPI.FilesFileGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**file** | **string** | The Id of file | 

### Other Parameters

Other parameters are passed through a pointer to a apiFilesFileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

