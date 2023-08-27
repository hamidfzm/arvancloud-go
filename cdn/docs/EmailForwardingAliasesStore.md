# EmailForwardingAliasesStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocalPart** | **string** |  | 
**Recipients** | **[]string** |  | 

## Methods

### NewEmailForwardingAliasesStore

`func NewEmailForwardingAliasesStore(localPart string, recipients []string, ) *EmailForwardingAliasesStore`

NewEmailForwardingAliasesStore instantiates a new EmailForwardingAliasesStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailForwardingAliasesStoreWithDefaults

`func NewEmailForwardingAliasesStoreWithDefaults() *EmailForwardingAliasesStore`

NewEmailForwardingAliasesStoreWithDefaults instantiates a new EmailForwardingAliasesStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocalPart

`func (o *EmailForwardingAliasesStore) GetLocalPart() string`

GetLocalPart returns the LocalPart field if non-nil, zero value otherwise.

### GetLocalPartOk

`func (o *EmailForwardingAliasesStore) GetLocalPartOk() (*string, bool)`

GetLocalPartOk returns a tuple with the LocalPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalPart

`func (o *EmailForwardingAliasesStore) SetLocalPart(v string)`

SetLocalPart sets LocalPart field to given value.


### GetRecipients

`func (o *EmailForwardingAliasesStore) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *EmailForwardingAliasesStore) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *EmailForwardingAliasesStore) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


