# V1Parameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of a parameter. Optional. | [optional] 
**DisplayName** | Pointer to **string** | Optional: The name that will show in UI instead of parameter &#39;Name&#39; | [optional] 
**From** | Pointer to **string** | From is an input value for the generator. Optional. | [optional] 
**Generate** | Pointer to **string** | generate specifies the generator to be used to generate random string from an input value specified by From field. The result string is stored into Value field. If empty, no generator is being used, leaving the result Value untouched. Optional.  The only supported generator is \&quot;expression\&quot;, which accepts a \&quot;from\&quot; value in the form of a simple regular expression containing the range expression \&quot;[a-zA-Z0-9]\&quot;, and the length expression \&quot;a{length}\&quot;.  Examples:  from             | value | [optional] 
**Name** | **string** | Name must be set and it can be referenced in Template Items using ${PARAMETER_NAME}. Required. | 
**Required** | Pointer to **bool** | Optional: Indicates the parameter must have a value.  Defaults to false. | [optional] 
**Value** | Pointer to **string** | Value holds the Parameter data. If specified, the generator will be ignored. The value replaces all occurrences of the Parameter ${Name} expression during the Template to Config transformation. Optional. | [optional] 

## Methods

### NewV1Parameter

`func NewV1Parameter(name string, ) *V1Parameter`

NewV1Parameter instantiates a new V1Parameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ParameterWithDefaults

`func NewV1ParameterWithDefaults() *V1Parameter`

NewV1ParameterWithDefaults instantiates a new V1Parameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1Parameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1Parameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1Parameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1Parameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *V1Parameter) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *V1Parameter) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *V1Parameter) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *V1Parameter) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFrom

`func (o *V1Parameter) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *V1Parameter) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *V1Parameter) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *V1Parameter) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetGenerate

`func (o *V1Parameter) GetGenerate() string`

GetGenerate returns the Generate field if non-nil, zero value otherwise.

### GetGenerateOk

`func (o *V1Parameter) GetGenerateOk() (*string, bool)`

GetGenerateOk returns a tuple with the Generate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerate

`func (o *V1Parameter) SetGenerate(v string)`

SetGenerate sets Generate field to given value.

### HasGenerate

`func (o *V1Parameter) HasGenerate() bool`

HasGenerate returns a boolean if a field has been set.

### GetName

`func (o *V1Parameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1Parameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1Parameter) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *V1Parameter) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *V1Parameter) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *V1Parameter) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *V1Parameter) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetValue

`func (o *V1Parameter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1Parameter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1Parameter) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *V1Parameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


