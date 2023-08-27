# AnalyticsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**AnalyticsId**](AnalyticsId.md) |  | 
**Filter** | Pointer to [**AnalyticsFilter**](AnalyticsFilter.md) |  | [optional] 
**StorageClassAnalysis** | [**StorageClassAnalysis**](StorageClassAnalysis.md) |  | 

## Methods

### NewAnalyticsConfiguration

`func NewAnalyticsConfiguration(id AnalyticsId, storageClassAnalysis StorageClassAnalysis, ) *AnalyticsConfiguration`

NewAnalyticsConfiguration instantiates a new AnalyticsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsConfigurationWithDefaults

`func NewAnalyticsConfigurationWithDefaults() *AnalyticsConfiguration`

NewAnalyticsConfigurationWithDefaults instantiates a new AnalyticsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsConfiguration) GetId() AnalyticsId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsConfiguration) GetIdOk() (*AnalyticsId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsConfiguration) SetId(v AnalyticsId)`

SetId sets Id field to given value.


### GetFilter

`func (o *AnalyticsConfiguration) GetFilter() AnalyticsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *AnalyticsConfiguration) GetFilterOk() (*AnalyticsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *AnalyticsConfiguration) SetFilter(v AnalyticsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *AnalyticsConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetStorageClassAnalysis

`func (o *AnalyticsConfiguration) GetStorageClassAnalysis() StorageClassAnalysis`

GetStorageClassAnalysis returns the StorageClassAnalysis field if non-nil, zero value otherwise.

### GetStorageClassAnalysisOk

`func (o *AnalyticsConfiguration) GetStorageClassAnalysisOk() (*StorageClassAnalysis, bool)`

GetStorageClassAnalysisOk returns a tuple with the StorageClassAnalysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageClassAnalysis

`func (o *AnalyticsConfiguration) SetStorageClassAnalysis(v StorageClassAnalysis)`

SetStorageClassAnalysis sets StorageClassAnalysis field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


