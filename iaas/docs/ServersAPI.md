# \ServersAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddServerPublicIP**](ServersAPI.md#AddServerPublicIP) | **Post** /regions/:region/servers/:id/add-public-ip | 
[**AttachServerRootVolume**](ServersAPI.md#AttachServerRootVolume) | **Put** /regions/:region/servers/:id/attachRoot | 
[**AttachServerToSecurityGroup**](ServersAPI.md#AttachServerToSecurityGroup) | **Post** /regions/:region/servers/:id/add-security-group | 
[**CreateServer**](ServersAPI.md#CreateServer) | **Post** /regions/:region/servers | 
[**DeleteServer**](ServersAPI.md#DeleteServer) | **Delete** /regions/:region/servers/:id | 
[**DetachServerFromSecurityGroup**](ServersAPI.md#DetachServerFromSecurityGroup) | **Post** /regions/:region/servers/:id/remove-security-group | 
[**DetachServerRootVolume**](ServersAPI.md#DetachServerRootVolume) | **Put** /regions/:region/servers/:id/detachRoot | 
[**GetAllServers**](ServersAPI.md#GetAllServers) | **Get** /regions/:region/servers | 
[**GetDeleteReasons**](ServersAPI.md#GetDeleteReasons) | **Get** /regions/:region/servers/delete-reasons | 
[**GetRegionServerCreationOptions**](ServersAPI.md#GetRegionServerCreationOptions) | **Get** /regions/:region/servers/options | 
[**GetServerAvailableActions**](ServersAPI.md#GetServerAvailableActions) | **Get** /regions/:region/servers/:id/actions | 
[**GetServerDetails**](ServersAPI.md#GetServerDetails) | **Get** /regions/:region/servers/:id | 
[**GetServerVNC**](ServersAPI.md#GetServerVNC) | **Get** /regions/:region/servers/:id/vnc | 
[**HardRebootServer**](ServersAPI.md#HardRebootServer) | **Post** /regions/:region/servers/:id/hard-reboot | 
[**PowerOffServer**](ServersAPI.md#PowerOffServer) | **Post** /regions/:region/servers/:id/power-off | 
[**PowerOnServer**](ServersAPI.md#PowerOnServer) | **Post** /regions/:region/servers/:id/power-on | 
[**RebootServer**](ServersAPI.md#RebootServer) | **Post** /regions/:region/servers/:id/reboot | 
[**RebuildServer**](ServersAPI.md#RebuildServer) | **Post** /regions/:region/servers/:id/rebuild | 
[**RenameServer**](ServersAPI.md#RenameServer) | **Post** /regions/:region/servers/:id/rename | 
[**RescueServer**](ServersAPI.md#RescueServer) | **Post** /regions/:region/servers/:id/rescue | 
[**ResetRootPassword**](ServersAPI.md#ResetRootPassword) | **Post** /regions/:region/servers/:id/reset-root-password | 
[**ResizeServer**](ServersAPI.md#ResizeServer) | **Post** /regions/:region/servers/:id/resize | 
[**ResizeServerRootVolume**](ServersAPI.md#ResizeServerRootVolume) | **Put** /regions/:region/servers/:id/resizeRoot | 
[**TakeServerSnapshot**](ServersAPI.md#TakeServerSnapshot) | **Post** /regions/:region/servers/:id/snapshot | 
[**ToggleInstanceHA**](ServersAPI.md#ToggleInstanceHA) | **Post** /regions/:region/servers/:id/instance-ha/:action | 
[**UnrescueServer**](ServersAPI.md#UnrescueServer) | **Post** /regions/:region/servers/:id/unrescue | 



## AddServerPublicIP

> MessageResponse AddServerPublicIP(ctx, region, id).AddServerPublicIPRequest(addServerPublicIPRequest).Execute()





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
    id := "id_example" // string | server id
    addServerPublicIPRequest := *openapiclient.NewAddServerPublicIPRequest() // AddServerPublicIPRequest | security groups to be applied to newly added public ip

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.AddServerPublicIP(context.Background(), region, id).AddServerPublicIPRequest(addServerPublicIPRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.AddServerPublicIP``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddServerPublicIP`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.AddServerPublicIP`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddServerPublicIPRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addServerPublicIPRequest** | [**AddServerPublicIPRequest**](AddServerPublicIPRequest.md) | security groups to be applied to newly added public ip | 

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


## AttachServerRootVolume

> MessageResponse AttachServerRootVolume(ctx, region, id).AttachServerRootVolumeRequest(attachServerRootVolumeRequest).Execute()





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
    id := "id_example" // string | server id
    attachServerRootVolumeRequest := *openapiclient.NewAttachServerRootVolumeRequest() // AttachServerRootVolumeRequest | volume id to be attached as root

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.AttachServerRootVolume(context.Background(), region, id).AttachServerRootVolumeRequest(attachServerRootVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.AttachServerRootVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachServerRootVolume`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.AttachServerRootVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachServerRootVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachServerRootVolumeRequest** | [**AttachServerRootVolumeRequest**](AttachServerRootVolumeRequest.md) | volume id to be attached as root | 

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


## AttachServerToSecurityGroup

> MessageResponse AttachServerToSecurityGroup(ctx, region, id).AttachServerToSecurityGroupRequest(attachServerToSecurityGroupRequest).Execute()





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
    id := "id_example" // string | server id
    attachServerToSecurityGroupRequest := *openapiclient.NewAttachServerToSecurityGroupRequest() // AttachServerToSecurityGroupRequest | security group info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.AttachServerToSecurityGroup(context.Background(), region, id).AttachServerToSecurityGroupRequest(attachServerToSecurityGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.AttachServerToSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AttachServerToSecurityGroup`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.AttachServerToSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAttachServerToSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachServerToSecurityGroupRequest** | [**AttachServerToSecurityGroupRequest**](AttachServerToSecurityGroupRequest.md) | security group info | 

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


## CreateServer

> ServerDetail CreateServer(ctx, region).CreateServerRequest(createServerRequest).Execute()





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
    createServerRequest := *openapiclient.NewCreateServerRequest() // CreateServerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.CreateServer(context.Background(), region).CreateServerRequest(createServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.CreateServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServer`: ServerDetail
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.CreateServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createServerRequest** | [**CreateServerRequest**](CreateServerRequest.md) |  | 

### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServer

> MessageResponse DeleteServer(ctx, region, id).ForceDelete(forceDelete).DeleteServerReasons(deleteServerReasons).Execute()





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
    id := "id_example" // string | server id
    forceDelete := true // bool | Force delete server. If true, the server will be forcibly deleted (optional)
    deleteServerReasons := map[string]interface{}{ ... } // map[string]interface{} | The reason why this server is being deleted (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.DeleteServer(context.Background(), region, id).ForceDelete(forceDelete).DeleteServerReasons(deleteServerReasons).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DeleteServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.DeleteServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **forceDelete** | **bool** | Force delete server. If true, the server will be forcibly deleted | 
 **deleteServerReasons** | [**map[string]interface{}**](map[string]interface{}.md) | The reason why this server is being deleted | 

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


## DetachServerFromSecurityGroup

> MessageResponse DetachServerFromSecurityGroup(ctx, region, id).AttachServerToSecurityGroupRequest(attachServerToSecurityGroupRequest).Execute()





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
    id := "id_example" // string | server id
    attachServerToSecurityGroupRequest := *openapiclient.NewAttachServerToSecurityGroupRequest() // AttachServerToSecurityGroupRequest | security group info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.DetachServerFromSecurityGroup(context.Background(), region, id).AttachServerToSecurityGroupRequest(attachServerToSecurityGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DetachServerFromSecurityGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachServerFromSecurityGroup`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.DetachServerFromSecurityGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachServerFromSecurityGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attachServerToSecurityGroupRequest** | [**AttachServerToSecurityGroupRequest**](AttachServerToSecurityGroupRequest.md) | security group info | 

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


## DetachServerRootVolume

> MessageResponse DetachServerRootVolume(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.DetachServerRootVolume(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.DetachServerRootVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DetachServerRootVolume`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.DetachServerRootVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDetachServerRootVolumeRequest struct via the builder pattern


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


## GetAllServers

> ServerDetail GetAllServers(ctx, region).Execute()





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
    resp, r, err := apiClient.ServersAPI.GetAllServers(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetAllServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllServers`: ServerDetail
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetAllServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeleteReasons

> DeleteServerReasons GetDeleteReasons(ctx, region).Execute()





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
    resp, r, err := apiClient.ServersAPI.GetDeleteReasons(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetDeleteReasons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeleteReasons`: DeleteServerReasons
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetDeleteReasons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeleteReasonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteServerReasons**](DeleteServerReasons.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegionServerCreationOptions

> RegionServerCreateOptions GetRegionServerCreationOptions(ctx, region).Execute()





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
    resp, r, err := apiClient.ServersAPI.GetRegionServerCreationOptions(context.Background(), region).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetRegionServerCreationOptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRegionServerCreationOptions`: RegionServerCreateOptions
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetRegionServerCreationOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionServerCreationOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegionServerCreateOptions**](RegionServerCreateOptions.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerAvailableActions

> GetServerActionsResponse GetServerAvailableActions(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerAvailableActions(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerAvailableActions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerAvailableActions`: GetServerActionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerAvailableActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerAvailableActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServerActionsResponse**](GetServerActionsResponse.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerDetails

> ServerDetail GetServerDetails(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerDetails(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerDetails`: ServerDetail
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServerVNC

> GetServerVNCRequest GetServerVNC(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.GetServerVNC(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.GetServerVNC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServerVNC`: GetServerVNCRequest
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.GetServerVNC`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServerVNCRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetServerVNCRequest**](GetServerVNCRequest.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HardRebootServer

> MessageResponse HardRebootServer(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.HardRebootServer(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.HardRebootServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `HardRebootServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.HardRebootServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiHardRebootServerRequest struct via the builder pattern


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


## PowerOffServer

> MessageResponse PowerOffServer(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.PowerOffServer(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.PowerOffServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PowerOffServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.PowerOffServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPowerOffServerRequest struct via the builder pattern


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


## PowerOnServer

> MessageResponse PowerOnServer(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.PowerOnServer(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.PowerOnServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PowerOnServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.PowerOnServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiPowerOnServerRequest struct via the builder pattern


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


## RebootServer

> MessageResponse RebootServer(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.RebootServer(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.RebootServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.RebootServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootServerRequest struct via the builder pattern


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


## RebuildServer

> ServerDetail RebuildServer(ctx, region, id).RebuildImageRequest(rebuildImageRequest).Execute()





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
    id := "id_example" // string | server id
    rebuildImageRequest := *openapiclient.NewRebuildImageRequest() // RebuildImageRequest | image info for the newly built server

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.RebuildServer(context.Background(), region, id).RebuildImageRequest(rebuildImageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.RebuildServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebuildServer`: ServerDetail
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.RebuildServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebuildServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebuildImageRequest** | [**RebuildImageRequest**](RebuildImageRequest.md) | image info for the newly built server | 

### Return type

[**ServerDetail**](ServerDetail.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RenameServer

> MessageResponse RenameServer(ctx, region, id).RenameServerRequest(renameServerRequest).Execute()





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
    id := "id_example" // string | server id
    renameServerRequest := *openapiclient.NewRenameServerRequest() // RenameServerRequest | server's new name

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.RenameServer(context.Background(), region, id).RenameServerRequest(renameServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.RenameServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RenameServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.RenameServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **renameServerRequest** | [**RenameServerRequest**](RenameServerRequest.md) | server&#39;s new name | 

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


## RescueServer

> MessageResponse RescueServer(ctx, region, id).RescueServerRequest(rescueServerRequest).Execute()





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
    id := "id_example" // string | server id
    rescueServerRequest := *openapiclient.NewRescueServerRequest() // RescueServerRequest | rescue image info (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.RescueServer(context.Background(), region, id).RescueServerRequest(rescueServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.RescueServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RescueServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.RescueServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRescueServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rescueServerRequest** | [**RescueServerRequest**](RescueServerRequest.md) | rescue image info | 

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


## ResetRootPassword

> MessageResponse ResetRootPassword(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.ResetRootPassword(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ResetRootPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetRootPassword`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ResetRootPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetRootPasswordRequest struct via the builder pattern


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


## ResizeServer

> MessageResponse ResizeServer(ctx, region, id).ResizeServerRequest(resizeServerRequest).Execute()





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
    id := "id_example" // string | server id
    resizeServerRequest := *openapiclient.NewResizeServerRequest() // ResizeServerRequest | new flavor info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.ResizeServer(context.Background(), region, id).ResizeServerRequest(resizeServerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ResizeServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ResizeServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeServerRequest** | [**ResizeServerRequest**](ResizeServerRequest.md) | new flavor info | 

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


## ResizeServerRootVolume

> MessageResponse ResizeServerRootVolume(ctx, region, id).ResizeRootVolumeRequest(resizeRootVolumeRequest).Execute()





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
    id := "id_example" // string | server id
    resizeRootVolumeRequest := *openapiclient.NewResizeRootVolumeRequest() // ResizeRootVolumeRequest | New root volume size (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.ResizeServerRootVolume(context.Background(), region, id).ResizeRootVolumeRequest(resizeRootVolumeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ResizeServerRootVolume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeServerRootVolume`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ResizeServerRootVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeServerRootVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeRootVolumeRequest** | [**ResizeRootVolumeRequest**](ResizeRootVolumeRequest.md) | New root volume size | 

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


## TakeServerSnapshot

> MessageResponse TakeServerSnapshot(ctx, region, id).Payload(payload).Execute()





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
    id := "id_example" // string | server id
    payload := *openapiclient.NewSnapshotCreateRequest() // SnapshotCreateRequest | snapshot info

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.TakeServerSnapshot(context.Background(), region, id).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.TakeServerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TakeServerSnapshot`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.TakeServerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiTakeServerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **payload** | [**SnapshotCreateRequest**](SnapshotCreateRequest.md) | snapshot info | 

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


## ToggleInstanceHA

> MessageResponse ToggleInstanceHA(ctx, region, id, action).Execute()





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
    id := "id_example" // string | server id
    action := "action_example" // string | action on the feature, either `enable` or `disable`

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.ToggleInstanceHA(context.Background(), region, id, action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.ToggleInstanceHA``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ToggleInstanceHA`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.ToggleInstanceHA`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 
**action** | **string** | action on the feature, either &#x60;enable&#x60; or &#x60;disable&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleInstanceHARequest struct via the builder pattern


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


## UnrescueServer

> MessageResponse UnrescueServer(ctx, region, id).Execute()





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
    id := "id_example" // string | server id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ServersAPI.UnrescueServer(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServersAPI.UnrescueServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnrescueServer`: MessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ServersAPI.UnrescueServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | server id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnrescueServerRequest struct via the builder pattern


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

