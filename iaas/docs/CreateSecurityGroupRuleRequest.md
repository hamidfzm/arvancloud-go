# CreateSecurityGroupRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** | connection direction, either &#x60;ingress&#x60; or &#x60;egress&#x60; | [optional] 
**Ips** | Pointer to **[]string** |  | [optional] 
**PortFrom** | Pointer to **map[string]interface{}** | starting port, ex. 443 | [optional] 
**PortTo** | Pointer to **map[string]interface{}** | ending port, ex. 8080 | [optional] 
**Protocol** | Pointer to **string** | &#x60;udp&#x60; or &#x60;tcp&#x60; | [optional] 

## Methods

### NewCreateSecurityGroupRuleRequest

`func NewCreateSecurityGroupRuleRequest() *CreateSecurityGroupRuleRequest`

NewCreateSecurityGroupRuleRequest instantiates a new CreateSecurityGroupRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSecurityGroupRuleRequestWithDefaults

`func NewCreateSecurityGroupRuleRequestWithDefaults() *CreateSecurityGroupRuleRequest`

NewCreateSecurityGroupRuleRequestWithDefaults instantiates a new CreateSecurityGroupRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateSecurityGroupRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSecurityGroupRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSecurityGroupRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSecurityGroupRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDirection

`func (o *CreateSecurityGroupRuleRequest) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CreateSecurityGroupRuleRequest) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CreateSecurityGroupRuleRequest) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CreateSecurityGroupRuleRequest) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetIps

`func (o *CreateSecurityGroupRuleRequest) GetIps() []string`

GetIps returns the Ips field if non-nil, zero value otherwise.

### GetIpsOk

`func (o *CreateSecurityGroupRuleRequest) GetIpsOk() (*[]string, bool)`

GetIpsOk returns a tuple with the Ips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIps

`func (o *CreateSecurityGroupRuleRequest) SetIps(v []string)`

SetIps sets Ips field to given value.

### HasIps

`func (o *CreateSecurityGroupRuleRequest) HasIps() bool`

HasIps returns a boolean if a field has been set.

### GetPortFrom

`func (o *CreateSecurityGroupRuleRequest) GetPortFrom() map[string]interface{}`

GetPortFrom returns the PortFrom field if non-nil, zero value otherwise.

### GetPortFromOk

`func (o *CreateSecurityGroupRuleRequest) GetPortFromOk() (*map[string]interface{}, bool)`

GetPortFromOk returns a tuple with the PortFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortFrom

`func (o *CreateSecurityGroupRuleRequest) SetPortFrom(v map[string]interface{})`

SetPortFrom sets PortFrom field to given value.

### HasPortFrom

`func (o *CreateSecurityGroupRuleRequest) HasPortFrom() bool`

HasPortFrom returns a boolean if a field has been set.

### GetPortTo

`func (o *CreateSecurityGroupRuleRequest) GetPortTo() map[string]interface{}`

GetPortTo returns the PortTo field if non-nil, zero value otherwise.

### GetPortToOk

`func (o *CreateSecurityGroupRuleRequest) GetPortToOk() (*map[string]interface{}, bool)`

GetPortToOk returns a tuple with the PortTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortTo

`func (o *CreateSecurityGroupRuleRequest) SetPortTo(v map[string]interface{})`

SetPortTo sets PortTo field to given value.

### HasPortTo

`func (o *CreateSecurityGroupRuleRequest) HasPortTo() bool`

HasPortTo returns a boolean if a field has been set.

### GetProtocol

`func (o *CreateSecurityGroupRuleRequest) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *CreateSecurityGroupRuleRequest) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *CreateSecurityGroupRuleRequest) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *CreateSecurityGroupRuleRequest) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


