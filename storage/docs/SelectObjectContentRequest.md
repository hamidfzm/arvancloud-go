# SelectObjectContentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** |  | 
**ExpressionType** | [**ExpressionType**](ExpressionType.md) |  | 
**RequestProgress** | Pointer to [**SelectObjectContentRequestRequestProgress**](SelectObjectContentRequestRequestProgress.md) |  | [optional] 
**InputSerialization** | [**SelectObjectContentRequestInputSerialization**](SelectObjectContentRequestInputSerialization.md) |  | 
**OutputSerialization** | [**SelectObjectContentRequestOutputSerialization**](SelectObjectContentRequestOutputSerialization.md) |  | 
**ScanRange** | Pointer to [**SelectObjectContentRequestScanRange**](SelectObjectContentRequestScanRange.md) |  | [optional] 

## Methods

### NewSelectObjectContentRequest

`func NewSelectObjectContentRequest(expression string, expressionType ExpressionType, inputSerialization SelectObjectContentRequestInputSerialization, outputSerialization SelectObjectContentRequestOutputSerialization, ) *SelectObjectContentRequest`

NewSelectObjectContentRequest instantiates a new SelectObjectContentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectObjectContentRequestWithDefaults

`func NewSelectObjectContentRequestWithDefaults() *SelectObjectContentRequest`

NewSelectObjectContentRequestWithDefaults instantiates a new SelectObjectContentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *SelectObjectContentRequest) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SelectObjectContentRequest) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SelectObjectContentRequest) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetExpressionType

`func (o *SelectObjectContentRequest) GetExpressionType() ExpressionType`

GetExpressionType returns the ExpressionType field if non-nil, zero value otherwise.

### GetExpressionTypeOk

`func (o *SelectObjectContentRequest) GetExpressionTypeOk() (*ExpressionType, bool)`

GetExpressionTypeOk returns a tuple with the ExpressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionType

`func (o *SelectObjectContentRequest) SetExpressionType(v ExpressionType)`

SetExpressionType sets ExpressionType field to given value.


### GetRequestProgress

`func (o *SelectObjectContentRequest) GetRequestProgress() SelectObjectContentRequestRequestProgress`

GetRequestProgress returns the RequestProgress field if non-nil, zero value otherwise.

### GetRequestProgressOk

`func (o *SelectObjectContentRequest) GetRequestProgressOk() (*SelectObjectContentRequestRequestProgress, bool)`

GetRequestProgressOk returns a tuple with the RequestProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestProgress

`func (o *SelectObjectContentRequest) SetRequestProgress(v SelectObjectContentRequestRequestProgress)`

SetRequestProgress sets RequestProgress field to given value.

### HasRequestProgress

`func (o *SelectObjectContentRequest) HasRequestProgress() bool`

HasRequestProgress returns a boolean if a field has been set.

### GetInputSerialization

`func (o *SelectObjectContentRequest) GetInputSerialization() SelectObjectContentRequestInputSerialization`

GetInputSerialization returns the InputSerialization field if non-nil, zero value otherwise.

### GetInputSerializationOk

`func (o *SelectObjectContentRequest) GetInputSerializationOk() (*SelectObjectContentRequestInputSerialization, bool)`

GetInputSerializationOk returns a tuple with the InputSerialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSerialization

`func (o *SelectObjectContentRequest) SetInputSerialization(v SelectObjectContentRequestInputSerialization)`

SetInputSerialization sets InputSerialization field to given value.


### GetOutputSerialization

`func (o *SelectObjectContentRequest) GetOutputSerialization() SelectObjectContentRequestOutputSerialization`

GetOutputSerialization returns the OutputSerialization field if non-nil, zero value otherwise.

### GetOutputSerializationOk

`func (o *SelectObjectContentRequest) GetOutputSerializationOk() (*SelectObjectContentRequestOutputSerialization, bool)`

GetOutputSerializationOk returns a tuple with the OutputSerialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSerialization

`func (o *SelectObjectContentRequest) SetOutputSerialization(v SelectObjectContentRequestOutputSerialization)`

SetOutputSerialization sets OutputSerialization field to given value.


### GetScanRange

`func (o *SelectObjectContentRequest) GetScanRange() SelectObjectContentRequestScanRange`

GetScanRange returns the ScanRange field if non-nil, zero value otherwise.

### GetScanRangeOk

`func (o *SelectObjectContentRequest) GetScanRangeOk() (*SelectObjectContentRequestScanRange, bool)`

GetScanRangeOk returns a tuple with the ScanRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanRange

`func (o *SelectObjectContentRequest) SetScanRange(v SelectObjectContentRequestScanRange)`

SetScanRange sets ScanRange field to given value.

### HasScanRange

`func (o *SelectObjectContentRequest) HasScanRange() bool`

HasScanRange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


