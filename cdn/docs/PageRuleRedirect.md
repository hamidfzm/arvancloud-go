# PageRuleRedirect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] [default to false]
**StatusCode** | Pointer to **int32** |  | [optional] [default to 301]
**Url** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPageRuleRedirect

`func NewPageRuleRedirect() *PageRuleRedirect`

NewPageRuleRedirect instantiates a new PageRuleRedirect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageRuleRedirectWithDefaults

`func NewPageRuleRedirectWithDefaults() *PageRuleRedirect`

NewPageRuleRedirectWithDefaults instantiates a new PageRuleRedirect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *PageRuleRedirect) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *PageRuleRedirect) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *PageRuleRedirect) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *PageRuleRedirect) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetStatusCode

`func (o *PageRuleRedirect) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *PageRuleRedirect) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *PageRuleRedirect) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *PageRuleRedirect) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetUrl

`func (o *PageRuleRedirect) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PageRuleRedirect) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PageRuleRedirect) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PageRuleRedirect) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *PageRuleRedirect) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *PageRuleRedirect) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


