# ProgressEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**ProgressEventDetails**](ProgressEventDetails.md) |  | [optional] 

## Methods

### NewProgressEvent

`func NewProgressEvent() *ProgressEvent`

NewProgressEvent instantiates a new ProgressEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProgressEventWithDefaults

`func NewProgressEventWithDefaults() *ProgressEvent`

NewProgressEventWithDefaults instantiates a new ProgressEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ProgressEvent) GetDetails() ProgressEventDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ProgressEvent) GetDetailsOk() (*ProgressEventDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ProgressEvent) SetDetails(v ProgressEventDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ProgressEvent) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


