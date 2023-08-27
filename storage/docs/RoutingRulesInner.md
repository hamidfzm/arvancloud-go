# RoutingRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | Pointer to [**RoutingRuleCondition**](RoutingRuleCondition.md) |  | [optional] 
**Redirect** | [**RoutingRuleRedirect**](RoutingRuleRedirect.md) |  | 

## Methods

### NewRoutingRulesInner

`func NewRoutingRulesInner(redirect RoutingRuleRedirect, ) *RoutingRulesInner`

NewRoutingRulesInner instantiates a new RoutingRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoutingRulesInnerWithDefaults

`func NewRoutingRulesInnerWithDefaults() *RoutingRulesInner`

NewRoutingRulesInnerWithDefaults instantiates a new RoutingRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *RoutingRulesInner) GetCondition() RoutingRuleCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *RoutingRulesInner) GetConditionOk() (*RoutingRuleCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *RoutingRulesInner) SetCondition(v RoutingRuleCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *RoutingRulesInner) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetRedirect

`func (o *RoutingRulesInner) GetRedirect() RoutingRuleRedirect`

GetRedirect returns the Redirect field if non-nil, zero value otherwise.

### GetRedirectOk

`func (o *RoutingRulesInner) GetRedirectOk() (*RoutingRuleRedirect, bool)`

GetRedirectOk returns a tuple with the Redirect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirect

`func (o *RoutingRulesInner) SetRedirect(v RoutingRuleRedirect)`

SetRedirect sets Redirect field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


