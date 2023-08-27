# PutBucketMetricsConfigurationRequestMetricsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**MetricsId**](MetricsId.md) |  | 
**Filter** | Pointer to [**MetricsFilter**](MetricsFilter.md) |  | [optional] 

## Methods

### NewPutBucketMetricsConfigurationRequestMetricsConfiguration

`func NewPutBucketMetricsConfigurationRequestMetricsConfiguration(id MetricsId, ) *PutBucketMetricsConfigurationRequestMetricsConfiguration`

NewPutBucketMetricsConfigurationRequestMetricsConfiguration instantiates a new PutBucketMetricsConfigurationRequestMetricsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutBucketMetricsConfigurationRequestMetricsConfigurationWithDefaults

`func NewPutBucketMetricsConfigurationRequestMetricsConfigurationWithDefaults() *PutBucketMetricsConfigurationRequestMetricsConfiguration`

NewPutBucketMetricsConfigurationRequestMetricsConfigurationWithDefaults instantiates a new PutBucketMetricsConfigurationRequestMetricsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) GetId() MetricsId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) GetIdOk() (*MetricsId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) SetId(v MetricsId)`

SetId sets Id field to given value.


### GetFilter

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) GetFilter() MetricsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) GetFilterOk() (*MetricsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) SetFilter(v MetricsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *PutBucketMetricsConfigurationRequestMetricsConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


