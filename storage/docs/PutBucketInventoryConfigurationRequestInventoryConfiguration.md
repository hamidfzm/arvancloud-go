# PutBucketInventoryConfigurationRequestInventoryConfiguration

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

### NewPutBucketInventoryConfigurationRequestInventoryConfiguration

`func NewPutBucketInventoryConfigurationRequestInventoryConfiguration(destination InventoryDestination, isEnabled IsEnabled, id InventoryId, includedObjectVersions InventoryIncludedObjectVersions, schedule InventorySchedule, ) *PutBucketInventoryConfigurationRequestInventoryConfiguration`

NewPutBucketInventoryConfigurationRequestInventoryConfiguration instantiates a new PutBucketInventoryConfigurationRequestInventoryConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketInventoryConfigurationRequestInventoryConfigurationWithDefaults

`func NewPutBucketInventoryConfigurationRequestInventoryConfigurationWithDefaults() *PutBucketInventoryConfigurationRequestInventoryConfiguration`

NewPutBucketInventoryConfigurationRequestInventoryConfigurationWithDefaults instantiates a new PutBucketInventoryConfigurationRequestInventoryConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestination

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetDestination() InventoryDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetDestinationOk() (*InventoryDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetDestination(v InventoryDestination)`

SetDestination sets Destination field to given value.


### GetIsEnabled

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetIsEnabled() IsEnabled`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetIsEnabledOk() (*IsEnabled, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetIsEnabled(v IsEnabled)`

SetIsEnabled sets IsEnabled field to given value.


### GetFilter

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetFilter() InventoryFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetFilterOk() (*InventoryFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetFilter(v InventoryFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetId() InventoryId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetIdOk() (*InventoryId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetId(v InventoryId)`

SetId sets Id field to given value.


### GetIncludedObjectVersions

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetIncludedObjectVersions() InventoryIncludedObjectVersions`

GetIncludedObjectVersions returns the IncludedObjectVersions field if non-nil, zero value otherwise.

### GetIncludedObjectVersionsOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetIncludedObjectVersionsOk() (*InventoryIncludedObjectVersions, bool)`

GetIncludedObjectVersionsOk returns a tuple with the IncludedObjectVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedObjectVersions

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetIncludedObjectVersions(v InventoryIncludedObjectVersions)`

SetIncludedObjectVersions sets IncludedObjectVersions field to given value.


### GetOptionalFields

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetOptionalFields() InventoryOptionalFields`

GetOptionalFields returns the OptionalFields field if non-nil, zero value otherwise.

### GetOptionalFieldsOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetOptionalFieldsOk() (*InventoryOptionalFields, bool)`

GetOptionalFieldsOk returns a tuple with the OptionalFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalFields

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetOptionalFields(v InventoryOptionalFields)`

SetOptionalFields sets OptionalFields field to given value.

### HasOptionalFields

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) HasOptionalFields() bool`

HasOptionalFields returns a boolean if a field has been set.

### GetSchedule

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetSchedule() InventorySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) GetScheduleOk() (*InventorySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PutBucketInventoryConfigurationRequestInventoryConfiguration) SetSchedule(v InventorySchedule)`

SetSchedule sets Schedule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


