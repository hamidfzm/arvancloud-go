# LoadBalancerPoolStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Status** | **bool** |  | 
**Priority** | Pointer to **int32** | Zero means the default pool | [optional] 
**Method** | **string** |  | 
**Keepalive** | **string** |  | [default to "off"]
**NextUpstreamTcp** | **string** |  | [default to "off"]
**Regions** | Pointer to **[]string** |  | [optional] 
**Origins** | Pointer to [**[]LoadBalancerOriginStore**](LoadBalancerOriginStore.md) |  | [optional] 

## Methods

### NewLoadBalancerPoolStore

`func NewLoadBalancerPoolStore(name string, status bool, method string, keepalive string, nextUpstreamTcp string, ) *LoadBalancerPoolStore`

NewLoadBalancerPoolStore instantiates a new LoadBalancerPoolStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolStoreWithDefaults

`func NewLoadBalancerPoolStoreWithDefaults() *LoadBalancerPoolStore`

NewLoadBalancerPoolStoreWithDefaults instantiates a new LoadBalancerPoolStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadBalancerPoolStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerPoolStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerPoolStore) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *LoadBalancerPoolStore) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoadBalancerPoolStore) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoadBalancerPoolStore) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoadBalancerPoolStore) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *LoadBalancerPoolStore) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerPoolStore) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerPoolStore) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetPriority

`func (o *LoadBalancerPoolStore) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *LoadBalancerPoolStore) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *LoadBalancerPoolStore) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *LoadBalancerPoolStore) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetMethod

`func (o *LoadBalancerPoolStore) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *LoadBalancerPoolStore) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *LoadBalancerPoolStore) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetKeepalive

`func (o *LoadBalancerPoolStore) GetKeepalive() string`

GetKeepalive returns the Keepalive field if non-nil, zero value otherwise.

### GetKeepaliveOk

`func (o *LoadBalancerPoolStore) GetKeepaliveOk() (*string, bool)`

GetKeepaliveOk returns a tuple with the Keepalive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepalive

`func (o *LoadBalancerPoolStore) SetKeepalive(v string)`

SetKeepalive sets Keepalive field to given value.


### GetNextUpstreamTcp

`func (o *LoadBalancerPoolStore) GetNextUpstreamTcp() string`

GetNextUpstreamTcp returns the NextUpstreamTcp field if non-nil, zero value otherwise.

### GetNextUpstreamTcpOk

`func (o *LoadBalancerPoolStore) GetNextUpstreamTcpOk() (*string, bool)`

GetNextUpstreamTcpOk returns a tuple with the NextUpstreamTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextUpstreamTcp

`func (o *LoadBalancerPoolStore) SetNextUpstreamTcp(v string)`

SetNextUpstreamTcp sets NextUpstreamTcp field to given value.


### GetRegions

`func (o *LoadBalancerPoolStore) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *LoadBalancerPoolStore) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *LoadBalancerPoolStore) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *LoadBalancerPoolStore) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetOrigins

`func (o *LoadBalancerPoolStore) GetOrigins() []LoadBalancerOriginStore`

GetOrigins returns the Origins field if non-nil, zero value otherwise.

### GetOriginsOk

`func (o *LoadBalancerPoolStore) GetOriginsOk() (*[]LoadBalancerOriginStore, bool)`

GetOriginsOk returns a tuple with the Origins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigins

`func (o *LoadBalancerPoolStore) SetOrigins(v []LoadBalancerOriginStore)`

SetOrigins sets Origins field to given value.

### HasOrigins

`func (o *LoadBalancerPoolStore) HasOrigins() bool`

HasOrigins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


