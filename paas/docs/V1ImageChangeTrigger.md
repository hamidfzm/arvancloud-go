# V1ImageChangeTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**LastTriggeredImageID** | Pointer to **string** | lastTriggeredImageID is used internally by the ImageChangeController to save last used image ID for build | [optional] 
**Paused** | Pointer to **bool** | paused is true if this trigger is temporarily disabled. Optional. | [optional] 

## Methods

### NewV1ImageChangeTrigger

`func NewV1ImageChangeTrigger() *V1ImageChangeTrigger`

NewV1ImageChangeTrigger instantiates a new V1ImageChangeTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ImageChangeTriggerWithDefaults

`func NewV1ImageChangeTriggerWithDefaults() *V1ImageChangeTrigger`

NewV1ImageChangeTriggerWithDefaults instantiates a new V1ImageChangeTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *V1ImageChangeTrigger) GetFrom() V1ObjectReference`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1ImageChangeTrigger) GetFromOk() (*V1ObjectReference, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1ImageChangeTrigger) SetFrom(v V1ObjectReference)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1ImageChangeTrigger) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLastTriggeredImageID

`func (o *V1ImageChangeTrigger) GetLastTriggeredImageID() string`

GetLastTriggeredImageID returns the LastTriggeredImageID field if non-nil, zero value otherwise.

### GetLastTriggeredImageIDOk

`func (o *V1ImageChangeTrigger) GetLastTriggeredImageIDOk() (*string, bool)`

GetLastTriggeredImageIDOk returns a tuple with the LastTriggeredImageID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastTriggeredImageID

`func (o *V1ImageChangeTrigger) SetLastTriggeredImageID(v string)`

SetLastTriggeredImageID sets LastTriggeredImageID field to given value.

### HasLastTriggeredImageID

`func (o *V1ImageChangeTrigger) HasLastTriggeredImageID() bool`

HasLastTriggeredImageID returns a boolean if a field has been set.

### GetPaused

`func (o *V1ImageChangeTrigger) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *V1ImageChangeTrigger) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *V1ImageChangeTrigger) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *V1ImageChangeTrigger) HasPaused() bool`

HasPaused returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


