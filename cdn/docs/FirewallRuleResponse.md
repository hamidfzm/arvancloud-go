# FirewallRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**FirewallRuleView**](FirewallRuleView.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFirewallRuleResponse

`func NewFirewallRuleResponse() *FirewallRuleResponse`

NewFirewallRuleResponse instantiates a new FirewallRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRuleResponseWithDefaults

`func NewFirewallRuleResponseWithDefaults() *FirewallRuleResponse`

NewFirewallRuleResponseWithDefaults instantiates a new FirewallRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FirewallRuleResponse) GetData() FirewallRuleView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FirewallRuleResponse) GetDataOk() (*FirewallRuleView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FirewallRuleResponse) SetData(v FirewallRuleView)`

SetData sets Data field to given value.

### HasData

`func (o *FirewallRuleResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallRuleResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallRuleResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallRuleResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallRuleResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *FirewallRuleResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *FirewallRuleResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


