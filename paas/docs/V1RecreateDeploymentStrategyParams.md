# V1RecreateDeploymentStrategyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mid** | Pointer to [**V1LifecycleHook**](V1LifecycleHook.md) |  | [optional] 
**Post** | Pointer to [**V1LifecycleHook**](V1LifecycleHook.md) |  | [optional] 
**Pre** | Pointer to [**V1LifecycleHook**](V1LifecycleHook.md) |  | [optional] 
**TimeoutSeconds** | Pointer to **int64** | TimeoutSeconds is the time to wait for updates before giving up. If the value is nil, a default will be used. | [optional] 

## Methods

### NewV1RecreateDeploymentStrategyParams

`func NewV1RecreateDeploymentStrategyParams() *V1RecreateDeploymentStrategyParams`

NewV1RecreateDeploymentStrategyParams instantiates a new V1RecreateDeploymentStrategyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RecreateDeploymentStrategyParamsWithDefaults

`func NewV1RecreateDeploymentStrategyParamsWithDefaults() *V1RecreateDeploymentStrategyParams`

NewV1RecreateDeploymentStrategyParamsWithDefaults instantiates a new V1RecreateDeploymentStrategyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMid

`func (o *V1RecreateDeploymentStrategyParams) GetMid() V1LifecycleHook`

GetMid returns the Mid field if non-nil, zero value otherwise.

### GetMidOk

`func (o *V1RecreateDeploymentStrategyParams) GetMidOk() (*V1LifecycleHook, bool)`

GetMidOk returns a tuple with the Mid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMid

`func (o *V1RecreateDeploymentStrategyParams) SetMid(v V1LifecycleHook)`

SetMid sets Mid field to given value.

### HasMid

`func (o *V1RecreateDeploymentStrategyParams) HasMid() bool`

HasMid returns a boolean if a field has been set.

### GetPost

`func (o *V1RecreateDeploymentStrategyParams) GetPost() V1LifecycleHook`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *V1RecreateDeploymentStrategyParams) GetPostOk() (*V1LifecycleHook, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *V1RecreateDeploymentStrategyParams) SetPost(v V1LifecycleHook)`

SetPost sets Post field to given value.

### HasPost

`func (o *V1RecreateDeploymentStrategyParams) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPre

`func (o *V1RecreateDeploymentStrategyParams) GetPre() V1LifecycleHook`

GetPre returns the Pre field if non-nil, zero value otherwise.

### GetPreOk

`func (o *V1RecreateDeploymentStrategyParams) GetPreOk() (*V1LifecycleHook, bool)`

GetPreOk returns a tuple with the Pre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPre

`func (o *V1RecreateDeploymentStrategyParams) SetPre(v V1LifecycleHook)`

SetPre sets Pre field to given value.

### HasPre

`func (o *V1RecreateDeploymentStrategyParams) HasPre() bool`

HasPre returns a boolean if a field has been set.

### GetTimeoutSeconds

`func (o *V1RecreateDeploymentStrategyParams) GetTimeoutSeconds() int64`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *V1RecreateDeploymentStrategyParams) GetTimeoutSecondsOk() (*int64, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *V1RecreateDeploymentStrategyParams) SetTimeoutSeconds(v int64)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *V1RecreateDeploymentStrategyParams) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


