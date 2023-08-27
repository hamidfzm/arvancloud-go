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

// checks if the DdosRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DdosRule{}

// DdosRule struct for DdosRule
type DdosRule struct {
	Id *string `json:"id,omitempty"`
	// - `?` matches any single character. - `*` matches any (possibly empty) sequence of characters. - `**` matches the current directory and arbitrary subdirectories. This sequence must form a single path component, so both `**a` and `b**` are invalid and will result in an error. A sequence of more than two consecutive `*` characters is also invalid. - `[...]` matches any character inside the brackets. Character sequences can also specify ranges of characters, as ordered by Unicode, so e.g. `[0-9]` specifies any character between 0 and 9 inclusive. An unclosed bracket is invalid. - `[!...]` is the negation of `[...]`, i.e. it matches any characters not in the brackets. - The metacharacters `?`, `*`, `[`, `] `can be matched by using brackets (e.g. `[?]`). When a `]` occurs immediately following `[` or `[!` then it is interpreted as being part of, rather then ending, the character set, so `]` and NOT `]` can be matched by `[]]` and `[!]]` respectively. The - character can be specified inside a character sequence pattern by placing it at the start or the end, e.g. `[abc-]`. 
	UrlPattern *string `json:"url_pattern,omitempty"`
	Sources []string `json:"sources,omitempty"`
	Description *string `json:"description,omitempty"`
	Action *string `json:"action,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
}

// NewDdosRule instantiates a new DdosRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDdosRule() *DdosRule {
	this := DdosRule{}
	return &this
}

// NewDdosRuleWithDefaults instantiates a new DdosRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDdosRuleWithDefaults() *DdosRule {
	this := DdosRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DdosRule) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRule) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DdosRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DdosRule) SetId(v string) {
	o.Id = &v
}

// GetUrlPattern returns the UrlPattern field value if set, zero value otherwise.
func (o *DdosRule) GetUrlPattern() string {
	if o == nil || IsNil(o.UrlPattern) {
		var ret string
		return ret
	}
	return *o.UrlPattern
}

// GetUrlPatternOk returns a tuple with the UrlPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRule) GetUrlPatternOk() (*string, bool) {
	if o == nil || IsNil(o.UrlPattern) {
		return nil, false
	}
	return o.UrlPattern, true
}

// HasUrlPattern returns a boolean if a field has been set.
func (o *DdosRule) HasUrlPattern() bool {
	if o != nil && !IsNil(o.UrlPattern) {
		return true
	}

	return false
}

// SetUrlPattern gets a reference to the given string and assigns it to the UrlPattern field.
func (o *DdosRule) SetUrlPattern(v string) {
	o.UrlPattern = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *DdosRule) GetSources() []string {
	if o == nil || IsNil(o.Sources) {
		var ret []string
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRule) GetSourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *DdosRule) HasSources() bool {
	if o != nil && !IsNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []string and assigns it to the Sources field.
func (o *DdosRule) SetSources(v []string) {
	o.Sources = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DdosRule) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRule) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DdosRule) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DdosRule) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DdosRule) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRule) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DdosRule) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DdosRule) SetAction(v string) {
	o.Action = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *DdosRule) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DdosRule) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *DdosRule) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *DdosRule) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

func (o DdosRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DdosRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.UrlPattern) {
		toSerialize["url_pattern"] = o.UrlPattern
	}
	if !IsNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	return toSerialize, nil
}

type NullableDdosRule struct {
	value *DdosRule
	isSet bool
}

func (v NullableDdosRule) Get() *DdosRule {
	return v.value
}

func (v *NullableDdosRule) Set(val *DdosRule) {
	v.value = val
	v.isSet = true
}

func (v NullableDdosRule) IsSet() bool {
	return v.isSet
}

func (v *NullableDdosRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDdosRule(val *DdosRule) *NullableDdosRule {
	return &NullableDdosRule{value: val, isSet: true}
}

func (v NullableDdosRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDdosRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


