/*
No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package iaas

import (
	"encoding/json"
)

// checks if the ServerImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServerImage{}

// ServerImage struct for ServerImage
type ServerImage struct {
	Created *string `json:"created,omitempty"`
	Documents []ImgDoc `json:"documents,omitempty"`
	Id *string `json:"id,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	MinDisk *int32 `json:"min_disk,omitempty"`
	MinRam *int32 `json:"min_ram,omitempty"`
	Name *string `json:"name,omitempty"`
	Os *string `json:"os,omitempty"`
	OsVersion *string `json:"os_version,omitempty"`
	Progress *int32 `json:"progress,omitempty"`
	Size *int64 `json:"size,omitempty"`
	Status *string `json:"status,omitempty"`
	Username *string `json:"username,omitempty"`
}

// NewServerImage instantiates a new ServerImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServerImage() *ServerImage {
	this := ServerImage{}
	return &this
}

// NewServerImageWithDefaults instantiates a new ServerImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServerImageWithDefaults() *ServerImage {
	this := ServerImage{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *ServerImage) GetCreated() string {
	if o == nil || IsNil(o.Created) {
		var ret string
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetCreatedOk() (*string, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *ServerImage) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given string and assigns it to the Created field.
func (o *ServerImage) SetCreated(v string) {
	o.Created = &v
}

// GetDocuments returns the Documents field value if set, zero value otherwise.
func (o *ServerImage) GetDocuments() []ImgDoc {
	if o == nil || IsNil(o.Documents) {
		var ret []ImgDoc
		return ret
	}
	return o.Documents
}

// GetDocumentsOk returns a tuple with the Documents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetDocumentsOk() ([]ImgDoc, bool) {
	if o == nil || IsNil(o.Documents) {
		return nil, false
	}
	return o.Documents, true
}

// HasDocuments returns a boolean if a field has been set.
func (o *ServerImage) HasDocuments() bool {
	if o != nil && !IsNil(o.Documents) {
		return true
	}

	return false
}

// SetDocuments gets a reference to the given []ImgDoc and assigns it to the Documents field.
func (o *ServerImage) SetDocuments(v []ImgDoc) {
	o.Documents = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServerImage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServerImage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServerImage) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServerImage) GetMetadata() map[string]string {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServerImage) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *ServerImage) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetMinDisk returns the MinDisk field value if set, zero value otherwise.
func (o *ServerImage) GetMinDisk() int32 {
	if o == nil || IsNil(o.MinDisk) {
		var ret int32
		return ret
	}
	return *o.MinDisk
}

// GetMinDiskOk returns a tuple with the MinDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetMinDiskOk() (*int32, bool) {
	if o == nil || IsNil(o.MinDisk) {
		return nil, false
	}
	return o.MinDisk, true
}

// HasMinDisk returns a boolean if a field has been set.
func (o *ServerImage) HasMinDisk() bool {
	if o != nil && !IsNil(o.MinDisk) {
		return true
	}

	return false
}

// SetMinDisk gets a reference to the given int32 and assigns it to the MinDisk field.
func (o *ServerImage) SetMinDisk(v int32) {
	o.MinDisk = &v
}

// GetMinRam returns the MinRam field value if set, zero value otherwise.
func (o *ServerImage) GetMinRam() int32 {
	if o == nil || IsNil(o.MinRam) {
		var ret int32
		return ret
	}
	return *o.MinRam
}

// GetMinRamOk returns a tuple with the MinRam field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetMinRamOk() (*int32, bool) {
	if o == nil || IsNil(o.MinRam) {
		return nil, false
	}
	return o.MinRam, true
}

// HasMinRam returns a boolean if a field has been set.
func (o *ServerImage) HasMinRam() bool {
	if o != nil && !IsNil(o.MinRam) {
		return true
	}

	return false
}

// SetMinRam gets a reference to the given int32 and assigns it to the MinRam field.
func (o *ServerImage) SetMinRam(v int32) {
	o.MinRam = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServerImage) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServerImage) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServerImage) SetName(v string) {
	o.Name = &v
}

// GetOs returns the Os field value if set, zero value otherwise.
func (o *ServerImage) GetOs() string {
	if o == nil || IsNil(o.Os) {
		var ret string
		return ret
	}
	return *o.Os
}

// GetOsOk returns a tuple with the Os field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetOsOk() (*string, bool) {
	if o == nil || IsNil(o.Os) {
		return nil, false
	}
	return o.Os, true
}

// HasOs returns a boolean if a field has been set.
func (o *ServerImage) HasOs() bool {
	if o != nil && !IsNil(o.Os) {
		return true
	}

	return false
}

// SetOs gets a reference to the given string and assigns it to the Os field.
func (o *ServerImage) SetOs(v string) {
	o.Os = &v
}

// GetOsVersion returns the OsVersion field value if set, zero value otherwise.
func (o *ServerImage) GetOsVersion() string {
	if o == nil || IsNil(o.OsVersion) {
		var ret string
		return ret
	}
	return *o.OsVersion
}

// GetOsVersionOk returns a tuple with the OsVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetOsVersionOk() (*string, bool) {
	if o == nil || IsNil(o.OsVersion) {
		return nil, false
	}
	return o.OsVersion, true
}

// HasOsVersion returns a boolean if a field has been set.
func (o *ServerImage) HasOsVersion() bool {
	if o != nil && !IsNil(o.OsVersion) {
		return true
	}

	return false
}

// SetOsVersion gets a reference to the given string and assigns it to the OsVersion field.
func (o *ServerImage) SetOsVersion(v string) {
	o.OsVersion = &v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *ServerImage) GetProgress() int32 {
	if o == nil || IsNil(o.Progress) {
		var ret int32
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetProgressOk() (*int32, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *ServerImage) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given int32 and assigns it to the Progress field.
func (o *ServerImage) SetProgress(v int32) {
	o.Progress = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ServerImage) GetSize() int64 {
	if o == nil || IsNil(o.Size) {
		var ret int64
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ServerImage) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int64 and assigns it to the Size field.
func (o *ServerImage) SetSize(v int64) {
	o.Size = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServerImage) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServerImage) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServerImage) SetStatus(v string) {
	o.Status = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ServerImage) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServerImage) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ServerImage) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ServerImage) SetUsername(v string) {
	o.Username = &v
}

func (o ServerImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServerImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Documents) {
		toSerialize["documents"] = o.Documents
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.MinDisk) {
		toSerialize["min_disk"] = o.MinDisk
	}
	if !IsNil(o.MinRam) {
		toSerialize["min_ram"] = o.MinRam
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Os) {
		toSerialize["os"] = o.Os
	}
	if !IsNil(o.OsVersion) {
		toSerialize["os_version"] = o.OsVersion
	}
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableServerImage struct {
	value *ServerImage
	isSet bool
}

func (v NullableServerImage) Get() *ServerImage {
	return v.value
}

func (v *NullableServerImage) Set(val *ServerImage) {
	v.value = val
	v.isSet = true
}

func (v NullableServerImage) IsSet() bool {
	return v.isSet
}

func (v *NullableServerImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerImage(val *ServerImage) *NullableServerImage {
	return &NullableServerImage{value: val, isSet: true}
}

func (v NullableServerImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


