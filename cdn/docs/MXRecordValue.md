# MXRecordValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** |  | 
**Priority** | **NullableInt32** |  | 

## Methods

### NewMXRecordValue

`func NewMXRecordValue(host string, priority NullableInt32, ) *MXRecordValue`

NewMXRecordValue instantiates a new MXRecordValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMXRecordValueWithDefaults

`func NewMXRecordValueWithDefaults() *MXRecordValue`

NewMXRecordValueWithDefaults instantiates a new MXRecordValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *MXRecordValue) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MXRecordValue) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MXRecordValue) SetHost(v string)`

SetHost sets Host field to given value.


### GetPriority

`func (o *MXRecordValue) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MXRecordValue) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MXRecordValue) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### SetPriorityNil

`func (o *MXRecordValue) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *MXRecordValue) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


