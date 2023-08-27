/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
)

// checks if the PaginatedResponseLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginatedResponseLinks{}

// PaginatedResponseLinks struct for PaginatedResponseLinks
type PaginatedResponseLinks struct {
	First string `json:"first"`
	Last NullableString `json:"last,omitempty"`
	Prev NullableString `json:"prev"`
	Next NullableString `json:"next"`
}

// NewPaginatedResponseLinks instantiates a new PaginatedResponseLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginatedResponseLinks(first string, prev NullableString, next NullableString) *PaginatedResponseLinks {
	this := PaginatedResponseLinks{}
	this.First = first
	this.Prev = prev
	this.Next = next
	return &this
}

// NewPaginatedResponseLinksWithDefaults instantiates a new PaginatedResponseLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginatedResponseLinksWithDefaults() *PaginatedResponseLinks {
	this := PaginatedResponseLinks{}
	return &this
}

// GetFirst returns the First field value
func (o *PaginatedResponseLinks) GetFirst() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.First
}

// GetFirstOk returns a tuple with the First field value
// and a boolean to check if the value has been set.
func (o *PaginatedResponseLinks) GetFirstOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.First, true
}

// SetFirst sets field value
func (o *PaginatedResponseLinks) SetFirst(v string) {
	o.First = v
}

// GetLast returns the Last field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaginatedResponseLinks) GetLast() string {
	if o == nil || IsNil(o.Last.Get()) {
		var ret string
		return ret
	}
	return *o.Last.Get()
}

// GetLastOk returns a tuple with the Last field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedResponseLinks) GetLastOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Last.Get(), o.Last.IsSet()
}

// HasLast returns a boolean if a field has been set.
func (o *PaginatedResponseLinks) HasLast() bool {
	if o != nil && o.Last.IsSet() {
		return true
	}

	return false
}

// SetLast gets a reference to the given NullableString and assigns it to the Last field.
func (o *PaginatedResponseLinks) SetLast(v string) {
	o.Last.Set(&v)
}
// SetLastNil sets the value for Last to be an explicit nil
func (o *PaginatedResponseLinks) SetLastNil() {
	o.Last.Set(nil)
}

// UnsetLast ensures that no value is present for Last, not even an explicit nil
func (o *PaginatedResponseLinks) UnsetLast() {
	o.Last.Unset()
}

// GetPrev returns the Prev field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedResponseLinks) GetPrev() string {
	if o == nil || o.Prev.Get() == nil {
		var ret string
		return ret
	}

	return *o.Prev.Get()
}

// GetPrevOk returns a tuple with the Prev field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedResponseLinks) GetPrevOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Prev.Get(), o.Prev.IsSet()
}

// SetPrev sets field value
func (o *PaginatedResponseLinks) SetPrev(v string) {
	o.Prev.Set(&v)
}

// GetNext returns the Next field value
// If the value is explicit nil, the zero value for string will be returned
func (o *PaginatedResponseLinks) GetNext() string {
	if o == nil || o.Next.Get() == nil {
		var ret string
		return ret
	}

	return *o.Next.Get()
}

// GetNextOk returns a tuple with the Next field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaginatedResponseLinks) GetNextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Next.Get(), o.Next.IsSet()
}

// SetNext sets field value
func (o *PaginatedResponseLinks) SetNext(v string) {
	o.Next.Set(&v)
}

func (o PaginatedResponseLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginatedResponseLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["first"] = o.First
	if o.Last.IsSet() {
		toSerialize["last"] = o.Last.Get()
	}
	toSerialize["prev"] = o.Prev.Get()
	toSerialize["next"] = o.Next.Get()
	return toSerialize, nil
}

type NullablePaginatedResponseLinks struct {
	value *PaginatedResponseLinks
	isSet bool
}

func (v NullablePaginatedResponseLinks) Get() *PaginatedResponseLinks {
	return v.value
}

func (v *NullablePaginatedResponseLinks) Set(val *PaginatedResponseLinks) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginatedResponseLinks) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginatedResponseLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginatedResponseLinks(val *PaginatedResponseLinks) *NullablePaginatedResponseLinks {
	return &NullablePaginatedResponseLinks{value: val, isSet: true}
}

func (v NullablePaginatedResponseLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginatedResponseLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

