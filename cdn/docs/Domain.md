# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**UserId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** | Deprecated in favor of name attribute | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PlanLevel** | Pointer to **int32** | - &#x60;0&#x60; - Traffic - &#x60;1&#x60; - Basic - &#x60;2&#x60; - Growth - &#x60;3&#x60; - Professional - &#x60;4&#x60; - Enterprise  | [optional] 
**NsKeys** | Pointer to **[]string** | Desired NS values for the domain | [optional] 
**CurrentNs** | Pointer to **[]string** | Current NS values for the domain | [optional] 
**TargetCname** | Pointer to **NullableString** | Current record for CNAME Setup of the domain | [optional] 
**CustomCname** | Pointer to **NullableString** | Domain&#39;s custom record for CNAME Setup | [optional] 
**Type** | Pointer to **string** | Partial domain is using CNAME Setup and full domain is using NS Setup | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DnsCloud** | Pointer to **bool** |  | [optional] [default to false]
**Restriction** | Pointer to **[]string** |  | [optional] 
**Transfer** | Pointer to [**DomainTransferData**](DomainTransferData.md) |  | [optional] 
**FingerprintStatus** | Pointer to **bool** |  | [optional] [readonly] [default to false]
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Domain) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Domain) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Domain) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Domain) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetDomain

`func (o *Domain) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Domain) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlanLevel

`func (o *Domain) GetPlanLevel() int32`

GetPlanLevel returns the PlanLevel field if non-nil, zero value otherwise.

### GetPlanLevelOk

`func (o *Domain) GetPlanLevelOk() (*int32, bool)`

GetPlanLevelOk returns a tuple with the PlanLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanLevel

`func (o *Domain) SetPlanLevel(v int32)`

SetPlanLevel sets PlanLevel field to given value.

### HasPlanLevel

`func (o *Domain) HasPlanLevel() bool`

HasPlanLevel returns a boolean if a field has been set.

### GetNsKeys

`func (o *Domain) GetNsKeys() []string`

GetNsKeys returns the NsKeys field if non-nil, zero value otherwise.

### GetNsKeysOk

`func (o *Domain) GetNsKeysOk() (*[]string, bool)`

GetNsKeysOk returns a tuple with the NsKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsKeys

`func (o *Domain) SetNsKeys(v []string)`

SetNsKeys sets NsKeys field to given value.

### HasNsKeys

`func (o *Domain) HasNsKeys() bool`

HasNsKeys returns a boolean if a field has been set.

### GetCurrentNs

`func (o *Domain) GetCurrentNs() []string`

GetCurrentNs returns the CurrentNs field if non-nil, zero value otherwise.

### GetCurrentNsOk

`func (o *Domain) GetCurrentNsOk() (*[]string, bool)`

GetCurrentNsOk returns a tuple with the CurrentNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNs

`func (o *Domain) SetCurrentNs(v []string)`

SetCurrentNs sets CurrentNs field to given value.

### HasCurrentNs

`func (o *Domain) HasCurrentNs() bool`

HasCurrentNs returns a boolean if a field has been set.

### GetTargetCname

`func (o *Domain) GetTargetCname() string`

GetTargetCname returns the TargetCname field if non-nil, zero value otherwise.

### GetTargetCnameOk

`func (o *Domain) GetTargetCnameOk() (*string, bool)`

GetTargetCnameOk returns a tuple with the TargetCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetCname

`func (o *Domain) SetTargetCname(v string)`

SetTargetCname sets TargetCname field to given value.

### HasTargetCname

`func (o *Domain) HasTargetCname() bool`

HasTargetCname returns a boolean if a field has been set.

### SetTargetCnameNil

`func (o *Domain) SetTargetCnameNil(b bool)`

 SetTargetCnameNil sets the value for TargetCname to be an explicit nil

### UnsetTargetCname
`func (o *Domain) UnsetTargetCname()`

UnsetTargetCname ensures that no value is present for TargetCname, not even an explicit nil
### GetCustomCname

`func (o *Domain) GetCustomCname() string`

GetCustomCname returns the CustomCname field if non-nil, zero value otherwise.

### GetCustomCnameOk

`func (o *Domain) GetCustomCnameOk() (*string, bool)`

GetCustomCnameOk returns a tuple with the CustomCname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCname

`func (o *Domain) SetCustomCname(v string)`

SetCustomCname sets CustomCname field to given value.

### HasCustomCname

`func (o *Domain) HasCustomCname() bool`

HasCustomCname returns a boolean if a field has been set.

### SetCustomCnameNil

`func (o *Domain) SetCustomCnameNil(b bool)`

 SetCustomCnameNil sets the value for CustomCname to be an explicit nil

### UnsetCustomCname
`func (o *Domain) UnsetCustomCname()`

UnsetCustomCname ensures that no value is present for CustomCname, not even an explicit nil
### GetType

`func (o *Domain) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Domain) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Domain) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Domain) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Domain) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Domain) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Domain) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Domain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDnsCloud

`func (o *Domain) GetDnsCloud() bool`

GetDnsCloud returns the DnsCloud field if non-nil, zero value otherwise.

### GetDnsCloudOk

`func (o *Domain) GetDnsCloudOk() (*bool, bool)`

GetDnsCloudOk returns a tuple with the DnsCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsCloud

`func (o *Domain) SetDnsCloud(v bool)`

SetDnsCloud sets DnsCloud field to given value.

### HasDnsCloud

`func (o *Domain) HasDnsCloud() bool`

HasDnsCloud returns a boolean if a field has been set.

### GetRestriction

`func (o *Domain) GetRestriction() []string`

GetRestriction returns the Restriction field if non-nil, zero value otherwise.

### GetRestrictionOk

`func (o *Domain) GetRestrictionOk() (*[]string, bool)`

GetRestrictionOk returns a tuple with the Restriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestriction

`func (o *Domain) SetRestriction(v []string)`

SetRestriction sets Restriction field to given value.

### HasRestriction

`func (o *Domain) HasRestriction() bool`

HasRestriction returns a boolean if a field has been set.

### GetTransfer

`func (o *Domain) GetTransfer() DomainTransferData`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *Domain) GetTransferOk() (*DomainTransferData, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *Domain) SetTransfer(v DomainTransferData)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *Domain) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetFingerprintStatus

`func (o *Domain) GetFingerprintStatus() bool`

GetFingerprintStatus returns the FingerprintStatus field if non-nil, zero value otherwise.

### GetFingerprintStatusOk

`func (o *Domain) GetFingerprintStatusOk() (*bool, bool)`

GetFingerprintStatusOk returns a tuple with the FingerprintStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprintStatus

`func (o *Domain) SetFingerprintStatus(v bool)`

SetFingerprintStatus sets FingerprintStatus field to given value.

### HasFingerprintStatus

`func (o *Domain) HasFingerprintStatus() bool`

HasFingerprintStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Domain) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Domain) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Domain) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Domain) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Domain) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Domain) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Domain) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Domain) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


