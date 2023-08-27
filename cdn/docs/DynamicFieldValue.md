# DynamicFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**DynamicFieldType**](DynamicFieldType.md) |  | [optional] 
**Desc** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDynamicFieldValue

`func NewDynamicFieldValue() *DynamicFieldValue`

NewDynamicFieldValue instantiates a new DynamicFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicFieldValueWithDefaults

`func NewDynamicFieldValueWithDefaults() *DynamicFieldValue`

NewDynamicFieldValueWithDefaults instantiates a new DynamicFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DynamicFieldValue) GetValue() DynamicFieldType`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DynamicFieldValue) GetValueOk() (*DynamicFieldType, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DynamicFieldValue) SetValue(v DynamicFieldType)`

SetValue sets Value field to given value.

### HasValue

`func (o *DynamicFieldValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetDesc

`func (o *DynamicFieldValue) GetDesc() string`

GetDesc returns the Desc field if non-nil, zero value otherwise.

### GetDescOk

`func (o *DynamicFieldValue) GetDescOk() (*string, bool)`

GetDescOk returns a tuple with the Desc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesc

`func (o *DynamicFieldValue) SetDesc(v string)`

SetDesc sets Desc field to given value.

### HasDesc

`func (o *DynamicFieldValue) HasDesc() bool`

HasDesc returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DynamicFieldValue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DynamicFieldValue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DynamicFieldValue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DynamicFieldValue) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


