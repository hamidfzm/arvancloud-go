# RateLimitRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActionDetails** | Pointer to [**ChallengeAction**](ChallengeAction.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Action** | Pointer to **string** |  | [optional] [default to "block"]
**IsEnabled** | Pointer to **bool** |  | [optional] 
**UrlPattern** | **string** | - &#x60;?&#x60; matches any single character. - &#x60;*&#x60; matches any (possibly empty) sequence of characters. - &#x60;**&#x60; matches the current directory and arbitrary subdirectories. This sequence must form a single path component, so both &#x60;**a&#x60; and &#x60;b**&#x60; are invalid and will result in an error. A sequence of more than two consecutive &#x60;*&#x60; characters is also invalid. - &#x60;[...]&#x60; matches any character inside the brackets. Character sequences can also specify ranges of characters, as ordered by Unicode, so e.g. &#x60;[0-9]&#x60; specifies any character between 0 and 9 inclusive. An unclosed bracket is invalid. - &#x60;[!...]&#x60; is the negation of &#x60;[...]&#x60;, i.e. it matches any characters not in the brackets. - The metacharacters &#x60;?&#x60;, &#x60;*&#x60;, &#x60;[&#x60;, &#x60;] &#x60;can be matched by using brackets (e.g. &#x60;[?]&#x60;). When a &#x60;]&#x60; occurs immediately following &#x60;[&#x60; or &#x60;[!&#x60; then it is interpreted as being part of, rather then ending, the character set, so &#x60;]&#x60; and NOT &#x60;]&#x60; can be matched by &#x60;[]]&#x60; and &#x60;[!]]&#x60; respectively. The - character can be specified inside a character sequence pattern by placing it at the start or the end, e.g. &#x60;[abc-]&#x60;.  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**ExcludeSources** | Pointer to **[]string** |  | [optional] 
**Rate** | **int32** |  | 
**Burst** | Pointer to **int32** |  | [optional] 
**BlockDuration** | Pointer to **int32** |  | [optional] 
**TimeDuration** | **int32** |  | 
**AllowedMethods** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRateLimitRule

`func NewRateLimitRule(urlPattern string, rate int32, timeDuration int32, ) *RateLimitRule`

NewRateLimitRule instantiates a new RateLimitRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitRuleWithDefaults

`func NewRateLimitRuleWithDefaults() *RateLimitRule`

NewRateLimitRuleWithDefaults instantiates a new RateLimitRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActionDetails

`func (o *RateLimitRule) GetActionDetails() ChallengeAction`

GetActionDetails returns the ActionDetails field if non-nil, zero value otherwise.

### GetActionDetailsOk

`func (o *RateLimitRule) GetActionDetailsOk() (*ChallengeAction, bool)`

GetActionDetailsOk returns a tuple with the ActionDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionDetails

`func (o *RateLimitRule) SetActionDetails(v ChallengeAction)`

SetActionDetails sets ActionDetails field to given value.

### HasActionDetails

`func (o *RateLimitRule) HasActionDetails() bool`

HasActionDetails returns a boolean if a field has been set.

### GetId

`func (o *RateLimitRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateLimitRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateLimitRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RateLimitRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAction

`func (o *RateLimitRule) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RateLimitRule) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RateLimitRule) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *RateLimitRule) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetIsEnabled

`func (o *RateLimitRule) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *RateLimitRule) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *RateLimitRule) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *RateLimitRule) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetUrlPattern

`func (o *RateLimitRule) GetUrlPattern() string`

GetUrlPattern returns the UrlPattern field if non-nil, zero value otherwise.

### GetUrlPatternOk

`func (o *RateLimitRule) GetUrlPatternOk() (*string, bool)`

GetUrlPatternOk returns a tuple with the UrlPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlPattern

`func (o *RateLimitRule) SetUrlPattern(v string)`

SetUrlPattern sets UrlPattern field to given value.


### GetDescription

`func (o *RateLimitRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RateLimitRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RateLimitRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RateLimitRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *RateLimitRule) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *RateLimitRule) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExcludeSources

`func (o *RateLimitRule) GetExcludeSources() []string`

GetExcludeSources returns the ExcludeSources field if non-nil, zero value otherwise.

### GetExcludeSourcesOk

`func (o *RateLimitRule) GetExcludeSourcesOk() (*[]string, bool)`

GetExcludeSourcesOk returns a tuple with the ExcludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSources

`func (o *RateLimitRule) SetExcludeSources(v []string)`

SetExcludeSources sets ExcludeSources field to given value.

### HasExcludeSources

`func (o *RateLimitRule) HasExcludeSources() bool`

HasExcludeSources returns a boolean if a field has been set.

### GetRate

`func (o *RateLimitRule) GetRate() int32`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *RateLimitRule) GetRateOk() (*int32, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *RateLimitRule) SetRate(v int32)`

SetRate sets Rate field to given value.


### GetBurst

`func (o *RateLimitRule) GetBurst() int32`

GetBurst returns the Burst field if non-nil, zero value otherwise.

### GetBurstOk

`func (o *RateLimitRule) GetBurstOk() (*int32, bool)`

GetBurstOk returns a tuple with the Burst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBurst

`func (o *RateLimitRule) SetBurst(v int32)`

SetBurst sets Burst field to given value.

### HasBurst

`func (o *RateLimitRule) HasBurst() bool`

HasBurst returns a boolean if a field has been set.

### GetBlockDuration

`func (o *RateLimitRule) GetBlockDuration() int32`

GetBlockDuration returns the BlockDuration field if non-nil, zero value otherwise.

### GetBlockDurationOk

`func (o *RateLimitRule) GetBlockDurationOk() (*int32, bool)`

GetBlockDurationOk returns a tuple with the BlockDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDuration

`func (o *RateLimitRule) SetBlockDuration(v int32)`

SetBlockDuration sets BlockDuration field to given value.

### HasBlockDuration

`func (o *RateLimitRule) HasBlockDuration() bool`

HasBlockDuration returns a boolean if a field has been set.

### GetTimeDuration

`func (o *RateLimitRule) GetTimeDuration() int32`

GetTimeDuration returns the TimeDuration field if non-nil, zero value otherwise.

### GetTimeDurationOk

`func (o *RateLimitRule) GetTimeDurationOk() (*int32, bool)`

GetTimeDurationOk returns a tuple with the TimeDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeDuration

`func (o *RateLimitRule) SetTimeDuration(v int32)`

SetTimeDuration sets TimeDuration field to given value.


### GetAllowedMethods

`func (o *RateLimitRule) GetAllowedMethods() []string`

GetAllowedMethods returns the AllowedMethods field if non-nil, zero value otherwise.

### GetAllowedMethodsOk

`func (o *RateLimitRule) GetAllowedMethodsOk() (*[]string, bool)`

GetAllowedMethodsOk returns a tuple with the AllowedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedMethods

`func (o *RateLimitRule) SetAllowedMethods(v []string)`

SetAllowedMethods sets AllowedMethods field to given value.

### HasAllowedMethods

`func (o *RateLimitRule) HasAllowedMethods() bool`

HasAllowedMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


