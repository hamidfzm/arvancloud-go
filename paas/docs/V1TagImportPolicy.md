# V1TagImportPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Insecure** | Pointer to **bool** | Insecure is true if the server may bypass certificate verification or connect directly over HTTP during image import. | [optional] 
**Scheduled** | Pointer to **bool** | Scheduled indicates to the server that this tag should be periodically checked to ensure it is up to date, and imported | [optional] 

## Methods

### NewV1TagImportPolicy

`func NewV1TagImportPolicy() *V1TagImportPolicy`

NewV1TagImportPolicy instantiates a new V1TagImportPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TagImportPolicyWithDefaults

`func NewV1TagImportPolicyWithDefaults() *V1TagImportPolicy`

NewV1TagImportPolicyWithDefaults instantiates a new V1TagImportPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInsecure

`func (o *V1TagImportPolicy) GetInsecure() bool`

GetInsecure returns the Insecure field if non-nil, zero value otherwise.

### GetInsecureOk

`func (o *V1TagImportPolicy) GetInsecureOk() (*bool, bool)`

GetInsecureOk returns a tuple with the Insecure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsecure

`func (o *V1TagImportPolicy) SetInsecure(v bool)`

SetInsecure sets Insecure field to given value.

### HasInsecure

`func (o *V1TagImportPolicy) HasInsecure() bool`

HasInsecure returns a boolean if a field has been set.

### GetScheduled

`func (o *V1TagImportPolicy) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *V1TagImportPolicy) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *V1TagImportPolicy) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *V1TagImportPolicy) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


