# SourceSelectionCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SseKmsEncryptedObjects** | Pointer to [**SourceSelectionCriteriaSseKmsEncryptedObjects**](SourceSelectionCriteriaSseKmsEncryptedObjects.md) |  | [optional] 
**ReplicaModifications** | Pointer to [**SourceSelectionCriteriaReplicaModifications**](SourceSelectionCriteriaReplicaModifications.md) |  | [optional] 

## Methods

### NewSourceSelectionCriteria

`func NewSourceSelectionCriteria() *SourceSelectionCriteria`

NewSourceSelectionCriteria instantiates a new SourceSelectionCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceSelectionCriteriaWithDefaults

`func NewSourceSelectionCriteriaWithDefaults() *SourceSelectionCriteria`

NewSourceSelectionCriteriaWithDefaults instantiates a new SourceSelectionCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSseKmsEncryptedObjects

`func (o *SourceSelectionCriteria) GetSseKmsEncryptedObjects() SourceSelectionCriteriaSseKmsEncryptedObjects`

GetSseKmsEncryptedObjects returns the SseKmsEncryptedObjects field if non-nil, zero value otherwise.

### GetSseKmsEncryptedObjectsOk

`func (o *SourceSelectionCriteria) GetSseKmsEncryptedObjectsOk() (*SourceSelectionCriteriaSseKmsEncryptedObjects, bool)`

GetSseKmsEncryptedObjectsOk returns a tuple with the SseKmsEncryptedObjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSseKmsEncryptedObjects

`func (o *SourceSelectionCriteria) SetSseKmsEncryptedObjects(v SourceSelectionCriteriaSseKmsEncryptedObjects)`

SetSseKmsEncryptedObjects sets SseKmsEncryptedObjects field to given value.

### HasSseKmsEncryptedObjects

`func (o *SourceSelectionCriteria) HasSseKmsEncryptedObjects() bool`

HasSseKmsEncryptedObjects returns a boolean if a field has been set.

### GetReplicaModifications

`func (o *SourceSelectionCriteria) GetReplicaModifications() SourceSelectionCriteriaReplicaModifications`

GetReplicaModifications returns the ReplicaModifications field if non-nil, zero value otherwise.

### GetReplicaModificationsOk

`func (o *SourceSelectionCriteria) GetReplicaModificationsOk() (*SourceSelectionCriteriaReplicaModifications, bool)`

GetReplicaModificationsOk returns a tuple with the ReplicaModifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicaModifications

`func (o *SourceSelectionCriteria) SetReplicaModifications(v SourceSelectionCriteriaReplicaModifications)`

SetReplicaModifications sets ReplicaModifications field to given value.

### HasReplicaModifications

`func (o *SourceSelectionCriteria) HasReplicaModifications() bool`

HasReplicaModifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


