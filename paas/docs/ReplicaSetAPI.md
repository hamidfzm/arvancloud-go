# \ReplicaSetAPI

All URIs are relative to *https://napi.arvancloud.ir/paas/v1/regions/ir-thr-at1/o*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNamespacedReplicaSet**](ReplicaSetAPI.md#CreateNamespacedReplicaSet) | **Post** /apis/apps/v1/namespaces/{namespace}/replicasets | create a ReplicaSet
[**DeleteNamespacedReplicaSet**](ReplicaSetAPI.md#DeleteNamespacedReplicaSet) | **Delete** /apis/apps/v1/namespaces/{namespace}/replicasets/{name} | delete a ReplicaSet
[**DeletecollectionNamespacedReplicaSet**](ReplicaSetAPI.md#DeletecollectionNamespacedReplicaSet) | **Delete** /apis/apps/v1/namespaces/{namespace}/replicasets | delete collection of ReplicaSet
[**ListNamespacedReplicaSet**](ReplicaSetAPI.md#ListNamespacedReplicaSet) | **Get** /apis/apps/v1/namespaces/{namespace}/replicasets | list or watch objects of kind ReplicaSet
[**PatchNamespacedReplicaSet**](ReplicaSetAPI.md#PatchNamespacedReplicaSet) | **Patch** /apis/apps/v1/namespaces/{namespace}/replicasets/{name} | partially update the specified ReplicaSet
[**PatchNamespacedReplicaSetScale**](ReplicaSetAPI.md#PatchNamespacedReplicaSetScale) | **Patch** /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/scale | partially update scale of the specified ReplicaSet
[**PatchNamespacedReplicaSetStatus**](ReplicaSetAPI.md#PatchNamespacedReplicaSetStatus) | **Patch** /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/status | partially update status of the specified ReplicaSet
[**ReadNamespacedReplicaSet**](ReplicaSetAPI.md#ReadNamespacedReplicaSet) | **Get** /apis/apps/v1/namespaces/{namespace}/replicasets/{name} | read the specified ReplicaSet
[**ReadNamespacedReplicaSetScale**](ReplicaSetAPI.md#ReadNamespacedReplicaSetScale) | **Get** /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/scale | read scale of the specified ReplicaSet
[**ReadNamespacedReplicaSetStatus**](ReplicaSetAPI.md#ReadNamespacedReplicaSetStatus) | **Get** /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/status | read status of the specified ReplicaSet
[**ReplaceNamespacedReplicaSet**](ReplicaSetAPI.md#ReplaceNamespacedReplicaSet) | **Put** /apis/apps/v1/namespaces/{namespace}/replicasets/{name} | replace the specified ReplicaSet
[**ReplaceNamespacedReplicaSetScale**](ReplicaSetAPI.md#ReplaceNamespacedReplicaSetScale) | **Put** /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/scale | replace scale of the specified ReplicaSet
[**ReplaceNamespacedReplicaSetStatus**](ReplicaSetAPI.md#ReplaceNamespacedReplicaSetStatus) | **Put** /apis/apps/v1/namespaces/{namespace}/replicasets/{name}/status | replace status of the specified ReplicaSet



## CreateNamespacedReplicaSet

> V1ReplicaSet CreateNamespacedReplicaSet(ctx, namespace).Body(body).Pretty(pretty).Execute()

create a ReplicaSet

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
    body := *openapiclient.NewV1ReplicaSet() // V1ReplicaSet | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.CreateNamespacedReplicaSet(context.Background(), namespace).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.CreateNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNamespacedReplicaSet`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.CreateNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedReplicaSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1ReplicaSet**](V1ReplicaSet.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedReplicaSet

> V1Status DeleteNamespacedReplicaSet(ctx, namespace, name).Body(body).Pretty(pretty).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

delete a ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    body := *openapiclient.NewV1DeleteOptions() // V1DeleteOptions | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    gracePeriodSeconds := int32(56) // int32 | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. (optional)
    orphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. (optional)
    propagationPolicy := "propagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.DeleteNamespacedReplicaSet(context.Background(), namespace, name).Body(body).Pretty(pretty).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.DeleteNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNamespacedReplicaSet`: V1Status
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.DeleteNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedReplicaSetRequest struct via the builder pattern


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


## DeletecollectionNamespacedReplicaSet

> V1Status DeletecollectionNamespacedReplicaSet(ctx, namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()

delete collection of ReplicaSet

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
    resp, r, err := apiClient.ReplicaSetAPI.DeletecollectionNamespacedReplicaSet(context.Background(), namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.DeletecollectionNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletecollectionNamespacedReplicaSet`: V1Status
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.DeletecollectionNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletecollectionNamespacedReplicaSetRequest struct via the builder pattern


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


## ListNamespacedReplicaSet

> V1ReplicaSetList ListNamespacedReplicaSet(ctx, namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()

list or watch objects of kind ReplicaSet

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
    resp, r, err := apiClient.ReplicaSetAPI.ListNamespacedReplicaSet(context.Background(), namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ListNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNamespacedReplicaSet`: V1ReplicaSetList
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ListNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedReplicaSetRequest struct via the builder pattern


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

[**V1ReplicaSetList**](V1ReplicaSetList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedReplicaSet

> V1ReplicaSet PatchNamespacedReplicaSet(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

partially update the specified ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.PatchNamespacedReplicaSet(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.PatchNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNamespacedReplicaSet`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.PatchNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedReplicaSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedReplicaSetScale

> V1Scale PatchNamespacedReplicaSetScale(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

partially update scale of the specified ReplicaSet

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
    name := "name_example" // string | name of the Scale
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.PatchNamespacedReplicaSetScale(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.PatchNamespacedReplicaSetScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNamespacedReplicaSetScale`: V1Scale
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.PatchNamespacedReplicaSetScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Scale | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedReplicaSetScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1Scale**](V1Scale.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedReplicaSetStatus

> V1ReplicaSet PatchNamespacedReplicaSetStatus(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

partially update status of the specified ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.PatchNamespacedReplicaSetStatus(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.PatchNamespacedReplicaSetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNamespacedReplicaSetStatus`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.PatchNamespacedReplicaSetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedReplicaSetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedReplicaSet

> V1ReplicaSet ReadNamespacedReplicaSet(ctx, namespace, name).Pretty(pretty).Export(export).Exact(exact).Execute()

read the specified ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    export := true // bool | Should this value be exported.  Export strips fields that a user can not specify. (optional)
    exact := true // bool | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.ReadNamespacedReplicaSet(context.Background(), namespace, name).Pretty(pretty).Export(export).Exact(exact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ReadNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNamespacedReplicaSet`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ReadNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedReplicaSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **export** | **bool** | Should this value be exported.  Export strips fields that a user can not specify. | 
 **exact** | **bool** | Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedReplicaSetScale

> V1Scale ReadNamespacedReplicaSetScale(ctx, namespace, name).Pretty(pretty).Execute()

read scale of the specified ReplicaSet

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
    name := "name_example" // string | name of the Scale
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.ReadNamespacedReplicaSetScale(context.Background(), namespace, name).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ReadNamespacedReplicaSetScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNamespacedReplicaSetScale`: V1Scale
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ReadNamespacedReplicaSetScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Scale | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedReplicaSetScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1Scale**](V1Scale.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedReplicaSetStatus

> V1ReplicaSet ReadNamespacedReplicaSetStatus(ctx, namespace, name).Pretty(pretty).Execute()

read status of the specified ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.ReadNamespacedReplicaSetStatus(context.Background(), namespace, name).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ReadNamespacedReplicaSetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNamespacedReplicaSetStatus`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ReadNamespacedReplicaSetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedReplicaSetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedReplicaSet

> V1ReplicaSet ReplaceNamespacedReplicaSet(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

replace the specified ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    body := *openapiclient.NewV1ReplicaSet() // V1ReplicaSet | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.ReplaceNamespacedReplicaSet(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ReplaceNamespacedReplicaSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNamespacedReplicaSet`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ReplaceNamespacedReplicaSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedReplicaSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1ReplicaSet**](V1ReplicaSet.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedReplicaSetScale

> V1Scale ReplaceNamespacedReplicaSetScale(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

replace scale of the specified ReplicaSet

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
    name := "name_example" // string | name of the Scale
    body := *openapiclient.NewV1Scale() // V1Scale | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.ReplaceNamespacedReplicaSetScale(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ReplaceNamespacedReplicaSetScale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNamespacedReplicaSetScale`: V1Scale
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ReplaceNamespacedReplicaSetScale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Scale | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedReplicaSetScaleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Scale**](V1Scale.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1Scale**](V1Scale.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedReplicaSetStatus

> V1ReplicaSet ReplaceNamespacedReplicaSetStatus(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

replace status of the specified ReplicaSet

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
    name := "name_example" // string | name of the ReplicaSet
    body := *openapiclient.NewV1ReplicaSet() // V1ReplicaSet | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicaSetAPI.ReplaceNamespacedReplicaSetStatus(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicaSetAPI.ReplaceNamespacedReplicaSetStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNamespacedReplicaSetStatus`: V1ReplicaSet
    fmt.Fprintf(os.Stdout, "Response from `ReplicaSetAPI.ReplaceNamespacedReplicaSetStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the ReplicaSet | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedReplicaSetStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1ReplicaSet**](V1ReplicaSet.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1ReplicaSet**](V1ReplicaSet.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

