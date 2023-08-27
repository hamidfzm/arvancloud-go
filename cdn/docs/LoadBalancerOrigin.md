# LoadBalancerOrigin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Weight** | Pointer to **int32** |  | [optional] 
**Protocol** | Pointer to **string** |  | [optional] [default to "auto"]
**HostHeader** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewLoadBalancerOrigin

`func NewLoadBalancerOrigin() *LoadBalancerOrigin`

NewLoadBalancerOrigin instantiates a new LoadBalancerOrigin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerOriginWithDefaults

`func NewLoadBalancerOriginWithDefaults() *LoadBalancerOrigin`

NewLoadBalancerOriginWithDefaults instantiates a new LoadBalancerOrigin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LoadBalancerOrigin) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoadBalancerOrigin) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoadBalancerOrigin) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoadBalancerOrigin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LoadBalancerOrigin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerOrigin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerOrigin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerOrigin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *LoadBalancerOrigin) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerOrigin) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerOrigin) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoadBalancerOrigin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAddress

`func (o *LoadBalancerOrigin) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoadBalancerOrigin) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoadBalancerOrigin) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *LoadBalancerOrigin) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetPort

`func (o *LoadBalancerOrigin) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerOrigin) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerOrigin) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoadBalancerOrigin) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetWeight

`func (o *LoadBalancerOrigin) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *LoadBalancerOrigin) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *LoadBalancerOrigin) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *LoadBalancerOrigin) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### GetProtocol

`func (o *LoadBalancerOrigin) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerOrigin) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerOrigin) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LoadBalancerOrigin) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetHostHeader

`func (o *LoadBalancerOrigin) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *LoadBalancerOrigin) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *LoadBalancerOrigin) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.

### HasHostHeader

`func (o *LoadBalancerOrigin) HasHostHeader() bool`

HasHostHeader returns a boolean if a field has been set.

### GetCreatedAt

`func (o *LoadBalancerOrigin) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LoadBalancerOrigin) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LoadBalancerOrigin) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LoadBalancerOrigin) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *LoadBalancerOrigin) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *LoadBalancerOrigin) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *LoadBalancerOrigin) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *LoadBalancerOrigin) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


