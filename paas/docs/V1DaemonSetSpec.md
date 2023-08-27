# V1DaemonSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinReadySeconds** | Pointer to **int32** | The minimum number of seconds for which a newly created DaemonSet pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready). | [optional] 
**RevisionHistoryLimit** | Pointer to **int32** | The number of old history to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified. Defaults to 10. | [optional] 
**Selector** | [**V1LabelSelector**](V1LabelSelector.md) |  | 
**Template** | [**V1PodTemplateSpec**](V1PodTemplateSpec.md) |  | 
**UpdateStrategy** | Pointer to [**V1DaemonSetUpdateStrategy**](V1DaemonSetUpdateStrategy.md) |  | [optional] 

## Methods

### NewV1DaemonSetSpec

`func NewV1DaemonSetSpec(selector V1LabelSelector, template V1PodTemplateSpec, ) *V1DaemonSetSpec`

NewV1DaemonSetSpec instantiates a new V1DaemonSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DaemonSetSpecWithDefaults

`func NewV1DaemonSetSpecWithDefaults() *V1DaemonSetSpec`

NewV1DaemonSetSpecWithDefaults instantiates a new V1DaemonSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinReadySeconds

`func (o *V1DaemonSetSpec) GetMinReadySeconds() int32`

GetMinReadySeconds returns the MinReadySeconds field if non-nil, zero value otherwise.

### GetMinReadySecondsOk

`func (o *V1DaemonSetSpec) GetMinReadySecondsOk() (*int32, bool)`

GetMinReadySecondsOk returns a tuple with the MinReadySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadySeconds

`func (o *V1DaemonSetSpec) SetMinReadySeconds(v int32)`

SetMinReadySeconds sets MinReadySeconds field to given value.

### HasMinReadySeconds

`func (o *V1DaemonSetSpec) HasMinReadySeconds() bool`

HasMinReadySeconds returns a boolean if a field has been set.

### GetRevisionHistoryLimit

`func (o *V1DaemonSetSpec) GetRevisionHistoryLimit() int32`

GetRevisionHistoryLimit returns the RevisionHistoryLimit field if non-nil, zero value otherwise.

### GetRevisionHistoryLimitOk

`func (o *V1DaemonSetSpec) GetRevisionHistoryLimitOk() (*int32, bool)`

GetRevisionHistoryLimitOk returns a tuple with the RevisionHistoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionHistoryLimit

`func (o *V1DaemonSetSpec) SetRevisionHistoryLimit(v int32)`

SetRevisionHistoryLimit sets RevisionHistoryLimit field to given value.

### HasRevisionHistoryLimit

`func (o *V1DaemonSetSpec) HasRevisionHistoryLimit() bool`

HasRevisionHistoryLimit returns a boolean if a field has been set.

### GetSelector

`func (o *V1DaemonSetSpec) GetSelector() V1LabelSelector`

GetSelector returns the Selector field if non-nil, zero value otherwise.

### GetSelectorOk

`func (o *V1DaemonSetSpec) GetSelectorOk() (*V1LabelSelector, bool)`

GetSelectorOk returns a tuple with the Selector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelector

`func (o *V1DaemonSetSpec) SetSelector(v V1LabelSelector)`

SetSelector sets Selector field to given value.


### GetTemplate

`func (o *V1DaemonSetSpec) GetTemplate() V1PodTemplateSpec`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *V1DaemonSetSpec) GetTemplateOk() (*V1PodTemplateSpec, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *V1DaemonSetSpec) SetTemplate(v V1PodTemplateSpec)`

SetTemplate sets Template field to given value.


### GetUpdateStrategy

`func (o *V1DaemonSetSpec) GetUpdateStrategy() V1DaemonSetUpdateStrategy`

GetUpdateStrategy returns the UpdateStrategy field if non-nil, zero value otherwise.

### GetUpdateStrategyOk

`func (o *V1DaemonSetSpec) GetUpdateStrategyOk() (*V1DaemonSetUpdateStrategy, bool)`

GetUpdateStrategyOk returns a tuple with the UpdateStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStrategy

`func (o *V1DaemonSetSpec) SetUpdateStrategy(v V1DaemonSetUpdateStrategy)`

SetUpdateStrategy sets UpdateStrategy field to given value.

### HasUpdateStrategy

`func (o *V1DaemonSetSpec) HasUpdateStrategy() bool`

HasUpdateStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


