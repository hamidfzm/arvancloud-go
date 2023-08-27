# InputSerializationCSV

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileHeaderInfo** | Pointer to [**FileHeaderInfo**](FileHeaderInfo.md) |  | [optional] 
**Comments** | Pointer to **string** |  | [optional] 
**QuoteEscapeCharacter** | Pointer to **string** |  | [optional] 
**RecordDelimiter** | Pointer to **string** |  | [optional] 
**FieldDelimiter** | Pointer to **string** |  | [optional] 
**QuoteCharacter** | Pointer to **string** |  | [optional] 
**AllowQuotedRecordDelimiter** | Pointer to **bool** |  | [optional] 

## Methods

### NewInputSerializationCSV

`func NewInputSerializationCSV() *InputSerializationCSV`

NewInputSerializationCSV instantiates a new InputSerializationCSV object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInputSerializationCSVWithDefaults

`func NewInputSerializationCSVWithDefaults() *InputSerializationCSV`

NewInputSerializationCSVWithDefaults instantiates a new InputSerializationCSV object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHeaderInfo

`func (o *InputSerializationCSV) GetFileHeaderInfo() FileHeaderInfo`

GetFileHeaderInfo returns the FileHeaderInfo field if non-nil, zero value otherwise.

### GetFileHeaderInfoOk

`func (o *InputSerializationCSV) GetFileHeaderInfoOk() (*FileHeaderInfo, bool)`

GetFileHeaderInfoOk returns a tuple with the FileHeaderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHeaderInfo

`func (o *InputSerializationCSV) SetFileHeaderInfo(v FileHeaderInfo)`

SetFileHeaderInfo sets FileHeaderInfo field to given value.

### HasFileHeaderInfo

`func (o *InputSerializationCSV) HasFileHeaderInfo() bool`

HasFileHeaderInfo returns a boolean if a field has been set.

### GetComments

`func (o *InputSerializationCSV) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InputSerializationCSV) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InputSerializationCSV) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *InputSerializationCSV) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetQuoteEscapeCharacter

`func (o *InputSerializationCSV) GetQuoteEscapeCharacter() string`

GetQuoteEscapeCharacter returns the QuoteEscapeCharacter field if non-nil, zero value otherwise.

### GetQuoteEscapeCharacterOk

`func (o *InputSerializationCSV) GetQuoteEscapeCharacterOk() (*string, bool)`

GetQuoteEscapeCharacterOk returns a tuple with the QuoteEscapeCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteEscapeCharacter

`func (o *InputSerializationCSV) SetQuoteEscapeCharacter(v string)`

SetQuoteEscapeCharacter sets QuoteEscapeCharacter field to given value.

### HasQuoteEscapeCharacter

`func (o *InputSerializationCSV) HasQuoteEscapeCharacter() bool`

HasQuoteEscapeCharacter returns a boolean if a field has been set.

### GetRecordDelimiter

`func (o *InputSerializationCSV) GetRecordDelimiter() string`

GetRecordDelimiter returns the RecordDelimiter field if non-nil, zero value otherwise.

### GetRecordDelimiterOk

`func (o *InputSerializationCSV) GetRecordDelimiterOk() (*string, bool)`

GetRecordDelimiterOk returns a tuple with the RecordDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDelimiter

`func (o *InputSerializationCSV) SetRecordDelimiter(v string)`

SetRecordDelimiter sets RecordDelimiter field to given value.

### HasRecordDelimiter

`func (o *InputSerializationCSV) HasRecordDelimiter() bool`

HasRecordDelimiter returns a boolean if a field has been set.

### GetFieldDelimiter

`func (o *InputSerializationCSV) GetFieldDelimiter() string`

GetFieldDelimiter returns the FieldDelimiter field if non-nil, zero value otherwise.

### GetFieldDelimiterOk

`func (o *InputSerializationCSV) GetFieldDelimiterOk() (*string, bool)`

GetFieldDelimiterOk returns a tuple with the FieldDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDelimiter

`func (o *InputSerializationCSV) SetFieldDelimiter(v string)`

SetFieldDelimiter sets FieldDelimiter field to given value.

### HasFieldDelimiter

`func (o *InputSerializationCSV) HasFieldDelimiter() bool`

HasFieldDelimiter returns a boolean if a field has been set.

### GetQuoteCharacter

`func (o *InputSerializationCSV) GetQuoteCharacter() string`

GetQuoteCharacter returns the QuoteCharacter field if non-nil, zero value otherwise.

### GetQuoteCharacterOk

`func (o *InputSerializationCSV) GetQuoteCharacterOk() (*string, bool)`

GetQuoteCharacterOk returns a tuple with the QuoteCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCharacter

`func (o *InputSerializationCSV) SetQuoteCharacter(v string)`

SetQuoteCharacter sets QuoteCharacter field to given value.

### HasQuoteCharacter

`func (o *InputSerializationCSV) HasQuoteCharacter() bool`

HasQuoteCharacter returns a boolean if a field has been set.

### GetAllowQuotedRecordDelimiter

`func (o *InputSerializationCSV) GetAllowQuotedRecordDelimiter() bool`

GetAllowQuotedRecordDelimiter returns the AllowQuotedRecordDelimiter field if non-nil, zero value otherwise.

### GetAllowQuotedRecordDelimiterOk

`func (o *InputSerializationCSV) GetAllowQuotedRecordDelimiterOk() (*bool, bool)`

GetAllowQuotedRecordDelimiterOk returns a tuple with the AllowQuotedRecordDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuotedRecordDelimiter

`func (o *InputSerializationCSV) SetAllowQuotedRecordDelimiter(v bool)`

SetAllowQuotedRecordDelimiter sets AllowQuotedRecordDelimiter field to given value.

### HasAllowQuotedRecordDelimiter

`func (o *InputSerializationCSV) HasAllowQuotedRecordDelimiter() bool`

HasAllowQuotedRecordDelimiter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


