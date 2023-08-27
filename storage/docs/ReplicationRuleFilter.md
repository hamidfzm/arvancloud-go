# ReplicationRuleFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefix** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to [**ReplicationRuleFilterTag**](ReplicationRuleFilterTag.md) |  | [optional] 
**And** | Pointer to [**ReplicationRuleFilterAnd**](ReplicationRuleFilterAnd.md) |  | [optional] 

## Methods

### NewReplicationRuleFilter

`func NewReplicationRuleFilter() *ReplicationRuleFilter`

NewReplicationRuleFilter instantiates a new ReplicationRuleFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRuleFilterWithDefaults

`func NewReplicationRuleFilterWithDefaults() *ReplicationRuleFilter`

NewReplicationRuleFilterWithDefaults instantiates a new ReplicationRuleFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefix

`func (o *ReplicationRuleFilter) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ReplicationRuleFilter) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ReplicationRuleFilter) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ReplicationRuleFilter) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetTag

`func (o *ReplicationRuleFilter) GetTag() ReplicationRuleFilterTag`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ReplicationRuleFilter) GetTagOk() (*ReplicationRuleFilterTag, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ReplicationRuleFilter) SetTag(v ReplicationRuleFilterTag)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ReplicationRuleFilter) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetAnd

`func (o *ReplicationRuleFilter) GetAnd() ReplicationRuleFilterAnd`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *ReplicationRuleFilter) GetAndOk() (*ReplicationRuleFilterAnd, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *ReplicationRuleFilter) SetAnd(v ReplicationRuleFilterAnd)`

SetAnd sets And field to given value.

### HasAnd

`func (o *ReplicationRuleFilter) HasAnd() bool`

HasAnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


