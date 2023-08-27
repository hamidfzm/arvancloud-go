# CSVInput

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

### NewCSVInput

`func NewCSVInput() *CSVInput`

NewCSVInput instantiates a new CSVInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSVInputWithDefaults

`func NewCSVInputWithDefaults() *CSVInput`

NewCSVInputWithDefaults instantiates a new CSVInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileHeaderInfo

`func (o *CSVInput) GetFileHeaderInfo() FileHeaderInfo`

GetFileHeaderInfo returns the FileHeaderInfo field if non-nil, zero value otherwise.

### GetFileHeaderInfoOk

`func (o *CSVInput) GetFileHeaderInfoOk() (*FileHeaderInfo, bool)`

GetFileHeaderInfoOk returns a tuple with the FileHeaderInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileHeaderInfo

`func (o *CSVInput) SetFileHeaderInfo(v FileHeaderInfo)`

SetFileHeaderInfo sets FileHeaderInfo field to given value.

### HasFileHeaderInfo

`func (o *CSVInput) HasFileHeaderInfo() bool`

HasFileHeaderInfo returns a boolean if a field has been set.

### GetComments

`func (o *CSVInput) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CSVInput) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CSVInput) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CSVInput) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetQuoteEscapeCharacter

`func (o *CSVInput) GetQuoteEscapeCharacter() string`

GetQuoteEscapeCharacter returns the QuoteEscapeCharacter field if non-nil, zero value otherwise.

### GetQuoteEscapeCharacterOk

`func (o *CSVInput) GetQuoteEscapeCharacterOk() (*string, bool)`

GetQuoteEscapeCharacterOk returns a tuple with the QuoteEscapeCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteEscapeCharacter

`func (o *CSVInput) SetQuoteEscapeCharacter(v string)`

SetQuoteEscapeCharacter sets QuoteEscapeCharacter field to given value.

### HasQuoteEscapeCharacter

`func (o *CSVInput) HasQuoteEscapeCharacter() bool`

HasQuoteEscapeCharacter returns a boolean if a field has been set.

### GetRecordDelimiter

`func (o *CSVInput) GetRecordDelimiter() string`

GetRecordDelimiter returns the RecordDelimiter field if non-nil, zero value otherwise.

### GetRecordDelimiterOk

`func (o *CSVInput) GetRecordDelimiterOk() (*string, bool)`

GetRecordDelimiterOk returns a tuple with the RecordDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordDelimiter

`func (o *CSVInput) SetRecordDelimiter(v string)`

SetRecordDelimiter sets RecordDelimiter field to given value.

### HasRecordDelimiter

`func (o *CSVInput) HasRecordDelimiter() bool`

HasRecordDelimiter returns a boolean if a field has been set.

### GetFieldDelimiter

`func (o *CSVInput) GetFieldDelimiter() string`

GetFieldDelimiter returns the FieldDelimiter field if non-nil, zero value otherwise.

### GetFieldDelimiterOk

`func (o *CSVInput) GetFieldDelimiterOk() (*string, bool)`

GetFieldDelimiterOk returns a tuple with the FieldDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldDelimiter

`func (o *CSVInput) SetFieldDelimiter(v string)`

SetFieldDelimiter sets FieldDelimiter field to given value.

### HasFieldDelimiter

`func (o *CSVInput) HasFieldDelimiter() bool`

HasFieldDelimiter returns a boolean if a field has been set.

### GetQuoteCharacter

`func (o *CSVInput) GetQuoteCharacter() string`

GetQuoteCharacter returns the QuoteCharacter field if non-nil, zero value otherwise.

### GetQuoteCharacterOk

`func (o *CSVInput) GetQuoteCharacterOk() (*string, bool)`

GetQuoteCharacterOk returns a tuple with the QuoteCharacter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCharacter

`func (o *CSVInput) SetQuoteCharacter(v string)`

SetQuoteCharacter sets QuoteCharacter field to given value.

### HasQuoteCharacter

`func (o *CSVInput) HasQuoteCharacter() bool`

HasQuoteCharacter returns a boolean if a field has been set.

### GetAllowQuotedRecordDelimiter

`func (o *CSVInput) GetAllowQuotedRecordDelimiter() bool`

GetAllowQuotedRecordDelimiter returns the AllowQuotedRecordDelimiter field if non-nil, zero value otherwise.

### GetAllowQuotedRecordDelimiterOk

`func (o *CSVInput) GetAllowQuotedRecordDelimiterOk() (*bool, bool)`

GetAllowQuotedRecordDelimiterOk returns a tuple with the AllowQuotedRecordDelimiter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuotedRecordDelimiter

`func (o *CSVInput) SetAllowQuotedRecordDelimiter(v bool)`

SetAllowQuotedRecordDelimiter sets AllowQuotedRecordDelimiter field to given value.

### HasAllowQuotedRecordDelimiter

`func (o *CSVInput) HasAllowQuotedRecordDelimiter() bool`

HasAllowQuotedRecordDelimiter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


