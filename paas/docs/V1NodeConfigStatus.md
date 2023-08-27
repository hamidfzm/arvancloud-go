# V1NodeConfigStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**V1NodeConfigSource**](V1NodeConfigSource.md) |  | [optional] 
**Assigned** | Pointer to [**V1NodeConfigSource**](V1NodeConfigSource.md) |  | [optional] 
**Error** | Pointer to **string** | Error describes any problems reconciling the Spec.ConfigSource to the Active config. Errors may occur, for example, attempting to checkpoint Spec.ConfigSource to the local Assigned record, attempting to checkpoint the payload associated with Spec.ConfigSource, attempting to load or validate the Assigned config, etc. Errors may occur at different points while syncing config. Earlier errors (e.g. download or checkpointing errors) will not result in a rollback to LastKnownGood, and may resolve across Kubelet retries. Later errors (e.g. loading or validating a checkpointed config) will result in a rollback to LastKnownGood. In the latter case, it is usually possible to resolve the error by fixing the config assigned in Spec.ConfigSource. You can find additional information for debugging by searching the error message in the Kubelet log. Error is a human-readable description of the error state; machines can check whether or not Error is empty, but should not rely on the stability of the Error text across Kubelet versions. | [optional] 
**LastKnownGood** | Pointer to [**V1NodeConfigSource**](V1NodeConfigSource.md) |  | [optional] 

## Methods

### NewV1NodeConfigStatus

`func NewV1NodeConfigStatus() *V1NodeConfigStatus`

NewV1NodeConfigStatus instantiates a new V1NodeConfigStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NodeConfigStatusWithDefaults

`func NewV1NodeConfigStatusWithDefaults() *V1NodeConfigStatus`

NewV1NodeConfigStatusWithDefaults instantiates a new V1NodeConfigStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *V1NodeConfigStatus) GetActive() V1NodeConfigSource`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *V1NodeConfigStatus) GetActiveOk() (*V1NodeConfigSource, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *V1NodeConfigStatus) SetActive(v V1NodeConfigSource)`

SetActive sets Active field to given value.

### HasActive

`func (o *V1NodeConfigStatus) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAssigned

`func (o *V1NodeConfigStatus) GetAssigned() V1NodeConfigSource`

GetAssigned returns the Assigned field if non-nil, zero value otherwise.

### GetAssignedOk

`func (o *V1NodeConfigStatus) GetAssignedOk() (*V1NodeConfigSource, bool)`

GetAssignedOk returns a tuple with the Assigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssigned

`func (o *V1NodeConfigStatus) SetAssigned(v V1NodeConfigSource)`

SetAssigned sets Assigned field to given value.

### HasAssigned

`func (o *V1NodeConfigStatus) HasAssigned() bool`

HasAssigned returns a boolean if a field has been set.

### GetError

`func (o *V1NodeConfigStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *V1NodeConfigStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *V1NodeConfigStatus) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *V1NodeConfigStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetLastKnownGood

`func (o *V1NodeConfigStatus) GetLastKnownGood() V1NodeConfigSource`

GetLastKnownGood returns the LastKnownGood field if non-nil, zero value otherwise.

### GetLastKnownGoodOk

`func (o *V1NodeConfigStatus) GetLastKnownGoodOk() (*V1NodeConfigSource, bool)`

GetLastKnownGoodOk returns a tuple with the LastKnownGood field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownGood

`func (o *V1NodeConfigStatus) SetLastKnownGood(v V1NodeConfigSource)`

SetLastKnownGood sets LastKnownGood field to given value.

### HasLastKnownGood

`func (o *V1NodeConfigStatus) HasLastKnownGood() bool`

HasLastKnownGood returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


