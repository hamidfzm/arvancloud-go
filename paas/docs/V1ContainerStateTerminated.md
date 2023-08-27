# V1ContainerStateTerminated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerID** | Pointer to **string** | Container&#39;s ID in the format &#39;docker://&lt;container_id&gt;&#39; | [optional] 
**ExitCode** | **int32** | Exit status from the last termination of the container | 
**FinishedAt** | Pointer to **string** | Time at which the container last terminated | [optional] 
**Message** | Pointer to **string** | Message regarding the last termination of the container | [optional] 
**Reason** | Pointer to **string** | (brief) reason from the last termination of the container | [optional] 
**Signal** | Pointer to **int32** | Signal from the last termination of the container | [optional] 
**StartedAt** | Pointer to **string** | Time at which previous execution of the container started | [optional] 

## Methods

### NewV1ContainerStateTerminated

`func NewV1ContainerStateTerminated(exitCode int32, ) *V1ContainerStateTerminated`

NewV1ContainerStateTerminated instantiates a new V1ContainerStateTerminated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerStateTerminatedWithDefaults

`func NewV1ContainerStateTerminatedWithDefaults() *V1ContainerStateTerminated`

NewV1ContainerStateTerminatedWithDefaults instantiates a new V1ContainerStateTerminated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerID

`func (o *V1ContainerStateTerminated) GetContainerID() string`

GetContainerID returns the ContainerID field if non-nil, zero value otherwise.

### GetContainerIDOk

`func (o *V1ContainerStateTerminated) GetContainerIDOk() (*string, bool)`

GetContainerIDOk returns a tuple with the ContainerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerID

`func (o *V1ContainerStateTerminated) SetContainerID(v string)`

SetContainerID sets ContainerID field to given value.

### HasContainerID

`func (o *V1ContainerStateTerminated) HasContainerID() bool`

HasContainerID returns a boolean if a field has been set.

### GetExitCode

`func (o *V1ContainerStateTerminated) GetExitCode() int32`

GetExitCode returns the ExitCode field if non-nil, zero value otherwise.

### GetExitCodeOk

`func (o *V1ContainerStateTerminated) GetExitCodeOk() (*int32, bool)`

GetExitCodeOk returns a tuple with the ExitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitCode

`func (o *V1ContainerStateTerminated) SetExitCode(v int32)`

SetExitCode sets ExitCode field to given value.


### GetFinishedAt

`func (o *V1ContainerStateTerminated) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *V1ContainerStateTerminated) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *V1ContainerStateTerminated) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *V1ContainerStateTerminated) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetMessage

`func (o *V1ContainerStateTerminated) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1ContainerStateTerminated) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1ContainerStateTerminated) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1ContainerStateTerminated) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *V1ContainerStateTerminated) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1ContainerStateTerminated) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1ContainerStateTerminated) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1ContainerStateTerminated) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetSignal

`func (o *V1ContainerStateTerminated) GetSignal() int32`

GetSignal returns the Signal field if non-nil, zero value otherwise.

### GetSignalOk

`func (o *V1ContainerStateTerminated) GetSignalOk() (*int32, bool)`

GetSignalOk returns a tuple with the Signal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignal

`func (o *V1ContainerStateTerminated) SetSignal(v int32)`

SetSignal sets Signal field to given value.

### HasSignal

`func (o *V1ContainerStateTerminated) HasSignal() bool`

HasSignal returns a boolean if a field has been set.

### GetStartedAt

`func (o *V1ContainerStateTerminated) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V1ContainerStateTerminated) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V1ContainerStateTerminated) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *V1ContainerStateTerminated) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


