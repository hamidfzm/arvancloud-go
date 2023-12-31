/*
ArvanCloud S3 Services

<p/>

API version: 2006-03-01
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"fmt"
)

// Event The bucket event for which to send notifications.
type Event string

// List of Event
const (
	EVENT_OBJECT_CREATED Event = "s3:ObjectCreated:*"
	EVENT_OBJECT_CREATEDPUT Event = "s3:ObjectCreated:Put"
	EVENT_OBJECT_CREATEDPOST Event = "s3:ObjectCreated:Post"
	EVENT_OBJECT_CREATEDCOPY Event = "s3:ObjectCreated:Copy"
	EVENT_OBJECT_CREATEDCOMPLETE_MULTIPART_UPLOAD Event = "s3:ObjectCreated:CompleteMultipartUpload"
	EVENT_OBJECT_REMOVED Event = "s3:ObjectRemoved:*"
	EVENT_OBJECT_REMOVEDDELETE Event = "s3:ObjectRemoved:Delete"
)

// All allowed values of Event enum
var AllowedEventEnumValues = []Event{
	"s3:ObjectCreated:*",
	"s3:ObjectCreated:Put",
	"s3:ObjectCreated:Post",
	"s3:ObjectCreated:Copy",
	"s3:ObjectCreated:CompleteMultipartUpload",
	"s3:ObjectRemoved:*",
	"s3:ObjectRemoved:Delete",
}

func (v *Event) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Event(value)
	for _, existing := range AllowedEventEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Event", value)
}

// NewEventFromValue returns a pointer to a valid Event
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventFromValue(v string) (*Event, error) {
	ev := Event(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Event: valid values are %v", v, AllowedEventEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Event) IsValid() bool {
	for _, existing := range AllowedEventEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Event value
func (v Event) Ptr() *Event {
	return &v
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

