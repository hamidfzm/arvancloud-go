# Tiering

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** |  | 
**AccessTier** | [**IntelligentTieringAccessTier**](IntelligentTieringAccessTier.md) |  | 

## Methods

### NewTiering

`func NewTiering(days int32, accessTier IntelligentTieringAccessTier, ) *Tiering`

NewTiering instantiates a new Tiering object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTieringWithDefaults

`func NewTieringWithDefaults() *Tiering`

NewTieringWithDefaults instantiates a new Tiering object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *Tiering) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *Tiering) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *Tiering) SetDays(v int32)`

SetDays sets Days field to given value.


### GetAccessTier

`func (o *Tiering) GetAccessTier() IntelligentTieringAccessTier`

GetAccessTier returns the AccessTier field if non-nil, zero value otherwise.

### GetAccessTierOk

`func (o *Tiering) GetAccessTierOk() (*IntelligentTieringAccessTier, bool)`

GetAccessTierOk returns a tuple with the AccessTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTier

`func (o *Tiering) SetAccessTier(v IntelligentTieringAccessTier)`

SetAccessTier sets AccessTier field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


