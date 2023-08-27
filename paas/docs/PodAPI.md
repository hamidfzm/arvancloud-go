# \PodAPI

All URIs are relative to *https://napi.arvancloud.ir/paas/v1/regions/ir-thr-at1/o*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectGetNamespacedPodExec**](PodAPI.md#ConnectGetNamespacedPodExec) | **Get** /api/v1/namespaces/{namespace}/pods/{name}/exec | connect GET requests to exec of Pod
[**ConnectPostNamespacedPodExec**](PodAPI.md#ConnectPostNamespacedPodExec) | **Post** /api/v1/namespaces/{namespace}/pods/{name}/exec | connect POST requests to exec of Pod
[**CreateNamespacedPod**](PodAPI.md#CreateNamespacedPod) | **Post** /api/v1/namespaces/{namespace}/pods | create a Pod
[**DeleteNamespacedPod**](PodAPI.md#DeleteNamespacedPod) | **Delete** /api/v1/namespaces/{namespace}/pods/{name} | delete a Pod
[**DeletecollectionNamespacedPod**](PodAPI.md#DeletecollectionNamespacedPod) | **Delete** /api/v1/namespaces/{namespace}/pods | delete collection of Pod
[**ListNamespacedPod**](PodAPI.md#ListNamespacedPod) | **Get** /api/v1/namespaces/{namespace}/pods | list or watch objects of kind Pod
[**PatchNamespacedPod**](PodAPI.md#PatchNamespacedPod) | **Patch** /api/v1/namespaces/{namespace}/pods/{name} | partially update the specified Pod
[**ReadNamespacedPod**](PodAPI.md#ReadNamespacedPod) | **Get** /api/v1/namespaces/{namespace}/pods/{name} | read the specified Pod
[**ReadNamespacedPodLog**](PodAPI.md#ReadNamespacedPodLog) | **Get** /api/v1/namespaces/{namespace}/pods/{name}/log | read log of the specified Pod
[**ReplaceNamespacedPod**](PodAPI.md#ReplaceNamespacedPod) | **Put** /api/v1/namespaces/{namespace}/pods/{name} | replace the specified Pod



## ConnectGetNamespacedPodExec

> string ConnectGetNamespacedPodExec(ctx, namespace, name).Stdin(stdin).Stdout(stdout).Stderr(stderr).Tty(tty).Container(container).Command(command).Execute()

connect GET requests to exec of Pod

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
    name := "name_example" // string | name of the Pod
    stdin := true // bool | Redirect the standard input stream of the pod for this call. Defaults to false. (optional)
    stdout := true // bool | Redirect the standard output stream of the pod for this call. Defaults to true. (optional)
    stderr := true // bool | Redirect the standard error stream of the pod for this call. Defaults to true. (optional)
    tty := true // bool | TTY if true indicates that a tty will be allocated for the exec call. Defaults to false. (optional)
    container := "container_example" // string | Container in which to execute the command. Defaults to only container if there is only one container in the pod. (optional)
    command := "command_example" // string | Command is the remote command to execute. argv array. Not executed within a shell. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.ConnectGetNamespacedPodExec(context.Background(), namespace, name).Stdin(stdin).Stdout(stdout).Stderr(stderr).Tty(tty).Container(container).Command(command).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.ConnectGetNamespacedPodExec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectGetNamespacedPodExec`: string
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.ConnectGetNamespacedPodExec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectGetNamespacedPodExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stdin** | **bool** | Redirect the standard input stream of the pod for this call. Defaults to false. | 
 **stdout** | **bool** | Redirect the standard output stream of the pod for this call. Defaults to true. | 
 **stderr** | **bool** | Redirect the standard error stream of the pod for this call. Defaults to true. | 
 **tty** | **bool** | TTY if true indicates that a tty will be allocated for the exec call. Defaults to false. | 
 **container** | **string** | Container in which to execute the command. Defaults to only container if there is only one container in the pod. | 
 **command** | **string** | Command is the remote command to execute. argv array. Not executed within a shell. | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ConnectPostNamespacedPodExec

> string ConnectPostNamespacedPodExec(ctx, namespace, name).Stdin(stdin).Stdout(stdout).Stderr(stderr).Tty(tty).Container(container).Command(command).Execute()

connect POST requests to exec of Pod

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
    name := "name_example" // string | name of the Pod
    stdin := true // bool | Redirect the standard input stream of the pod for this call. Defaults to false. (optional)
    stdout := true // bool | Redirect the standard output stream of the pod for this call. Defaults to true. (optional)
    stderr := true // bool | Redirect the standard error stream of the pod for this call. Defaults to true. (optional)
    tty := true // bool | TTY if true indicates that a tty will be allocated for the exec call. Defaults to false. (optional)
    container := "container_example" // string | Container in which to execute the command. Defaults to only container if there is only one container in the pod. (optional)
    command := "command_example" // string | Command is the remote command to execute. argv array. Not executed within a shell. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.ConnectPostNamespacedPodExec(context.Background(), namespace, name).Stdin(stdin).Stdout(stdout).Stderr(stderr).Tty(tty).Container(container).Command(command).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.ConnectPostNamespacedPodExec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ConnectPostNamespacedPodExec`: string
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.ConnectPostNamespacedPodExec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectPostNamespacedPodExecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stdin** | **bool** | Redirect the standard input stream of the pod for this call. Defaults to false. | 
 **stdout** | **bool** | Redirect the standard output stream of the pod for this call. Defaults to true. | 
 **stderr** | **bool** | Redirect the standard error stream of the pod for this call. Defaults to true. | 
 **tty** | **bool** | TTY if true indicates that a tty will be allocated for the exec call. Defaults to false. | 
 **container** | **string** | Container in which to execute the command. Defaults to only container if there is only one container in the pod. | 
 **command** | **string** | Command is the remote command to execute. argv array. Not executed within a shell. | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNamespacedPod

> V1Pod CreateNamespacedPod(ctx, namespace).Body(body).Pretty(pretty).Execute()

create a Pod

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
    body := *openapiclient.NewV1Pod() // V1Pod | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.CreateNamespacedPod(context.Background(), namespace).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.CreateNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNamespacedPod`: V1Pod
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.CreateNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNamespacedPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**V1Pod**](V1Pod.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1Pod**](V1Pod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNamespacedPod

> V1Status DeleteNamespacedPod(ctx, namespace, name).Body(body).Pretty(pretty).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()

delete a Pod

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
    name := "name_example" // string | name of the Pod
    body := *openapiclient.NewV1DeleteOptions() // V1DeleteOptions | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    gracePeriodSeconds := int32(56) // int32 | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately. (optional)
    orphanDependents := true // bool | Deprecated: please use the PropagationPolicy, this field will be deprecated in 1.7. Should the dependent objects be orphaned. If true/false, the \"orphan\" finalizer will be added to/removed from the object's finalizers list. Either this field or PropagationPolicy may be set, but not both. (optional)
    propagationPolicy := "propagationPolicy_example" // string | Whether and how garbage collection will be performed. Either this field or OrphanDependents may be set, but not both. The default policy is decided by the existing finalizer set in the metadata.finalizers and the resource-specific default policy. Acceptable values are: 'Orphan' - orphan the dependents; 'Background' - allow the garbage collector to delete the dependents in the background; 'Foreground' - a cascading policy that deletes all dependents in the foreground. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.DeleteNamespacedPod(context.Background(), namespace, name).Body(body).Pretty(pretty).GracePeriodSeconds(gracePeriodSeconds).OrphanDependents(orphanDependents).PropagationPolicy(propagationPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.DeleteNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNamespacedPod`: V1Status
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.DeleteNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNamespacedPodRequest struct via the builder pattern


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


## DeletecollectionNamespacedPod

> V1Status DeletecollectionNamespacedPod(ctx, namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()

delete collection of Pod

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
    resp, r, err := apiClient.PodAPI.DeletecollectionNamespacedPod(context.Background(), namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.DeletecollectionNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletecollectionNamespacedPod`: V1Status
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.DeletecollectionNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletecollectionNamespacedPodRequest struct via the builder pattern


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


## ListNamespacedPod

> V1PodList ListNamespacedPod(ctx, namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()

list or watch objects of kind Pod

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
    resp, r, err := apiClient.PodAPI.ListNamespacedPod(context.Background(), namespace).Pretty(pretty).LabelSelector(labelSelector).FieldSelector(fieldSelector).IncludeUninitialized(includeUninitialized).Watch(watch).ResourceVersion(resourceVersion).TimeoutSeconds(timeoutSeconds).Limit(limit).Continue_(continue_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.ListNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNamespacedPod`: V1PodList
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.ListNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNamespacedPodRequest struct via the builder pattern


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

[**V1PodList**](V1PodList.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf, application/json;stream=watch, application/vnd.kubernetes.protobuf;stream=watch

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchNamespacedPod

> V1Pod PatchNamespacedPod(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

partially update the specified Pod

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
    name := "name_example" // string | name of the Pod
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.PatchNamespacedPod(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.PatchNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchNamespacedPod`: V1Pod
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.PatchNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchNamespacedPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1Pod**](V1Pod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json, application/merge-patch+json, application/strategic-merge-patch+json
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedPod

> V1Pod ReadNamespacedPod(ctx, namespace, name).Pretty(pretty).Export(export).Exact(exact).Execute()

read the specified Pod

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
    name := "name_example" // string | name of the Pod
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    export := true // bool | Should this value be exported.  Export strips fields that a user can not specify. (optional)
    exact := true // bool | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.ReadNamespacedPod(context.Background(), namespace, name).Pretty(pretty).Export(export).Exact(exact).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.ReadNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNamespacedPod`: V1Pod
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.ReadNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **export** | **bool** | Should this value be exported.  Export strips fields that a user can not specify. | 
 **exact** | **bool** | Should the export be exact.  Exact export maintains cluster-specific fields like &#39;Namespace&#39;. | 

### Return type

[**V1Pod**](V1Pod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadNamespacedPodLog

> string ReadNamespacedPodLog(ctx, namespace, name).Pretty(pretty).Container(container).Follow(follow).Previous(previous).SinceSeconds(sinceSeconds).Timestamps(timestamps).TailLines(tailLines).LimitBytes(limitBytes).Execute()

read log of the specified Pod

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
    name := "name_example" // string | name of the Pod
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)
    container := "container_example" // string | The container for which to stream logs. Defaults to only container if there is one container in the pod. (optional)
    follow := true // bool | Follow the log stream of the pod. Defaults to false. (optional)
    previous := true // bool | Return previous terminated container logs. Defaults to false. (optional)
    sinceSeconds := int32(56) // int32 | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. (optional)
    timestamps := true // bool | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. (optional)
    tailLines := int32(56) // int32 | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime (optional)
    limitBytes := int32(56) // int32 | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.ReadNamespacedPodLog(context.Background(), namespace, name).Pretty(pretty).Container(container).Follow(follow).Previous(previous).SinceSeconds(sinceSeconds).Timestamps(timestamps).TailLines(tailLines).LimitBytes(limitBytes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.ReadNamespacedPodLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadNamespacedPodLog`: string
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.ReadNamespacedPodLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadNamespacedPodLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 
 **container** | **string** | The container for which to stream logs. Defaults to only container if there is one container in the pod. | 
 **follow** | **bool** | Follow the log stream of the pod. Defaults to false. | 
 **previous** | **bool** | Return previous terminated container logs. Defaults to false. | 
 **sinceSeconds** | **int32** | A relative time in seconds before the current time from which to show logs. If this value precedes the time a pod was started, only logs since the pod start will be returned. If this value is in the future, no logs will be returned. Only one of sinceSeconds or sinceTime may be specified. | 
 **timestamps** | **bool** | If true, add an RFC3339 or RFC3339Nano timestamp at the beginning of every line of log output. Defaults to false. | 
 **tailLines** | **int32** | If set, the number of lines from the end of the logs to show. If not specified, logs are shown from the creation of the container or sinceSeconds or sinceTime | 
 **limitBytes** | **int32** | If set, the number of bytes to read from the server before terminating the log output. This may not display a complete final line of logging, and may return slightly more or slightly less than the specified limit. | 

### Return type

**string**

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceNamespacedPod

> V1Pod ReplaceNamespacedPod(ctx, namespace, name).Body(body).Pretty(pretty).Execute()

replace the specified Pod

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
    name := "name_example" // string | name of the Pod
    body := *openapiclient.NewV1Pod() // V1Pod | 
    pretty := "pretty_example" // string | If 'true', then the output is pretty printed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PodAPI.ReplaceNamespacedPod(context.Background(), namespace, name).Body(body).Pretty(pretty).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PodAPI.ReplaceNamespacedPod``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplaceNamespacedPod`: V1Pod
    fmt.Fprintf(os.Stdout, "Response from `PodAPI.ReplaceNamespacedPod`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | object name and auth scope, such as for teams and projects | 
**name** | **string** | name of the Pod | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceNamespacedPodRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**V1Pod**](V1Pod.md) |  | 
 **pretty** | **string** | If &#39;true&#39;, then the output is pretty printed. | 

### Return type

[**V1Pod**](V1Pod.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/yaml, application/vnd.kubernetes.protobuf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

