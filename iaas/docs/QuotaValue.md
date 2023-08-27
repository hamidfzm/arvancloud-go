# QuotaValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unlimited** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **int64** |  | [optional] 

## Methods

### NewQuotaValue

`func NewQuotaValue() *QuotaValue`

NewQuotaValue instantiates a new QuotaValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuotaValueWithDefaults

`func NewQuotaValueWithDefaults() *QuotaValue`

NewQuotaValueWithDefaults instantiates a new QuotaValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnlimited

`func (o *QuotaValue) GetUnlimited() bool`

GetUnlimited returns the Unlimited field if non-nil, zero value otherwise.

### GetUnlimitedOk

`func (o *QuotaValue) GetUnlimitedOk() (*bool, bool)`

GetUnlimitedOk returns a tuple with the Unlimited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlimited

`func (o *QuotaValue) SetUnlimited(v bool)`

SetUnlimited sets Unlimited field to given value.

### HasUnlimited

`func (o *QuotaValue) HasUnlimited() bool`

HasUnlimited returns a boolean if a field has been set.

### GetValue

`func (o *QuotaValue) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QuotaValue) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QuotaValue) SetValue(v int64)`

SetValue sets Value field to given value.

### HasValue

`func (o *QuotaValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


