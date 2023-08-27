# WafRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**UrlPattern** | Pointer to **string** | - &#x60;?&#x60; matches any single character. - &#x60;*&#x60; matches any (possibly empty) sequence of characters. - &#x60;**&#x60; matches the current directory and arbitrary subdirectories. This sequence must form a single path component, so both &#x60;**a&#x60; and &#x60;b**&#x60; are invalid and will result in an error. A sequence of more than two consecutive &#x60;*&#x60; characters is also invalid. - &#x60;[...]&#x60; matches any character inside the brackets. Character sequences can also specify ranges of characters, as ordered by Unicode, so e.g. &#x60;[0-9]&#x60; specifies any character between 0 and 9 inclusive. An unclosed bracket is invalid. - &#x60;[!...]&#x60; is the negation of &#x60;[...]&#x60;, i.e. it matches any characters not in the brackets. - The metacharacters &#x60;?&#x60;, &#x60;*&#x60;, &#x60;[&#x60;, &#x60;] &#x60;can be matched by using brackets (e.g. &#x60;[?]&#x60;). When a &#x60;]&#x60; occurs immediately following &#x60;[&#x60; or &#x60;[!&#x60; then it is interpreted as being part of, rather then ending, the character set, so &#x60;]&#x60; and NOT &#x60;]&#x60; can be matched by &#x60;[]]&#x60; and &#x60;[!]]&#x60; respectively. The - character can be specified inside a character sequence pattern by placing it at the start or the end, e.g. &#x60;[abc-]&#x60;.  | [optional] 
**Sources** | Pointer to **[]string** |  | [optional] 
**Action** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**IsEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewWafRule

`func NewWafRule() *WafRule`

NewWafRule instantiates a new WafRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafRuleWithDefaults

`func NewWafRuleWithDefaults() *WafRule`

NewWafRuleWithDefaults instantiates a new WafRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WafRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WafRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WafRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WafRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrlPattern

`func (o *WafRule) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *WafRule) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *WafRule) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.

### HasUrlPattern

`func (o *WafRule) HasUrlPattern() bool`

HasUrlPattern returns a boolean if a field has been set.

### GetSources

`func (o *WafRule) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *WafRule) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *WafRule) SetSources(v []string)`

SetSources sets Sources field to given value.

### HasSources

`func (o *WafRule) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetAction

`func (o *WafRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WafRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WafRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WafRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetDescription

`func (o *WafRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WafRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WafRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WafRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsEnabled

`func (o *WafRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *WafRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *WafRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *WafRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


