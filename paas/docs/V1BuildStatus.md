# V1BuildStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cancelled** | Pointer to **bool** | cancelled describes if a cancel event was triggered for the build. | [optional] 
**CompletionTimestamp** | Pointer to **string** | completionTimestamp is a timestamp representing the server time when this Build was finished, whether that build failed or succeeded.  It reflects the time at which the Pod running the Build terminated. It is represented in RFC3339 form and is in UTC. | [optional] 
**Config** | Pointer to [**V1ObjectReference**](V1ObjectReference.md) |  | [optional] 
**LogSnippet** | Pointer to **string** | logSnippet is the last few lines of the build log.  This value is only set for builds that failed. | [optional] 
**Message** | Pointer to **string** | message is a human-readable message indicating details about why the build has this status. | [optional] 
**Output** | Pointer to [**V1BuildStatusOutput**](V1BuildStatusOutput.md) |  | [optional] 
**OutputDockerImageReference** | Pointer to **string** | outputDockerImageReference contains a reference to the Docker image that will be built by this build. Its value is computed from Build.Spec.Output.To, and should include the registry address, so that it can be used to push and pull the image. | [optional] 
**Phase** | **string** | phase is the point in the build lifecycle. Possible values are \&quot;New\&quot;, \&quot;Pending\&quot;, \&quot;Running\&quot;, \&quot;Complete\&quot;, \&quot;Failed\&quot;, \&quot;Error\&quot;, and \&quot;Cancelled\&quot;. | 
**Reason** | Pointer to **string** | reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI. | [optional] 
**Stages** | Pointer to [**[]V1StageInfo**](V1StageInfo.md) | stages contains details about each stage that occurs during the build including start time, duration (in milliseconds), and the steps that occured within each stage. | [optional] 
**StartTimestamp** | Pointer to **string** | startTimestamp is a timestamp representing the server time when this Build started running in a Pod. It is represented in RFC3339 form and is in UTC. | [optional] 

## Methods

### NewV1BuildStatus

`func NewV1BuildStatus(phase string, ) *V1BuildStatus`

NewV1BuildStatus instantiates a new V1BuildStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildStatusWithDefaults

`func NewV1BuildStatusWithDefaults() *V1BuildStatus`

NewV1BuildStatusWithDefaults instantiates a new V1BuildStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelled

`func (o *V1BuildStatus) GetCancelled() bool`

GetCancelled returns the Cancelled field if non-nil, zero value otherwise.

### GetCancelledOk

`func (o *V1BuildStatus) GetCancelledOk() (*bool, bool)`

GetCancelledOk returns a tuple with the Cancelled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelled

`func (o *V1BuildStatus) SetCancelled(v bool)`

SetCancelled sets Cancelled field to given value.

### HasCancelled

`func (o *V1BuildStatus) HasCancelled() bool`

HasCancelled returns a boolean if a field has been set.

### GetCompletionTimestamp

`func (o *V1BuildStatus) GetCompletionTimestamp() string`

GetCompletionTimestamp returns the CompletionTimestamp field if non-nil, zero value otherwise.

### GetCompletionTimestampOk

`func (o *V1BuildStatus) GetCompletionTimestampOk() (*string, bool)`

GetCompletionTimestampOk returns a tuple with the CompletionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTimestamp

`func (o *V1BuildStatus) SetCompletionTimestamp(v string)`

SetCompletionTimestamp sets CompletionTimestamp field to given value.

### HasCompletionTimestamp

`func (o *V1BuildStatus) HasCompletionTimestamp() bool`

HasCompletionTimestamp returns a boolean if a field has been set.

### GetConfig

`func (o *V1BuildStatus) GetConfig() V1ObjectReference`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1BuildStatus) GetConfigOk() (*V1ObjectReference, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1BuildStatus) SetConfig(v V1ObjectReference)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1BuildStatus) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLogSnippet

`func (o *V1BuildStatus) GetLogSnippet() string`

GetLogSnippet returns the LogSnippet field if non-nil, zero value otherwise.

### GetLogSnippetOk

`func (o *V1BuildStatus) GetLogSnippetOk() (*string, bool)`

GetLogSnippetOk returns a tuple with the LogSnippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSnippet

`func (o *V1BuildStatus) SetLogSnippet(v string)`

SetLogSnippet sets LogSnippet field to given value.

### HasLogSnippet

`func (o *V1BuildStatus) HasLogSnippet() bool`

HasLogSnippet returns a boolean if a field has been set.

### GetMessage

`func (o *V1BuildStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *V1BuildStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *V1BuildStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *V1BuildStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *V1BuildStatus) GetOutput() V1BuildStatusOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *V1BuildStatus) GetOutputOk() (*V1BuildStatusOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *V1BuildStatus) SetOutput(v V1BuildStatusOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *V1BuildStatus) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetOutputDockerImageReference

`func (o *V1BuildStatus) GetOutputDockerImageReference() string`

GetOutputDockerImageReference returns the OutputDockerImageReference field if non-nil, zero value otherwise.

### GetOutputDockerImageReferenceOk

`func (o *V1BuildStatus) GetOutputDockerImageReferenceOk() (*string, bool)`

GetOutputDockerImageReferenceOk returns a tuple with the OutputDockerImageReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputDockerImageReference

`func (o *V1BuildStatus) SetOutputDockerImageReference(v string)`

SetOutputDockerImageReference sets OutputDockerImageReference field to given value.

### HasOutputDockerImageReference

`func (o *V1BuildStatus) HasOutputDockerImageReference() bool`

HasOutputDockerImageReference returns a boolean if a field has been set.

### GetPhase

`func (o *V1BuildStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1BuildStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1BuildStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.


### GetReason

`func (o *V1BuildStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *V1BuildStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *V1BuildStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *V1BuildStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStages

`func (o *V1BuildStatus) GetStages() []V1StageInfo`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *V1BuildStatus) GetStagesOk() (*[]V1StageInfo, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *V1BuildStatus) SetStages(v []V1StageInfo)`

SetStages sets Stages field to given value.

### HasStages

`func (o *V1BuildStatus) HasStages() bool`

HasStages returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *V1BuildStatus) GetStartTimestamp() string`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *V1BuildStatus) GetStartTimestampOk() (*string, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *V1BuildStatus) SetStartTimestamp(v string)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *V1BuildStatus) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


