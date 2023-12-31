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

// checks if the ListMultipartUploadsOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListMultipartUploadsOutput{}

// ListMultipartUploadsOutput struct for ListMultipartUploadsOutput
type ListMultipartUploadsOutput struct {
	Bucket *string `json:"Bucket,omitempty"`
	KeyMarker *string `json:"KeyMarker,omitempty"`
	UploadIdMarker *string `json:"UploadIdMarker,omitempty"`
	NextKeyMarker *string `json:"NextKeyMarker,omitempty"`
	Prefix *string `json:"Prefix,omitempty"`
	Delimiter *string `json:"Delimiter,omitempty"`
	NextUploadIdMarker *string `json:"NextUploadIdMarker,omitempty"`
	MaxUploads *int32 `json:"MaxUploads,omitempty"`
	IsTruncated *bool `json:"IsTruncated,omitempty"`
	Uploads *Array `json:"Uploads,omitempty"`
	CommonPrefixes *Array `json:"CommonPrefixes,omitempty"`
	EncodingType *EncodingType `json:"EncodingType,omitempty"`
}

// NewListMultipartUploadsOutput instantiates a new ListMultipartUploadsOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListMultipartUploadsOutput() *ListMultipartUploadsOutput {
	this := ListMultipartUploadsOutput{}
	return &this
}

// NewListMultipartUploadsOutputWithDefaults instantiates a new ListMultipartUploadsOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListMultipartUploadsOutputWithDefaults() *ListMultipartUploadsOutput {
	this := ListMultipartUploadsOutput{}
	return &this
}

// GetBucket returns the Bucket field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetBucket() string {
	if o == nil || IsNil(o.Bucket) {
		var ret string
		return ret
	}
	return *o.Bucket
}

// GetBucketOk returns a tuple with the Bucket field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetBucketOk() (*string, bool) {
	if o == nil || IsNil(o.Bucket) {
		return nil, false
	}
	return o.Bucket, true
}

// HasBucket returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasBucket() bool {
	if o != nil && !IsNil(o.Bucket) {
		return true
	}

	return false
}

// SetBucket gets a reference to the given string and assigns it to the Bucket field.
func (o *ListMultipartUploadsOutput) SetBucket(v string) {
	o.Bucket = &v
}

// GetKeyMarker returns the KeyMarker field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetKeyMarker() string {
	if o == nil || IsNil(o.KeyMarker) {
		var ret string
		return ret
	}
	return *o.KeyMarker
}

// GetKeyMarkerOk returns a tuple with the KeyMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetKeyMarkerOk() (*string, bool) {
	if o == nil || IsNil(o.KeyMarker) {
		return nil, false
	}
	return o.KeyMarker, true
}

// HasKeyMarker returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasKeyMarker() bool {
	if o != nil && !IsNil(o.KeyMarker) {
		return true
	}

	return false
}

// SetKeyMarker gets a reference to the given string and assigns it to the KeyMarker field.
func (o *ListMultipartUploadsOutput) SetKeyMarker(v string) {
	o.KeyMarker = &v
}

// GetUploadIdMarker returns the UploadIdMarker field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetUploadIdMarker() string {
	if o == nil || IsNil(o.UploadIdMarker) {
		var ret string
		return ret
	}
	return *o.UploadIdMarker
}

// GetUploadIdMarkerOk returns a tuple with the UploadIdMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetUploadIdMarkerOk() (*string, bool) {
	if o == nil || IsNil(o.UploadIdMarker) {
		return nil, false
	}
	return o.UploadIdMarker, true
}

// HasUploadIdMarker returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasUploadIdMarker() bool {
	if o != nil && !IsNil(o.UploadIdMarker) {
		return true
	}

	return false
}

// SetUploadIdMarker gets a reference to the given string and assigns it to the UploadIdMarker field.
func (o *ListMultipartUploadsOutput) SetUploadIdMarker(v string) {
	o.UploadIdMarker = &v
}

// GetNextKeyMarker returns the NextKeyMarker field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetNextKeyMarker() string {
	if o == nil || IsNil(o.NextKeyMarker) {
		var ret string
		return ret
	}
	return *o.NextKeyMarker
}

// GetNextKeyMarkerOk returns a tuple with the NextKeyMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetNextKeyMarkerOk() (*string, bool) {
	if o == nil || IsNil(o.NextKeyMarker) {
		return nil, false
	}
	return o.NextKeyMarker, true
}

