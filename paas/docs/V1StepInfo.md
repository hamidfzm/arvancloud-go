# V1StepInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMilliseconds** | Pointer to **int64** | durationMilliseconds identifies how long the step took to complete in milliseconds. | [optional] 
**Name** | Pointer to **string** | name is a unique identifier for each build step. | [optional] 
**StartTime** | Pointer to **string** | startTime is a timestamp representing the server time when this Step started. it is represented in RFC3339 form and is in UTC. | [optional] 

## Methods

### NewV1StepInfo

`func NewV1StepInfo() *V1StepInfo`

NewV1StepInfo instantiates a new V1StepInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StepInfoWithDefaults

`func NewV1StepInfoWithDefaults() *V1StepInfo`

NewV1StepInfoWithDefaults instantiates a new V1StepInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMilliseconds

`func (o *V1StepInfo) GetDurationMilliseconds() int64`

GetDurationMilliseconds returns the DurationMilliseconds field if non-nil, zero value otherwise.

### GetDurationMillisecondsOk

`func (o *V1StepInfo) GetDurationMillisecondsOk() (*int64, bool)`

GetDurationMillisecondsOk returns a tuple with the DurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMilliseconds

`func (o *V1StepInfo) SetDurationMilliseconds(v int64)`

SetDurationMilliseconds sets DurationMilliseconds field to given value.

### HasDurationMilliseconds

`func (o *V1StepInfo) HasDurationMilliseconds() bool`

HasDurationMilliseconds returns a boolean if a field has been set.

### GetName

`func (o *V1StepInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1StepInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1StepInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1StepInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *V1StepInfo) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V1StepInfo) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V1StepInfo) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V1StepInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


