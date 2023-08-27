# DdosRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**DdosRule**](DdosRule.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewDdosRuleResponse

`func NewDdosRuleResponse() *DdosRuleResponse`

NewDdosRuleResponse instantiates a new DdosRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDdosRuleResponseWithDefaults

`func NewDdosRuleResponseWithDefaults() *DdosRuleResponse`

NewDdosRuleResponseWithDefaults instantiates a new DdosRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DdosRuleResponse) GetData() DdosRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DdosRuleResponse) GetDataOk() (*DdosRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DdosRuleResponse) SetData(v DdosRule)`

SetData sets Data field to given value.

### HasData

`func (o *DdosRuleResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *DdosRuleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DdosRuleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DdosRuleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DdosRuleResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *DdosRuleResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *DdosRuleResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


