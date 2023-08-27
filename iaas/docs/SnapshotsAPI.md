# \SnapshotsAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateImageFromSnapshot**](SnapshotsAPI.md#CreateImageFromSnapshot) | **Post** /regions/:region/snapshots/:id/images/ | 
[**CreateSnapshotFromServer**](SnapshotsAPI.md#CreateSnapshotFromServer) | **Post** /regions/:region/snapshots/servers/:id | 
[**CreateSnapshotFromVolume**](SnapshotsAPI.md#CreateSnapshotFromVolume) | **Post** /regions/:region/snapshots/volumes/:id/ | 
[**CreateVolumeFromSnapshot**](SnapshotsAPI.md#CreateVolumeFromSnapshot) | **Post** /regions/:region/snapshots/:id/volumes/ | 
[**DeleteSnapshot**](SnapshotsAPI.md#DeleteSnapshot) | **Delete** /regions/:region/snapshots/:id | 
[**GetAllSnapshots**](SnapshotsAPI.md#GetAllSnapshots) | **Get** /regions/:region/snapshots | 
[**RevertSnapshot**](SnapshotsAPI.md#RevertSnapshot) | **Put** /regions/:region/snapshots/:id/revert | 
[**UpdateSnapshot**](SnapshotsAPI.md#UpdateSnapshot) | **Put** /regions/:region/snapshots/:id | 



## CreateImageFromSnapshot

> MessageResponse CreateImageFromSnapshot(ctx, region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | snapshot id
    createVolumeFromSnapshotRequest := *openapiclient.NewCreateVolumeFromSnapshotRequest() // CreateVolumeFromSnapshotRequest | snapshot info to be reverted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.CreateImageFromSnapshot(context.Background(), region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateImageFromSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateImageFromSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateImageFromSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateImageFromSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVolumeFromSnapshotRequest** | [**CreateVolumeFromSnapshotRequest**](CreateVolumeFromSnapshotRequest.md) | snapshot info to be reverted | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshotFromServer

> MessageResponse CreateSnapshotFromServer(ctx, region, id).CreateSnapshotFromVolumeRequest(createSnapshotFromVolumeRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | volume id
    createSnapshotFromVolumeRequest := *openapiclient.NewCreateSnapshotFromVolumeRequest() // CreateSnapshotFromVolumeRequest | info about the snapshot to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.CreateSnapshotFromServer(context.Background(), region, id).CreateSnapshotFromVolumeRequest(createSnapshotFromVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateSnapshotFromServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshotFromServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateSnapshotFromServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotFromServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createSnapshotFromVolumeRequest** | [**CreateSnapshotFromVolumeRequest**](CreateSnapshotFromVolumeRequest.md) | info about the snapshot to be created | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSnapshotFromVolume

> MessageResponse CreateSnapshotFromVolume(ctx, region, id).CreateSnapshotFromVolumeRequest(createSnapshotFromVolumeRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | volume id
    createSnapshotFromVolumeRequest := *openapiclient.NewCreateSnapshotFromVolumeRequest() // CreateSnapshotFromVolumeRequest | info about the snapshot to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.CreateSnapshotFromVolume(context.Background(), region, id).CreateSnapshotFromVolumeRequest(createSnapshotFromVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateSnapshotFromVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSnapshotFromVolume`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateSnapshotFromVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotFromVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createSnapshotFromVolumeRequest** | [**CreateSnapshotFromVolumeRequest**](CreateSnapshotFromVolumeRequest.md) | info about the snapshot to be created | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVolumeFromSnapshot

> MessageResponse CreateVolumeFromSnapshot(ctx, region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | snapshot id
    createVolumeFromSnapshotRequest := *openapiclient.NewCreateVolumeFromSnapshotRequest() // CreateVolumeFromSnapshotRequest | snapshot info to be reverted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.CreateVolumeFromSnapshot(context.Background(), region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.CreateVolumeFromSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumeFromSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.CreateVolumeFromSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeFromSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVolumeFromSnapshotRequest** | [**CreateVolumeFromSnapshotRequest**](CreateVolumeFromSnapshotRequest.md) | snapshot info to be reverted | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSnapshot

> MessageResponse DeleteSnapshot(ctx, region, id).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | snapshot id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.DeleteSnapshot(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.DeleteSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.DeleteSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSnapshots

> VolumeSnapshot GetAllSnapshots(ctx, region).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.GetAllSnapshots(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.GetAllSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSnapshots`: VolumeSnapshot
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.GetAllSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeSnapshot**](VolumeSnapshot.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertSnapshot

> MessageResponse RevertSnapshot(ctx, region, id).RevertSnapshotRequest(revertSnapshotRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | server or volume id which the snapshot belongs to
    revertSnapshotRequest := *openapiclient.NewRevertSnapshotRequest() // RevertSnapshotRequest | snapshot info to be reverted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.RevertSnapshot(context.Background(), region, id).RevertSnapshotRequest(revertSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.RevertSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.RevertSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server or volume id which the snapshot belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revertSnapshotRequest** | [**RevertSnapshotRequest**](RevertSnapshotRequest.md) | snapshot info to be reverted | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSnapshot

> MessageResponse UpdateSnapshot(ctx, region, id).UpdateVolumeSnapshotRequest(updateVolumeSnapshotRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/hamidfzm/arvancloud-go/iaas"
)

func main() {
    region := "region_example" // string | region code
    id := "id_example" // string | snapshot id
    updateVolumeSnapshotRequest := *openapiclient.NewUpdateVolumeSnapshotRequest() // UpdateVolumeSnapshotRequest | new info for volume snapshot

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotsAPI.UpdateSnapshot(context.Background(), region, id).UpdateVolumeSnapshotRequest(updateVolumeSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotsAPI.UpdateSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotsAPI.UpdateSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVolumeSnapshotRequest** | [**UpdateVolumeSnapshotRequest**](UpdateVolumeSnapshotRequest.md) | new info for volume snapshot | 

### Return type

[**MessageResponse**](MessageResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

