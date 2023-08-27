# V1RollingDeploymentStrategyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntervalSeconds** | Pointer to **int64** | IntervalSeconds is the time to wait between polling deployment status after update. If the value is nil, a default will be used. | [optional] 
**MaxSurge** | Pointer to **string** | MaxSurge is the maximum number of pods that can be scheduled above the original number of pods. Value can be an absolute number (ex: 5) or a percentage of total pods at the start of the update (ex: 10%). Absolute number is calculated from percentage by rounding up.  This cannot be 0 if MaxUnavailable is 0. By default, 25% is used.  Example: when this is set to 30%, the new RC can be scaled up by 30% immediately when the rolling update starts. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of original pods. | [optional] 
**MaxUnavailable** | Pointer to **string** | MaxUnavailable is the maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of total pods at the start of update (ex: 10%). Absolute number is calculated from percentage by rounding down.  This cannot be 0 if MaxSurge is 0. By default, 25% is used.  Example: when this is set to 30%, the old RC can be scaled down by 30% immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that at least 70% of original number of pods are available at all times during the update. | [optional] 
**Post** | Pointer to [**V1LifecycleHook**](V1LifecycleHook.md) |  | [optional] 
**Pre** | Pointer to [**V1LifecycleHook**](V1LifecycleHook.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** | TimeoutSeconds is the time to wait for updates before giving up. If the value is nil, a default will be used. | [optional] 
**UpdatePeriodSeconds** | Pointer to **int64** | UpdatePeriodSeconds is the time to wait between individual pod updates. If the value is nil, a default will be used. | [optional] 

## Methods

### NewV1RollingDeploymentStrategyParams

`func NewV1RollingDeploymentStrategyParams() *V1RollingDeploymentStrategyParams`

NewV1RollingDeploymentStrategyParams instantiates a new V1RollingDeploymentStrategyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RollingDeploymentStrategyParamsWithDefaults

`func NewV1RollingDeploymentStrategyParamsWithDefaults() *V1RollingDeploymentStrategyParams`

NewV1RollingDeploymentStrategyParamsWithDefaults instantiates a new V1RollingDeploymentStrategyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntervalSeconds

`func (o *V1RollingDeploymentStrategyParams) GetIntervalSeconds() int64`

GetIntervalSeconds returns the IntervalSeconds field if non-nil, zero value otherwise.

### GetIntervalSecondsOk

`func (o *V1RollingDeploymentStrategyParams) GetIntervalSecondsOk() (*int64, bool)`

GetIntervalSecondsOk returns a tuple with the IntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntervalSeconds

`func (o *V1RollingDeploymentStrategyParams) SetIntervalSeconds(v int64)`

SetIntervalSeconds sets IntervalSeconds field to given value.

### HasIntervalSeconds

`func (o *V1RollingDeploymentStrategyParams) HasIntervalSeconds() bool`

HasIntervalSeconds returns a boolean if a field has been set.

### GetMaxSurge

`func (o *V1RollingDeploymentStrategyParams) GetMaxSurge() string`

GetMaxSurge returns the MaxSurge field if non-nil, zero value otherwise.

### GetMaxSurgeOk

`func (o *V1RollingDeploymentStrategyParams) GetMaxSurgeOk() (*string, bool)`

GetMaxSurgeOk returns a tuple with the MaxSurge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSurge

`func (o *V1RollingDeploymentStrategyParams) SetMaxSurge(v string)`

SetMaxSurge sets MaxSurge field to given value.

### HasMaxSurge

`func (o *V1RollingDeploymentStrategyParams) HasMaxSurge() bool`

HasMaxSurge returns a boolean if a field has been set.

### GetMaxUnavailable

`func (o *V1RollingDeploymentStrategyParams) GetMaxUnavailable() string`

GetMaxUnavailable returns the MaxUnavailable field if non-nil, zero value otherwise.

### GetMaxUnavailableOk

`func (o *V1RollingDeploymentStrategyParams) GetMaxUnavailableOk() (*string, bool)`

GetMaxUnavailableOk returns a tuple with the MaxUnavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUnavailable

`func (o *V1RollingDeploymentStrategyParams) SetMaxUnavailable(v string)`

SetMaxUnavailable sets MaxUnavailable field to given value.

### HasMaxUnavailable

`func (o *V1RollingDeploymentStrategyParams) HasMaxUnavailable() bool`

HasMaxUnavailable returns a boolean if a field has been set.

### GetPost

`func (o *V1RollingDeploymentStrategyParams) GetPost() V1LifecycleHook`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *V1RollingDeploymentStrategyParams) GetPostOk() (*V1LifecycleHook, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *V1RollingDeploymentStrategyParams) SetPost(v V1LifecycleHook)`

SetPost sets Post field to given value.

### HasPost

`func (o *V1RollingDeploymentStrategyParams) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPre

`func (o *V1RollingDeploymentStrategyParams) GetPre() V1LifecycleHook`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *V1RollingDeploymentStrategyParams) GetPreOk() (*V1LifecycleHook, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *V1RollingDeploymentStrategyParams) SetPre(v V1LifecycleHook)`

SetPre sets Pre field to given value.

### HasPre

`func (o *V1RollingDeploymentStrategyParams) HasPre() bool`

HasPre returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *V1RollingDeploymentStrategyParams) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *V1RollingDeploymentStrategyParams) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *V1RollingDeploymentStrategyParams) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *V1RollingDeploymentStrategyParams) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### GetUpdatePeriodSeconds

`func (o *V1RollingDeploymentStrategyParams) GetUpdatePeriodSeconds() int64`

GetUpdatePeriodSeconds returns the UpdatePeriodSeconds field if non-nil, zero value otherwise.

### GetUpdatePeriodSecondsOk

`func (o *V1RollingDeploymentStrategyParams) GetUpdatePeriodSecondsOk() (*int64, bool)`

GetUpdatePeriodSecondsOk returns a tuple with the UpdatePeriodSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatePeriodSeconds

`func (o *V1RollingDeploymentStrategyParams) SetUpdatePeriodSeconds(v int64)`

SetUpdatePeriodSeconds sets UpdatePeriodSeconds field to given value.

### HasUpdatePeriodSeconds

`func (o *V1RollingDeploymentStrategyParams) HasUpdatePeriodSeconds() bool`

HasUpdatePeriodSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


