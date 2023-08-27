# SRVRecordValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | **string** |  | 
**Port** | **NullableInt32** |  | 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**Priority** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewSRVRecordValue

`func NewSRVRecordValue(target string, port NullableInt32, ) *SRVRecordValue`

NewSRVRecordValue instantiates a new SRVRecordValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSRVRecordValueWithDefaults

`func NewSRVRecordValueWithDefaults() *SRVRecordValue`

NewSRVRecordValueWithDefaults instantiates a new SRVRecordValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *SRVRecordValue) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *SRVRecordValue) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *SRVRecordValue) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetPort

`func (o *SRVRecordValue) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SRVRecordValue) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SRVRecordValue) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *SRVRecordValue) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *SRVRecordValue) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetWeight

`func (o *SRVRecordValue) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *SRVRecordValue) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *SRVRecordValue) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *SRVRecordValue) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *SRVRecordValue) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *SRVRecordValue) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetPriority

`func (o *SRVRecordValue) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SRVRecordValue) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SRVRecordValue) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *SRVRecordValue) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *SRVRecordValue) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *SRVRecordValue) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


