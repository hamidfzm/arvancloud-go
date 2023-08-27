# SelectParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputSerialization** | [**SelectParametersInputSerialization**](SelectParametersInputSerialization.md) |  | 
**ExpressionType** | [**ExpressionType**](ExpressionType.md) |  | 
**Expression** | **string** |  | 
**OutputSerialization** | [**SelectParametersOutputSerialization**](SelectParametersOutputSerialization.md) |  | 

## Methods

### NewSelectParameters

`func NewSelectParameters(inputSerialization SelectParametersInputSerialization, expressionType ExpressionType, expression string, outputSerialization SelectParametersOutputSerialization, ) *SelectParameters`

NewSelectParameters instantiates a new SelectParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectParametersWithDefaults

`func NewSelectParametersWithDefaults() *SelectParameters`

NewSelectParametersWithDefaults instantiates a new SelectParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputSerialization

`func (o *SelectParameters) GetInputSerialization() SelectParametersInputSerialization`

GetInputSerialization returns the InputSerialization field if non-nil, zero value otherwise.

### GetInputSerializationOk

`func (o *SelectParameters) GetInputSerializationOk() (*SelectParametersInputSerialization, bool)`

GetInputSerializationOk returns a tuple with the InputSerialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSerialization

`func (o *SelectParameters) SetInputSerialization(v SelectParametersInputSerialization)`

SetInputSerialization sets InputSerialization field to given value.


### GetExpressionType

`func (o *SelectParameters) GetExpressionType() ExpressionType`

GetExpressionType returns the ExpressionType field if non-nil, zero value otherwise.

### GetExpressionTypeOk

`func (o *SelectParameters) GetExpressionTypeOk() (*ExpressionType, bool)`

GetExpressionTypeOk returns a tuple with the ExpressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressionType

`func (o *SelectParameters) SetExpressionType(v ExpressionType)`

SetExpressionType sets ExpressionType field to given value.


### GetExpression

`func (o *SelectParameters) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *SelectParameters) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *SelectParameters) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetOutputSerialization

`func (o *SelectParameters) GetOutputSerialization() SelectParametersOutputSerialization`

GetOutputSerialization returns the OutputSerialization field if non-nil, zero value otherwise.

### GetOutputSerializationOk

`func (o *SelectParameters) GetOutputSerializationOk() (*SelectParametersOutputSerialization, bool)`

GetOutputSerializationOk returns a tuple with the OutputSerialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSerialization

`func (o *SelectParameters) SetOutputSerialization(v SelectParametersOutputSerialization)`

SetOutputSerialization sets OutputSerialization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


