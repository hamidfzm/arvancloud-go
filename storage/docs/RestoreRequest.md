# RestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | Pointer to **int32** |  | [optional] 
**GlacierJobParameters** | Pointer to [**RestoreRequestGlacierJobParameters**](RestoreRequestGlacierJobParameters.md) |  | [optional] 
**Type** | Pointer to [**RestoreRequestType**](RestoreRequestType.md) |  | [optional] 
**Tier** | Pointer to [**Tier**](Tier.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SelectParameters** | Pointer to [**RestoreRequestSelectParameters**](RestoreRequestSelectParameters.md) |  | [optional] 
**OutputLocation** | Pointer to [**RestoreRequestOutputLocation**](RestoreRequestOutputLocation.md) |  | [optional] 

## Methods

### NewRestoreRequest

`func NewRestoreRequest() *RestoreRequest`

NewRestoreRequest instantiates a new RestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreRequestWithDefaults

`func NewRestoreRequestWithDefaults() *RestoreRequest`

NewRestoreRequestWithDefaults instantiates a new RestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *RestoreRequest) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *RestoreRequest) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *RestoreRequest) SetDays(v int32)`

SetDays sets Days field to given value.

### HasDays

`func (o *RestoreRequest) HasDays() bool`

HasDays returns a boolean if a field has been set.

### GetGlacierJobParameters

`func (o *RestoreRequest) GetGlacierJobParameters() RestoreRequestGlacierJobParameters`

GetGlacierJobParameters returns the GlacierJobParameters field if non-nil, zero value otherwise.

### GetGlacierJobParametersOk

`func (o *RestoreRequest) GetGlacierJobParametersOk() (*RestoreRequestGlacierJobParameters, bool)`

GetGlacierJobParametersOk returns a tuple with the GlacierJobParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlacierJobParameters

`func (o *RestoreRequest) SetGlacierJobParameters(v RestoreRequestGlacierJobParameters)`

SetGlacierJobParameters sets GlacierJobParameters field to given value.

### HasGlacierJobParameters

`func (o *RestoreRequest) HasGlacierJobParameters() bool`

HasGlacierJobParameters returns a boolean if a field has been set.

### GetType

`func (o *RestoreRequest) GetType() RestoreRequestType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RestoreRequest) GetTypeOk() (*RestoreRequestType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RestoreRequest) SetType(v RestoreRequestType)`

SetType sets Type field to given value.

### HasType

`func (o *RestoreRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTier

`func (o *RestoreRequest) GetTier() Tier`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *RestoreRequest) GetTierOk() (*Tier, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *RestoreRequest) SetTier(v Tier)`

SetTier sets Tier field to given value.

### HasTier

`func (o *RestoreRequest) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetDescription

`func (o *RestoreRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestoreRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestoreRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestoreRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSelectParameters

`func (o *RestoreRequest) GetSelectParameters() RestoreRequestSelectParameters`

GetSelectParameters returns the SelectParameters field if non-nil, zero value otherwise.

### GetSelectParametersOk

`func (o *RestoreRequest) GetSelectParametersOk() (*RestoreRequestSelectParameters, bool)`

GetSelectParametersOk returns a tuple with the SelectParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectParameters

`func (o *RestoreRequest) SetSelectParameters(v RestoreRequestSelectParameters)`

SetSelectParameters sets SelectParameters field to given value.

### HasSelectParameters

`func (o *RestoreRequest) HasSelectParameters() bool`

HasSelectParameters returns a boolean if a field has been set.

### GetOutputLocation

`func (o *RestoreRequest) GetOutputLocation() RestoreRequestOutputLocation`

GetOutputLocation returns the OutputLocation field if non-nil, zero value otherwise.

### GetOutputLocationOk

`func (o *RestoreRequest) GetOutputLocationOk() (*RestoreRequestOutputLocation, bool)`

GetOutputLocationOk returns a tuple with the OutputLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputLocation

`func (o *RestoreRequest) SetOutputLocation(v RestoreRequestOutputLocation)`

SetOutputLocation sets OutputLocation field to given value.

### HasOutputLocation

`func (o *RestoreRequest) HasOutputLocation() bool`

HasOutputLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


