/*
Arvancloud PaaS REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paas

import (
	"encoding/json"
)

// checks if the V1GitSourceRevision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1GitSourceRevision{}

// V1GitSourceRevision GitSourceRevision is the commit information from a git source for a build
type V1GitSourceRevision struct {
	Author *V1SourceControlUser `json:"author,omitempty"`
	// commit is the commit hash identifying a specific commit
	Commit *string `json:"commit,omitempty"`
	Committer *V1SourceControlUser `json:"committer,omitempty"`
	// message is the description of a specific commit
	Message *string `json:"message,omitempty"`
}

// NewV1GitSourceRevision instantiates a new V1GitSourceRevision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1GitSourceRevision() *V1GitSourceRevision {
	this := V1GitSourceRevision{}
	return &this
}

// NewV1GitSourceRevisionWithDefaults instantiates a new V1GitSourceRevision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1GitSourceRevisionWithDefaults() *V1GitSourceRevision {
	this := V1GitSourceRevision{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *V1GitSourceRevision) GetAuthor() V1SourceControlUser {
	if o == nil || IsNil(o.Author) {
		var ret V1SourceControlUser
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GitSourceRevision) GetAuthorOk() (*V1SourceControlUser, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *V1GitSourceRevision) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given V1SourceControlUser and assigns it to the Author field.
func (o *V1GitSourceRevision) SetAuthor(v V1SourceControlUser) {
	o.Author = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *V1GitSourceRevision) GetCommit() string {
	if o == nil || IsNil(o.Commit) {
		var ret string
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GitSourceRevision) GetCommitOk() (*string, bool) {
	if o == nil || IsNil(o.Commit) {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *V1GitSourceRevision) HasCommit() bool {
	if o != nil && !IsNil(o.Commit) {
		return true
	}

	return false
}

// SetCommit gets a reference to the given string and assigns it to the Commit field.
func (o *V1GitSourceRevision) SetCommit(v string) {
	o.Commit = &v
}

// GetCommitter returns the Committer field value if set, zero value otherwise.
func (o *V1GitSourceRevision) GetCommitter() V1SourceControlUser {
	if o == nil || IsNil(o.Committer) {
		var ret V1SourceControlUser
		return ret
	}
	return *o.Committer
}

// GetCommitterOk returns a tuple with the Committer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GitSourceRevision) GetCommitterOk() (*V1SourceControlUser, bool) {
	if o == nil || IsNil(o.Committer) {
		return nil, false
	}
	return o.Committer, true
}

// HasCommitter returns a boolean if a field has been set.
func (o *V1GitSourceRevision) HasCommitter() bool {
	if o != nil && !IsNil(o.Committer) {
		return true
	}

	return false
}

// SetCommitter gets a reference to the given V1SourceControlUser and assigns it to the Committer field.
func (o *V1GitSourceRevision) SetCommitter(v V1SourceControlUser) {
	o.Committer = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1GitSourceRevision) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1GitSourceRevision) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1GitSourceRevision) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1GitSourceRevision) SetMessage(v string) {
	o.Message = &v
}

func (o V1GitSourceRevision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1GitSourceRevision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Commit) {
		toSerialize["commit"] = o.Commit
	}
	if !IsNil(o.Committer) {
		toSerialize["committer"] = o.Committer
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableV1GitSourceRevision struct {
	value *V1GitSourceRevision
	isSet bool
}

func (v NullableV1GitSourceRevision) Get() *V1GitSourceRevision {
	return v.value
}

func (v *NullableV1GitSourceRevision) Set(val *V1GitSourceRevision) {
	v.value = val
	v.isSet = true
}

func (v NullableV1GitSourceRevision) IsSet() bool {
	return v.isSet
}

func (v *NullableV1GitSourceRevision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1GitSourceRevision(val *V1GitSourceRevision) *NullableV1GitSourceRevision {
	return &NullableV1GitSourceRevision{value: val, isSet: true}
}

func (v NullableV1GitSourceRevision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1GitSourceRevision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


