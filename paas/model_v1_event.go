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

// checks if the V1Event type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1Event{}

// V1Event Event is a report of an event somewhere in the cluster.
type V1Event struct {
	// What action was taken/failed regarding to the Regarding object.
	Action *string `json:"action,omitempty"`
	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
	ApiVersion *string `json:"apiVersion,omitempty"`
	// The number of times this event has occurred.
	Count *int32 `json:"count,omitempty"`
	// Time when this Event was first observed.
	EventTime *string `json:"eventTime,omitempty"`
	// The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
	FirstTimestamp *string `json:"firstTimestamp,omitempty"`
	InvolvedObject V1ObjectReference `json:"involvedObject"`
	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
	Kind *string `json:"kind,omitempty"`
	// The time at which the most recent occurrence of this event was recorded.
	LastTimestamp *string `json:"lastTimestamp,omitempty"`
	// A human-readable description of the status of this operation.
	Message *string `json:"message,omitempty"`
	Metadata V1ObjectMeta `json:"metadata"`
	// This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
	Reason *string `json:"reason,omitempty"`
	Related *V1ObjectReference `json:"related,omitempty"`
	// Name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`.
	ReportingComponent string `json:"reportingComponent"`
	// ID of the controller instance, e.g. `kubelet-xyzf`.
	ReportingInstance string `json:"reportingInstance"`
	Series *V1EventSeries `json:"series,omitempty"`
	Source *V1EventSource `json:"source,omitempty"`
	// Type of this event (Normal, Warning), new types could be added in the future
	Type *string `json:"type,omitempty"`
}

// NewV1Event instantiates a new V1Event object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Event(involvedObject V1ObjectReference, metadata V1ObjectMeta, reportingComponent string, reportingInstance string) *V1Event {
	this := V1Event{}
	this.InvolvedObject = involvedObject
	this.Metadata = metadata
	this.ReportingComponent = reportingComponent
	this.ReportingInstance = reportingInstance
	return &this
}

// NewV1EventWithDefaults instantiates a new V1Event object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EventWithDefaults() *V1Event {
	this := V1Event{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *V1Event) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *V1Event) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *V1Event) SetAction(v string) {
	o.Action = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *V1Event) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *V1Event) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *V1Event) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V1Event) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V1Event) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *V1Event) SetCount(v int32) {
	o.Count = &v
}

// GetEventTime returns the EventTime field value if set, zero value otherwise.
func (o *V1Event) GetEventTime() string {
	if o == nil || IsNil(o.EventTime) {
		var ret string
		return ret
	}
	return *o.EventTime
}

// GetEventTimeOk returns a tuple with the EventTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetEventTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EventTime) {
		return nil, false
	}
	return o.EventTime, true
}

// HasEventTime returns a boolean if a field has been set.
func (o *V1Event) HasEventTime() bool {
	if o != nil && !IsNil(o.EventTime) {
		return true
	}

	return false
}

// SetEventTime gets a reference to the given string and assigns it to the EventTime field.
func (o *V1Event) SetEventTime(v string) {
	o.EventTime = &v
}

// GetFirstTimestamp returns the FirstTimestamp field value if set, zero value otherwise.
func (o *V1Event) GetFirstTimestamp() string {
	if o == nil || IsNil(o.FirstTimestamp) {
		var ret string
		return ret
	}
	return *o.FirstTimestamp
}

// GetFirstTimestampOk returns a tuple with the FirstTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetFirstTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.FirstTimestamp) {
		return nil, false
	}
	return o.FirstTimestamp, true
}

// HasFirstTimestamp returns a boolean if a field has been set.
func (o *V1Event) HasFirstTimestamp() bool {
	if o != nil && !IsNil(o.FirstTimestamp) {
		return true
	}

	return false
}

// SetFirstTimestamp gets a reference to the given string and assigns it to the FirstTimestamp field.
func (o *V1Event) SetFirstTimestamp(v string) {
	o.FirstTimestamp = &v
}

// GetInvolvedObject returns the InvolvedObject field value
func (o *V1Event) GetInvolvedObject() V1ObjectReference {
	if o == nil {
		var ret V1ObjectReference
		return ret
	}

	return o.InvolvedObject
}

// GetInvolvedObjectOk returns a tuple with the InvolvedObject field value
// and a boolean to check if the value has been set.
func (o *V1Event) GetInvolvedObjectOk() (*V1ObjectReference, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvolvedObject, true
}

// SetInvolvedObject sets field value
func (o *V1Event) SetInvolvedObject(v V1ObjectReference) {
	o.InvolvedObject = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *V1Event) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *V1Event) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *V1Event) SetKind(v string) {
	o.Kind = &v
}

