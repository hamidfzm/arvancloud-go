# MapTrafficsData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Lists** | Pointer to [**[]CountryList**](CountryList.md) |  | [optional] 

## Methods

### NewMapTrafficsData

`func NewMapTrafficsData() *MapTrafficsData`

NewMapTrafficsData instantiates a new MapTrafficsData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMapTrafficsDataWithDefaults

`func NewMapTrafficsDataWithDefaults() *MapTrafficsData`

NewMapTrafficsDataWithDefaults instantiates a new MapTrafficsData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MapTrafficsData) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MapTrafficsData) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MapTrafficsData) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *MapTrafficsData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLists

`func (o *MapTrafficsData) GetLists() []CountryList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *MapTrafficsData) GetListsOk() (*[]CountryList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *MapTrafficsData) SetLists(v []CountryList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *MapTrafficsData) HasLists() bool`

HasLists returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


