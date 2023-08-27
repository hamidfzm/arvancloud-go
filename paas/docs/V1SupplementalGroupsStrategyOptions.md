# V1SupplementalGroupsStrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ranges** | Pointer to [**[]V1IDRange**](V1IDRange.md) | Ranges are the allowed ranges of supplemental groups.  If you would like to force a single supplemental group then supply a single range with the same start and end. | [optional] 
**Type** | Pointer to **string** | Type is the strategy that will dictate what supplemental groups is used in the SecurityContext. | [optional] 

## Methods

### NewV1SupplementalGroupsStrategyOptions

`func NewV1SupplementalGroupsStrategyOptions() *V1SupplementalGroupsStrategyOptions`

NewV1SupplementalGroupsStrategyOptions instantiates a new V1SupplementalGroupsStrategyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SupplementalGroupsStrategyOptionsWithDefaults

`func NewV1SupplementalGroupsStrategyOptionsWithDefaults() *V1SupplementalGroupsStrategyOptions`

NewV1SupplementalGroupsStrategyOptionsWithDefaults instantiates a new V1SupplementalGroupsStrategyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRanges

`func (o *V1SupplementalGroupsStrategyOptions) GetRanges() []V1IDRange`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *V1SupplementalGroupsStrategyOptions) GetRangesOk() (*[]V1IDRange, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *V1SupplementalGroupsStrategyOptions) SetRanges(v []V1IDRange)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *V1SupplementalGroupsStrategyOptions) HasRanges() bool`

HasRanges returns a boolean if a field has been set.

### GetType

`func (o *V1SupplementalGroupsStrategyOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1SupplementalGroupsStrategyOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1SupplementalGroupsStrategyOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1SupplementalGroupsStrategyOptions) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


