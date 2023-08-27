# V1BuildSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionDeadlineSeconds** | Pointer to **int64** | completionDeadlineSeconds is an optional duration in seconds, counted from the time when a build pod gets scheduled in the system, that the build may be active on a node before the system actively tries to terminate the build; value must be positive integer | [optional] 
**NodeSelector** | **map[string]interface{}** | nodeSelector is a selector which must be true for the build pod to fit on a node If nil, it can be overridden by default build nodeselector values for the cluster. If set to an empty map or a map with any values, default build nodeselector values are ignored. | 
**Output** | Pointer to [**V1BuildOutput**](V1BuildOutput.md) |  | [optional] 
**PostCommit** | Pointer to [**V1BuildPostCommitSpec**](V1BuildPostCommitSpec.md) |  | [optional] 
**Resources** | Pointer to [**V1ResourceRequirements**](V1ResourceRequirements.md) |  | [optional] 
**Revision** | Pointer to [**V1SourceRevision**](V1SourceRevision.md) |  | [optional] 
**ServiceAccount** | Pointer to **string** | serviceAccount is the name of the ServiceAccount to use to run the pod created by this build. The pod will be allowed to use secrets referenced by the ServiceAccount | [optional] 
**Source** | Pointer to [**V1BuildSource**](V1BuildSource.md) |  | [optional] 
**Strategy** | [**V1BuildStrategy**](V1BuildStrategy.md) |  | 
**TriggeredBy** | [**[]V1BuildTriggerCause**](V1BuildTriggerCause.md) | triggeredBy describes which triggers started the most recent update to the build configuration and contains information about those triggers. | 

## Methods

### NewV1BuildSpec

`func NewV1BuildSpec(nodeSelector map[string]interface{}, strategy V1BuildStrategy, triggeredBy []V1BuildTriggerCause, ) *V1BuildSpec`

NewV1BuildSpec instantiates a new V1BuildSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildSpecWithDefaults

`func NewV1BuildSpecWithDefaults() *V1BuildSpec`

NewV1BuildSpecWithDefaults instantiates a new V1BuildSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionDeadlineSeconds

`func (o *V1BuildSpec) GetCompletionDeadlineSeconds() int64`

GetCompletionDeadlineSeconds returns the CompletionDeadlineSeconds field if non-nil, zero value otherwise.

### GetCompletionDeadlineSecondsOk

`func (o *V1BuildSpec) GetCompletionDeadlineSecondsOk() (*int64, bool)`

GetCompletionDeadlineSecondsOk returns a tuple with the CompletionDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDeadlineSeconds

`func (o *V1BuildSpec) SetCompletionDeadlineSeconds(v int64)`

SetCompletionDeadlineSeconds sets CompletionDeadlineSeconds field to given value.

### HasCompletionDeadlineSeconds

`func (o *V1BuildSpec) HasCompletionDeadlineSeconds() bool`

HasCompletionDeadlineSeconds returns a boolean if a field has been set.

### GetNodeSelector

`func (o *V1BuildSpec) GetNodeSelector() map[string]interface{}`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *V1BuildSpec) GetNodeSelectorOk() (*map[string]interface{}, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *V1BuildSpec) SetNodeSelector(v map[string]interface{})`

SetNodeSelector sets NodeSelector field to given value.


### GetOutput

`func (o *V1BuildSpec) GetOutput() V1BuildOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *V1BuildSpec) GetOutputOk() (*V1BuildOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *V1BuildSpec) SetOutput(v V1BuildOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *V1BuildSpec) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPostCommit

`func (o *V1BuildSpec) GetPostCommit() V1BuildPostCommitSpec`

GetPostCommit returns the PostCommit field if non-nil, zero value otherwise.

### GetPostCommitOk

`func (o *V1BuildSpec) GetPostCommitOk() (*V1BuildPostCommitSpec, bool)`

GetPostCommitOk returns a tuple with the PostCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCommit

`func (o *V1BuildSpec) SetPostCommit(v V1BuildPostCommitSpec)`

SetPostCommit sets PostCommit field to given value.

### HasPostCommit

`func (o *V1BuildSpec) HasPostCommit() bool`

HasPostCommit returns a boolean if a field has been set.

### GetResources

`func (o *V1BuildSpec) GetResources() V1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1BuildSpec) GetResourcesOk() (*V1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1BuildSpec) SetResources(v V1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1BuildSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *V1BuildSpec) GetRevision() V1SourceRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1BuildSpec) GetRevisionOk() (*V1SourceRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1BuildSpec) SetRevision(v V1SourceRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1BuildSpec) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetServiceAccount

`func (o *V1BuildSpec) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *V1BuildSpec) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *V1BuildSpec) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *V1BuildSpec) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetSource

`func (o *V1BuildSpec) GetSource() V1BuildSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1BuildSpec) GetSourceOk() (*V1BuildSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1BuildSpec) SetSource(v V1BuildSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1BuildSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStrategy

`func (o *V1BuildSpec) GetStrategy() V1BuildStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *V1BuildSpec) GetStrategyOk() (*V1BuildStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *V1BuildSpec) SetStrategy(v V1BuildStrategy)`

SetStrategy sets Strategy field to given value.


### GetTriggeredBy

`func (o *V1BuildSpec) GetTriggeredBy() []V1BuildTriggerCause`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *V1BuildSpec) GetTriggeredByOk() (*[]V1BuildTriggerCause, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *V1BuildSpec) SetTriggeredBy(v []V1BuildTriggerCause)`

SetTriggeredBy sets TriggeredBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


