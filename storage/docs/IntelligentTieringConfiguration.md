# IntelligentTieringConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**IntelligentTieringId**](IntelligentTieringId.md) |  | 
**Filter** | Pointer to [**IntelligentTieringFilter**](IntelligentTieringFilter.md) |  | [optional] 
**Status** | [**IntelligentTieringStatus**](IntelligentTieringStatus.md) |  | 
**Tierings** | [**TieringList**](TieringList.md) |  | 

## Methods

### NewIntelligentTieringConfiguration

`func NewIntelligentTieringConfiguration(id IntelligentTieringId, status IntelligentTieringStatus, tierings TieringList, ) *IntelligentTieringConfiguration`

NewIntelligentTieringConfiguration instantiates a new IntelligentTieringConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntelligentTieringConfigurationWithDefaults

`func NewIntelligentTieringConfigurationWithDefaults() *IntelligentTieringConfiguration`

NewIntelligentTieringConfigurationWithDefaults instantiates a new IntelligentTieringConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntelligentTieringConfiguration) GetId() IntelligentTieringId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntelligentTieringConfiguration) GetIdOk() (*IntelligentTieringId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntelligentTieringConfiguration) SetId(v IntelligentTieringId)`

SetId sets Id field to given value.


### GetFilter

`func (o *IntelligentTieringConfiguration) GetFilter() IntelligentTieringFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *IntelligentTieringConfiguration) GetFilterOk() (*IntelligentTieringFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *IntelligentTieringConfiguration) SetFilter(v IntelligentTieringFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *IntelligentTieringConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetStatus

`func (o *IntelligentTieringConfiguration) GetStatus() IntelligentTieringStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntelligentTieringConfiguration) GetStatusOk() (*IntelligentTieringStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntelligentTieringConfiguration) SetStatus(v IntelligentTieringStatus)`

SetStatus sets Status field to given value.


### GetTierings

`func (o *IntelligentTieringConfiguration) GetTierings() TieringList`

GetTierings returns the Tierings field if non-nil, zero value otherwise.

### GetTieringsOk

`func (o *IntelligentTieringConfiguration) GetTieringsOk() (*TieringList, bool)`

GetTieringsOk returns a tuple with the Tierings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierings

`func (o *IntelligentTieringConfiguration) SetTierings(v TieringList)`

SetTierings sets Tierings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


