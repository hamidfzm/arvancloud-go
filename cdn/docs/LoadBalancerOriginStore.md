# LoadBalancerOriginStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Status** | **bool** |  | 
**Address** | **string** |  | 
**Port** | **int32** |  | 
**Weight** | **int32** |  | 
**Protocol** | **string** |  | [default to "auto"]
**HostHeader** | Pointer to **string** |  | [optional] 

## Methods

### NewLoadBalancerOriginStore

`func NewLoadBalancerOriginStore(status bool, address string, port int32, weight int32, protocol string, ) *LoadBalancerOriginStore`

NewLoadBalancerOriginStore instantiates a new LoadBalancerOriginStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerOriginStoreWithDefaults

`func NewLoadBalancerOriginStoreWithDefaults() *LoadBalancerOriginStore`

NewLoadBalancerOriginStoreWithDefaults instantiates a new LoadBalancerOriginStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LoadBalancerOriginStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoadBalancerOriginStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoadBalancerOriginStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoadBalancerOriginStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *LoadBalancerOriginStore) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoadBalancerOriginStore) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoadBalancerOriginStore) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetAddress

`func (o *LoadBalancerOriginStore) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *LoadBalancerOriginStore) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *LoadBalancerOriginStore) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPort

`func (o *LoadBalancerOriginStore) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerOriginStore) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerOriginStore) SetPort(v int32)`

SetPort sets Port field to given value.


### GetWeight

`func (o *LoadBalancerOriginStore) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *LoadBalancerOriginStore) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *LoadBalancerOriginStore) SetWeight(v int32)`

SetWeight sets Weight field to given value.


### GetProtocol

`func (o *LoadBalancerOriginStore) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LoadBalancerOriginStore) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LoadBalancerOriginStore) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.


### GetHostHeader

`func (o *LoadBalancerOriginStore) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *LoadBalancerOriginStore) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *LoadBalancerOriginStore) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.

### HasHostHeader

`func (o *LoadBalancerOriginStore) HasHostHeader() bool`

HasHostHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


