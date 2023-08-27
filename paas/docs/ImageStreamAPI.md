# \ImageStreamAPI

All URIs are relative to *https://napi.arvancloud.ir/paas/v1/regions/ir-thr-at1/o*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedImageStream**](ImageStreamAPI.md#CreateNamespacedImageStream) | **Post** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams | create an ImageStream
[**DeleteNamespacedImageStream**](ImageStreamAPI.md#DeleteNamespacedImageStream) | **Delete** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams/{name} | delete an ImageStream
[**DeletecollectionNamespacedImageStream**](ImageStreamAPI.md#DeletecollectionNamespacedImageStream) | **Delete** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams | delete collection of ImageStream
[**ListNamespacedImageStream**](ImageStreamAPI.md#ListNamespacedImageStream) | **Get** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams | list or watch objects of kind ImageStream
[**PatchNamespacedImageStream**](ImageStreamAPI.md#PatchNamespacedImageStream) | **Patch** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams/{name} | partially update the specified ImageStream
[**ReadNamespacedImageStream**](ImageStreamAPI.md#ReadNamespacedImageStream) | **Get** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams/{name} | read the specified ImageStream
[**ReplaceNamespacedImageStream**](ImageStreamAPI.md#ReplaceNamespacedImageStream) | **Put** /apis/image.openshift.io/v1/namespaces/{namespace}/imagestreams/{name} | replace the specified ImageStream



## CreateNamespacedImageStream

> V1ImageStream CreateNamespacedImageStream(ctx, namespace).Body(body).Pretty(pretty).Execute()

create an ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    body := *openapiclient.NewV1ImageStream(*openapiclient.NewV1ImageStreamSpec()) // V1ImageStream | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.CreateNamespacedImageStream(context.Background(), namespace).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.CreateNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNamespacedImageStream`: V1ImageStream
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.CreateNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1ImageStream**](V1ImageStream.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ImageStream**](V1ImageStream.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedImageStream

> V1Status DeleteNamespacedImageStream(ctx, namespace, name).Body(body).Pretty(pretty).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

delete an ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    name := "name_example" // string | name of the ImageStream
    body := *openapiclient.NewV1DeleteOptions() // V1DeleteOptions | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    gracePeriodSeconds := int32(56) // int32 | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. (optional)
    orphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. (optional)
    propagationPolicy := "propagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.DeleteNamespacedImageStream(context.Background(), namespace, name).Body(body).Pretty(pretty).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.DeleteNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNamespacedImageStream`: V1Status
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.DeleteNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ImageStream | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1DeleteOptions**](V1DeleteOptions.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **gracePeriodSeconds** | **int32** | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. | 
 **orphanDependents** | **bool** | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \&quot;orphan\&quot; finalizer will be added to/removed from the object&#39;s finalizers list. Either this field or PropagationPolicy may be set, but not both. | 
 **propagationPolicy** | **string** | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: &#39;Orphan&#39; - orphan the dependents; &#39;Background&#39; - allow the garbage collector to delete the dependents in the background; &#39;Foreground&#39; - a cascading policy that deletes all dependents in the foreground. | 

### Return type

[**V1Status**](V1Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletecollectionNamespacedImageStream

> V1Status DeletecollectionNamespacedImageStream(ctx, namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()

delete collection of ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    labelSelector := "labelSelector_example" // string | A selector to restrict the list of returned objects by their labels. Defaults to everything. (optional)
    fieldSelector := "fieldSelector_example" // string | A selector to restrict the list of returned objects by their fields. Defaults to everything. (optional)
    includeUninitialized := true // bool | If true, partially initialized resources are included in the response. (optional)
    watch := true // bool | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. (optional)
    resourceVersion := "resourceVersion_example" // string | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. (optional)
    timeoutSeconds := int32(56) // int32 | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. (optional)
    limit := int32(56) // int32 | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. (optional)
    continue_ := "continue__example" // string | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.DeletecollectionNamespacedImageStream(context.Background(), namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.DeletecollectionNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletecollectionNamespacedImageStream`: V1Status
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.DeletecollectionNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletecollectionNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **labelSelector** | **string** | A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **fieldSelector** | **string** | A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **includeUninitialized** | **bool** | If true, partially initialized resources are included in the response. | 
 **watch** | **bool** | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 
 **resourceVersion** | **string** | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **int32** | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **limit** | **int32** | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **continue_** | **string** | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 

### Return type

[**V1Status**](V1Status.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNamespacedImageStream

> V1ImageStreamList ListNamespacedImageStream(ctx, namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()

list or watch objects of kind ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    labelSelector := "labelSelector_example" // string | A selector to restrict the list of returned objects by their labels. Defaults to everything. (optional)
    fieldSelector := "fieldSelector_example" // string | A selector to restrict the list of returned objects by their fields. Defaults to everything. (optional)
    includeUninitialized := true // bool | If true, partially initialized resources are included in the response. (optional)
    watch := true // bool | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. (optional)
    resourceVersion := "resourceVersion_example" // string | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it's 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. (optional)
    timeoutSeconds := int32(56) // int32 | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. (optional)
    limit := int32(56) // int32 | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the `continue` field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. (optional)
    continue_ := "continue__example" // string | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.ListNamespacedImageStream(context.Background(), namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.ListNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNamespacedImageStream`: V1ImageStreamList
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.ListNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **labelSelector** | **string** | A selector to restrict the list of returned objects by their labels. Defaults to everything. | 
 **fieldSelector** | **string** | A selector to restrict the list of returned objects by their fields. Defaults to everything. | 
 **includeUninitialized** | **bool** | If true, partially initialized resources are included in the response. | 
 **watch** | **bool** | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | 
 **resourceVersion** | **string** | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. When specified for list: - if unset, then the result is returned from remote storage based on quorum-read flag; - if it&#39;s 0, then we simply return what we currently have in cache, no guarantee; - if set to non zero, then the result is at least as fresh as given rv. | 
 **timeoutSeconds** | **int32** | Timeout for the list/watch call. This limits the duration of the call, regardless of any activity or inactivity. | 
 **limit** | **int32** | limit is a maximum number of responses to return for a list call. If more items exist, the server will set the &#x60;continue&#x60; field on the list metadata to a value that can be used with the same initial query to retrieve the next set of results. Setting a limit may return fewer than the requested amount of items (up to zero items) in the event all requested objects are filtered out and clients should only use the presence of the continue field to determine whether more results are available. Servers may choose not to support the limit argument and will return all of the available results. If limit is specified and the continue field is empty, clients may assume that no more results are available. This field is not supported if watch is true.  The server guarantees that the objects returned when using continue will be identical to issuing a single list call without a limit - that is, no objects created, modified, or deleted after the first request is issued will be included in any subsequent continued requests. This is sometimes referred to as a consistent snapshot, and ensures that a client that is using limit to receive smaller chunks of a very large result can ensure they see all possible objects. If objects are updated during a chunked list the version of the object that was present at the time the first list result was calculated is returned. | 
 **continue_** | **string** | The continue option should be set when retrieving more results from the server. Since this value is server defined, clients may only use the continue value from a previous query result with identical query parameters (except for the value of continue) and the server may reject a continue value it does not recognize. If the specified continue value is no longer valid whether due to expiration (generally five to fifteen minutes) or a configuration change on the server the server will respond with a 410 ResourceExpired error indicating the client must restart their list without the continue field. This field is not supported when watch is true. Clients may start a watch from the last resourceVersion value returned by the server and not miss any modifications. | 

### Return type

[**V1ImageStreamList**](V1ImageStreamList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedImageStream

> V1ImageStream PatchNamespacedImageStream(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

partially update the specified ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    name := "name_example" // string | name of the ImageStream
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.PatchNamespacedImageStream(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.PatchNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNamespacedImageStream`: V1ImageStream
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.PatchNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ImageStream | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ImageStream**](V1ImageStream.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedImageStream

> V1ImageStream ReadNamespacedImageStream(ctx, namespace, name).Pretty(pretty).Export(export).Exact(exact).Execute()

read the specified ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    name := "name_example" // string | name of the ImageStream
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    export := true // bool | Should this value be exported.  Export strips fields that a user can not specify. (optional)
    exact := true // bool | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.ReadNamespacedImageStream(context.Background(), namespace, name).Pretty(pretty).Export(export).Exact(exact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.ReadNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNamespacedImageStream`: V1ImageStream
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.ReadNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ImageStream | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **export** | **bool** | Should this value be exported.  Export strips fields that a user can not specify. | 
 **exact** | **bool** | Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 

### Return type

[**V1ImageStream**](V1ImageStream.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedImageStream

> V1ImageStream ReplaceNamespacedImageStream(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

replace the specified ImageStream

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/paas"
)

func main() {
    namespace := "namespace_example" // string | object name and auth scope, such as for teams and projects
    name := "name_example" // string | name of the ImageStream
    body := *openapiclient.NewV1ImageStream(*openapiclient.NewV1ImageStreamSpec()) // V1ImageStream | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ImageStreamAPI.ReplaceNamespacedImageStream(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ImageStreamAPI.ReplaceNamespacedImageStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNamespacedImageStream`: V1ImageStream
    fmt.Fprintf(os.Stdout, "Response from `ImageStreamAPI.ReplaceNamespacedImageStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ImageStream | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedImageStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1ImageStream**](V1ImageStream.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ImageStream**](V1ImageStream.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

