# CertificateOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**OrderId** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to **string** | - &#x60;unprocessed&#x60; - Order is in the process queue - &#x60;canceled&#x60; - Order is canceled in favor of a new one with updated subject names - &#x60;pending&#x60; - Authorization Challenges are set, Validating authorization challenges... - &#x60;ready&#x60; - Challenges are validated, ready to issue the certificate - &#x60;processing&#x60; - Issuing Certificate... - &#x60;valid&#x60; - Certificate is issued successfully, this is the final stage - &#x60;invalid&#x60; - An Error Occurred, this order cannot proceed anymore, a new order will be created automatically - &#x60;terminated&#x60; - An Unknown Error occurred, this order cannot proceed anymore, a new order will be created automatically - &#x60;killed&#x60; - Order failed despite many retries, will not proceed anymore nor retry, needs manual intervention  | [optional] [readonly] 
**DomainNames** | Pointer to **[]string** |  | [optional] [readonly] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**ExpiryDate** | Pointer to **time.Time** | Expired order is treated as invalid order | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewCertificateOrder

`func NewCertificateOrder() *CertificateOrder`

NewCertificateOrder instantiates a new CertificateOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateOrderWithDefaults

`func NewCertificateOrderWithDefaults() *CertificateOrder`

NewCertificateOrderWithDefaults instantiates a new CertificateOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrderId

`func (o *CertificateOrder) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CertificateOrder) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CertificateOrder) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CertificateOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetStatus

`func (o *CertificateOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CertificateOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDomainNames

`func (o *CertificateOrder) GetDomainNames() []string`

GetDomainNames returns the DomainNames field if non-nil, zero value otherwise.

### GetDomainNamesOk

`func (o *CertificateOrder) GetDomainNamesOk() (*[]string, bool)`

GetDomainNamesOk returns a tuple with the DomainNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainNames

`func (o *CertificateOrder) SetDomainNames(v []string)`

SetDomainNames sets DomainNames field to given value.

### HasDomainNames

`func (o *CertificateOrder) HasDomainNames() bool`

HasDomainNames returns a boolean if a field has been set.

### GetErrors

`func (o *CertificateOrder) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *CertificateOrder) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *CertificateOrder) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *CertificateOrder) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetExpiryDate

`func (o *CertificateOrder) GetExpiryDate() time.Time`

GetExpiryDate returns the ExpiryDate field if non-nil, zero value otherwise.

### GetExpiryDateOk

`func (o *CertificateOrder) GetExpiryDateOk() (*time.Time, bool)`

GetExpiryDateOk returns a tuple with the ExpiryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryDate

`func (o *CertificateOrder) SetExpiryDate(v time.Time)`

SetExpiryDate sets ExpiryDate field to given value.

### HasExpiryDate

`func (o *CertificateOrder) HasExpiryDate() bool`

HasExpiryDate returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CertificateOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CertificateOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CertificateOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CertificateOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CertificateOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CertificateOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CertificateOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CertificateOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


