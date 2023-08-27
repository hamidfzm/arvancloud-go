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

// RequestPayer Confirms that the requester knows that they will be charged for the request. Bucket owners need not specify this parameter in their requests. For information about downloading objects from requester pays buckets, see <a href=\"https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html\">Downloading Objects in Requestor Pays Buckets</a> in the <i>S3 User Guide</i>.
type RequestPayer string

// List of RequestPayer
const (
	REQUESTPAYER_REQUESTER RequestPayer = "requester"
)

// All allowed values of RequestPayer enum
var AllowedRequestPayerEnumValues = []RequestPayer{
	"requester",
}

func (v *RequestPayer) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestPayer(value)
	for _, existing := range AllowedRequestPayerEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestPayer", value)
}

// NewRequestPayerFromValue returns a pointer to a valid RequestPayer
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestPayerFromValue(v string) (*RequestPayer, error) {
	ev := RequestPayer(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestPayer: valid values are %v", v, AllowedRequestPayerEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestPayer) IsValid() bool {
	for _, existing := range AllowedRequestPayerEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestPayer value
func (v RequestPayer) Ptr() *RequestPayer {
	return &v
}

type NullableRequestPayer struct {
	value *RequestPayer
	isSet bool
}

func (v NullableRequestPayer) Get() *RequestPayer {
	return v.value
}

func (v *NullableRequestPayer) Set(val *RequestPayer) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPayer) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPayer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPayer(val *RequestPayer) *NullableRequestPayer {
	return &NullableRequestPayer{value: val, isSet: true}
}

func (v NullableRequestPayer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPayer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

