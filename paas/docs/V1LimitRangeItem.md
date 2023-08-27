# V1LimitRangeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | Pointer to **map[string]interface{}** | Default resource requirement limit value by resource name if resource limit is omitted. | [optional] 
**DefaultRequest** | Pointer to **map[string]interface{}** | DefaultRequest is the default resource requirement request value by resource name if resource request is omitted. | [optional] 
**Max** | Pointer to **map[string]interface{}** | Max usage constraints on this kind by resource name. | [optional] 
**MaxLimitRequestRatio** | Pointer to **map[string]interface{}** | MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource. | [optional] 
**Min** | Pointer to **map[string]interface{}** | Min usage constraints on this kind by resource name. | [optional] 
**Type** | Pointer to **string** | Type of resource that this limit applies to. | [optional] 

## Methods

### NewV1LimitRangeItem

`func NewV1LimitRangeItem() *V1LimitRangeItem`

NewV1LimitRangeItem instantiates a new V1LimitRangeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LimitRangeItemWithDefaults

`func NewV1LimitRangeItemWithDefaults() *V1LimitRangeItem`

NewV1LimitRangeItemWithDefaults instantiates a new V1LimitRangeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *V1LimitRangeItem) GetDefault() map[string]interface{}`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *V1LimitRangeItem) GetDefaultOk() (*map[string]interface{}, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *V1LimitRangeItem) SetDefault(v map[string]interface{})`

SetDefault sets Default field to given value.

### HasDefault

`func (o *V1LimitRangeItem) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDefaultRequest

`func (o *V1LimitRangeItem) GetDefaultRequest() map[string]interface{}`

GetDefaultRequest returns the DefaultRequest field if non-nil, zero value otherwise.

### GetDefaultRequestOk

`func (o *V1LimitRangeItem) GetDefaultRequestOk() (*map[string]interface{}, bool)`

GetDefaultRequestOk returns a tuple with the DefaultRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRequest

`func (o *V1LimitRangeItem) SetDefaultRequest(v map[string]interface{})`

SetDefaultRequest sets DefaultRequest field to given value.

### HasDefaultRequest

`func (o *V1LimitRangeItem) HasDefaultRequest() bool`

HasDefaultRequest returns a boolean if a field has been set.

### GetMax

`func (o *V1LimitRangeItem) GetMax() map[string]interface{}`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *V1LimitRangeItem) GetMaxOk() (*map[string]interface{}, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *V1LimitRangeItem) SetMax(v map[string]interface{})`

SetMax sets Max field to given value.

### HasMax

`func (o *V1LimitRangeItem) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMaxLimitRequestRatio

`func (o *V1LimitRangeItem) GetMaxLimitRequestRatio() map[string]interface{}`

GetMaxLimitRequestRatio returns the MaxLimitRequestRatio field if non-nil, zero value otherwise.

### GetMaxLimitRequestRatioOk

`func (o *V1LimitRangeItem) GetMaxLimitRequestRatioOk() (*map[string]interface{}, bool)`

GetMaxLimitRequestRatioOk returns a tuple with the MaxLimitRequestRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLimitRequestRatio

`func (o *V1LimitRangeItem) SetMaxLimitRequestRatio(v map[string]interface{})`

SetMaxLimitRequestRatio sets MaxLimitRequestRatio field to given value.

### HasMaxLimitRequestRatio

`func (o *V1LimitRangeItem) HasMaxLimitRequestRatio() bool`

HasMaxLimitRequestRatio returns a boolean if a field has been set.

### GetMin

`func (o *V1LimitRangeItem) GetMin() map[string]interface{}`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *V1LimitRangeItem) GetMinOk() (*map[string]interface{}, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *V1LimitRangeItem) SetMin(v map[string]interface{})`

SetMin sets Min field to given value.

### HasMin

`func (o *V1LimitRangeItem) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetType

`func (o *V1LimitRangeItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1LimitRangeItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1LimitRangeItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1LimitRangeItem) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


