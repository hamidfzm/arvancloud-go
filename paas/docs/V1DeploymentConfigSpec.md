# V1DeploymentConfigSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | MinReadySeconds is the minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready) | [optional] 
**Paused** | Pointer to **bool** | Paused indicates that the deployment config is paused resulting in no new deployments on template changes or changes in the template caused by other triggers. | [optional] 
**Replicas** | **int32** | Replicas is the number of desired replicas. | 
**RevisionHistoryLimit** | Pointer to **int32** | RevisionHistoryLimit is the number of old ReplicationControllers to retain to allow for rollbacks. This field is a pointer to allow for differentiation between an explicit zero and not specified. Defaults to 10. (This only applies to DeploymentConfigs created via the new group API resource, not the legacy resource.) | [optional] 
**Selector** | Pointer to **map[string]interface{}** | Selector is a label query over pods that should match the Replicas count. | [optional] 
**Strategy** | [**V1DeploymentStrategy**](V1DeploymentStrategy.md) |  | 
**Template** | Pointer to [**V1PodTemplateSpec**](V1PodTemplateSpec.md) |  | [optional] 
**Test** | **bool** | Test ensures that this deployment config will have zero replicas except while a deployment is running. This allows the deployment config to be used as a continuous deployment test - triggering on images, running the deployment, and then succeeding or failing. Post strategy hooks and After actions can be used to integrate successful deployment with an action. | 
**Triggers** | [**[]V1DeploymentTriggerPolicy**](V1DeploymentTriggerPolicy.md) | Triggers determine how updates to a DeploymentConfig result in new deployments. If no triggers are defined, a new deployment can only occur as a result of an explicit client update to the DeploymentConfig with a new LatestVersion. If null, defaults to having a config change trigger. | 

## Methods

### NewV1DeploymentConfigSpec

`func NewV1DeploymentConfigSpec(replicas int32, strategy V1DeploymentStrategy, test bool, triggers []V1DeploymentTriggerPolicy, ) *V1DeploymentConfigSpec`

NewV1DeploymentConfigSpec instantiates a new V1DeploymentConfigSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DeploymentConfigSpecWithDefaults

`func NewV1DeploymentConfigSpecWithDefaults() *V1DeploymentConfigSpec`

NewV1DeploymentConfigSpecWithDefaults instantiates a new V1DeploymentConfigSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *V1DeploymentConfigSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *V1DeploymentConfigSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *V1DeploymentConfigSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *V1DeploymentConfigSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetPaused

`func (o *V1DeploymentConfigSpec) GetPaused() bool`

GetPaused returns the Paused field if non-nil, zero value otherwise.

### GetPausedOk

`func (o *V1DeploymentConfigSpec) GetPausedOk() (*bool, bool)`

GetPausedOk returns a tuple with the Paused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaused

`func (o *V1DeploymentConfigSpec) SetPaused(v bool)`

SetPaused sets Paused field to given value.

### HasPaused

`func (o *V1DeploymentConfigSpec) HasPaused() bool`

HasPaused returns a boolean if a field has been set.

### GetReplicas

`func (o *V1DeploymentConfigSpec) GetReplicas() int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *V1DeploymentConfigSpec) GetReplicasOk() (*int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *V1DeploymentConfigSpec) SetReplicas(v int32)`

SetReplicas sets Replicas field to given value.


### GetRevisionHistoryLimit

`func (o *V1DeploymentConfigSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *V1DeploymentConfigSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *V1DeploymentConfigSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *V1DeploymentConfigSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *V1DeploymentConfigSpec) GetSelector() map[string]interface{}`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1DeploymentConfigSpec) GetSelectorOk() (*map[string]interface{}, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1DeploymentConfigSpec) SetSelector(v map[string]interface{})`

SetSelector sets Selector field to given value.

### HasSelector

`func (o *V1DeploymentConfigSpec) HasSelector() bool`

HasSelector returns a boolean if a field has been set.

### GetStrategy

`func (o *V1DeploymentConfigSpec) GetStrategy() V1DeploymentStrategy`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *V1DeploymentConfigSpec) GetStrategyOk() (*V1DeploymentStrategy, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *V1DeploymentConfigSpec) SetStrategy(v V1DeploymentStrategy)`

SetStrategy sets Strategy field to given value.


### GetTemplate

`func (o *V1DeploymentConfigSpec) GetTemplate() V1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1DeploymentConfigSpec) GetTemplateOk() (*V1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1DeploymentConfigSpec) SetTemplate(v V1PodTemplateSpec)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *V1DeploymentConfigSpec) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTest

`func (o *V1DeploymentConfigSpec) GetTest() bool`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *V1DeploymentConfigSpec) GetTestOk() (*bool, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *V1DeploymentConfigSpec) SetTest(v bool)`

SetTest sets Test field to given value.


### GetTriggers

`func (o *V1DeploymentConfigSpec) GetTriggers() []V1DeploymentTriggerPolicy`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *V1DeploymentConfigSpec) GetTriggersOk() (*[]V1DeploymentTriggerPolicy, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *V1DeploymentConfigSpec) SetTriggers(v []V1DeploymentTriggerPolicy)`

SetTriggers sets Triggers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


