# V1EndpointPort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of this port (corresponds to ServicePort.Name). Must be a DNS_LABEL. Optional only if one port is defined. | [optional] 
**Port** | **int32** | The port number of the endpoint. | 
**Protocol** | Pointer to **string** | The IP protocol for this port. Must be UDP or TCP. Default is TCP. | [optional] 

## Methods

### NewV1EndpointPort

`func NewV1EndpointPort(port int32, ) *V1EndpointPort`

NewV1EndpointPort instantiates a new V1EndpointPort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1EndpointPortWithDefaults

`func NewV1EndpointPortWithDefaults() *V1EndpointPort`

NewV1EndpointPortWithDefaults instantiates a new V1EndpointPort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1EndpointPort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1EndpointPort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1EndpointPort) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1EndpointPort) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *V1EndpointPort) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1EndpointPort) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1EndpointPort) SetPort(v int32)`

SetPort sets Port field to given value.


### GetProtocol

`func (o *V1EndpointPort) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *V1EndpointPort) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *V1EndpointPort) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *V1EndpointPort) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


