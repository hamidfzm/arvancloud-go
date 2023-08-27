# CORSRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**AllowedHeaders** | Pointer to [**Array**](array.md) |  | [optional] 
**AllowedMethods** | [**Array**](array.md) |  | 
**AllowedOrigins** | [**Array**](array.md) |  | 
**ExposeHeaders** | Pointer to [**Array**](array.md) |  | [optional] 
**MaxAgeSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewCORSRule

`func NewCORSRule(allowedMethods Array, allowedOrigins Array, ) *CORSRule`

NewCORSRule instantiates a new CORSRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCORSRuleWithDefaults

`func NewCORSRuleWithDefaults() *CORSRule`

NewCORSRuleWithDefaults instantiates a new CORSRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *CORSRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *CORSRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *CORSRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *CORSRule) HasID() bool`

HasID returns a boolean if a field has been set.

### GetAllowedHeaders

`func (o *CORSRule) GetAllowedHeaders() Array`

GetAllowedHeaders returns the AllowedHeaders field if non-nil, zero value otherwise.

### GetAllowedHeadersOk

`func (o *CORSRule) GetAllowedHeadersOk() (*Array, bool)`

GetAllowedHeadersOk returns a tuple with the AllowedHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedHeaders

`func (o *CORSRule) SetAllowedHeaders(v Array)`

SetAllowedHeaders sets AllowedHeaders field to given value.

### HasAllowedHeaders

`func (o *CORSRule) HasAllowedHeaders() bool`

HasAllowedHeaders returns a boolean if a field has been set.

### GetAllowedMethods

`func (o *CORSRule) GetAllowedMethods() Array`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *CORSRule) GetAllowedMethodsOk() (*Array, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *CORSRule) SetAllowedMethods(v Array)`

SetAllowedMethods sets AllowedMethods field to given value.


### GetAllowedOrigins

`func (o *CORSRule) GetAllowedOrigins() Array`

GetAllowedOrigins returns the AllowedOrigins field if non-nil, zero value otherwise.

### GetAllowedOriginsOk

`func (o *CORSRule) GetAllowedOriginsOk() (*Array, bool)`

GetAllowedOriginsOk returns a tuple with the AllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedOrigins

`func (o *CORSRule) SetAllowedOrigins(v Array)`

SetAllowedOrigins sets AllowedOrigins field to given value.


### GetExposeHeaders

`func (o *CORSRule) GetExposeHeaders() Array`

GetExposeHeaders returns the ExposeHeaders field if non-nil, zero value otherwise.

### GetExposeHeadersOk

`func (o *CORSRule) GetExposeHeadersOk() (*Array, bool)`

GetExposeHeadersOk returns a tuple with the ExposeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposeHeaders

`func (o *CORSRule) SetExposeHeaders(v Array)`

SetExposeHeaders sets ExposeHeaders field to given value.

### HasExposeHeaders

`func (o *CORSRule) HasExposeHeaders() bool`

HasExposeHeaders returns a boolean if a field has been set.

### GetMaxAgeSeconds

`func (o *CORSRule) GetMaxAgeSeconds() int32`

GetMaxAgeSeconds returns the MaxAgeSeconds field if non-nil, zero value otherwise.

### GetMaxAgeSecondsOk

`func (o *CORSRule) GetMaxAgeSecondsOk() (*int32, bool)`

GetMaxAgeSecondsOk returns a tuple with the MaxAgeSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeSeconds

`func (o *CORSRule) SetMaxAgeSeconds(v int32)`

SetMaxAgeSeconds sets MaxAgeSeconds field to given value.

### HasMaxAgeSeconds

`func (o *CORSRule) HasMaxAgeSeconds() bool`

HasMaxAgeSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


