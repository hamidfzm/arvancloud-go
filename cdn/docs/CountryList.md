# CountryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** | The name of the country | [optional] 
**Code** | Pointer to **string** | The 2-letter country code | [optional] 
**Requests** | Pointer to **int32** | The number of requests from the country | [optional] 
**Traffics** | Pointer to **int32** | The amount of traffic from the country | [optional] 

## Methods

### NewCountryList

`func NewCountryList() *CountryList`

NewCountryList instantiates a new CountryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryListWithDefaults

`func NewCountryListWithDefaults() *CountryList`

NewCountryListWithDefaults instantiates a new CountryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CountryList) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CountryList) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CountryList) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CountryList) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCode

`func (o *CountryList) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CountryList) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CountryList) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CountryList) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetRequests

`func (o *CountryList) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CountryList) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CountryList) SetRequests(v int32)`

SetRequests sets Requests field to given value.

### HasRequests

`func (o *CountryList) HasRequests() bool`

HasRequests returns a boolean if a field has been set.

### GetTraffics

`func (o *CountryList) GetTraffics() int32`

GetTraffics returns the Traffics field if non-nil, zero value otherwise.

### GetTrafficsOk

`func (o *CountryList) GetTrafficsOk() (*int32, bool)`

GetTrafficsOk returns a tuple with the Traffics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffics

`func (o *CountryList) SetTraffics(v int32)`

SetTraffics sets Traffics field to given value.

### HasTraffics

`func (o *CountryList) HasTraffics() bool`

HasTraffics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


