# ReplicationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | Pointer to **string** |  | [optional] 
**Priority** | Pointer to **int32** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Filter** | Pointer to [**ReplicationRuleFilter**](ReplicationRuleFilter.md) |  | [optional] 
**Status** | [**ReplicationRuleStatus**](ReplicationRuleStatus.md) |  | 
**SourceSelectionCriteria** | Pointer to [**ReplicationRuleSourceSelectionCriteria**](ReplicationRuleSourceSelectionCriteria.md) |  | [optional] 
**ExistingObjectReplication** | Pointer to [**ReplicationRuleExistingObjectReplication**](ReplicationRuleExistingObjectReplication.md) |  | [optional] 
**Destination** | [**ReplicationRuleDestination**](ReplicationRuleDestination.md) |  | 
**DeleteMarkerReplication** | Pointer to [**DeleteMarkerReplication**](DeleteMarkerReplication.md) |  | [optional] 

## Methods

### NewReplicationRule

`func NewReplicationRule(status ReplicationRuleStatus, destination ReplicationRuleDestination, ) *ReplicationRule`

NewReplicationRule instantiates a new ReplicationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRuleWithDefaults

`func NewReplicationRuleWithDefaults() *ReplicationRule`

NewReplicationRuleWithDefaults instantiates a new ReplicationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *ReplicationRule) GetID() string`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *ReplicationRule) GetIDOk() (*string, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *ReplicationRule) SetID(v string)`

SetID sets ID field to given value.

### HasID

`func (o *ReplicationRule) HasID() bool`

HasID returns a boolean if a field has been set.

### GetPriority

`func (o *ReplicationRule) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *ReplicationRule) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *ReplicationRule) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *ReplicationRule) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### GetPrefix

`func (o *ReplicationRule) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *ReplicationRule) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *ReplicationRule) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *ReplicationRule) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetFilter

`func (o *ReplicationRule) GetFilter() ReplicationRuleFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *ReplicationRule) GetFilterOk() (*ReplicationRuleFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *ReplicationRule) SetFilter(v ReplicationRuleFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *ReplicationRule) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetStatus

`func (o *ReplicationRule) GetStatus() ReplicationRuleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReplicationRule) GetStatusOk() (*ReplicationRuleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReplicationRule) SetStatus(v ReplicationRuleStatus)`

SetStatus sets Status field to given value.


### GetSourceSelectionCriteria

`func (o *ReplicationRule) GetSourceSelectionCriteria() ReplicationRuleSourceSelectionCriteria`

GetSourceSelectionCriteria returns the SourceSelectionCriteria field if non-nil, zero value otherwise.

### GetSourceSelectionCriteriaOk

`func (o *ReplicationRule) GetSourceSelectionCriteriaOk() (*ReplicationRuleSourceSelectionCriteria, bool)`

GetSourceSelectionCriteriaOk returns a tuple with the SourceSelectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceSelectionCriteria

`func (o *ReplicationRule) SetSourceSelectionCriteria(v ReplicationRuleSourceSelectionCriteria)`

SetSourceSelectionCriteria sets SourceSelectionCriteria field to given value.

### HasSourceSelectionCriteria

`func (o *ReplicationRule) HasSourceSelectionCriteria() bool`

HasSourceSelectionCriteria returns a boolean if a field has been set.

### GetExistingObjectReplication

`func (o *ReplicationRule) GetExistingObjectReplication() ReplicationRuleExistingObjectReplication`

GetExistingObjectReplication returns the ExistingObjectReplication field if non-nil, zero value otherwise.

### GetExistingObjectReplicationOk

`func (o *ReplicationRule) GetExistingObjectReplicationOk() (*ReplicationRuleExistingObjectReplication, bool)`

GetExistingObjectReplicationOk returns a tuple with the ExistingObjectReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingObjectReplication

`func (o *ReplicationRule) SetExistingObjectReplication(v ReplicationRuleExistingObjectReplication)`

SetExistingObjectReplication sets ExistingObjectReplication field to given value.

### HasExistingObjectReplication

`func (o *ReplicationRule) HasExistingObjectReplication() bool`

HasExistingObjectReplication returns a boolean if a field has been set.

### GetDestination

`func (o *ReplicationRule) GetDestination() ReplicationRuleDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ReplicationRule) GetDestinationOk() (*ReplicationRuleDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ReplicationRule) SetDestination(v ReplicationRuleDestination)`

SetDestination sets Destination field to given value.


### GetDeleteMarkerReplication

`func (o *ReplicationRule) GetDeleteMarkerReplication() DeleteMarkerReplication`

GetDeleteMarkerReplication returns the DeleteMarkerReplication field if non-nil, zero value otherwise.

### GetDeleteMarkerReplicationOk

`func (o *ReplicationRule) GetDeleteMarkerReplicationOk() (*DeleteMarkerReplication, bool)`

GetDeleteMarkerReplicationOk returns a tuple with the DeleteMarkerReplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteMarkerReplication

`func (o *ReplicationRule) SetDeleteMarkerReplication(v DeleteMarkerReplication)`

SetDeleteMarkerReplication sets DeleteMarkerReplication field to given value.

### HasDeleteMarkerReplication

`func (o *ReplicationRule) HasDeleteMarkerReplication() bool`

HasDeleteMarkerReplication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


