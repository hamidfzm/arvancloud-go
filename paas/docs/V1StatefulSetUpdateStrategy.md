# V1StatefulSetUpdateStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollingUpdate** | Pointer to [**V1RollingUpdateStatefulSetStrategy**](V1RollingUpdateStatefulSetStrategy.md) |  | [optional] 
**Type** | Pointer to **string** | Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate. | [optional] 

## Methods

### NewV1StatefulSetUpdateStrategy

`func NewV1StatefulSetUpdateStrategy() *V1StatefulSetUpdateStrategy`

NewV1StatefulSetUpdateStrategy instantiates a new V1StatefulSetUpdateStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StatefulSetUpdateStrategyWithDefaults

`func NewV1StatefulSetUpdateStrategyWithDefaults() *V1StatefulSetUpdateStrategy`

NewV1StatefulSetUpdateStrategyWithDefaults instantiates a new V1StatefulSetUpdateStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollingUpdate

`func (o *V1StatefulSetUpdateStrategy) GetRollingUpdate() V1RollingUpdateStatefulSetStrategy`

GetRollingUpdate returns the RollingUpdate field if non-nil, zero value otherwise.

### GetRollingUpdateOk

`func (o *V1StatefulSetUpdateStrategy) GetRollingUpdateOk() (*V1RollingUpdateStatefulSetStrategy, bool)`

GetRollingUpdateOk returns a tuple with the RollingUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollingUpdate

`func (o *V1StatefulSetUpdateStrategy) SetRollingUpdate(v V1RollingUpdateStatefulSetStrategy)`

SetRollingUpdate sets RollingUpdate field to given value.

### HasRollingUpdate

`func (o *V1StatefulSetUpdateStrategy) HasRollingUpdate() bool`

HasRollingUpdate returns a boolean if a field has been set.

### GetType

`func (o *V1StatefulSetUpdateStrategy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1StatefulSetUpdateStrategy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1StatefulSetUpdateStrategy) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1StatefulSetUpdateStrategy) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


