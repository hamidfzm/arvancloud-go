# FirewallSettingsIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**FirewallSettingsView**](FirewallSettingsView.md) |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewFirewallSettingsIndex200Response

`func NewFirewallSettingsIndex200Response() *FirewallSettingsIndex200Response`

NewFirewallSettingsIndex200Response instantiates a new FirewallSettingsIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallSettingsIndex200ResponseWithDefaults

`func NewFirewallSettingsIndex200ResponseWithDefaults() *FirewallSettingsIndex200Response`

NewFirewallSettingsIndex200ResponseWithDefaults instantiates a new FirewallSettingsIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FirewallSettingsIndex200Response) GetData() FirewallSettingsView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FirewallSettingsIndex200Response) GetDataOk() (*FirewallSettingsView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FirewallSettingsIndex200Response) SetData(v FirewallSettingsView)`

SetData sets Data field to given value.

### HasData

`func (o *FirewallSettingsIndex200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *FirewallSettingsIndex200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FirewallSettingsIndex200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FirewallSettingsIndex200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *FirewallSettingsIndex200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *FirewallSettingsIndex200Response) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *FirewallSettingsIndex200Response) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


