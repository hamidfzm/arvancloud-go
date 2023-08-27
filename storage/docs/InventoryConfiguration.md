# InventoryConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Destination** | [**InventoryDestination**](InventoryDestination.md) |  | 
**IsEnabled** | [**IsEnabled**](IsEnabled.md) |  | 
**Filter** | Pointer to [**InventoryFilter**](InventoryFilter.md) |  | [optional] 
**Id** | [**InventoryId**](InventoryId.md) |  | 
**IncludedObjectVersions** | [**InventoryIncludedObjectVersions**](InventoryIncludedObjectVersions.md) |  | 
**OptionalFields** | Pointer to [**InventoryOptionalFields**](InventoryOptionalFields.md) |  | [optional] 
**Schedule** | [**InventorySchedule**](InventorySchedule.md) |  | 

## Methods

### NewInventoryConfiguration

`func NewInventoryConfiguration(destination InventoryDestination, isEnabled IsEnabled, id InventoryId, includedObjectVersions InventoryIncludedObjectVersions, schedule InventorySchedule, ) *InventoryConfiguration`

NewInventoryConfiguration instantiates a new InventoryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInventoryConfigurationWithDefaults

`func NewInventoryConfigurationWithDefaults() *InventoryConfiguration`

NewInventoryConfigurationWithDefaults instantiates a new InventoryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *InventoryConfiguration) GetDestination() InventoryDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *InventoryConfiguration) GetDestinationOk() (*InventoryDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *InventoryConfiguration) SetDestination(v InventoryDestination)`

SetDestination sets Destination field to given value.


### GetIsEnabled

`func (o *InventoryConfiguration) GetIsEnabled() IsEnabled`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *InventoryConfiguration) GetIsEnabledOk() (*IsEnabled, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *InventoryConfiguration) SetIsEnabled(v IsEnabled)`

SetIsEnabled sets IsEnabled field to given value.


### GetFilter

`func (o *InventoryConfiguration) GetFilter() InventoryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *InventoryConfiguration) GetFilterOk() (*InventoryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *InventoryConfiguration) SetFilter(v InventoryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *InventoryConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *InventoryConfiguration) GetId() InventoryId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InventoryConfiguration) GetIdOk() (*InventoryId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InventoryConfiguration) SetId(v InventoryId)`

SetId sets Id field to given value.


### GetIncludedObjectVersions

`func (o *InventoryConfiguration) GetIncludedObjectVersions() InventoryIncludedObjectVersions`

GetIncludedObjectVersions returns the IncludedObjectVersions field if non-nil, zero value otherwise.

### GetIncludedObjectVersionsOk

`func (o *InventoryConfiguration) GetIncludedObjectVersionsOk() (*InventoryIncludedObjectVersions, bool)`

GetIncludedObjectVersionsOk returns a tuple with the IncludedObjectVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedObjectVersions

`func (o *InventoryConfiguration) SetIncludedObjectVersions(v InventoryIncludedObjectVersions)`

SetIncludedObjectVersions sets IncludedObjectVersions field to given value.


### GetOptionalFields

`func (o *InventoryConfiguration) GetOptionalFields() InventoryOptionalFields`

GetOptionalFields returns the OptionalFields field if non-nil, zero value otherwise.

### GetOptionalFieldsOk

`func (o *InventoryConfiguration) GetOptionalFieldsOk() (*InventoryOptionalFields, bool)`

GetOptionalFieldsOk returns a tuple with the OptionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalFields

`func (o *InventoryConfiguration) SetOptionalFields(v InventoryOptionalFields)`

SetOptionalFields sets OptionalFields field to given value.

### HasOptionalFields

`func (o *InventoryConfiguration) HasOptionalFields() bool`

HasOptionalFields returns a boolean if a field has been set.

### GetSchedule

`func (o *InventoryConfiguration) GetSchedule() InventorySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *InventoryConfiguration) GetScheduleOk() (*InventorySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *InventoryConfiguration) SetSchedule(v InventorySchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


