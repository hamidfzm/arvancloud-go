# CreateBucketConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationConstraint** | Pointer to [**BucketLocationConstraint**](BucketLocationConstraint.md) |  | [optional] 

## Methods

### NewCreateBucketConfiguration

`func NewCreateBucketConfiguration() *CreateBucketConfiguration`

NewCreateBucketConfiguration instantiates a new CreateBucketConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBucketConfigurationWithDefaults

`func NewCreateBucketConfigurationWithDefaults() *CreateBucketConfiguration`

NewCreateBucketConfigurationWithDefaults instantiates a new CreateBucketConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationConstraint

`func (o *CreateBucketConfiguration) GetLocationConstraint() BucketLocationConstraint`

GetLocationConstraint returns the LocationConstraint field if non-nil, zero value otherwise.

### GetLocationConstraintOk

`func (o *CreateBucketConfiguration) GetLocationConstraintOk() (*BucketLocationConstraint, bool)`

GetLocationConstraintOk returns a tuple with the LocationConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationConstraint

`func (o *CreateBucketConfiguration) SetLocationConstraint(v BucketLocationConstraint)`

SetLocationConstraint sets LocationConstraint field to given value.

### HasLocationConstraint

`func (o *CreateBucketConfiguration) HasLocationConstraint() bool`

HasLocationConstraint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


