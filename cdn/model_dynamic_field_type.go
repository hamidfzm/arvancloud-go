/*
ArvanCloud CDN Services

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.107.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cdn

import (
	"encoding/json"
	"fmt"
)

// DynamicFieldType - struct for DynamicFieldType
type DynamicFieldType struct {
	Float32 *float32
	String *string
}

// float32AsDynamicFieldType is a convenience function that returns float32 wrapped in DynamicFieldType
func Float32AsDynamicFieldType(v *float32) DynamicFieldType {
	return DynamicFieldType{
		Float32: v,
	}
}

// stringAsDynamicFieldType is a convenience function that returns string wrapped in DynamicFieldType
func StringAsDynamicFieldType(v *string) DynamicFieldType {
	return DynamicFieldType{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DynamicFieldType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			match++
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DynamicFieldType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DynamicFieldType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DynamicFieldType) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DynamicFieldType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDynamicFieldType struct {
	value *DynamicFieldType
	isSet bool
}

func (v NullableDynamicFieldType) Get() *DynamicFieldType {
	return v.value
}

func (v *NullableDynamicFieldType) Set(val *DynamicFieldType) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicFieldType) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicFieldType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicFieldType(val *DynamicFieldType) *NullableDynamicFieldType {
	return &NullableDynamicFieldType{value: val, isSet: true}
}

func (v NullableDynamicFieldType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicFieldType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


