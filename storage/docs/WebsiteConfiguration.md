# WebsiteConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorDocument** | Pointer to [**PutBucketWebsiteRequestWebsiteConfigurationErrorDocument**](PutBucketWebsiteRequestWebsiteConfigurationErrorDocument.md) |  | [optional] 
**IndexDocument** | Pointer to [**PutBucketWebsiteRequestWebsiteConfigurationIndexDocument**](PutBucketWebsiteRequestWebsiteConfigurationIndexDocument.md) |  | [optional] 
**RedirectAllRequestsTo** | Pointer to [**PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo**](PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo.md) |  | [optional] 
**RoutingRules** | Pointer to [**Array**](array.md) |  | [optional] 

## Methods

### NewWebsiteConfiguration

`func NewWebsiteConfiguration() *WebsiteConfiguration`

NewWebsiteConfiguration instantiates a new WebsiteConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebsiteConfigurationWithDefaults

`func NewWebsiteConfigurationWithDefaults() *WebsiteConfiguration`

NewWebsiteConfigurationWithDefaults instantiates a new WebsiteConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorDocument

`func (o *WebsiteConfiguration) GetErrorDocument() PutBucketWebsiteRequestWebsiteConfigurationErrorDocument`

GetErrorDocument returns the ErrorDocument field if non-nil, zero value otherwise.

### GetErrorDocumentOk

`func (o *WebsiteConfiguration) GetErrorDocumentOk() (*PutBucketWebsiteRequestWebsiteConfigurationErrorDocument, bool)`

GetErrorDocumentOk returns a tuple with the ErrorDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDocument

`func (o *WebsiteConfiguration) SetErrorDocument(v PutBucketWebsiteRequestWebsiteConfigurationErrorDocument)`

SetErrorDocument sets ErrorDocument field to given value.

### HasErrorDocument

`func (o *WebsiteConfiguration) HasErrorDocument() bool`

HasErrorDocument returns a boolean if a field has been set.

### GetIndexDocument

`func (o *WebsiteConfiguration) GetIndexDocument() PutBucketWebsiteRequestWebsiteConfigurationIndexDocument`

GetIndexDocument returns the IndexDocument field if non-nil, zero value otherwise.

### GetIndexDocumentOk

`func (o *WebsiteConfiguration) GetIndexDocumentOk() (*PutBucketWebsiteRequestWebsiteConfigurationIndexDocument, bool)`

GetIndexDocumentOk returns a tuple with the IndexDocument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexDocument

`func (o *WebsiteConfiguration) SetIndexDocument(v PutBucketWebsiteRequestWebsiteConfigurationIndexDocument)`

SetIndexDocument sets IndexDocument field to given value.

### HasIndexDocument

`func (o *WebsiteConfiguration) HasIndexDocument() bool`

HasIndexDocument returns a boolean if a field has been set.

### GetRedirectAllRequestsTo

`func (o *WebsiteConfiguration) GetRedirectAllRequestsTo() PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo`

GetRedirectAllRequestsTo returns the RedirectAllRequestsTo field if non-nil, zero value otherwise.

### GetRedirectAllRequestsToOk

`func (o *WebsiteConfiguration) GetRedirectAllRequestsToOk() (*PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo, bool)`

GetRedirectAllRequestsToOk returns a tuple with the RedirectAllRequestsTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectAllRequestsTo

`func (o *WebsiteConfiguration) SetRedirectAllRequestsTo(v PutBucketWebsiteRequestWebsiteConfigurationRedirectAllRequestsTo)`

SetRedirectAllRequestsTo sets RedirectAllRequestsTo field to given value.

### HasRedirectAllRequestsTo

`func (o *WebsiteConfiguration) HasRedirectAllRequestsTo() bool`

HasRedirectAllRequestsTo returns a boolean if a field has been set.

### GetRoutingRules

`func (o *WebsiteConfiguration) GetRoutingRules() Array`

GetRoutingRules returns the RoutingRules field if non-nil, zero value otherwise.

### GetRoutingRulesOk

`func (o *WebsiteConfiguration) GetRoutingRulesOk() (*Array, bool)`

GetRoutingRulesOk returns a tuple with the RoutingRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingRules

`func (o *WebsiteConfiguration) SetRoutingRules(v Array)`

SetRoutingRules sets RoutingRules field to given value.

### HasRoutingRules

`func (o *WebsiteConfiguration) HasRoutingRules() bool`

HasRoutingRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


