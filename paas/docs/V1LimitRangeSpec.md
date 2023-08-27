# V1LimitRangeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limits** | [**[]V1LimitRangeItem**](V1LimitRangeItem.md) | Limits is the list of LimitRangeItem objects that are enforced. | 

## Methods

### NewV1LimitRangeSpec

`func NewV1LimitRangeSpec(limits []V1LimitRangeItem, ) *V1LimitRangeSpec`

NewV1LimitRangeSpec instantiates a new V1LimitRangeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LimitRangeSpecWithDefaults

`func NewV1LimitRangeSpecWithDefaults() *V1LimitRangeSpec`

NewV1LimitRangeSpecWithDefaults instantiates a new V1LimitRangeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimits

`func (o *V1LimitRangeSpec) GetLimits() []V1LimitRangeItem`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *V1LimitRangeSpec) GetLimitsOk() (*[]V1LimitRangeItem, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *V1LimitRangeSpec) SetLimits(v []V1LimitRangeItem)`

SetLimits sets Limits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


