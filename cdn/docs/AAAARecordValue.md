# AAAARecordValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** |  | 
**Port** | Pointer to **NullableInt32** |  | [optional] 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**OriginalWeight** | Pointer to **int32** | This key shows itself here if the weights have been changed by Health Check. | [optional] [readonly] 
**Country** | Pointer to **NullableString** | ISO 3166 alpha-2 country code | [optional] 

## Methods

### NewAAAARecordValue

`func NewAAAARecordValue(ip string, ) *AAAARecordValue`

NewAAAARecordValue instantiates a new AAAARecordValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAAAARecordValueWithDefaults

`func NewAAAARecordValueWithDefaults() *AAAARecordValue`

NewAAAARecordValueWithDefaults instantiates a new AAAARecordValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *AAAARecordValue) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AAAARecordValue) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AAAARecordValue) SetIp(v string)`

SetIp sets Ip field to given value.


### GetPort

`func (o *AAAARecordValue) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AAAARecordValue) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AAAARecordValue) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *AAAARecordValue) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *AAAARecordValue) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *AAAARecordValue) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetWeight

`func (o *AAAARecordValue) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *AAAARecordValue) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *AAAARecordValue) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *AAAARecordValue) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *AAAARecordValue) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *AAAARecordValue) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetOriginalWeight

`func (o *AAAARecordValue) GetOriginalWeight() int32`

GetOriginalWeight returns the OriginalWeight field if non-nil, zero value otherwise.

### GetOriginalWeightOk

`func (o *AAAARecordValue) GetOriginalWeightOk() (*int32, bool)`

GetOriginalWeightOk returns a tuple with the OriginalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalWeight

`func (o *AAAARecordValue) SetOriginalWeight(v int32)`

SetOriginalWeight sets OriginalWeight field to given value.

### HasOriginalWeight

`func (o *AAAARecordValue) HasOriginalWeight() bool`

HasOriginalWeight returns a boolean if a field has been set.

### GetCountry

`func (o *AAAARecordValue) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AAAARecordValue) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AAAARecordValue) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AAAARecordValue) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *AAAARecordValue) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AAAARecordValue) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


