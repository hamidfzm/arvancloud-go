# V1SecretLocalReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name is the name of the resource in the same namespace being referenced | 

## Methods

### NewV1SecretLocalReference

`func NewV1SecretLocalReference(name string, ) *V1SecretLocalReference`

NewV1SecretLocalReference instantiates a new V1SecretLocalReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SecretLocalReferenceWithDefaults

`func NewV1SecretLocalReferenceWithDefaults() *V1SecretLocalReference`

NewV1SecretLocalReferenceWithDefaults instantiates a new V1SecretLocalReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1SecretLocalReference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1SecretLocalReference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1SecretLocalReference) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


