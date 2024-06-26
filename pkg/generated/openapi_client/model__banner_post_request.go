/*
Сервис баннеров

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_client

import (
	"encoding/json"
)

// checks if the BannerPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BannerPostRequest{}

// BannerPostRequest struct for BannerPostRequest
type BannerPostRequest struct {
	// Идентификаторы тэгов
	TagIds []int32 `json:"tag_ids,omitempty"`
	// Идентификатор фичи
	FeatureId *int32 `json:"feature_id,omitempty"`
	// Содержимое баннера
	Content map[string]interface{} `json:"content,omitempty"`
	// Флаг активности баннера
	IsActive *bool `json:"is_active,omitempty"`
}

// NewBannerPostRequest instantiates a new BannerPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBannerPostRequest() *BannerPostRequest {
	this := BannerPostRequest{}
	return &this
}

// NewBannerPostRequestWithDefaults instantiates a new BannerPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBannerPostRequestWithDefaults() *BannerPostRequest {
	this := BannerPostRequest{}
	return &this
}

// GetTagIds returns the TagIds field value if set, zero value otherwise.
func (o *BannerPostRequest) GetTagIds() []int32 {
	if o == nil || IsNil(o.TagIds) {
		var ret []int32
		return ret
	}
	return o.TagIds
}

// GetTagIdsOk returns a tuple with the TagIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BannerPostRequest) GetTagIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.TagIds) {
		return nil, false
	}
	return o.TagIds, true
}

// HasTagIds returns a boolean if a field has been set.
func (o *BannerPostRequest) HasTagIds() bool {
	if o != nil && !IsNil(o.TagIds) {
		return true
	}

	return false
}

// SetTagIds gets a reference to the given []int32 and assigns it to the TagIds field.
func (o *BannerPostRequest) SetTagIds(v []int32) {
	o.TagIds = v
}

// GetFeatureId returns the FeatureId field value if set, zero value otherwise.
func (o *BannerPostRequest) GetFeatureId() int32 {
	if o == nil || IsNil(o.FeatureId) {
		var ret int32
		return ret
	}
	return *o.FeatureId
}

// GetFeatureIdOk returns a tuple with the FeatureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BannerPostRequest) GetFeatureIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FeatureId) {
		return nil, false
	}
	return o.FeatureId, true
}

// HasFeatureId returns a boolean if a field has been set.
func (o *BannerPostRequest) HasFeatureId() bool {
	if o != nil && !IsNil(o.FeatureId) {
		return true
	}

	return false
}

// SetFeatureId gets a reference to the given int32 and assigns it to the FeatureId field.
func (o *BannerPostRequest) SetFeatureId(v int32) {
	o.FeatureId = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *BannerPostRequest) GetContent() map[string]interface{} {
	if o == nil || IsNil(o.Content) {
		var ret map[string]interface{}
		return ret
	}
	return o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BannerPostRequest) GetContentOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Content) {
		return map[string]interface{}{}, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *BannerPostRequest) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given map[string]interface{} and assigns it to the Content field.
func (o *BannerPostRequest) SetContent(v map[string]interface{}) {
	o.Content = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *BannerPostRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BannerPostRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *BannerPostRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *BannerPostRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

func (o BannerPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BannerPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TagIds) {
		toSerialize["tag_ids"] = o.TagIds
	}
	if !IsNil(o.FeatureId) {
		toSerialize["feature_id"] = o.FeatureId
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	return toSerialize, nil
}

type NullableBannerPostRequest struct {
	value *BannerPostRequest
	isSet bool
}

func (v NullableBannerPostRequest) Get() *BannerPostRequest {
	return v.value
}

func (v *NullableBannerPostRequest) Set(val *BannerPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBannerPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBannerPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBannerPostRequest(val *BannerPostRequest) *NullableBannerPostRequest {
	return &NullableBannerPostRequest{value: val, isSet: true}
}

func (v NullableBannerPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBannerPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


