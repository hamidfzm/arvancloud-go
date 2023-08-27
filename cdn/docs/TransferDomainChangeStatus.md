# TransferDomainChangeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** |  | 
**Status** | **string** |  | 
**PreserveState** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewTransferDomainChangeStatus

`func NewTransferDomainChangeStatus(domain string, status string, ) *TransferDomainChangeStatus`

NewTransferDomainChangeStatus instantiates a new TransferDomainChangeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferDomainChangeStatusWithDefaults

`func NewTransferDomainChangeStatusWithDefaults() *TransferDomainChangeStatus`

NewTransferDomainChangeStatusWithDefaults instantiates a new TransferDomainChangeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *TransferDomainChangeStatus) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TransferDomainChangeStatus) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TransferDomainChangeStatus) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetStatus

`func (o *TransferDomainChangeStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferDomainChangeStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferDomainChangeStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPreserveState

`func (o *TransferDomainChangeStatus) GetPreserveState() bool`

GetPreserveState returns the PreserveState field if non-nil, zero value otherwise.

### GetPreserveStateOk

`func (o *TransferDomainChangeStatus) GetPreserveStateOk() (*bool, bool)`

GetPreserveStateOk returns a tuple with the PreserveState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveState

`func (o *TransferDomainChangeStatus) SetPreserveState(v bool)`

SetPreserveState sets PreserveState field to given value.

### HasPreserveState

`func (o *TransferDomainChangeStatus) HasPreserveState() bool`

HasPreserveState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


