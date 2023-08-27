# V1NamedTagEventList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]V1TagEventCondition**](V1TagEventCondition.md) | Conditions is an array of conditions that apply to the tag event list. | [optional] 
**Items** | [**[]V1TagEvent**](V1TagEvent.md) | Standard object&#39;s metadata. | 
**Tag** | **string** | Tag is the tag for which the history is recorded | 

## Methods

### NewV1NamedTagEventList

`func NewV1NamedTagEventList(items []V1TagEvent, tag string, ) *V1NamedTagEventList`

NewV1NamedTagEventList instantiates a new V1NamedTagEventList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NamedTagEventListWithDefaults

`func NewV1NamedTagEventListWithDefaults() *V1NamedTagEventList`

NewV1NamedTagEventListWithDefaults instantiates a new V1NamedTagEventList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *V1NamedTagEventList) GetConditions() []V1TagEventCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1NamedTagEventList) GetConditionsOk() (*[]V1TagEventCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1NamedTagEventList) SetConditions(v []V1TagEventCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1NamedTagEventList) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetItems

`func (o *V1NamedTagEventList) GetItems() []V1TagEvent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *V1NamedTagEventList) GetItemsOk() (*[]V1TagEvent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *V1NamedTagEventList) SetItems(v []V1TagEvent)`

SetItems sets Items field to given value.


### GetTag

`func (o *V1NamedTagEventList) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1NamedTagEventList) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1NamedTagEventList) SetTag(v string)`

SetTag sets Tag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


