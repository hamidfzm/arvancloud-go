# SelectParametersInputSerialization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CSV** | Pointer to [**InputSerializationCSV**](InputSerializationCSV.md) |  | [optional] 
**CompressionType** | Pointer to [**CompressionType**](CompressionType.md) |  | [optional] 
**JSON** | Pointer to [**InputSerializationJSON**](InputSerializationJSON.md) |  | [optional] 
**Parquet** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSelectParametersInputSerialization

`func NewSelectParametersInputSerialization() *SelectParametersInputSerialization`

NewSelectParametersInputSerialization instantiates a new SelectParametersInputSerialization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectParametersInputSerializationWithDefaults

`func NewSelectParametersInputSerializationWithDefaults() *SelectParametersInputSerialization`

NewSelectParametersInputSerializationWithDefaults instantiates a new SelectParametersInputSerialization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCSV

`func (o *SelectParametersInputSerialization) GetCSV() InputSerializationCSV`

GetCSV returns the CSV field if non-nil, zero value otherwise.

### GetCSVOk

`func (o *SelectParametersInputSerialization) GetCSVOk() (*InputSerializationCSV, bool)`

GetCSVOk returns a tuple with the CSV field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSV

`func (o *SelectParametersInputSerialization) SetCSV(v InputSerializationCSV)`

SetCSV sets CSV field to given value.

### HasCSV

`func (o *SelectParametersInputSerialization) HasCSV() bool`

HasCSV returns a boolean if a field has been set.

### GetCompressionType

`func (o *SelectParametersInputSerialization) GetCompressionType() CompressionType`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *SelectParametersInputSerialization) GetCompressionTypeOk() (*CompressionType, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *SelectParametersInputSerialization) SetCompressionType(v CompressionType)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *SelectParametersInputSerialization) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### GetJSON

`func (o *SelectParametersInputSerialization) GetJSON() InputSerializationJSON`

GetJSON returns the JSON field if non-nil, zero value otherwise.

### GetJSONOk

`func (o *SelectParametersInputSerialization) GetJSONOk() (*InputSerializationJSON, bool)`

GetJSONOk returns a tuple with the JSON field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJSON

`func (o *SelectParametersInputSerialization) SetJSON(v InputSerializationJSON)`

SetJSON sets JSON field to given value.

### HasJSON

`func (o *SelectParametersInputSerialization) HasJSON() bool`

HasJSON returns a boolean if a field has been set.

### GetParquet

`func (o *SelectParametersInputSerialization) GetParquet() interface{}`

GetParquet returns the Parquet field if non-nil, zero value otherwise.

### GetParquetOk

`func (o *SelectParametersInputSerialization) GetParquetOk() (*interface{}, bool)`

GetParquetOk returns a tuple with the Parquet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParquet

`func (o *SelectParametersInputSerialization) SetParquet(v interface{})`

SetParquet sets Parquet field to given value.

### HasParquet

`func (o *SelectParametersInputSerialization) HasParquet() bool`

HasParquet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


