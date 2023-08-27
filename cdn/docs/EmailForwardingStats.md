# EmailForwardingStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DnsActivation** | Pointer to **bool** |  | [optional] 
**RecipientsCount** | Pointer to **int32** |  | [optional] 
**AliasesCount** | Pointer to **int32** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**IsConfigured** | Pointer to **bool** |  | [optional] 
**EmailsForwarded** | Pointer to **int32** |  | [optional] 
**EmailsBlocked** | Pointer to **int32** |  | [optional] 
**EmailsReplied** | Pointer to **int32** |  | [optional] 

## Methods

### NewEmailForwardingStats

`func NewEmailForwardingStats() *EmailForwardingStats`

NewEmailForwardingStats instantiates a new EmailForwardingStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailForwardingStatsWithDefaults

`func NewEmailForwardingStatsWithDefaults() *EmailForwardingStats`

NewEmailForwardingStatsWithDefaults instantiates a new EmailForwardingStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EmailForwardingStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EmailForwardingStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EmailForwardingStats) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EmailForwardingStats) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDnsActivation

`func (o *EmailForwardingStats) GetDnsActivation() bool`

GetDnsActivation returns the DnsActivation field if non-nil, zero value otherwise.

### GetDnsActivationOk

`func (o *EmailForwardingStats) GetDnsActivationOk() (*bool, bool)`

GetDnsActivationOk returns a tuple with the DnsActivation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsActivation

`func (o *EmailForwardingStats) SetDnsActivation(v bool)`

SetDnsActivation sets DnsActivation field to given value.

### HasDnsActivation

`func (o *EmailForwardingStats) HasDnsActivation() bool`

HasDnsActivation returns a boolean if a field has been set.

### GetRecipientsCount

`func (o *EmailForwardingStats) GetRecipientsCount() int32`

GetRecipientsCount returns the RecipientsCount field if non-nil, zero value otherwise.

### GetRecipientsCountOk

`func (o *EmailForwardingStats) GetRecipientsCountOk() (*int32, bool)`

GetRecipientsCountOk returns a tuple with the RecipientsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientsCount

`func (o *EmailForwardingStats) SetRecipientsCount(v int32)`

SetRecipientsCount sets RecipientsCount field to given value.

### HasRecipientsCount

`func (o *EmailForwardingStats) HasRecipientsCount() bool`

HasRecipientsCount returns a boolean if a field has been set.

### GetAliasesCount

`func (o *EmailForwardingStats) GetAliasesCount() int32`

GetAliasesCount returns the AliasesCount field if non-nil, zero value otherwise.

### GetAliasesCountOk

`func (o *EmailForwardingStats) GetAliasesCountOk() (*int32, bool)`

GetAliasesCountOk returns a tuple with the AliasesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasesCount

`func (o *EmailForwardingStats) SetAliasesCount(v int32)`

SetAliasesCount sets AliasesCount field to given value.

### HasAliasesCount

`func (o *EmailForwardingStats) HasAliasesCount() bool`

HasAliasesCount returns a boolean if a field has been set.

### GetIsActive

`func (o *EmailForwardingStats) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *EmailForwardingStats) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *EmailForwardingStats) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *EmailForwardingStats) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsConfigured

`func (o *EmailForwardingStats) GetIsConfigured() bool`

GetIsConfigured returns the IsConfigured field if non-nil, zero value otherwise.

### GetIsConfiguredOk

`func (o *EmailForwardingStats) GetIsConfiguredOk() (*bool, bool)`

GetIsConfiguredOk returns a tuple with the IsConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConfigured

`func (o *EmailForwardingStats) SetIsConfigured(v bool)`

SetIsConfigured sets IsConfigured field to given value.

### HasIsConfigured

`func (o *EmailForwardingStats) HasIsConfigured() bool`

HasIsConfigured returns a boolean if a field has been set.

### GetEmailsForwarded

`func (o *EmailForwardingStats) GetEmailsForwarded() int32`

GetEmailsForwarded returns the EmailsForwarded field if non-nil, zero value otherwise.

### GetEmailsForwardedOk

`func (o *EmailForwardingStats) GetEmailsForwardedOk() (*int32, bool)`

GetEmailsForwardedOk returns a tuple with the EmailsForwarded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsForwarded

`func (o *EmailForwardingStats) SetEmailsForwarded(v int32)`

SetEmailsForwarded sets EmailsForwarded field to given value.

### HasEmailsForwarded

`func (o *EmailForwardingStats) HasEmailsForwarded() bool`

HasEmailsForwarded returns a boolean if a field has been set.

### GetEmailsBlocked

`func (o *EmailForwardingStats) GetEmailsBlocked() int32`

GetEmailsBlocked returns the EmailsBlocked field if non-nil, zero value otherwise.

### GetEmailsBlockedOk

`func (o *EmailForwardingStats) GetEmailsBlockedOk() (*int32, bool)`

GetEmailsBlockedOk returns a tuple with the EmailsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsBlocked

`func (o *EmailForwardingStats) SetEmailsBlocked(v int32)`

SetEmailsBlocked sets EmailsBlocked field to given value.

### HasEmailsBlocked

`func (o *EmailForwardingStats) HasEmailsBlocked() bool`

HasEmailsBlocked returns a boolean if a field has been set.

### GetEmailsReplied

`func (o *EmailForwardingStats) GetEmailsReplied() int32`

GetEmailsReplied returns the EmailsReplied field if non-nil, zero value otherwise.

### GetEmailsRepliedOk

`func (o *EmailForwardingStats) GetEmailsRepliedOk() (*int32, bool)`

GetEmailsRepliedOk returns a tuple with the EmailsReplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailsReplied

`func (o *EmailForwardingStats) SetEmailsReplied(v int32)`

SetEmailsReplied sets EmailsReplied field to given value.

### HasEmailsReplied

`func (o *EmailForwardingStats) HasEmailsReplied() bool`

HasEmailsReplied returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


