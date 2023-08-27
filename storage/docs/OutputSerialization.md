# OutputSerialization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CSV** | Pointer to [**OutputSerializationCSV**](OutputSerializationCSV.md) |  | [optional] 
**JSON** | Pointer to [**OutputSerializationJSON**](OutputSerializationJSON.md) |  | [optional] 

## Methods

### NewOutputSerialization

`func NewOutputSerialization() *OutputSerialization`

NewOutputSerialization instantiates a new OutputSerialization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputSerializationWithDefaults

`func NewOutputSerializationWithDefaults() *OutputSerialization`

NewOutputSerializationWithDefaults instantiates a new OutputSerialization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCSV

`func (o *OutputSerialization) GetCSV() OutputSerializationCSV`

GetCSV returns the CSV field if non-nil, zero value otherwise.

### GetCSVOk

`func (o *OutputSerialization) GetCSVOk() (*OutputSerializationCSV, bool)`

GetCSVOk returns a tuple with the CSV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSV

`func (o *OutputSerialization) SetCSV(v OutputSerializationCSV)`

SetCSV sets CSV field to given value.

### HasCSV

`func (o *OutputSerialization) HasCSV() bool`

HasCSV returns a boolean if a field has been set.

### GetJSON

`func (o *OutputSerialization) GetJSON() OutputSerializationJSON`

GetJSON returns the JSON field if non-nil, zero value otherwise.

### GetJSONOk

`func (o *OutputSerialization) GetJSONOk() (*OutputSerializationJSON, bool)`

GetJSONOk returns a tuple with the JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJSON

`func (o *OutputSerialization) SetJSON(v OutputSerializationJSON)`

SetJSON sets JSON field to given value.

### HasJSON

`func (o *OutputSerialization) HasJSON() bool`

HasJSON returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


