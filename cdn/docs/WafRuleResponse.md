# WafRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**WafRule**](WafRule.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewWafRuleResponse

`func NewWafRuleResponse() *WafRuleResponse`

NewWafRuleResponse instantiates a new WafRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleResponseWithDefaults

`func NewWafRuleResponseWithDefaults() *WafRuleResponse`

NewWafRuleResponseWithDefaults instantiates a new WafRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WafRuleResponse) GetData() WafRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WafRuleResponse) GetDataOk() (*WafRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WafRuleResponse) SetData(v WafRule)`

SetData sets Data field to given value.

### HasData

`func (o *WafRuleResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *WafRuleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WafRuleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WafRuleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WafRuleResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *WafRuleResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *WafRuleResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