// GetLastTimestamp returns the LastTimestamp field value if set, zero value otherwise.
func (o *V1Event) GetLastTimestamp() string {
	if o == nil || IsNil(o.LastTimestamp) {
		var ret string
		return ret
	}
	return *o.LastTimestamp
}

// GetLastTimestampOk returns a tuple with the LastTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetLastTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.LastTimestamp) {
		return nil, false
	}
	return o.LastTimestamp, true
}

// HasLastTimestamp returns a boolean if a field has been set.
func (o *V1Event) HasLastTimestamp() bool {
	if o != nil && !IsNil(o.LastTimestamp) {
		return true
	}

	return false
}

// SetLastTimestamp gets a reference to the given string and assigns it to the LastTimestamp field.
func (o *V1Event) SetLastTimestamp(v string) {
	o.LastTimestamp = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1Event) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1Event) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1Event) SetMessage(v string) {
	o.Message = &v
}

// GetMetadata returns the Metadata field value
func (o *V1Event) GetMetadata() V1ObjectMeta {
	if o == nil {
		var ret V1ObjectMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *V1Event) GetMetadataOk() (*V1ObjectMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *V1Event) SetMetadata(v V1ObjectMeta) {
	o.Metadata = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *V1Event) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *V1Event) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *V1Event) SetReason(v string) {
	o.Reason = &v
}

// GetRelated returns the Related field value if set, zero value otherwise.
func (o *V1Event) GetRelated() V1ObjectReference {
	if o == nil || IsNil(o.Related) {
		var ret V1ObjectReference
		return ret
	}
	return *o.Related
}

// GetRelatedOk returns a tuple with the Related field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetRelatedOk() (*V1ObjectReference, bool) {
	if o == nil || IsNil(o.Related) {
		return nil, false
	}
	return o.Related, true
}

// HasRelated returns a boolean if a field has been set.
func (o *V1Event) HasRelated() bool {
	if o != nil && !IsNil(o.Related) {
		return true
	}

	return false
}

// SetRelated gets a reference to the given V1ObjectReference and assigns it to the Related field.
func (o *V1Event) SetRelated(v V1ObjectReference) {
	o.Related = &v
}

// GetReportingComponent returns the ReportingComponent field value
func (o *V1Event) GetReportingComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportingComponent
}

// GetReportingComponentOk returns a tuple with the ReportingComponent field value
// and a boolean to check if the value has been set.
func (o *V1Event) GetReportingComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportingComponent, true
}

// SetReportingComponent sets field value
func (o *V1Event) SetReportingComponent(v string) {
	o.ReportingComponent = v
}

// GetReportingInstance returns the ReportingInstance field value
func (o *V1Event) GetReportingInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReportingInstance
}

// GetReportingInstanceOk returns a tuple with the ReportingInstance field value
// and a boolean to check if the value has been set.
func (o *V1Event) GetReportingInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReportingInstance, true
}

// SetReportingInstance sets field value
func (o *V1Event) SetReportingInstance(v string) {
	o.ReportingInstance = v
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *V1Event) GetSeries() V1EventSeries {
	if o == nil || IsNil(o.Series) {
		var ret V1EventSeries
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetSeriesOk() (*V1EventSeries, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *V1Event) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given V1EventSeries and assigns it to the Series field.
func (o *V1Event) SetSeries(v V1EventSeries) {
	o.Series = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *V1Event) GetSource() V1EventSource {
	if o == nil || IsNil(o.Source) {
		var ret V1EventSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetSourceOk() (*V1EventSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *V1Event) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given V1EventSource and assigns it to the Source field.
func (o *V1Event) SetSource(v V1EventSource) {
	o.Source = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1Event) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Event) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1Event) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V1Event) SetType(v string) {
	o.Type = &v
}

func (o V1Event) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1Event) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.EventTime) {
		toSerialize["eventTime"] = o.EventTime
	}
	if !IsNil(o.FirstTimestamp) {
		toSerialize["firstTimestamp"] = o.FirstTimestamp
	}
	toSerialize["involvedObject"] = o.InvolvedObject
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.LastTimestamp) {
		toSerialize["lastTimestamp"] = o.LastTimestamp
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	toSerialize["metadata"] = o.Metadata
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Related) {
		toSerialize["related"] = o.Related
	}
	toSerialize["reportingComponent"] = o.ReportingComponent
	toSerialize["reportingInstance"] = o.ReportingInstance
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableV1Event struct {
	value *V1Event
	isSet bool
}

func (v NullableV1Event) Get() *V1Event {
	return v.value
}

func (v *NullableV1Event) Set(val *V1Event) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Event) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Event) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Event(val *V1Event) *NullableV1Event {
	return &NullableV1Event{value: val, isSet: true}
}

func (v NullableV1Event) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Event) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


