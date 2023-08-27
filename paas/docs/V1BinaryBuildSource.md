# V1BinaryBuildSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsFile** | Pointer to **string** | asFile indicates that the provided binary input should be considered a single file within the build input. For example, specifying \&quot;webapp.war\&quot; would place the provided binary as &#x60;/webapp.war&#x60; for the builder. If left empty, the Docker and Source build strategies assume this file is a zip, tar, or tar.gz file and extract it as the source. The custom strategy receives this binary as standard input. This filename may not contain slashes or be &#39;..&#39; or &#39;.&#39;. | [optional] 

## Methods

### NewV1BinaryBuildSource

`func NewV1BinaryBuildSource() *V1BinaryBuildSource`

NewV1BinaryBuildSource instantiates a new V1BinaryBuildSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BinaryBuildSourceWithDefaults

`func NewV1BinaryBuildSourceWithDefaults() *V1BinaryBuildSource`

NewV1BinaryBuildSourceWithDefaults instantiates a new V1BinaryBuildSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsFile

`func (o *V1BinaryBuildSource) GetAsFile() string`

GetAsFile returns the AsFile field if non-nil, zero value otherwise.

### GetAsFileOk

`func (o *V1BinaryBuildSource) GetAsFileOk() (*string, bool)`

GetAsFileOk returns a tuple with the AsFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsFile

`func (o *V1BinaryBuildSource) SetAsFile(v string)`

SetAsFile sets AsFile field to given value.

### HasAsFile

`func (o *V1BinaryBuildSource) HasAsFile() bool`

HasAsFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


