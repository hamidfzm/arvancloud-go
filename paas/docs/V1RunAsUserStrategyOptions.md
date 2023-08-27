# V1RunAsUserStrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Type is the strategy that will dictate what RunAsUser is used in the SecurityContext. | [optional] 
**Uid** | Pointer to **int64** | UID is the user id that containers must run as.  Required for the MustRunAs strategy if not using namespace/service account allocated uids. | [optional] 
**UidRangeMax** | Pointer to **int64** | UIDRangeMax defines the max value for a strategy that allocates by range. | [optional] 
**UidRangeMin** | Pointer to **int64** | UIDRangeMin defines the min value for a strategy that allocates by range. | [optional] 

## Methods

### NewV1RunAsUserStrategyOptions

`func NewV1RunAsUserStrategyOptions() *V1RunAsUserStrategyOptions`

NewV1RunAsUserStrategyOptions instantiates a new V1RunAsUserStrategyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1RunAsUserStrategyOptionsWithDefaults

`func NewV1RunAsUserStrategyOptionsWithDefaults() *V1RunAsUserStrategyOptions`

NewV1RunAsUserStrategyOptionsWithDefaults instantiates a new V1RunAsUserStrategyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1RunAsUserStrategyOptions) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1RunAsUserStrategyOptions) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1RunAsUserStrategyOptions) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1RunAsUserStrategyOptions) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUid

`func (o *V1RunAsUserStrategyOptions) GetUid() int64`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *V1RunAsUserStrategyOptions) GetUidOk() (*int64, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *V1RunAsUserStrategyOptions) SetUid(v int64)`

SetUid sets Uid field to given value.

### HasUid

`func (o *V1RunAsUserStrategyOptions) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUidRangeMax

`func (o *V1RunAsUserStrategyOptions) GetUidRangeMax() int64`

GetUidRangeMax returns the UidRangeMax field if non-nil, zero value otherwise.

### GetUidRangeMaxOk

`func (o *V1RunAsUserStrategyOptions) GetUidRangeMaxOk() (*int64, bool)`

GetUidRangeMaxOk returns a tuple with the UidRangeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidRangeMax

`func (o *V1RunAsUserStrategyOptions) SetUidRangeMax(v int64)`

SetUidRangeMax sets UidRangeMax field to given value.

### HasUidRangeMax

`func (o *V1RunAsUserStrategyOptions) HasUidRangeMax() bool`

HasUidRangeMax returns a boolean if a field has been set.

### GetUidRangeMin

`func (o *V1RunAsUserStrategyOptions) GetUidRangeMin() int64`

GetUidRangeMin returns the UidRangeMin field if non-nil, zero value otherwise.

### GetUidRangeMinOk

`func (o *V1RunAsUserStrategyOptions) GetUidRangeMinOk() (*int64, bool)`

GetUidRangeMinOk returns a tuple with the UidRangeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidRangeMin

`func (o *V1RunAsUserStrategyOptions) SetUidRangeMin(v int64)`

SetUidRangeMin sets UidRangeMin field to given value.

### HasUidRangeMin

`func (o *V1RunAsUserStrategyOptions) HasUidRangeMin() bool`

HasUidRangeMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


