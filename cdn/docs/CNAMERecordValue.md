# CNAMERecordValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | A fully-qualified domain name (FQDN) | 
**HostHeader** | **NullableString** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCNAMERecordValue

`func NewCNAMERecordValue(host string, hostHeader NullableString, ) *CNAMERecordValue`

NewCNAMERecordValue instantiates a new CNAMERecordValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCNAMERecordValueWithDefaults

`func NewCNAMERecordValueWithDefaults() *CNAMERecordValue`

NewCNAMERecordValueWithDefaults instantiates a new CNAMERecordValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CNAMERecordValue) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CNAMERecordValue) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CNAMERecordValue) SetHost(v string)`

SetHost sets Host field to given value.


### GetHostHeader

`func (o *CNAMERecordValue) GetHostHeader() string`

GetHostHeader returns the HostHeader field if non-nil, zero value otherwise.

### GetHostHeaderOk

`func (o *CNAMERecordValue) GetHostHeaderOk() (*string, bool)`

GetHostHeaderOk returns a tuple with the HostHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostHeader

`func (o *CNAMERecordValue) SetHostHeader(v string)`

SetHostHeader sets HostHeader field to given value.


### SetHostHeaderNil

`func (o *CNAMERecordValue) SetHostHeaderNil(b bool)`

 SetHostHeaderNil sets the value for HostHeader to be an explicit nil

### UnsetHostHeader
`func (o *CNAMERecordValue) UnsetHostHeader()`

UnsetHostHeader ensures that no value is present for HostHeader, not even an explicit nil
### GetPort

`func (o *CNAMERecordValue) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CNAMERecordValue) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CNAMERecordValue) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CNAMERecordValue) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *CNAMERecordValue) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *CNAMERecordValue) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


