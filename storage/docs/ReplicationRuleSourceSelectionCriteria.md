# ReplicationRuleSourceSelectionCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SseKmsEncryptedObjects** | Pointer to [**SourceSelectionCriteriaSseKmsEncryptedObjects**](SourceSelectionCriteriaSseKmsEncryptedObjects.md) |  | [optional] 
**ReplicaModifications** | Pointer to [**SourceSelectionCriteriaReplicaModifications**](SourceSelectionCriteriaReplicaModifications.md) |  | [optional] 

## Methods

### NewReplicationRuleSourceSelectionCriteria

`func NewReplicationRuleSourceSelectionCriteria() *ReplicationRuleSourceSelectionCriteria`

NewReplicationRuleSourceSelectionCriteria instantiates a new ReplicationRuleSourceSelectionCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplicationRuleSourceSelectionCriteriaWithDefaults

`func NewReplicationRuleSourceSelectionCriteriaWithDefaults() *ReplicationRuleSourceSelectionCriteria`

NewReplicationRuleSourceSelectionCriteriaWithDefaults instantiates a new ReplicationRuleSourceSelectionCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSseKmsEncryptedObjects

`func (o *ReplicationRuleSourceSelectionCriteria) GetSseKmsEncryptedObjects() SourceSelectionCriteriaSseKmsEncryptedObjects`

GetSseKmsEncryptedObjects returns the SseKmsEncryptedObjects field if non-nil, zero value otherwise.

### GetSseKmsEncryptedObjectsOk

`func (o *ReplicationRuleSourceSelectionCriteria) GetSseKmsEncryptedObjectsOk() (*SourceSelectionCriteriaSseKmsEncryptedObjects, bool)`

GetSseKmsEncryptedObjectsOk returns a tuple with the SseKmsEncryptedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSseKmsEncryptedObjects

`func (o *ReplicationRuleSourceSelectionCriteria) SetSseKmsEncryptedObjects(v SourceSelectionCriteriaSseKmsEncryptedObjects)`

SetSseKmsEncryptedObjects sets SseKmsEncryptedObjects field to given value.

### HasSseKmsEncryptedObjects

`func (o *ReplicationRuleSourceSelectionCriteria) HasSseKmsEncryptedObjects() bool`

HasSseKmsEncryptedObjects returns a boolean if a field has been set.

### GetReplicaModifications

`func (o *ReplicationRuleSourceSelectionCriteria) GetReplicaModifications() SourceSelectionCriteriaReplicaModifications`

GetReplicaModifications returns the ReplicaModifications field if non-nil, zero value otherwise.

### GetReplicaModificationsOk

`func (o *ReplicationRuleSourceSelectionCriteria) GetReplicaModificationsOk() (*SourceSelectionCriteriaReplicaModifications, bool)`

GetReplicaModificationsOk returns a tuple with the ReplicaModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaModifications

`func (o *ReplicationRuleSourceSelectionCriteria) SetReplicaModifications(v SourceSelectionCriteriaReplicaModifications)`

SetReplicaModifications sets ReplicaModifications field to given value.

### HasReplicaModifications

`func (o *ReplicationRuleSourceSelectionCriteria) HasReplicaModifications() bool`

HasReplicaModifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


