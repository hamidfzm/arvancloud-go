# V1Secret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources | [optional] 
**Data** | Pointer to **map[string]interface{}** | Data contains the secret data. Each key must consist of alphanumeric characters, &#39;-&#39;, &#39;_&#39; or &#39;.&#39;. The serialized form of the secret data is a base64 encoded string, representing the arbitrary (possibly non-string) data value here. Described in https://tools.ietf.org/html/rfc4648#section-4 | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**StringData** | Pointer to **map[string]interface{}** | stringData allows specifying non-binary secret data in string form. It is provided as a write-only convenience method. All keys and values are merged into the data field on write, overwriting any existing values. It is never output when reading from the API. | [optional] 
**Type** | Pointer to **string** | Used to facilitate programmatic handling of secret data. | [optional] 

## Methods

### NewV1Secret

`func NewV1Secret() *V1Secret`

NewV1Secret instantiates a new V1Secret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1SecretWithDefaults

`func NewV1SecretWithDefaults() *V1Secret`

NewV1SecretWithDefaults instantiates a new V1Secret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *V1Secret) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1Secret) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1Secret) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1Secret) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetData

`func (o *V1Secret) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1Secret) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1Secret) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *V1Secret) HasData() bool`

HasData returns a boolean if a field has been set.

### GetKind

`func (o *V1Secret) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1Secret) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1Secret) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1Secret) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1Secret) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1Secret) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1Secret) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1Secret) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStringData

`func (o *V1Secret) GetStringData() map[string]interface{}`

GetStringData returns the StringData field if non-nil, zero value otherwise.

### GetStringDataOk

`func (o *V1Secret) GetStringDataOk() (*map[string]interface{}, bool)`

GetStringDataOk returns a tuple with the StringData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringData

`func (o *V1Secret) SetStringData(v map[string]interface{})`

SetStringData sets StringData field to given value.

### HasStringData

`func (o *V1Secret) HasStringData() bool`

HasStringData returns a boolean if a field has been set.

### GetType

`func (o *V1Secret) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1Secret) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1Secret) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *V1Secret) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


