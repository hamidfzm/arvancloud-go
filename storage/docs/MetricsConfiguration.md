# MetricsConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**MetricsId**](MetricsId.md) |  | 
**Filter** | Pointer to [**MetricsFilter**](MetricsFilter.md) |  | [optional] 

## Methods

### NewMetricsConfiguration

`func NewMetricsConfiguration(id MetricsId, ) *MetricsConfiguration`

NewMetricsConfiguration instantiates a new MetricsConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsConfigurationWithDefaults

`func NewMetricsConfigurationWithDefaults() *MetricsConfiguration`

NewMetricsConfigurationWithDefaults instantiates a new MetricsConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MetricsConfiguration) GetId() MetricsId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsConfiguration) GetIdOk() (*MetricsId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsConfiguration) SetId(v MetricsId)`

SetId sets Id field to given value.


### GetFilter

`func (o *MetricsConfiguration) GetFilter() MetricsFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MetricsConfiguration) GetFilterOk() (*MetricsFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MetricsConfiguration) SetFilter(v MetricsFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MetricsConfiguration) HasFilter() bool`

HasFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


