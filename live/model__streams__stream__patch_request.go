/*
Arvan LIVE

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package live

import (
	"encoding/json"
)

// checks if the StreamsStreamPatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamsStreamPatchRequest{}

// StreamsStreamPatchRequest struct for StreamsStreamPatchRequest
type StreamsStreamPatchRequest struct {
	// Title of the stream
	Title *string `json:"title,omitempty"`
	// Description of the stream
	Description *string `json:"description,omitempty"`
	// Stream latency type
	Type *string `json:"type,omitempty"`
	// Way to receive input stream
	Mode *string `json:"mode,omitempty"`
	// Public URL of stream input in PULL mode
	InputUrl *string `json:"input_url,omitempty"`
	// Unique slug for your stream. Only contain lower case letters and digits
	Slug *string `json:"slug,omitempty"`
	// Timeshift (DVR) in minutes to watch the earlier content
	Timeshift *int32 `json:"timeshift,omitempty"`
	// stream fps mode
	FpsMode *string `json:"fps_mode,omitempty"`
	// Input steam frame per second
	Fps *int32 `json:"fps,omitempty"`
	// Set this value to True to archive the stream as a VoD
	ArchiveEnabled *bool `json:"archive_enabled,omitempty"`
	// Set this value to True to enable catchup for the stream
	CatchupEnabled *bool `json:"catchup_enabled,omitempty"`
	// Hours of catchup must be available for the stream
	CatchupPeriod *int32 `json:"catchup_period,omitempty"`
	// Way to archive stream
	ArchiveMode *string `json:"archive_mode,omitempty"`
	// Channel Id to save archive
	ChannelId *string `json:"channel_id,omitempty"`
	// If you want to use watermark for a video, use this ID
	WatermarkId *string `json:"watermark_id,omitempty"`
	// Area of the watermark if watermark_id presents
	WatermarkArea *string `json:"watermark_area,omitempty"`
	// Array of convert details
	ConvertInfo []StreamsPostRequestConvertInfoInner `json:"convert_info,omitempty"`
	// Enable or disable secure link for all videos in channel
	SecureLinkEnabled *bool `json:"secure_link_enabled,omitempty"`
	// Key for generate secure links
	SecureLinkKey *string `json:"secure_link_key,omitempty"`
	// IP can be considered as an optional parameter
	SecureLinkWithIp *bool `json:"secure_link_with_ip,omitempty"`
	// Enable or disable Ads
	AdsEnabled *bool `json:"ads_enabled,omitempty"`
	// Ads present method
	PresentType *string `json:"present_type,omitempty"`
	// Created CampaignId in Ads
	CampaignId *string `json:"campaign_id,omitempty"`
}

// NewStreamsStreamPatchRequest instantiates a new StreamsStreamPatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamsStreamPatchRequest() *StreamsStreamPatchRequest {
	this := StreamsStreamPatchRequest{}
	var archiveEnabled bool = false
	this.ArchiveEnabled = &archiveEnabled
	var catchupEnabled bool = false
	this.CatchupEnabled = &catchupEnabled
	return &this
}

// NewStreamsStreamPatchRequestWithDefaults instantiates a new StreamsStreamPatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamsStreamPatchRequestWithDefaults() *StreamsStreamPatchRequest {
	this := StreamsStreamPatchRequest{}
	var archiveEnabled bool = false
	this.ArchiveEnabled = &archiveEnabled
	var catchupEnabled bool = false
	this.CatchupEnabled = &catchupEnabled
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *StreamsStreamPatchRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StreamsStreamPatchRequest) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *StreamsStreamPatchRequest) SetType(v string) {
	o.Type = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *StreamsStreamPatchRequest) SetMode(v string) {
	o.Mode = &v
}

// GetInputUrl returns the InputUrl field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetInputUrl() string {
	if o == nil || IsNil(o.InputUrl) {
		var ret string
		return ret
	}
	return *o.InputUrl
}

// GetInputUrlOk returns a tuple with the InputUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetInputUrlOk() (*string, bool) {
	if o == nil || IsNil(o.InputUrl) {
		return nil, false
	}
	return o.InputUrl, true
}

// HasInputUrl returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasInputUrl() bool {
	if o != nil && !IsNil(o.InputUrl) {
		return true
	}

	return false
}

// SetInputUrl gets a reference to the given string and assigns it to the InputUrl field.
func (o *StreamsStreamPatchRequest) SetInputUrl(v string) {
	o.InputUrl = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetSlug() string {
	if o == nil || IsNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetSlugOk() (*string, bool) {
	if o == nil || IsNil(o.Slug) {
		return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasSlug() bool {
	if o != nil && !IsNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *StreamsStreamPatchRequest) SetSlug(v string) {
	o.Slug = &v
}

// GetTimeshift returns the Timeshift field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetTimeshift() int32 {
	if o == nil || IsNil(o.Timeshift) {
		var ret int32
		return ret
	}
	return *o.Timeshift
}

// GetTimeshiftOk returns a tuple with the Timeshift field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetTimeshiftOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeshift) {
		return nil, false
	}
	return o.Timeshift, true
}

// HasTimeshift returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasTimeshift() bool {
	if o != nil && !IsNil(o.Timeshift) {
		return true
	}

	return false
}

// SetTimeshift gets a reference to the given int32 and assigns it to the Timeshift field.
func (o *StreamsStreamPatchRequest) SetTimeshift(v int32) {
	o.Timeshift = &v
}

// GetFpsMode returns the FpsMode field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetFpsMode() string {
	if o == nil || IsNil(o.FpsMode) {
		var ret string
		return ret
	}
	return *o.FpsMode
}

// GetFpsModeOk returns a tuple with the FpsMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetFpsModeOk() (*string, bool) {
	if o == nil || IsNil(o.FpsMode) {
		return nil, false
	}
	return o.FpsMode, true
}

// HasFpsMode returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasFpsMode() bool {
	if o != nil && !IsNil(o.FpsMode) {
		return true
	}

	return false
}

// SetFpsMode gets a reference to the given string and assigns it to the FpsMode field.
func (o *StreamsStreamPatchRequest) SetFpsMode(v string) {
	o.FpsMode = &v
}

// GetFps returns the Fps field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetFps() int32 {
	if o == nil || IsNil(o.Fps) {
		var ret int32
		return ret
	}
	return *o.Fps
}

// GetFpsOk returns a tuple with the Fps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetFpsOk() (*int32, bool) {
	if o == nil || IsNil(o.Fps) {
		return nil, false
	}
	return o.Fps, true
}

// HasFps returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasFps() bool {
	if o != nil && !IsNil(o.Fps) {
		return true
	}

	return false
}

// SetFps gets a reference to the given int32 and assigns it to the Fps field.
func (o *StreamsStreamPatchRequest) SetFps(v int32) {
	o.Fps = &v
}

// GetArchiveEnabled returns the ArchiveEnabled field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetArchiveEnabled() bool {
	if o == nil || IsNil(o.ArchiveEnabled) {
		var ret bool
		return ret
	}
	return *o.ArchiveEnabled
}

// GetArchiveEnabledOk returns a tuple with the ArchiveEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetArchiveEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ArchiveEnabled) {
		return nil, false
	}
	return o.ArchiveEnabled, true
}

// HasArchiveEnabled returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasArchiveEnabled() bool {
	if o != nil && !IsNil(o.ArchiveEnabled) {
		return true
	}

	return false
}

// SetArchiveEnabled gets a reference to the given bool and assigns it to the ArchiveEnabled field.
func (o *StreamsStreamPatchRequest) SetArchiveEnabled(v bool) {
	o.ArchiveEnabled = &v
}

// GetCatchupEnabled returns the CatchupEnabled field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetCatchupEnabled() bool {
	if o == nil || IsNil(o.CatchupEnabled) {
		var ret bool
		return ret
	}
	return *o.CatchupEnabled
}

// GetCatchupEnabledOk returns a tuple with the CatchupEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetCatchupEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CatchupEnabled) {
		return nil, false
	}
	return o.CatchupEnabled, true
}

// HasCatchupEnabled returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasCatchupEnabled() bool {
	if o != nil && !IsNil(o.CatchupEnabled) {
		return true
	}

	return false
}

// SetCatchupEnabled gets a reference to the given bool and assigns it to the CatchupEnabled field.
func (o *StreamsStreamPatchRequest) SetCatchupEnabled(v bool) {
	o.CatchupEnabled = &v
}

// GetCatchupPeriod returns the CatchupPeriod field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetCatchupPeriod() int32 {
	if o == nil || IsNil(o.CatchupPeriod) {
		var ret int32
		return ret
	}
	return *o.CatchupPeriod
}

// GetCatchupPeriodOk returns a tuple with the CatchupPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetCatchupPeriodOk() (*int32, bool) {
	if o == nil || IsNil(o.CatchupPeriod) {
		return nil, false
	}
	return o.CatchupPeriod, true
}

// HasCatchupPeriod returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasCatchupPeriod() bool {
	if o != nil && !IsNil(o.CatchupPeriod) {
		return true
	}

	return false
}

// SetCatchupPeriod gets a reference to the given int32 and assigns it to the CatchupPeriod field.
func (o *StreamsStreamPatchRequest) SetCatchupPeriod(v int32) {
	o.CatchupPeriod = &v
}

// GetArchiveMode returns the ArchiveMode field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetArchiveMode() string {
	if o == nil || IsNil(o.ArchiveMode) {
		var ret string
		return ret
	}
	return *o.ArchiveMode
}

// GetArchiveModeOk returns a tuple with the ArchiveMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetArchiveModeOk() (*string, bool) {
	if o == nil || IsNil(o.ArchiveMode) {
		return nil, false
	}
	return o.ArchiveMode, true
}

// HasArchiveMode returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasArchiveMode() bool {
	if o != nil && !IsNil(o.ArchiveMode) {
		return true
	}

	return false
}

// SetArchiveMode gets a reference to the given string and assigns it to the ArchiveMode field.
func (o *StreamsStreamPatchRequest) SetArchiveMode(v string) {
	o.ArchiveMode = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetChannelId() string {
	if o == nil || IsNil(o.ChannelId) {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetChannelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ChannelId) {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasChannelId() bool {
	if o != nil && !IsNil(o.ChannelId) {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *StreamsStreamPatchRequest) SetChannelId(v string) {
	o.ChannelId = &v
}

// GetWatermarkId returns the WatermarkId field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetWatermarkId() string {
	if o == nil || IsNil(o.WatermarkId) {
		var ret string
		return ret
	}
	return *o.WatermarkId
}

// GetWatermarkIdOk returns a tuple with the WatermarkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetWatermarkIdOk() (*string, bool) {
	if o == nil || IsNil(o.WatermarkId) {
		return nil, false
	}
	return o.WatermarkId, true
}

// HasWatermarkId returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasWatermarkId() bool {
	if o != nil && !IsNil(o.WatermarkId) {
		return true
	}

	return false
}

// SetWatermarkId gets a reference to the given string and assigns it to the WatermarkId field.
func (o *StreamsStreamPatchRequest) SetWatermarkId(v string) {
	o.WatermarkId = &v
}

// GetWatermarkArea returns the WatermarkArea field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetWatermarkArea() string {
	if o == nil || IsNil(o.WatermarkArea) {
		var ret string
		return ret
	}
	return *o.WatermarkArea
}

// GetWatermarkAreaOk returns a tuple with the WatermarkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetWatermarkAreaOk() (*string, bool) {
	if o == nil || IsNil(o.WatermarkArea) {
		return nil, false
	}
	return o.WatermarkArea, true
}

// HasWatermarkArea returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasWatermarkArea() bool {
	if o != nil && !IsNil(o.WatermarkArea) {
		return true
	}

	return false
}

// SetWatermarkArea gets a reference to the given string and assigns it to the WatermarkArea field.
func (o *StreamsStreamPatchRequest) SetWatermarkArea(v string) {
	o.WatermarkArea = &v
}

// GetConvertInfo returns the ConvertInfo field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetConvertInfo() []StreamsPostRequestConvertInfoInner {
	if o == nil || IsNil(o.ConvertInfo) {
		var ret []StreamsPostRequestConvertInfoInner
		return ret
	}
	return o.ConvertInfo
}

// GetConvertInfoOk returns a tuple with the ConvertInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetConvertInfoOk() ([]StreamsPostRequestConvertInfoInner, bool) {
	if o == nil || IsNil(o.ConvertInfo) {
		return nil, false
	}
	return o.ConvertInfo, true
}

// HasConvertInfo returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasConvertInfo() bool {
	if o != nil && !IsNil(o.ConvertInfo) {
		return true
	}

	return false
}

// SetConvertInfo gets a reference to the given []StreamsPostRequestConvertInfoInner and assigns it to the ConvertInfo field.
func (o *StreamsStreamPatchRequest) SetConvertInfo(v []StreamsPostRequestConvertInfoInner) {
	o.ConvertInfo = v
}

// GetSecureLinkEnabled returns the SecureLinkEnabled field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetSecureLinkEnabled() bool {
	if o == nil || IsNil(o.SecureLinkEnabled) {
		var ret bool
		return ret
	}
	return *o.SecureLinkEnabled
}

// GetSecureLinkEnabledOk returns a tuple with the SecureLinkEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetSecureLinkEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureLinkEnabled) {
		return nil, false
	}
	return o.SecureLinkEnabled, true
}

// HasSecureLinkEnabled returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasSecureLinkEnabled() bool {
	if o != nil && !IsNil(o.SecureLinkEnabled) {
		return true
	}

	return false
}

// SetSecureLinkEnabled gets a reference to the given bool and assigns it to the SecureLinkEnabled field.
func (o *StreamsStreamPatchRequest) SetSecureLinkEnabled(v bool) {
	o.SecureLinkEnabled = &v
}

// GetSecureLinkKey returns the SecureLinkKey field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetSecureLinkKey() string {
	if o == nil || IsNil(o.SecureLinkKey) {
		var ret string
		return ret
	}
	return *o.SecureLinkKey
}

// GetSecureLinkKeyOk returns a tuple with the SecureLinkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetSecureLinkKeyOk() (*string, bool) {
	if o == nil || IsNil(o.SecureLinkKey) {
		return nil, false
	}
	return o.SecureLinkKey, true
}

// HasSecureLinkKey returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasSecureLinkKey() bool {
	if o != nil && !IsNil(o.SecureLinkKey) {
		return true
	}

	return false
}

// SetSecureLinkKey gets a reference to the given string and assigns it to the SecureLinkKey field.
func (o *StreamsStreamPatchRequest) SetSecureLinkKey(v string) {
	o.SecureLinkKey = &v
}

// GetSecureLinkWithIp returns the SecureLinkWithIp field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetSecureLinkWithIp() bool {
	if o == nil || IsNil(o.SecureLinkWithIp) {
		var ret bool
		return ret
	}
	return *o.SecureLinkWithIp
}

// GetSecureLinkWithIpOk returns a tuple with the SecureLinkWithIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetSecureLinkWithIpOk() (*bool, bool) {
	if o == nil || IsNil(o.SecureLinkWithIp) {
		return nil, false
	}
	return o.SecureLinkWithIp, true
}

// HasSecureLinkWithIp returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasSecureLinkWithIp() bool {
	if o != nil && !IsNil(o.SecureLinkWithIp) {
		return true
	}

	return false
}

// SetSecureLinkWithIp gets a reference to the given bool and assigns it to the SecureLinkWithIp field.
func (o *StreamsStreamPatchRequest) SetSecureLinkWithIp(v bool) {
	o.SecureLinkWithIp = &v
}

// GetAdsEnabled returns the AdsEnabled field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetAdsEnabled() bool {
	if o == nil || IsNil(o.AdsEnabled) {
		var ret bool
		return ret
	}
	return *o.AdsEnabled
}

// GetAdsEnabledOk returns a tuple with the AdsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetAdsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AdsEnabled) {
		return nil, false
	}
	return o.AdsEnabled, true
}

// HasAdsEnabled returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasAdsEnabled() bool {
	if o != nil && !IsNil(o.AdsEnabled) {
		return true
	}

	return false
}

// SetAdsEnabled gets a reference to the given bool and assigns it to the AdsEnabled field.
func (o *StreamsStreamPatchRequest) SetAdsEnabled(v bool) {
	o.AdsEnabled = &v
}

// GetPresentType returns the PresentType field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetPresentType() string {
	if o == nil || IsNil(o.PresentType) {
		var ret string
		return ret
	}
	return *o.PresentType
}

// GetPresentTypeOk returns a tuple with the PresentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetPresentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PresentType) {
		return nil, false
	}
	return o.PresentType, true
}

// HasPresentType returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasPresentType() bool {
	if o != nil && !IsNil(o.PresentType) {
		return true
	}

	return false
}

// SetPresentType gets a reference to the given string and assigns it to the PresentType field.
func (o *StreamsStreamPatchRequest) SetPresentType(v string) {
	o.PresentType = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *StreamsStreamPatchRequest) GetCampaignId() string {
	if o == nil || IsNil(o.CampaignId) {
		var ret string
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamsStreamPatchRequest) GetCampaignIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *StreamsStreamPatchRequest) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given string and assigns it to the CampaignId field.
func (o *StreamsStreamPatchRequest) SetCampaignId(v string) {
	o.CampaignId = &v
}

func (o StreamsStreamPatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamsStreamPatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.InputUrl) {
		toSerialize["input_url"] = o.InputUrl
	}
	if !IsNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !IsNil(o.Timeshift) {
		toSerialize["timeshift"] = o.Timeshift
	}
	if !IsNil(o.FpsMode) {
		toSerialize["fps_mode"] = o.FpsMode
	}
	if !IsNil(o.Fps) {
		toSerialize["fps"] = o.Fps
	}
	if !IsNil(o.ArchiveEnabled) {
		toSerialize["archive_enabled"] = o.ArchiveEnabled
	}
	if !IsNil(o.CatchupEnabled) {
		toSerialize["catchup_enabled"] = o.CatchupEnabled
	}
	if !IsNil(o.CatchupPeriod) {
		toSerialize["catchup_period"] = o.CatchupPeriod
	}
	if !IsNil(o.ArchiveMode) {
		toSerialize["archive_mode"] = o.ArchiveMode
	}
	if !IsNil(o.ChannelId) {
		toSerialize["channel_id"] = o.ChannelId
	}
	if !IsNil(o.WatermarkId) {
		toSerialize["watermark_id"] = o.WatermarkId
	}
	if !IsNil(o.WatermarkArea) {
		toSerialize["watermark_area"] = o.WatermarkArea
	}
	if !IsNil(o.ConvertInfo) {
		toSerialize["convert_info"] = o.ConvertInfo
	}
	if !IsNil(o.SecureLinkEnabled) {
		toSerialize["secure_link_enabled"] = o.SecureLinkEnabled
	}
	if !IsNil(o.SecureLinkKey) {
		toSerialize["secure_link_key"] = o.SecureLinkKey
	}
	if !IsNil(o.SecureLinkWithIp) {
		toSerialize["secure_link_with_ip"] = o.SecureLinkWithIp
	}
	if !IsNil(o.AdsEnabled) {
		toSerialize["ads_enabled"] = o.AdsEnabled
	}
	if !IsNil(o.PresentType) {
		toSerialize["present_type"] = o.PresentType
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaign_id"] = o.CampaignId
	}
	return toSerialize, nil
}

type NullableStreamsStreamPatchRequest struct {
	value *StreamsStreamPatchRequest
	isSet bool
}

func (v NullableStreamsStreamPatchRequest) Get() *StreamsStreamPatchRequest {
	return v.value
}

func (v *NullableStreamsStreamPatchRequest) Set(val *StreamsStreamPatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamsStreamPatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamsStreamPatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamsStreamPatchRequest(val *StreamsStreamPatchRequest) *NullableStreamsStreamPatchRequest {
	return &NullableStreamsStreamPatchRequest{value: val, isSet: true}
}

func (v NullableStreamsStreamPatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamsStreamPatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


