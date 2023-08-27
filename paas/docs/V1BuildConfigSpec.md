# V1BuildConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionDeadlineSeconds** | Pointer to **int64** | completionDeadlineSeconds is an optional duration in seconds, counted from the time when a build pod gets scheduled in the system, that the build may be active on a node before the system actively tries to terminate the build; value must be positive integer | [optional] 
**FailedBuildsHistoryLimit** | Pointer to **int32** | failedBuildsHistoryLimit is the number of old failed builds to retain. If not specified, all failed builds are retained. | [optional] 
**NodeSelector** | **map[string]interface{}** | nodeSelector is a selector which must be true for the build pod to fit on a node If nil, it can be overridden by default build nodeselector values for the cluster. If set to an empty map or a map with any values, default build nodeselector values are ignored. | 
**Output** | Pointer to [**V1BuildOutput**](V1BuildOutput.md) |  | [optional] 
**PostCommit** | Pointer to [**V1BuildPostCommitSpec**](V1BuildPostCommitSpec.md) |  | [optional] 
**Resources** | Pointer to [**V1ResourceRequirements**](V1ResourceRequirements.md) |  | [optional] 
**Revision** | Pointer to [**V1SourceRevision**](V1SourceRevision.md) |  | [optional] 
**RunPolicy** | Pointer to **string** | RunPolicy describes how the new build created from this build configuration will be scheduled for execution. This is optional, if not specified we default to \&quot;Serial\&quot;. | [optional] 
**ServiceAccount** | Pointer to **string** | serviceAccount is the name of the ServiceAccount to use to run the pod created by this build. The pod will be allowed to use secrets referenced by the ServiceAccount | [optional] 
**Source** | Pointer to [**V1BuildSource**](V1BuildSource.md) |  | [optional] 
**Strategy** | [**V1BuildStrategy**](V1BuildStrategy.md) |  | 
**SuccessfulBuildsHistoryLimit** | Pointer to **int32** | successfulBuildsHistoryLimit is the number of old successful builds to retain. If not specified, all successful builds are retained. | [optional] 
**Triggers** | [**[]V1BuildTriggerPolicy**](V1BuildTriggerPolicy.md) | triggers determine how new Builds can be launched from a BuildConfig. If no triggers are defined, a new build can only occur as a result of an explicit client build creation. | 

## Methods

### NewV1BuildConfigSpec

`func NewV1BuildConfigSpec(nodeSelector map[string]interface{}, strategy V1BuildStrategy, triggers []V1BuildTriggerPolicy, ) *V1BuildConfigSpec`

NewV1BuildConfigSpec instantiates a new V1BuildConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildConfigSpecWithDefaults

`func NewV1BuildConfigSpecWithDefaults() *V1BuildConfigSpec`

NewV1BuildConfigSpecWithDefaults instantiates a new V1BuildConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionDeadlineSeconds

`func (o *V1BuildConfigSpec) GetCompletionDeadlineSeconds() int64`

GetCompletionDeadlineSeconds returns the CompletionDeadlineSeconds field if non-nil, zero value otherwise.

### GetCompletionDeadlineSecondsOk

`func (o *V1BuildConfigSpec) GetCompletionDeadlineSecondsOk() (*int64, bool)`

GetCompletionDeadlineSecondsOk returns a tuple with the CompletionDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionDeadlineSeconds

`func (o *V1BuildConfigSpec) SetCompletionDeadlineSeconds(v int64)`

SetCompletionDeadlineSeconds sets CompletionDeadlineSeconds field to given value.

### HasCompletionDeadlineSeconds

`func (o *V1BuildConfigSpec) HasCompletionDeadlineSeconds() bool`

HasCompletionDeadlineSeconds returns a boolean if a field has been set.

### GetFailedBuildsHistoryLimit

`func (o *V1BuildConfigSpec) GetFailedBuildsHistoryLimit() int32`

GetFailedBuildsHistoryLimit returns the FailedBuildsHistoryLimit field if non-nil, zero value otherwise.

### GetFailedBuildsHistoryLimitOk

`func (o *V1BuildConfigSpec) GetFailedBuildsHistoryLimitOk() (*int32, bool)`

GetFailedBuildsHistoryLimitOk returns a tuple with the FailedBuildsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedBuildsHistoryLimit

`func (o *V1BuildConfigSpec) SetFailedBuildsHistoryLimit(v int32)`

SetFailedBuildsHistoryLimit sets FailedBuildsHistoryLimit field to given value.

### HasFailedBuildsHistoryLimit

`func (o *V1BuildConfigSpec) HasFailedBuildsHistoryLimit() bool`

HasFailedBuildsHistoryLimit returns a boolean if a field has been set.

### GetNodeSelector

`func (o *V1BuildConfigSpec) GetNodeSelector() map[string]interface{}`

GetNodeSelector returns the NodeSelector field if non-nil, zero value otherwise.

