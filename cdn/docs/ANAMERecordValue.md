# ANAMERecordValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **string** | A fully-qualified domain name (FQDN) | 
**HostHeader** | **NullableString** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewANAMERecordValue

`func NewANAMERecordValue(location string, hostHeader NullableString, ) *ANAMERecordValue`

NewANAMERecordValue instantiates a new ANAMERecordValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewANAMERecordValueWithDefaults

`func NewANAMERecordValueWithDefaults() *ANAMERecordValue`

NewANAMERecordValueWithDefaults instantiates a new ANAMERecordValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ANAMERecordValue) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ANAMERecordValue) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ANAMERecordValue) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetHostHeader

`func (o *ANAMERecordValue) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *ANAMERecordValue) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *ANAMERecordValue) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.


### SetHostHeaderNil

`func (o *ANAMERecordValue) SetHostHeaderNil(b bool)`

 SetHostHeaderNil sets the value for HostHeader to be an explicit nil

### UnsetHostHeader
`func (o *ANAMERecordValue) UnsetHostHeader()`

UnsetHostHeader ensures that no value is present for HostHeader, not even an explicit nil
### GetPort

`func (o *ANAMERecordValue) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ANAMERecordValue) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ANAMERecordValue) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ANAMERecordValue) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *ANAMERecordValue) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *ANAMERecordValue) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


