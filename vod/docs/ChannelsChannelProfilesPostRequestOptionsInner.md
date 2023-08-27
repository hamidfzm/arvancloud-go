# ChannelsChannelProfilesPostRequestOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bframe** | Pointer to **int32** | minimum: 0 maximum: 16, Sets the amount of b frames. | [optional] 
**Level** | Pointer to **string** | If three-pass encoding is used and a level is set for the encoder,      *                          the bitrate for some segments may exceed the bitrate limit which is      *                          defined by the level.      *                          Accepted values: 1,2,3,4,5,1.1,1.2,1.3,2.1,2.2,3.1,3.2,4.1,4.2,5.1,5.2 | [optional] 
**Cabac** | Pointer to **bool** | Enable or disable CABAC and it should be boolean | [optional] 
**Crf** | Pointer to **int32** | Sets the constant rate factor for quality-based variable bitrate.      *                          Either bitrate or crf is required.      *                          Minimum value should be 0 and the maximum would be 51 | [optional] 
**MinGop** | Pointer to **int32** | Minimum GOP length, the minimum distance between I-frames.      *                          Minimum value should be 1 and the maximum would be 6 | [optional] 
**MinKeyframeInterval** | Pointer to **int32** | Minimum interval in seconds between key frames.      *                          Minimum value should be 1 and the maximum would be 6 | [optional] 
**BitrateTolerance** | Pointer to **string** | bitrate_tolerance | [optional] 
**Fps** | Pointer to **int32** | Minimum value should be 1 and the maximum would be 60 | [optional] 
**Profile** | Pointer to **string** | When setting a profile, all other settings must not exceed the limits      *                          which are defined in the profile. Otherwise, a higher profile may      *                          be automatically chosen. | [optional] 

## Methods

### NewChannelsChannelProfilesPostRequestOptionsInner

`func NewChannelsChannelProfilesPostRequestOptionsInner() *ChannelsChannelProfilesPostRequestOptionsInner`

NewChannelsChannelProfilesPostRequestOptionsInner instantiates a new ChannelsChannelProfilesPostRequestOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelsChannelProfilesPostRequestOptionsInnerWithDefaults

`func NewChannelsChannelProfilesPostRequestOptionsInnerWithDefaults() *ChannelsChannelProfilesPostRequestOptionsInner`

NewChannelsChannelProfilesPostRequestOptionsInnerWithDefaults instantiates a new ChannelsChannelProfilesPostRequestOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBframe

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetBframe() int32`

GetBframe returns the Bframe field if non-nil, zero value otherwise.

### GetBframeOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetBframeOk() (*int32, bool)`

GetBframeOk returns a tuple with the Bframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBframe

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetBframe(v int32)`

SetBframe sets Bframe field to given value.

### HasBframe

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasBframe() bool`

HasBframe returns a boolean if a field has been set.

### GetLevel

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetCabac

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetCabac() bool`

GetCabac returns the Cabac field if non-nil, zero value otherwise.

### GetCabacOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetCabacOk() (*bool, bool)`

GetCabacOk returns a tuple with the Cabac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCabac

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetCabac(v bool)`

SetCabac sets Cabac field to given value.

### HasCabac

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasCabac() bool`

HasCabac returns a boolean if a field has been set.

### GetCrf

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetCrf() int32`

GetCrf returns the Crf field if non-nil, zero value otherwise.

### GetCrfOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetCrfOk() (*int32, bool)`

GetCrfOk returns a tuple with the Crf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrf

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetCrf(v int32)`

SetCrf sets Crf field to given value.

### HasCrf

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasCrf() bool`

HasCrf returns a boolean if a field has been set.

### GetMinGop

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetMinGop() int32`

GetMinGop returns the MinGop field if non-nil, zero value otherwise.

### GetMinGopOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetMinGopOk() (*int32, bool)`

GetMinGopOk returns a tuple with the MinGop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinGop

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetMinGop(v int32)`

SetMinGop sets MinGop field to given value.

### HasMinGop

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasMinGop() bool`

HasMinGop returns a boolean if a field has been set.

### GetMinKeyframeInterval

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetMinKeyframeInterval() int32`

GetMinKeyframeInterval returns the MinKeyframeInterval field if non-nil, zero value otherwise.

### GetMinKeyframeIntervalOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetMinKeyframeIntervalOk() (*int32, bool)`

GetMinKeyframeIntervalOk returns a tuple with the MinKeyframeInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinKeyframeInterval

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetMinKeyframeInterval(v int32)`

SetMinKeyframeInterval sets MinKeyframeInterval field to given value.

### HasMinKeyframeInterval

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasMinKeyframeInterval() bool`

HasMinKeyframeInterval returns a boolean if a field has been set.

### GetBitrateTolerance

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetBitrateTolerance() string`

GetBitrateTolerance returns the BitrateTolerance field if non-nil, zero value otherwise.

### GetBitrateToleranceOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetBitrateToleranceOk() (*string, bool)`

GetBitrateToleranceOk returns a tuple with the BitrateTolerance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBitrateTolerance

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetBitrateTolerance(v string)`

SetBitrateTolerance sets BitrateTolerance field to given value.

### HasBitrateTolerance

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasBitrateTolerance() bool`

HasBitrateTolerance returns a boolean if a field has been set.

### GetFps

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetFps() int32`

GetFps returns the Fps field if non-nil, zero value otherwise.

### GetFpsOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetFpsOk() (*int32, bool)`

GetFpsOk returns a tuple with the Fps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFps

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetFps(v int32)`

SetFps sets Fps field to given value.

### HasFps

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasFps() bool`

HasFps returns a boolean if a field has been set.

### GetProfile

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ChannelsChannelProfilesPostRequestOptionsInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


