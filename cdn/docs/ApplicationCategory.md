# ApplicationCategory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Order** | Pointer to **int32** |  | [optional] 
**NameTranslation** | Pointer to [**NullableApplicationCategoryNameTranslation**](ApplicationCategoryNameTranslation.md) |  | [optional] 
**Applications** | Pointer to [**[]ApplicationCategoryApplicationsInner**](ApplicationCategoryApplicationsInner.md) |  | [optional] 

## Methods

### NewApplicationCategory

`func NewApplicationCategory() *ApplicationCategory`

NewApplicationCategory instantiates a new ApplicationCategory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCategoryWithDefaults

`func NewApplicationCategoryWithDefaults() *ApplicationCategory`

NewApplicationCategoryWithDefaults instantiates a new ApplicationCategory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationCategory) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationCategory) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationCategory) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationCategory) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationCategory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCategory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCategory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationCategory) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *ApplicationCategory) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApplicationCategory) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApplicationCategory) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApplicationCategory) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetOrder

`func (o *ApplicationCategory) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *ApplicationCategory) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *ApplicationCategory) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *ApplicationCategory) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetNameTranslation

`func (o *ApplicationCategory) GetNameTranslation() ApplicationCategoryNameTranslation`

GetNameTranslation returns the NameTranslation field if non-nil, zero value otherwise.

### GetNameTranslationOk

`func (o *ApplicationCategory) GetNameTranslationOk() (*ApplicationCategoryNameTranslation, bool)`

GetNameTranslationOk returns a tuple with the NameTranslation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameTranslation

`func (o *ApplicationCategory) SetNameTranslation(v ApplicationCategoryNameTranslation)`

SetNameTranslation sets NameTranslation field to given value.

### HasNameTranslation

`func (o *ApplicationCategory) HasNameTranslation() bool`

HasNameTranslation returns a boolean if a field has been set.

### SetNameTranslationNil

`func (o *ApplicationCategory) SetNameTranslationNil(b bool)`

 SetNameTranslationNil sets the value for NameTranslation to be an explicit nil

### UnsetNameTranslation
`func (o *ApplicationCategory) UnsetNameTranslation()`

UnsetNameTranslation ensures that no value is present for NameTranslation, not even an explicit nil
### GetApplications

`func (o *ApplicationCategory) GetApplications() []ApplicationCategoryApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ApplicationCategory) GetApplicationsOk() (*[]ApplicationCategoryApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ApplicationCategory) SetApplications(v []ApplicationCategoryApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *ApplicationCategory) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