### GetNodeSelectorOk

`func (o *V1BuildConfigSpec) GetNodeSelectorOk() (*map[string]interface{}, bool)`

GetNodeSelectorOk returns a tuple with the NodeSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeSelector

`func (o *V1BuildConfigSpec) SetNodeSelector(v map[string]interface{})`

SetNodeSelector sets NodeSelector field to given value.


### GetOutput

`func (o *V1BuildConfigSpec) GetOutput() V1BuildOutput`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *V1BuildConfigSpec) GetOutputOk() (*V1BuildOutput, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *V1BuildConfigSpec) SetOutput(v V1BuildOutput)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *V1BuildConfigSpec) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetPostCommit

`func (o *V1BuildConfigSpec) GetPostCommit() V1BuildPostCommitSpec`

GetPostCommit returns the PostCommit field if non-nil, zero value otherwise.

### GetPostCommitOk

`func (o *V1BuildConfigSpec) GetPostCommitOk() (*V1BuildPostCommitSpec, bool)`

GetPostCommitOk returns a tuple with the PostCommit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCommit

`func (o *V1BuildConfigSpec) SetPostCommit(v V1BuildPostCommitSpec)`

SetPostCommit sets PostCommit field to given value.

### HasPostCommit

`func (o *V1BuildConfigSpec) HasPostCommit() bool`

HasPostCommit returns a boolean if a field has been set.

### GetResources

`func (o *V1BuildConfigSpec) GetResources() V1ResourceRequirements`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *V1BuildConfigSpec) GetResourcesOk() (*V1ResourceRequirements, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *V1BuildConfigSpec) SetResources(v V1ResourceRequirements)`

SetResources sets Resources field to given value.

### HasResources

`func (o *V1BuildConfigSpec) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetRevision

`func (o *V1BuildConfigSpec) GetRevision() V1SourceRevision`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *V1BuildConfigSpec) GetRevisionOk() (*V1SourceRevision, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *V1BuildConfigSpec) SetRevision(v V1SourceRevision)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *V1BuildConfigSpec) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetRunPolicy

`func (o *V1BuildConfigSpec) GetRunPolicy() string`

GetRunPolicy returns the RunPolicy field if non-nil, zero value otherwise.

### GetRunPolicyOk

`func (o *V1BuildConfigSpec) GetRunPolicyOk() (*string, bool)`

GetRunPolicyOk returns a tuple with the RunPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunPolicy

`func (o *V1BuildConfigSpec) SetRunPolicy(v string)`

SetRunPolicy sets RunPolicy field to given value.

### HasRunPolicy

`func (o *V1BuildConfigSpec) HasRunPolicy() bool`

HasRunPolicy returns a boolean if a field has been set.

### GetServiceAccount

`func (o *V1BuildConfigSpec) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *V1BuildConfigSpec) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *V1BuildConfigSpec) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *V1BuildConfigSpec) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetSource

`func (o *V1BuildConfigSpec) GetSource() V1BuildSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1BuildConfigSpec) GetSourceOk() (*V1BuildSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1BuildConfigSpec) SetSource(v V1BuildSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1BuildConfigSpec) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStrategy

`func (o *V1BuildConfigSpec) GetStrategy() V1BuildStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *V1BuildConfigSpec) GetStrategyOk() (*V1BuildStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *V1BuildConfigSpec) SetStrategy(v V1BuildStrategy)`

SetStrategy sets Strategy field to given value.


### GetSuccessfulBuildsHistoryLimit

`func (o *V1BuildConfigSpec) GetSuccessfulBuildsHistoryLimit() int32`

GetSuccessfulBuildsHistoryLimit returns the SuccessfulBuildsHistoryLimit field if non-nil, zero value otherwise.

### GetSuccessfulBuildsHistoryLimitOk

`func (o *V1BuildConfigSpec) GetSuccessfulBuildsHistoryLimitOk() (*int32, bool)`

GetSuccessfulBuildsHistoryLimitOk returns a tuple with the SuccessfulBuildsHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessfulBuildsHistoryLimit

`func (o *V1BuildConfigSpec) SetSuccessfulBuildsHistoryLimit(v int32)`

SetSuccessfulBuildsHistoryLimit sets SuccessfulBuildsHistoryLimit field to given value.

### HasSuccessfulBuildsHistoryLimit

`func (o *V1BuildConfigSpec) HasSuccessfulBuildsHistoryLimit() bool`

HasSuccessfulBuildsHistoryLimit returns a boolean if a field has been set.

### GetTriggers

`func (o *V1BuildConfigSpec) GetTriggers() []V1BuildTriggerPolicy`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *V1BuildConfigSpec) GetTriggersOk() (*[]V1BuildTriggerPolicy, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *V1BuildConfigSpec) SetTriggers(v []V1BuildTriggerPolicy)`

SetTriggers sets Triggers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


