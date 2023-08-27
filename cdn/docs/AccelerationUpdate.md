# AccelerationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Extensions** | Pointer to **[]string** |  | [optional] [default to []]

## Methods

### NewAccelerationUpdate

`func NewAccelerationUpdate() *AccelerationUpdate`

NewAccelerationUpdate instantiates a new AccelerationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccelerationUpdateWithDefaults

`func NewAccelerationUpdateWithDefaults() *AccelerationUpdate`

NewAccelerationUpdateWithDefaults instantiates a new AccelerationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccelerationUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccelerationUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccelerationUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AccelerationUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExtensions

`func (o *AccelerationUpdate) GetExtensions() []string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *AccelerationUpdate) GetExtensionsOk() (*[]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *AccelerationUpdate) SetExtensions(v []string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *AccelerationUpdate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


