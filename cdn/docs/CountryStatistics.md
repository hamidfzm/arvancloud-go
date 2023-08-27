# CountryStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The name of the country | [optional] 
**Requests** | Pointer to **int32** | The number of requests from the country | [optional] 
**Traffics** | Pointer to **int32** | The amount of traffic from the country | [optional] 

## Methods

### NewCountryStatistics

`func NewCountryStatistics() *CountryStatistics`

NewCountryStatistics instantiates a new CountryStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryStatisticsWithDefaults

`func NewCountryStatisticsWithDefaults() *CountryStatistics`

NewCountryStatisticsWithDefaults instantiates a new CountryStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CountryStatistics) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CountryStatistics) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CountryStatistics) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CountryStatistics) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRequests

`func (o *CountryStatistics) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CountryStatistics) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CountryStatistics) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CountryStatistics) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTraffics

`func (o *CountryStatistics) GetTraffics() int32`

GetTraffics returns the Traffics field if non-nil, zero value otherwise.

### GetTrafficsOk

`func (o *CountryStatistics) GetTrafficsOk() (*int32, bool)`

GetTrafficsOk returns a tuple with the Traffics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffics

`func (o *CountryStatistics) SetTraffics(v int32)`

SetTraffics sets Traffics field to given value.

### HasTraffics

`func (o *CountryStatistics) HasTraffics() bool`

HasTraffics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


