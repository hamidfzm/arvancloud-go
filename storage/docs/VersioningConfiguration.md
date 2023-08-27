# VersioningConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MFADelete** | Pointer to [**MFADelete**](MFADelete.md) |  | [optional] 
**Status** | Pointer to [**BucketVersioningStatus**](BucketVersioningStatus.md) |  | [optional] 

## Methods

### NewVersioningConfiguration

`func NewVersioningConfiguration() *VersioningConfiguration`

NewVersioningConfiguration instantiates a new VersioningConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersioningConfigurationWithDefaults

`func NewVersioningConfigurationWithDefaults() *VersioningConfiguration`

NewVersioningConfigurationWithDefaults instantiates a new VersioningConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMFADelete

`func (o *VersioningConfiguration) GetMFADelete() MFADelete`

GetMFADelete returns the MFADelete field if non-nil, zero value otherwise.

### GetMFADeleteOk

`func (o *VersioningConfiguration) GetMFADeleteOk() (*MFADelete, bool)`

GetMFADeleteOk returns a tuple with the MFADelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMFADelete

`func (o *VersioningConfiguration) SetMFADelete(v MFADelete)`

SetMFADelete sets MFADelete field to given value.

### HasMFADelete

`func (o *VersioningConfiguration) HasMFADelete() bool`

HasMFADelete returns a boolean if a field has been set.

### GetStatus

`func (o *VersioningConfiguration) GetStatus() BucketVersioningStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersioningConfiguration) GetStatusOk() (*BucketVersioningStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersioningConfiguration) SetStatus(v BucketVersioningStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VersioningConfiguration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