// HasNextKeyMarker returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasNextKeyMarker() bool {
	if o != nil && !IsNil(o.NextKeyMarker) {
		return true
	}

	return false
}

// SetNextKeyMarker gets a reference to the given string and assigns it to the NextKeyMarker field.
func (o *ListMultipartUploadsOutput) SetNextKeyMarker(v string) {
	o.NextKeyMarker = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *ListMultipartUploadsOutput) SetPrefix(v string) {
	o.Prefix = &v
}

// GetDelimiter returns the Delimiter field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetDelimiter() string {
	if o == nil || IsNil(o.Delimiter) {
		var ret string
		return ret
	}
	return *o.Delimiter
}

// GetDelimiterOk returns a tuple with the Delimiter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetDelimiterOk() (*string, bool) {
	if o == nil || IsNil(o.Delimiter) {
		return nil, false
	}
	return o.Delimiter, true
}

// HasDelimiter returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasDelimiter() bool {
	if o != nil && !IsNil(o.Delimiter) {
		return true
	}

	return false
}

// SetDelimiter gets a reference to the given string and assigns it to the Delimiter field.
func (o *ListMultipartUploadsOutput) SetDelimiter(v string) {
	o.Delimiter = &v
}

// GetNextUploadIdMarker returns the NextUploadIdMarker field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetNextUploadIdMarker() string {
	if o == nil || IsNil(o.NextUploadIdMarker) {
		var ret string
		return ret
	}
	return *o.NextUploadIdMarker
}

// GetNextUploadIdMarkerOk returns a tuple with the NextUploadIdMarker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetNextUploadIdMarkerOk() (*string, bool) {
	if o == nil || IsNil(o.NextUploadIdMarker) {
		return nil, false
	}
	return o.NextUploadIdMarker, true
}

// HasNextUploadIdMarker returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasNextUploadIdMarker() bool {
	if o != nil && !IsNil(o.NextUploadIdMarker) {
		return true
	}

	return false
}

// SetNextUploadIdMarker gets a reference to the given string and assigns it to the NextUploadIdMarker field.
func (o *ListMultipartUploadsOutput) SetNextUploadIdMarker(v string) {
	o.NextUploadIdMarker = &v
}

// GetMaxUploads returns the MaxUploads field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetMaxUploads() int32 {
	if o == nil || IsNil(o.MaxUploads) {
		var ret int32
		return ret
	}
	return *o.MaxUploads
}

// GetMaxUploadsOk returns a tuple with the MaxUploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetMaxUploadsOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxUploads) {
		return nil, false
	}
	return o.MaxUploads, true
}

// HasMaxUploads returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasMaxUploads() bool {
	if o != nil && !IsNil(o.MaxUploads) {
		return true
	}

	return false
}

// SetMaxUploads gets a reference to the given int32 and assigns it to the MaxUploads field.
func (o *ListMultipartUploadsOutput) SetMaxUploads(v int32) {
	o.MaxUploads = &v
}

// GetIsTruncated returns the IsTruncated field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetIsTruncated() bool {
	if o == nil || IsNil(o.IsTruncated) {
		var ret bool
		return ret
	}
	return *o.IsTruncated
}

// GetIsTruncatedOk returns a tuple with the IsTruncated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetIsTruncatedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTruncated) {
		return nil, false
	}
	return o.IsTruncated, true
}

// HasIsTruncated returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasIsTruncated() bool {
	if o != nil && !IsNil(o.IsTruncated) {
		return true
	}

	return false
}

// SetIsTruncated gets a reference to the given bool and assigns it to the IsTruncated field.
func (o *ListMultipartUploadsOutput) SetIsTruncated(v bool) {
	o.IsTruncated = &v
}

// GetUploads returns the Uploads field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetUploads() Array {
	if o == nil || IsNil(o.Uploads) {
		var ret Array
		return ret
	}
	return *o.Uploads
}

// GetUploadsOk returns a tuple with the Uploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetUploadsOk() (*Array, bool) {
	if o == nil || IsNil(o.Uploads) {
		return nil, false
	}
	return o.Uploads, true
}

// HasUploads returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasUploads() bool {
	if o != nil && !IsNil(o.Uploads) {
		return true
	}

	return false
}

