# \VolumesAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AttachVolumeToServer**](VolumesAPI.md#AttachVolumeToServer) | **Patch** /regions/:region/volumes/attach | 
[**CreateOSVolumeFromSnapshot**](VolumesAPI.md#CreateOSVolumeFromSnapshot) | **Post** /regions/:region/volumes/snapshots/:id/os-volume | 
[**CreateVolume**](VolumesAPI.md#CreateVolume) | **Post** /regions/:region/volumes | 
[**CreateVolumeFromSnapshot**](VolumesAPI.md#CreateVolumeFromSnapshot) | **Post** /regions/:region/volumes/snapshots/:id/create-volume | 
[**DeleteVolume**](VolumesAPI.md#DeleteVolume) | **Delete** /regions/:region/volumes/:id | 
[**DeleteVolumeSnapshot**](VolumesAPI.md#DeleteVolumeSnapshot) | **Delete** /regions/:region/volumes/:id/snapshot | 
[**DetachVolumeFromServer**](VolumesAPI.md#DetachVolumeFromServer) | **Patch** /regions/:region/volumes/detach | 
[**GetAllOSVolumes**](VolumesAPI.md#GetAllOSVolumes) | **Get** /regions/:region/volumes/os-volumes | 
[**GetAllVolumes**](VolumesAPI.md#GetAllVolumes) | **Get** /regions/:region/volumes | 
[**GetVolumeLimits**](VolumesAPI.md#GetVolumeLimits) | **Get** /regions/:region/volumes/limits | 
[**GetVolumeOptions**](VolumesAPI.md#GetVolumeOptions) | **Get** /regions/:region/volumes/options | 
[**GetVolumeSnapshots**](VolumesAPI.md#GetVolumeSnapshots) | **Get** /regions/:region/volumes/snapshots | 
[**RevertVolumeSnapshot**](VolumesAPI.md#RevertVolumeSnapshot) | **Put** /regions/:region/volumes/:id/snapshot/revert | 
[**UpdateVolume**](VolumesAPI.md#UpdateVolume) | **Patch** /regions/:region/volumes/:id | 
[**UpdateVolumeSnapshot**](VolumesAPI.md#UpdateVolumeSnapshot) | **Put** /regions/:region/volumes/:id/snapshot | 



## AttachVolumeToServer

> MessageResponse AttachVolumeToServer(ctx, region).VolumeAttachDetachRequest(volumeAttachDetachRequest).Execute()





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
    volumeAttachDetachRequest := *openapiclient.NewVolumeAttachDetachRequest() // VolumeAttachDetachRequest | volume and instance id for attachment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.AttachVolumeToServer(context.Background(), region).VolumeAttachDetachRequest(volumeAttachDetachRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.AttachVolumeToServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachVolumeToServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.AttachVolumeToServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachVolumeToServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeAttachDetachRequest** | [**VolumeAttachDetachRequest**](VolumeAttachDetachRequest.md) | volume and instance id for attachment | 

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


## CreateOSVolumeFromSnapshot

> MessageResponse CreateOSVolumeFromSnapshot(ctx, region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()





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
    createVolumeFromSnapshotRequest := *openapiclient.NewCreateVolumeFromSnapshotRequest() // CreateVolumeFromSnapshotRequest | new volume info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.CreateOSVolumeFromSnapshot(context.Background(), region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.CreateOSVolumeFromSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOSVolumeFromSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.CreateOSVolumeFromSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOSVolumeFromSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVolumeFromSnapshotRequest** | [**CreateVolumeFromSnapshotRequest**](CreateVolumeFromSnapshotRequest.md) | new volume info | 

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


## CreateVolume

> Volume CreateVolume(ctx, region, id).CreateVolumeRequest(createVolumeRequest).Execute()





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
    id := "id_example" // string | tag id
    createVolumeRequest := *openapiclient.NewCreateVolumeRequest() // CreateVolumeRequest | info about new volume to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.CreateVolume(context.Background(), region, id).CreateVolumeRequest(createVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.CreateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolume`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.CreateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | tag id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVolumeRequest** | [**CreateVolumeRequest**](CreateVolumeRequest.md) | info about new volume to be created | 

### Return type

[**Volume**](Volume.md)

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
    createVolumeFromSnapshotRequest := *openapiclient.NewCreateVolumeFromSnapshotRequest() // CreateVolumeFromSnapshotRequest | new volume info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.CreateVolumeFromSnapshot(context.Background(), region, id).CreateVolumeFromSnapshotRequest(createVolumeFromSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.CreateVolumeFromSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVolumeFromSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.CreateVolumeFromSnapshot`: %v\n", resp)
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


 **createVolumeFromSnapshotRequest** | [**CreateVolumeFromSnapshotRequest**](CreateVolumeFromSnapshotRequest.md) | new volume info | 

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


## DeleteVolume

> MessageResponse DeleteVolume(ctx, region, id).Execute()





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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.DeleteVolume(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.DeleteVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVolume`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## DeleteVolumeSnapshot

> MessageResponse DeleteVolumeSnapshot(ctx, region, id).Execute()





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
    resp, r, err := apiClient.VolumesAPI.DeleteVolumeSnapshot(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.DeleteVolumeSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVolumeSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.DeleteVolumeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeSnapshotRequest struct via the builder pattern


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


## DetachVolumeFromServer

> MessageResponse DetachVolumeFromServer(ctx, region).VolumeAttachDetachRequest(volumeAttachDetachRequest).Execute()





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
    volumeAttachDetachRequest := *openapiclient.NewVolumeAttachDetachRequest() // VolumeAttachDetachRequest | volume and instance id for attachment

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.DetachVolumeFromServer(context.Background(), region).VolumeAttachDetachRequest(volumeAttachDetachRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.DetachVolumeFromServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachVolumeFromServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.DetachVolumeFromServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachVolumeFromServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **volumeAttachDetachRequest** | [**VolumeAttachDetachRequest**](VolumeAttachDetachRequest.md) | volume and instance id for attachment | 

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


## GetAllOSVolumes

> Volume GetAllOSVolumes(ctx, region).Execute()





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
    resp, r, err := apiClient.VolumesAPI.GetAllOSVolumes(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetAllOSVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllOSVolumes`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetAllOSVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllOSVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllVolumes

> Volume GetAllVolumes(ctx, region).Execute()





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
    resp, r, err := apiClient.VolumesAPI.GetAllVolumes(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetAllVolumes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllVolumes`: Volume
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetAllVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Volume**](Volume.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeLimits

> VolumeLimits GetVolumeLimits(ctx, region).Execute()





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
    resp, r, err := apiClient.VolumesAPI.GetVolumeLimits(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolumeLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumeLimits`: VolumeLimits
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolumeLimits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeLimitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeLimits**](VolumeLimits.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeOptions

> VolumeOptions GetVolumeOptions(ctx, region).Execute()





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
    resp, r, err := apiClient.VolumesAPI.GetVolumeOptions(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolumeOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVolumeOptions`: VolumeOptions
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolumeOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VolumeOptions**](VolumeOptions.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeSnapshots

> GetVolumeSnapshots(ctx, region).Execute()





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
    r, err := apiClient.VolumesAPI.GetVolumeSnapshots(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolumeSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevertVolumeSnapshot

> MessageResponse RevertVolumeSnapshot(ctx, region, id).RevertSnapshotRequest(revertSnapshotRequest).Execute()





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
    revertSnapshotRequest := *openapiclient.NewRevertSnapshotRequest() // RevertSnapshotRequest | new volume info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.RevertVolumeSnapshot(context.Background(), region, id).RevertSnapshotRequest(revertSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.RevertVolumeSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RevertVolumeSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.RevertVolumeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevertVolumeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **revertSnapshotRequest** | [**RevertSnapshotRequest**](RevertSnapshotRequest.md) | new volume info | 

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


## UpdateVolume

> MessageResponse UpdateVolume(ctx, region, id).UpdateVolumeRequest(updateVolumeRequest).Execute()





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
    updateVolumeRequest := *openapiclient.NewUpdateVolumeRequest() // UpdateVolumeRequest | new volume info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.UpdateVolume(context.Background(), region, id).UpdateVolumeRequest(updateVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.UpdateVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolume`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.UpdateVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | volume id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVolumeRequest** | [**UpdateVolumeRequest**](UpdateVolumeRequest.md) | new volume info | 

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


## UpdateVolumeSnapshot

> MessageResponse UpdateVolumeSnapshot(ctx, region, id).UpdateVolumeSnapshotRequest(updateVolumeSnapshotRequest).Execute()





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
    updateVolumeSnapshotRequest := *openapiclient.NewUpdateVolumeSnapshotRequest() // UpdateVolumeSnapshotRequest | new info for snapshot

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumesAPI.UpdateVolumeSnapshot(context.Background(), region, id).UpdateVolumeSnapshotRequest(updateVolumeSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.UpdateVolumeSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVolumeSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.UpdateVolumeSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | snapshot id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVolumeSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVolumeSnapshotRequest** | [**UpdateVolumeSnapshotRequest**](UpdateVolumeSnapshotRequest.md) | new info for snapshot | 

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

