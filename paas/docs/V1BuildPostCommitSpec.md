# V1BuildPostCommitSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Args** | Pointer to **[]string** | args is a list of arguments that are provided to either Command, Script or the Docker image&#39;s default entrypoint. The arguments are placed immediately after the command to be run. | [optional] 
**Command** | Pointer to **[]string** | command is the command to run. It may not be specified with Script. This might be needed if the image doesn&#39;t have &#x60;/bin/sh&#x60;, or if you do not want to use a shell. In all other cases, using Script might be more convenient. | [optional] 
**Script** | Pointer to **string** | script is a shell script to be run with &#x60;/bin/sh -ic&#x60;. It may not be specified with Command. Use Script when a shell script is appropriate to execute the post build hook, for example for running unit tests with &#x60;rake test&#x60;. If you need control over the image entrypoint, or if the image does not have &#x60;/bin/sh&#x60;, use Command and/or Args. The &#x60;-i&#x60; flag is needed to support CentOS and RHEL images that use Software Collections (SCL), in order to have the appropriate collections enabled in the shell. E.g., in the Ruby image, this is necessary to make &#x60;ruby&#x60;, &#x60;bundle&#x60; and other binaries available in the PATH. | [optional] 

## Methods

### NewV1BuildPostCommitSpec

`func NewV1BuildPostCommitSpec() *V1BuildPostCommitSpec`

NewV1BuildPostCommitSpec instantiates a new V1BuildPostCommitSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1BuildPostCommitSpecWithDefaults

`func NewV1BuildPostCommitSpecWithDefaults() *V1BuildPostCommitSpec`

NewV1BuildPostCommitSpecWithDefaults instantiates a new V1BuildPostCommitSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgs

`func (o *V1BuildPostCommitSpec) GetArgs() []string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *V1BuildPostCommitSpec) GetArgsOk() (*[]string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *V1BuildPostCommitSpec) SetArgs(v []string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *V1BuildPostCommitSpec) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetCommand

`func (o *V1BuildPostCommitSpec) GetCommand() []string`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *V1BuildPostCommitSpec) GetCommandOk() (*[]string, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *V1BuildPostCommitSpec) SetCommand(v []string)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *V1BuildPostCommitSpec) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetScript

`func (o *V1BuildPostCommitSpec) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *V1BuildPostCommitSpec) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *V1BuildPostCommitSpec) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *V1BuildPostCommitSpec) HasScript() bool`

HasScript returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


