# \PlansAPI

All URIs are relative to *https://napi.arvancloud.ir/ecc/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllRegionPlans**](PlansAPI.md#GetAllRegionPlans) | **Get** /regions/:region/sizes | 
[**GetPlanByID**](PlansAPI.md#GetPlanByID) | **Get** /regions/:region/sizes/:id | 



## GetAllRegionPlans

> Plan GetAllRegionPlans(ctx, region).Cpu(cpu).Ram(ram).Disk(disk).Execute()





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
    cpu := "cpu_example" // string | filters based on cpu core count if exists (optional)
    ram := "ram_example" // string | filters based on RAM MB if exists (optional)
    disk := "disk_example" // string | filters based on disk GB if exists (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansAPI.GetAllRegionPlans(context.Background(), region).Cpu(cpu).Ram(ram).Disk(disk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetAllRegionPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRegionPlans`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetAllRegionPlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRegionPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cpu** | **string** | filters based on cpu core count if exists | 
 **ram** | **string** | filters based on RAM MB if exists | 
 **disk** | **string** | filters based on disk GB if exists | 

### Return type

[**Plan**](Plan.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanByID

> SinglePlan GetPlanByID(ctx, region, id).Execute()





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
    id := "id_example" // string | plan id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlansAPI.GetPlanByID(context.Background(), region, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlanByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPlanByID`: SinglePlan
    fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlanByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**region** | **string** | region code | 
**id** | **string** | plan id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SinglePlan**](SinglePlan.md)

### Authorization

[key](../README.md#key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

