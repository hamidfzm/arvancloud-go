# V1DeploymentSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Paused** | Pointer to **bool** | Indicates that the deployment is paused. | [optional] 
**ProgressDeadlineSeconds** | Pointer to **int32** | The maximum time in seconds for a deployment to make progress before it is considered to be failed. The deployment controller will continue to process failed deployments and a condition with a ProgressDeadlineExceeded reason will be surfaced in the deployment status. Note that progress will not be estimated during the time a deployment is paused. Defaults to 600s. | [optional] 
**Replicas** | Pointer to **int32** | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1. | [optional] 
**RevisionHistoryLimit** | Pointer to **int32** | The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. | [optional] 
**Selector** | [**V1LabelSelector**](V1LabelSelector.md) |  | 
**Strategy** | Pointer to [**V1DeploymentStrategy**](V1DeploymentStrategy.md) |  | [optional] 
**Template** | [**V1PodTemplateSpec**](V1PodTemplateSpec.md) |  | 

## Methods

### NewV1DeploymentSpec

`func NewV1DeploymentSpec(selector V1LabelSelector, template V1PodTemplateSpec, ) *V1DeploymentSpec`

NewV1DeploymentSpec instantiates a new V1DeploymentSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentSpecWithDefaults

`func NewV1DeploymentSpecWithDefaults() *V1DeploymentSpec`

NewV1DeploymentSpecWithDefaults instantiates a new V1DeploymentSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *V1DeploymentSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *V1DeploymentSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *V1DeploymentSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *V1DeploymentSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetPaused

`func (o *V1DeploymentSpec) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *V1DeploymentSpec) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *V1DeploymentSpec) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *V1DeploymentSpec) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetProgressDeadlineSeconds

`func (o *V1DeploymentSpec) GetProgressDeadlineSeconds() int32`

GetProgressDeadlineSeconds returns the ProgressDeadlineSeconds field if non-nil, zero value otherwise.

### GetProgressDeadlineSecondsOk

`func (o *V1DeploymentSpec) GetProgressDeadlineSecondsOk() (*int32, bool)`

GetProgressDeadlineSecondsOk returns a tuple with the ProgressDeadlineSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressDeadlineSeconds

`func (o *V1DeploymentSpec) SetProgressDeadlineSeconds(v int32)`

SetProgressDeadlineSeconds sets ProgressDeadlineSeconds field to given value.

### HasProgressDeadlineSeconds

`func (o *V1DeploymentSpec) HasProgressDeadlineSeconds() bool`

HasProgressDeadlineSeconds returns a boolean if a field has been set.

### GetReplicas

`func (o *V1DeploymentSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1DeploymentSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1DeploymentSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *V1DeploymentSpec) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *V1DeploymentSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *V1DeploymentSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *V1DeploymentSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *V1DeploymentSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *V1DeploymentSpec) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1DeploymentSpec) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1DeploymentSpec) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.


### GetStrategy

`func (o *V1DeploymentSpec) GetStrategy() V1DeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *V1DeploymentSpec) GetStrategyOk() (*V1DeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *V1DeploymentSpec) SetStrategy(v V1DeploymentStrategy)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *V1DeploymentSpec) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetTemplate

`func (o *V1DeploymentSpec) GetTemplate() V1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1DeploymentSpec) GetTemplateOk() (*V1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1DeploymentSpec) SetTemplate(v V1PodTemplateSpec)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


