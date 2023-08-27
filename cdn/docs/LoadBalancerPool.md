# LoadBalancerPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Priority** | Pointer to **int32** | Zero means the default pool | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Keepalive** | Pointer to **string** |  | [optional] [default to "off"]
**NextUpstreamTcp** | Pointer to **string** |  | [optional] [default to "off"]
**Regions** | Pointer to **[]string** |  | [optional] 
**Origins** | Pointer to [**[]LoadBalancerOrigin**](LoadBalancerOrigin.md) |  | [optional] [readonly] 
**MonitoringStatus** | Pointer to [**NullableMonitoringStatus**](MonitoringStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewLoadBalancerPool

`func NewLoadBalancerPool() *LoadBalancerPool`

NewLoadBalancerPool instantiates a new LoadBalancerPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolWithDefaults

`func NewLoadBalancerPoolWithDefaults() *LoadBalancerPool`

NewLoadBalancerPoolWithDefaults instantiates a new LoadBalancerPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerPool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerPool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerPool) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancerPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *LoadBalancerPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *LoadBalancerPool) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerPool) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerPool) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoadBalancerPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriority

`func (o *LoadBalancerPool) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LoadBalancerPool) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LoadBalancerPool) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LoadBalancerPool) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMethod

`func (o *LoadBalancerPool) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoadBalancerPool) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoadBalancerPool) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *LoadBalancerPool) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetKeepalive

`func (o *LoadBalancerPool) GetKeepalive() string`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *LoadBalancerPool) GetKeepaliveOk() (*string, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *LoadBalancerPool) SetKeepalive(v string)`

SetKeepalive sets Keepalive field to given value.

### HasKeepalive

`func (o *LoadBalancerPool) HasKeepalive() bool`

HasKeepalive returns a boolean if a field has been set.

### GetNextUpstreamTcp

`func (o *LoadBalancerPool) GetNextUpstreamTcp() string`

GetNextUpstreamTcp returns the NextUpstreamTcp field if non-nil, zero value otherwise.

### GetNextUpstreamTcpOk

`func (o *LoadBalancerPool) GetNextUpstreamTcpOk() (*string, bool)`

GetNextUpstreamTcpOk returns a tuple with the NextUpstreamTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpstreamTcp

`func (o *LoadBalancerPool) SetNextUpstreamTcp(v string)`

SetNextUpstreamTcp sets NextUpstreamTcp field to given value.

### HasNextUpstreamTcp

`func (o *LoadBalancerPool) HasNextUpstreamTcp() bool`

HasNextUpstreamTcp returns a boolean if a field has been set.

### GetRegions

`func (o *LoadBalancerPool) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *LoadBalancerPool) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *LoadBalancerPool) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *LoadBalancerPool) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetOrigins

`func (o *LoadBalancerPool) GetOrigins() []LoadBalancerOrigin`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *LoadBalancerPool) GetOriginsOk() (*[]LoadBalancerOrigin, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *LoadBalancerPool) SetOrigins(v []LoadBalancerOrigin)`

SetOrigins sets Origins field to given value.

### HasOrigins

`func (o *LoadBalancerPool) HasOrigins() bool`

HasOrigins returns a boolean if a field has been set.

### GetMonitoringStatus

`func (o *LoadBalancerPool) GetMonitoringStatus() MonitoringStatus`

GetMonitoringStatus returns the MonitoringStatus field if non-nil, zero value otherwise.

### GetMonitoringStatusOk

`func (o *LoadBalancerPool) GetMonitoringStatusOk() (*MonitoringStatus, bool)`

GetMonitoringStatusOk returns a tuple with the MonitoringStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitoringStatus

`func (o *LoadBalancerPool) SetMonitoringStatus(v MonitoringStatus)`

SetMonitoringStatus sets MonitoringStatus field to given value.

### HasMonitoringStatus

`func (o *LoadBalancerPool) HasMonitoringStatus() bool`

HasMonitoringStatus returns a boolean if a field has been set.

### SetMonitoringStatusNil

`func (o *LoadBalancerPool) SetMonitoringStatusNil(b bool)`

 SetMonitoringStatusNil sets the value for MonitoringStatus to be an explicit nil

### UnsetMonitoringStatus
`func (o *LoadBalancerPool) UnsetMonitoringStatus()`

UnsetMonitoringStatus ensures that no value is present for MonitoringStatus, not even an explicit nil
### GetCreatedAt

`func (o *LoadBalancerPool) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerPool) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerPool) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerPool) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LoadBalancerPool) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerPool) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerPool) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerPool) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


