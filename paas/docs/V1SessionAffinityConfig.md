# V1SessionAffinityConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientIP** | Pointer to [**V1ClientIPConfig**](V1ClientIPConfig.md) |  | [optional] 

## Methods

### NewV1SessionAffinityConfig

`func NewV1SessionAffinityConfig() *V1SessionAffinityConfig`

NewV1SessionAffinityConfig instantiates a new V1SessionAffinityConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SessionAffinityConfigWithDefaults

`func NewV1SessionAffinityConfigWithDefaults() *V1SessionAffinityConfig`

NewV1SessionAffinityConfigWithDefaults instantiates a new V1SessionAffinityConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientIP

`func (o *V1SessionAffinityConfig) GetClientIP() V1ClientIPConfig`

GetClientIP returns the ClientIP field if non-nil, zero value otherwise.

### GetClientIPOk

`func (o *V1SessionAffinityConfig) GetClientIPOk() (*V1ClientIPConfig, bool)`

GetClientIPOk returns a tuple with the ClientIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIP

`func (o *V1SessionAffinityConfig) SetClientIP(v V1ClientIPConfig)`

SetClientIP sets ClientIP field to given value.

### HasClientIP

`func (o *V1SessionAffinityConfig) HasClientIP() bool`

HasClientIP returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