// SetUploads gets a reference to the given Array and assigns it to the Uploads field.
func (o *ListMultipartUploadsOutput) SetUploads(v Array) {
	o.Uploads = &v
}

// GetCommonPrefixes returns the CommonPrefixes field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetCommonPrefixes() Array {
	if o == nil || IsNil(o.CommonPrefixes) {
		var ret Array
		return ret
	}
	return *o.CommonPrefixes
}

// GetCommonPrefixesOk returns a tuple with the CommonPrefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetCommonPrefixesOk() (*Array, bool) {
	if o == nil || IsNil(o.CommonPrefixes) {
		return nil, false
	}
	return o.CommonPrefixes, true
}

// HasCommonPrefixes returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasCommonPrefixes() bool {
	if o != nil && !IsNil(o.CommonPrefixes) {
		return true
	}

	return false
}

// SetCommonPrefixes gets a reference to the given Array and assigns it to the CommonPrefixes field.
func (o *ListMultipartUploadsOutput) SetCommonPrefixes(v Array) {
	o.CommonPrefixes = &v
}

// GetEncodingType returns the EncodingType field value if set, zero value otherwise.
func (o *ListMultipartUploadsOutput) GetEncodingType() EncodingType {
	if o == nil || IsNil(o.EncodingType) {
		var ret EncodingType
		return ret
	}
	return *o.EncodingType
}

// GetEncodingTypeOk returns a tuple with the EncodingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListMultipartUploadsOutput) GetEncodingTypeOk() (*EncodingType, bool) {
	if o == nil || IsNil(o.EncodingType) {
		return nil, false
	}
	return o.EncodingType, true
}

// HasEncodingType returns a boolean if a field has been set.
func (o *ListMultipartUploadsOutput) HasEncodingType() bool {
	if o != nil && !IsNil(o.EncodingType) {
		return true
	}

	return false
}

// SetEncodingType gets a reference to the given EncodingType and assigns it to the EncodingType field.
func (o *ListMultipartUploadsOutput) SetEncodingType(v EncodingType) {
	o.EncodingType = &v
}

func (o ListMultipartUploadsOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListMultipartUploadsOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bucket) {
		toSerialize["Bucket"] = o.Bucket
	}
	if !IsNil(o.KeyMarker) {
		toSerialize["KeyMarker"] = o.KeyMarker
	}
	if !IsNil(o.UploadIdMarker) {
		toSerialize["UploadIdMarker"] = o.UploadIdMarker
	}
	if !IsNil(o.NextKeyMarker) {
		toSerialize["NextKeyMarker"] = o.NextKeyMarker
	}
	if !IsNil(o.Prefix) {
		toSerialize["Prefix"] = o.Prefix
	}
	if !IsNil(o.Delimiter) {
		toSerialize["Delimiter"] = o.Delimiter
	}
	if !IsNil(o.NextUploadIdMarker) {
		toSerialize["NextUploadIdMarker"] = o.NextUploadIdMarker
	}
	if !IsNil(o.MaxUploads) {
		toSerialize["MaxUploads"] = o.MaxUploads
	}
	if !IsNil(o.IsTruncated) {
		toSerialize["IsTruncated"] = o.IsTruncated
	}
	if !IsNil(o.Uploads) {
		toSerialize["Uploads"] = o.Uploads
	}
	if !IsNil(o.CommonPrefixes) {
		toSerialize["CommonPrefixes"] = o.CommonPrefixes
	}
	if !IsNil(o.EncodingType) {
		toSerialize["EncodingType"] = o.EncodingType
	}
	return toSerialize, nil
}

type NullableListMultipartUploadsOutput struct {
	value *ListMultipartUploadsOutput
	isSet bool
}

func (v NullableListMultipartUploadsOutput) Get() *ListMultipartUploadsOutput {
	return v.value
}

func (v *NullableListMultipartUploadsOutput) Set(val *ListMultipartUploadsOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableListMultipartUploadsOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableListMultipartUploadsOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListMultipartUploadsOutput(val *ListMultipartUploadsOutput) *NullableListMultipartUploadsOutput {
	return &NullableListMultipartUploadsOutput{value: val, isSet: true}
}

func (v NullableListMultipartUploadsOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListMultipartUploadsOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


