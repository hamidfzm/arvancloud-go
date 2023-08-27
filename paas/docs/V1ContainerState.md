# V1ContainerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Running** | Pointer to [**V1ContainerStateRunning**](V1ContainerStateRunning.md) |  | [optional] 
**Terminated** | Pointer to [**V1ContainerStateTerminated**](V1ContainerStateTerminated.md) |  | [optional] 
**Waiting** | Pointer to [**V1ContainerStateWaiting**](V1ContainerStateWaiting.md) |  | [optional] 

## Methods

### NewV1ContainerState

`func NewV1ContainerState() *V1ContainerState`

NewV1ContainerState instantiates a new V1ContainerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ContainerStateWithDefaults

`func NewV1ContainerStateWithDefaults() *V1ContainerState`

NewV1ContainerStateWithDefaults instantiates a new V1ContainerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunning

`func (o *V1ContainerState) GetRunning() V1ContainerStateRunning`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *V1ContainerState) GetRunningOk() (*V1ContainerStateRunning, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *V1ContainerState) SetRunning(v V1ContainerStateRunning)`

SetRunning sets Running field to given value.

### HasRunning

`func (o *V1ContainerState) HasRunning() bool`

HasRunning returns a boolean if a field has been set.

### GetTerminated

`func (o *V1ContainerState) GetTerminated() V1ContainerStateTerminated`

GetTerminated returns the Terminated field if non-nil, zero value otherwise.

### GetTerminatedOk

`func (o *V1ContainerState) GetTerminatedOk() (*V1ContainerStateTerminated, bool)`

GetTerminatedOk returns a tuple with the Terminated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminated

`func (o *V1ContainerState) SetTerminated(v V1ContainerStateTerminated)`

SetTerminated sets Terminated field to given value.

### HasTerminated

`func (o *V1ContainerState) HasTerminated() bool`

HasTerminated returns a boolean if a field has been set.

### GetWaiting

`func (o *V1ContainerState) GetWaiting() V1ContainerStateWaiting`

GetWaiting returns the Waiting field if non-nil, zero value otherwise.

### GetWaitingOk

`func (o *V1ContainerState) GetWaitingOk() (*V1ContainerStateWaiting, bool)`

GetWaitingOk returns a tuple with the Waiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaiting

`func (o *V1ContainerState) SetWaiting(v V1ContainerStateWaiting)`

SetWaiting sets Waiting field to given value.

### HasWaiting

`func (o *V1ContainerState) HasWaiting() bool`

HasWaiting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


