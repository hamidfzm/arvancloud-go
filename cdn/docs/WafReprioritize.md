# WafReprioritize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageId** | **string** | ID of the package you want to move | 
**AfterPackageId** | Pointer to **string** | ID of the package you want to be prior to the selected package | [optional] 
**BeforePackageId** | Pointer to **string** | ID of the package you want to follow the selected package | [optional] 

## Methods

### NewWafReprioritize

`func NewWafReprioritize(packageId string, ) *WafReprioritize`

NewWafReprioritize instantiates a new WafReprioritize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWafReprioritizeWithDefaults

`func NewWafReprioritizeWithDefaults() *WafReprioritize`

NewWafReprioritizeWithDefaults instantiates a new WafReprioritize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageId

`func (o *WafReprioritize) GetPackageId() string`

GetPackageId returns the PackageId field if non-nil, zero value otherwise.

### GetPackageIdOk

`func (o *WafReprioritize) GetPackageIdOk() (*string, bool)`

GetPackageIdOk returns a tuple with the PackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageId

`func (o *WafReprioritize) SetPackageId(v string)`

SetPackageId sets PackageId field to given value.


### GetAfterPackageId

`func (o *WafReprioritize) GetAfterPackageId() string`

GetAfterPackageId returns the AfterPackageId field if non-nil, zero value otherwise.

### GetAfterPackageIdOk

`func (o *WafReprioritize) GetAfterPackageIdOk() (*string, bool)`

GetAfterPackageIdOk returns a tuple with the AfterPackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterPackageId

`func (o *WafReprioritize) SetAfterPackageId(v string)`

SetAfterPackageId sets AfterPackageId field to given value.

### HasAfterPackageId

`func (o *WafReprioritize) HasAfterPackageId() bool`

HasAfterPackageId returns a boolean if a field has been set.

### GetBeforePackageId

`func (o *WafReprioritize) GetBeforePackageId() string`

GetBeforePackageId returns the BeforePackageId field if non-nil, zero value otherwise.

### GetBeforePackageIdOk

`func (o *WafReprioritize) GetBeforePackageIdOk() (*string, bool)`

GetBeforePackageIdOk returns a tuple with the BeforePackageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforePackageId

`func (o *WafReprioritize) SetBeforePackageId(v string)`

SetBeforePackageId sets BeforePackageId field to given value.

### HasBeforePackageId

`func (o *WafReprioritize) HasBeforePackageId() bool`

HasBeforePackageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


