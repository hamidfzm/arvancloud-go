# OutputLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3** | Pointer to [**OutputLocationS3**](OutputLocationS3.md) |  | [optional] 

## Methods

### NewOutputLocation

`func NewOutputLocation() *OutputLocation`

NewOutputLocation instantiates a new OutputLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutputLocationWithDefaults

`func NewOutputLocationWithDefaults() *OutputLocation`

NewOutputLocationWithDefaults instantiates a new OutputLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3

`func (o *OutputLocation) GetS3() OutputLocationS3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *OutputLocation) GetS3Ok() (*OutputLocationS3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *OutputLocation) SetS3(v OutputLocationS3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *OutputLocation) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


