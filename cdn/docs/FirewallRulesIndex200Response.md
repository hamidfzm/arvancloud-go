# FirewallRulesIndex200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]FirewallRuleView**](FirewallRuleView.md) |  | [optional] 
**Links** | Pointer to [**PaginatedResponseLinks**](PaginatedResponseLinks.md) |  | [optional] 
**Meta** | Pointer to [**PaginatedResponseMeta**](PaginatedResponseMeta.md) |  | [optional] 

## Methods

### NewFirewallRulesIndex200Response

`func NewFirewallRulesIndex200Response() *FirewallRulesIndex200Response`

NewFirewallRulesIndex200Response instantiates a new FirewallRulesIndex200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallRulesIndex200ResponseWithDefaults

`func NewFirewallRulesIndex200ResponseWithDefaults() *FirewallRulesIndex200Response`

NewFirewallRulesIndex200ResponseWithDefaults instantiates a new FirewallRulesIndex200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FirewallRulesIndex200Response) GetData() []FirewallRuleView`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FirewallRulesIndex200Response) GetDataOk() (*[]FirewallRuleView, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FirewallRulesIndex200Response) SetData(v []FirewallRuleView)`

SetData sets Data field to given value.

### HasData

`func (o *FirewallRulesIndex200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *FirewallRulesIndex200Response) GetLinks() PaginatedResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *FirewallRulesIndex200Response) GetLinksOk() (*PaginatedResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *FirewallRulesIndex200Response) SetLinks(v PaginatedResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *FirewallRulesIndex200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMeta

`func (o *FirewallRulesIndex200Response) GetMeta() PaginatedResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *FirewallRulesIndex200Response) GetMetaOk() (*PaginatedResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *FirewallRulesIndex200Response) SetMeta(v PaginatedResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *FirewallRulesIndex200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


