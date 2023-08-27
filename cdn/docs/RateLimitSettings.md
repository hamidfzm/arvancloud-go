# RateLimitSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DdosDetection** | Pointer to **bool** |  | [optional] 
**ExcludeSources** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRateLimitSettings

`func NewRateLimitSettings() *RateLimitSettings`

NewRateLimitSettings instantiates a new RateLimitSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitSettingsWithDefaults

`func NewRateLimitSettingsWithDefaults() *RateLimitSettings`

NewRateLimitSettingsWithDefaults instantiates a new RateLimitSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDdosDetection

`func (o *RateLimitSettings) GetDdosDetection() bool`

GetDdosDetection returns the DdosDetection field if non-nil, zero value otherwise.

### GetDdosDetectionOk

`func (o *RateLimitSettings) GetDdosDetectionOk() (*bool, bool)`

GetDdosDetectionOk returns a tuple with the DdosDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDdosDetection

`func (o *RateLimitSettings) SetDdosDetection(v bool)`

SetDdosDetection sets DdosDetection field to given value.

### HasDdosDetection

`func (o *RateLimitSettings) HasDdosDetection() bool`

HasDdosDetection returns a boolean if a field has been set.

### GetExcludeSources

`func (o *RateLimitSettings) GetExcludeSources() []string`

GetExcludeSources returns the ExcludeSources field if non-nil, zero value otherwise.

### GetExcludeSourcesOk

`func (o *RateLimitSettings) GetExcludeSourcesOk() (*[]string, bool)`

GetExcludeSourcesOk returns a tuple with the ExcludeSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeSources

`func (o *RateLimitSettings) SetExcludeSources(v []string)`

SetExcludeSources sets ExcludeSources field to given value.

### HasExcludeSources

`func (o *RateLimitSettings) HasExcludeSources() bool`

HasExcludeSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


