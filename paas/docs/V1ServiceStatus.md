# V1ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoadBalancer** | Pointer to [**V1LoadBalancerStatus**](V1LoadBalancerStatus.md) |  | [optional] 

## Methods

### NewV1ServiceStatus

`func NewV1ServiceStatus() *V1ServiceStatus`

NewV1ServiceStatus instantiates a new V1ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ServiceStatusWithDefaults

`func NewV1ServiceStatusWithDefaults() *V1ServiceStatus`

NewV1ServiceStatusWithDefaults instantiates a new V1ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoadBalancer

`func (o *V1ServiceStatus) GetLoadBalancer() V1LoadBalancerStatus`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *V1ServiceStatus) GetLoadBalancerOk() (*V1LoadBalancerStatus, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *V1ServiceStatus) SetLoadBalancer(v V1LoadBalancerStatus)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *V1ServiceStatus) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


