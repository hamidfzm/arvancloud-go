# V1EndpointSubset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**[]V1EndpointAddress**](V1EndpointAddress.md) | IP addresses which offer the related ports that are marked as ready. These endpoints should be considered safe for load balancers and clients to utilize. | [optional] 
**NotReadyAddresses** | Pointer to [**[]V1EndpointAddress**](V1EndpointAddress.md) | IP addresses which offer the related ports but are not currently marked as ready because they have not yet finished starting, have recently failed a readiness check, or have recently failed a liveness check. | [optional] 
**Ports** | Pointer to [**[]V1EndpointPort**](V1EndpointPort.md) | Port numbers available on the related IP addresses. | [optional] 

## Methods

### NewV1EndpointSubset

`func NewV1EndpointSubset() *V1EndpointSubset`

NewV1EndpointSubset instantiates a new V1EndpointSubset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EndpointSubsetWithDefaults

`func NewV1EndpointSubsetWithDefaults() *V1EndpointSubset`

NewV1EndpointSubsetWithDefaults instantiates a new V1EndpointSubset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *V1EndpointSubset) GetAddresses() []V1EndpointAddress`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *V1EndpointSubset) GetAddressesOk() (*[]V1EndpointAddress, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *V1EndpointSubset) SetAddresses(v []V1EndpointAddress)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *V1EndpointSubset) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetNotReadyAddresses

`func (o *V1EndpointSubset) GetNotReadyAddresses() []V1EndpointAddress`

GetNotReadyAddresses returns the NotReadyAddresses field if non-nil, zero value otherwise.

### GetNotReadyAddressesOk

`func (o *V1EndpointSubset) GetNotReadyAddressesOk() (*[]V1EndpointAddress, bool)`

GetNotReadyAddressesOk returns a tuple with the NotReadyAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotReadyAddresses

`func (o *V1EndpointSubset) SetNotReadyAddresses(v []V1EndpointAddress)`

SetNotReadyAddresses sets NotReadyAddresses field to given value.

### HasNotReadyAddresses

`func (o *V1EndpointSubset) HasNotReadyAddresses() bool`

HasNotReadyAddresses returns a boolean if a field has been set.

### GetPorts

`func (o *V1EndpointSubset) GetPorts() []V1EndpointPort`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *V1EndpointSubset) GetPortsOk() (*[]V1EndpointPort, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *V1EndpointSubset) SetPorts(v []V1EndpointPort)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *V1EndpointSubset) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


