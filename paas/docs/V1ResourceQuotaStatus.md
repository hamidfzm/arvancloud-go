# V1ResourceQuotaStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hard** | Pointer to **map[string]interface{}** | Hard is the set of enforced hard limits for each named resource. More info: https://kubernetes.io/docs/concepts/policy/resource-quotas/ | [optional] 
**Used** | Pointer to **map[string]interface{}** | Used is the current observed total usage of the resource in the namespace. | [optional] 

## Methods

### NewV1ResourceQuotaStatus

`func NewV1ResourceQuotaStatus() *V1ResourceQuotaStatus`

NewV1ResourceQuotaStatus instantiates a new V1ResourceQuotaStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ResourceQuotaStatusWithDefaults

`func NewV1ResourceQuotaStatusWithDefaults() *V1ResourceQuotaStatus`

NewV1ResourceQuotaStatusWithDefaults instantiates a new V1ResourceQuotaStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHard

`func (o *V1ResourceQuotaStatus) GetHard() map[string]interface{}`

GetHard returns the Hard field if non-nil, zero value otherwise.

### GetHardOk

`func (o *V1ResourceQuotaStatus) GetHardOk() (*map[string]interface{}, bool)`

GetHardOk returns a tuple with the Hard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHard

`func (o *V1ResourceQuotaStatus) SetHard(v map[string]interface{})`

SetHard sets Hard field to given value.

### HasHard

`func (o *V1ResourceQuotaStatus) HasHard() bool`

HasHard returns a boolean if a field has been set.

### GetUsed

`func (o *V1ResourceQuotaStatus) GetUsed() map[string]interface{}`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *V1ResourceQuotaStatus) GetUsedOk() (*map[string]interface{}, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *V1ResourceQuotaStatus) SetUsed(v map[string]interface{})`

SetUsed sets Used field to given value.

### HasUsed

`func (o *V1ResourceQuotaStatus) HasUsed() bool`

HasUsed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


