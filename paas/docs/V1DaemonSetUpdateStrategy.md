# V1DaemonSetUpdateStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpdate** | Pointer to [**V1RollingUpdateDaemonSet**](V1RollingUpdateDaemonSet.md) |  | [optional] 
**Type** | Pointer to **string** | Type of daemon set update. Can be \&quot;RollingUpdate\&quot; or \&quot;OnDelete\&quot;. Default is RollingUpdate. | [optional] 

## Methods

### NewV1DaemonSetUpdateStrategy

`func NewV1DaemonSetUpdateStrategy() *V1DaemonSetUpdateStrategy`

NewV1DaemonSetUpdateStrategy instantiates a new V1DaemonSetUpdateStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1DaemonSetUpdateStrategyWithDefaults

`func NewV1DaemonSetUpdateStrategyWithDefaults() *V1DaemonSetUpdateStrategy`

NewV1DaemonSetUpdateStrategyWithDefaults instantiates a new V1DaemonSetUpdateStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpdate

`func (o *V1DaemonSetUpdateStrategy) GetRollingUpdate() V1RollingUpdateDaemonSet`

GetRollingUpdate returns the RollingUpdate field if non-nil, zero value otherwise.

### GetRollingUpdateOk

`func (o *V1DaemonSetUpdateStrategy) GetRollingUpdateOk() (*V1RollingUpdateDaemonSet, bool)`

GetRollingUpdateOk returns a tuple with the RollingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpdate

`func (o *V1DaemonSetUpdateStrategy) SetRollingUpdate(v V1RollingUpdateDaemonSet)`

SetRollingUpdate sets RollingUpdate field to given value.

### HasRollingUpdate

`func (o *V1DaemonSetUpdateStrategy) HasRollingUpdate() bool`

HasRollingUpdate returns a boolean if a field has been set.

### GetType

`func (o *V1DaemonSetUpdateStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1DaemonSetUpdateStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1DaemonSetUpdateStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1DaemonSetUpdateStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


