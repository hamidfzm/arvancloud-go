# V1StageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMilliseconds** | Pointer to **int64** | durationMilliseconds identifies how long the stage took to complete in milliseconds. Note: the duration of a stage can exceed the sum of the duration of the steps within the stage as not all actions are accounted for in explicit build steps. | [optional] 
**Name** | Pointer to **string** | name is a unique identifier for each build stage that occurs. | [optional] 
**StartTime** | Pointer to **string** | startTime is a timestamp representing the server time when this Stage started. It is represented in RFC3339 form and is in UTC. | [optional] 
**Steps** | Pointer to [**[]V1StepInfo**](V1StepInfo.md) | steps contains details about each step that occurs during a build stage including start time and duration in milliseconds. | [optional] 

## Methods

### NewV1StageInfo

`func NewV1StageInfo() *V1StageInfo`

NewV1StageInfo instantiates a new V1StageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StageInfoWithDefaults

`func NewV1StageInfoWithDefaults() *V1StageInfo`

NewV1StageInfoWithDefaults instantiates a new V1StageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationMilliseconds

`func (o *V1StageInfo) GetDurationMilliseconds() int64`

GetDurationMilliseconds returns the DurationMilliseconds field if non-nil, zero value otherwise.

### GetDurationMillisecondsOk

`func (o *V1StageInfo) GetDurationMillisecondsOk() (*int64, bool)`

GetDurationMillisecondsOk returns a tuple with the DurationMilliseconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMilliseconds

`func (o *V1StageInfo) SetDurationMilliseconds(v int64)`

SetDurationMilliseconds sets DurationMilliseconds field to given value.

### HasDurationMilliseconds

`func (o *V1StageInfo) HasDurationMilliseconds() bool`

HasDurationMilliseconds returns a boolean if a field has been set.

### GetName

`func (o *V1StageInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1StageInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1StageInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1StageInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStartTime

`func (o *V1StageInfo) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *V1StageInfo) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *V1StageInfo) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *V1StageInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetSteps

`func (o *V1StageInfo) GetSteps() []V1StepInfo`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *V1StageInfo) GetStepsOk() (*[]V1StepInfo, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *V1StageInfo) SetSteps(v []V1StepInfo)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *V1StageInfo) HasSteps() bool`

HasSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


