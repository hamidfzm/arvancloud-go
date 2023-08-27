# SelectObjectContentRequestInputSerialization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CSV** | Pointer to [**InputSerializationCSV**](InputSerializationCSV.md) |  | [optional] 
**CompressionType** | Pointer to [**CompressionType**](CompressionType.md) |  | [optional] 
**JSON** | Pointer to [**InputSerializationJSON**](InputSerializationJSON.md) |  | [optional] 
**Parquet** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSelectObjectContentRequestInputSerialization

`func NewSelectObjectContentRequestInputSerialization() *SelectObjectContentRequestInputSerialization`

NewSelectObjectContentRequestInputSerialization instantiates a new SelectObjectContentRequestInputSerialization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectObjectContentRequestInputSerializationWithDefaults

`func NewSelectObjectContentRequestInputSerializationWithDefaults() *SelectObjectContentRequestInputSerialization`

NewSelectObjectContentRequestInputSerializationWithDefaults instantiates a new SelectObjectContentRequestInputSerialization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCSV

`func (o *SelectObjectContentRequestInputSerialization) GetCSV() InputSerializationCSV`

GetCSV returns the CSV field if non-nil, zero value otherwise.

### GetCSVOk

`func (o *SelectObjectContentRequestInputSerialization) GetCSVOk() (*InputSerializationCSV, bool)`

GetCSVOk returns a tuple with the CSV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSV

`func (o *SelectObjectContentRequestInputSerialization) SetCSV(v InputSerializationCSV)`

SetCSV sets CSV field to given value.

### HasCSV

`func (o *SelectObjectContentRequestInputSerialization) HasCSV() bool`

HasCSV returns a boolean if a field has been set.

### GetCompressionType

`func (o *SelectObjectContentRequestInputSerialization) GetCompressionType() CompressionType`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *SelectObjectContentRequestInputSerialization) GetCompressionTypeOk() (*CompressionType, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *SelectObjectContentRequestInputSerialization) SetCompressionType(v CompressionType)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *SelectObjectContentRequestInputSerialization) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### GetJSON

`func (o *SelectObjectContentRequestInputSerialization) GetJSON() InputSerializationJSON`

GetJSON returns the JSON field if non-nil, zero value otherwise.

### GetJSONOk

`func (o *SelectObjectContentRequestInputSerialization) GetJSONOk() (*InputSerializationJSON, bool)`

GetJSONOk returns a tuple with the JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJSON

`func (o *SelectObjectContentRequestInputSerialization) SetJSON(v InputSerializationJSON)`

SetJSON sets JSON field to given value.

### HasJSON

`func (o *SelectObjectContentRequestInputSerialization) HasJSON() bool`

HasJSON returns a boolean if a field has been set.

### GetParquet

`func (o *SelectObjectContentRequestInputSerialization) GetParquet() interface{}`

GetParquet returns the Parquet field if non-nil, zero value otherwise.

### GetParquetOk

`func (o *SelectObjectContentRequestInputSerialization) GetParquetOk() (*interface{}, bool)`

GetParquetOk returns a tuple with the Parquet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParquet

`func (o *SelectObjectContentRequestInputSerialization) SetParquet(v interface{})`

SetParquet sets Parquet field to given value.

### HasParquet

`func (o *SelectObjectContentRequestInputSerialization) HasParquet() bool`

HasParquet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


