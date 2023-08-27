# V1SourceStrategyOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incremental** | Pointer to **bool** | incremental overrides the source-strategy incremental option in the build config | [optional] 

## Methods

### NewV1SourceStrategyOptions

`func NewV1SourceStrategyOptions() *V1SourceStrategyOptions`

NewV1SourceStrategyOptions instantiates a new V1SourceStrategyOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SourceStrategyOptionsWithDefaults

`func NewV1SourceStrategyOptionsWithDefaults() *V1SourceStrategyOptions`

NewV1SourceStrategyOptionsWithDefaults instantiates a new V1SourceStrategyOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncremental

`func (o *V1SourceStrategyOptions) GetIncremental() bool`

GetIncremental returns the Incremental field if non-nil, zero value otherwise.

### GetIncrementalOk

`func (o *V1SourceStrategyOptions) GetIncrementalOk() (*bool, bool)`

GetIncrementalOk returns a tuple with the Incremental field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncremental

`func (o *V1SourceStrategyOptions) SetIncremental(v bool)`

SetIncremental sets Incremental field to given value.

### HasIncremental

`func (o *V1SourceStrategyOptions) HasIncremental() bool`

HasIncremental returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


