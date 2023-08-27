# Troubleshoot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]TroubleshootDetailsInner**](TroubleshootDetailsInner.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTroubleshoot

`func NewTroubleshoot() *Troubleshoot`

NewTroubleshoot instantiates a new Troubleshoot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTroubleshootWithDefaults

`func NewTroubleshootWithDefaults() *Troubleshoot`

NewTroubleshootWithDefaults instantiates a new Troubleshoot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Troubleshoot) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Troubleshoot) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Troubleshoot) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Troubleshoot) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDetails

`func (o *Troubleshoot) GetDetails() []TroubleshootDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *Troubleshoot) GetDetailsOk() (*[]TroubleshootDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *Troubleshoot) SetDetails(v []TroubleshootDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *Troubleshoot) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Troubleshoot) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Troubleshoot) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Troubleshoot) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Troubleshoot) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


