/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
)

// checks if the SourceSelectionCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SourceSelectionCriteria{}

// SourceSelectionCriteria A container that describes additional filters for identifying the source objects that you want to replicate. You can choose to enable or disable the replication of these objects. Currently, ArvanCloud S3 supports only the filter that you can specify for objects created with server-side encryption using a customer managed key stored in Amazon Web Services Key Management Service (SSE-KMS).
type SourceSelectionCriteria struct {
	SseKmsEncryptedObjects *SourceSelectionCriteriaSseKmsEncryptedObjects `json:"SseKmsEncryptedObjects,omitempty"`
	ReplicaModifications *SourceSelectionCriteriaReplicaModifications `json:"ReplicaModifications,omitempty"`
}

// NewSourceSelectionCriteria instantiates a new SourceSelectionCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSourceSelectionCriteria() *SourceSelectionCriteria {
	this := SourceSelectionCriteria{}
	return &this
}

// NewSourceSelectionCriteriaWithDefaults instantiates a new SourceSelectionCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceSelectionCriteriaWithDefaults() *SourceSelectionCriteria {
	this := SourceSelectionCriteria{}
	return &this
}

// GetSseKmsEncryptedObjects returns the SseKmsEncryptedObjects field value if set, zero value otherwise.
func (o *SourceSelectionCriteria) GetSseKmsEncryptedObjects() SourceSelectionCriteriaSseKmsEncryptedObjects {
	if o == nil || IsNil(o.SseKmsEncryptedObjects) {
		var ret SourceSelectionCriteriaSseKmsEncryptedObjects
		return ret
	}
	return *o.SseKmsEncryptedObjects
}

// GetSseKmsEncryptedObjectsOk returns a tuple with the SseKmsEncryptedObjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSelectionCriteria) GetSseKmsEncryptedObjectsOk() (*SourceSelectionCriteriaSseKmsEncryptedObjects, bool) {
	if o == nil || IsNil(o.SseKmsEncryptedObjects) {
		return nil, false
	}
	return o.SseKmsEncryptedObjects, true
}

// HasSseKmsEncryptedObjects returns a boolean if a field has been set.
func (o *SourceSelectionCriteria) HasSseKmsEncryptedObjects() bool {
	if o != nil && !IsNil(o.SseKmsEncryptedObjects) {
		return true
	}

	return false
}

// SetSseKmsEncryptedObjects gets a reference to the given SourceSelectionCriteriaSseKmsEncryptedObjects and assigns it to the SseKmsEncryptedObjects field.
func (o *SourceSelectionCriteria) SetSseKmsEncryptedObjects(v SourceSelectionCriteriaSseKmsEncryptedObjects) {
	o.SseKmsEncryptedObjects = &v
}

// GetReplicaModifications returns the ReplicaModifications field value if set, zero value otherwise.
func (o *SourceSelectionCriteria) GetReplicaModifications() SourceSelectionCriteriaReplicaModifications {
	if o == nil || IsNil(o.ReplicaModifications) {
		var ret SourceSelectionCriteriaReplicaModifications
		return ret
	}
	return *o.ReplicaModifications
}

// GetReplicaModificationsOk returns a tuple with the ReplicaModifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceSelectionCriteria) GetReplicaModificationsOk() (*SourceSelectionCriteriaReplicaModifications, bool) {
	if o == nil || IsNil(o.ReplicaModifications) {
		return nil, false
	}
	return o.ReplicaModifications, true
}

// HasReplicaModifications returns a boolean if a field has been set.
func (o *SourceSelectionCriteria) HasReplicaModifications() bool {
	if o != nil && !IsNil(o.ReplicaModifications) {
		return true
	}

	return false
}

// SetReplicaModifications gets a reference to the given SourceSelectionCriteriaReplicaModifications and assigns it to the ReplicaModifications field.
func (o *SourceSelectionCriteria) SetReplicaModifications(v SourceSelectionCriteriaReplicaModifications) {
	o.ReplicaModifications = &v
}

func (o SourceSelectionCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceSelectionCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SseKmsEncryptedObjects) {
		toSerialize["SseKmsEncryptedObjects"] = o.SseKmsEncryptedObjects
	}
	if !IsNil(o.ReplicaModifications) {
		toSerialize["ReplicaModifications"] = o.ReplicaModifications
	}
	return toSerialize, nil
}

type NullableSourceSelectionCriteria struct {
	value *SourceSelectionCriteria
	isSet bool
}

func (v NullableSourceSelectionCriteria) Get() *SourceSelectionCriteria {
	return v.value
}

func (v *NullableSourceSelectionCriteria) Set(val *SourceSelectionCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSourceSelectionCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceSelectionCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceSelectionCriteria(val *SourceSelectionCriteria) *NullableSourceSelectionCriteria {
	return &NullableSourceSelectionCriteria{value: val, isSet: true}
}

func (v NullableSourceSelectionCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceSelectionCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

