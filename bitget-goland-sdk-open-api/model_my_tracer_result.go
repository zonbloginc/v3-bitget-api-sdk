/*
Bitget Open API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MyTracerResult struct for MyTracerResult
type MyTracerResult struct {
	AccountEquity        *string `json:"accountEquity,omitempty"`
	CanRemoveTraceUser   *bool   `json:"canRemoveTraceUser,omitempty"`
	TracerHeadPic        *string `json:"tracerHeadPic,omitempty"`
	TracerNickName       *string `json:"tracerNickName,omitempty"`
	TracerUserId         *string `json:"tracerUserId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MyTracerResult MyTracerResult

// NewMyTracerResult instantiates a new MyTracerResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyTracerResult() *MyTracerResult {
	this := MyTracerResult{}
	return &this
}

// NewMyTracerResultWithDefaults instantiates a new MyTracerResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyTracerResultWithDefaults() *MyTracerResult {
	this := MyTracerResult{}
	return &this
}

// GetAccountEquity returns the AccountEquity field value if set, zero value otherwise.
func (o *MyTracerResult) GetAccountEquity() string {
	if o == nil || isNil(o.AccountEquity) {
		var ret string
		return ret
	}
	return *o.AccountEquity
}

// GetAccountEquityOk returns a tuple with the AccountEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTracerResult) GetAccountEquityOk() (*string, bool) {
	if o == nil || isNil(o.AccountEquity) {
		return nil, false
	}
	return o.AccountEquity, true
}

// HasAccountEquity returns a boolean if a field has been set.
func (o *MyTracerResult) HasAccountEquity() bool {
	if o != nil && !isNil(o.AccountEquity) {
		return true
	}

	return false
}

// SetAccountEquity gets a reference to the given string and assigns it to the AccountEquity field.
func (o *MyTracerResult) SetAccountEquity(v string) {
	o.AccountEquity = &v
}

// GetCanRemoveTraceUser returns the CanRemoveTraceUser field value if set, zero value otherwise.
func (o *MyTracerResult) GetCanRemoveTraceUser() bool {
	if o == nil || isNil(o.CanRemoveTraceUser) {
		var ret bool
		return ret
	}
	return *o.CanRemoveTraceUser
}

// GetCanRemoveTraceUserOk returns a tuple with the CanRemoveTraceUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTracerResult) GetCanRemoveTraceUserOk() (*bool, bool) {
	if o == nil || isNil(o.CanRemoveTraceUser) {
		return nil, false
	}
	return o.CanRemoveTraceUser, true
}

// HasCanRemoveTraceUser returns a boolean if a field has been set.
func (o *MyTracerResult) HasCanRemoveTraceUser() bool {
	if o != nil && !isNil(o.CanRemoveTraceUser) {
		return true
	}

	return false
}

// SetCanRemoveTraceUser gets a reference to the given bool and assigns it to the CanRemoveTraceUser field.
func (o *MyTracerResult) SetCanRemoveTraceUser(v bool) {
	o.CanRemoveTraceUser = &v
}

// GetTracerHeadPic returns the TracerHeadPic field value if set, zero value otherwise.
func (o *MyTracerResult) GetTracerHeadPic() string {
	if o == nil || isNil(o.TracerHeadPic) {
		var ret string
		return ret
	}
	return *o.TracerHeadPic
}

// GetTracerHeadPicOk returns a tuple with the TracerHeadPic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTracerResult) GetTracerHeadPicOk() (*string, bool) {
	if o == nil || isNil(o.TracerHeadPic) {
		return nil, false
	}
	return o.TracerHeadPic, true
}

// HasTracerHeadPic returns a boolean if a field has been set.
func (o *MyTracerResult) HasTracerHeadPic() bool {
	if o != nil && !isNil(o.TracerHeadPic) {
		return true
	}

	return false
}

// SetTracerHeadPic gets a reference to the given string and assigns it to the TracerHeadPic field.
func (o *MyTracerResult) SetTracerHeadPic(v string) {
	o.TracerHeadPic = &v
}

// GetTracerNickName returns the TracerNickName field value if set, zero value otherwise.
func (o *MyTracerResult) GetTracerNickName() string {
	if o == nil || isNil(o.TracerNickName) {
		var ret string
		return ret
	}
	return *o.TracerNickName
}

// GetTracerNickNameOk returns a tuple with the TracerNickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTracerResult) GetTracerNickNameOk() (*string, bool) {
	if o == nil || isNil(o.TracerNickName) {
		return nil, false
	}
	return o.TracerNickName, true
}

// HasTracerNickName returns a boolean if a field has been set.
func (o *MyTracerResult) HasTracerNickName() bool {
	if o != nil && !isNil(o.TracerNickName) {
		return true
	}

	return false
}

// SetTracerNickName gets a reference to the given string and assigns it to the TracerNickName field.
func (o *MyTracerResult) SetTracerNickName(v string) {
	o.TracerNickName = &v
}

// GetTracerUserId returns the TracerUserId field value if set, zero value otherwise.
func (o *MyTracerResult) GetTracerUserId() string {
	if o == nil || isNil(o.TracerUserId) {
		var ret string
		return ret
	}
	return *o.TracerUserId
}

// GetTracerUserIdOk returns a tuple with the TracerUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyTracerResult) GetTracerUserIdOk() (*string, bool) {
	if o == nil || isNil(o.TracerUserId) {
		return nil, false
	}
	return o.TracerUserId, true
}

// HasTracerUserId returns a boolean if a field has been set.
func (o *MyTracerResult) HasTracerUserId() bool {
	if o != nil && !isNil(o.TracerUserId) {
		return true
	}

	return false
}

// SetTracerUserId gets a reference to the given string and assigns it to the TracerUserId field.
func (o *MyTracerResult) SetTracerUserId(v string) {
	o.TracerUserId = &v
}

func (o MyTracerResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AccountEquity) {
		toSerialize["accountEquity"] = o.AccountEquity
	}
	if !isNil(o.CanRemoveTraceUser) {
		toSerialize["canRemoveTraceUser"] = o.CanRemoveTraceUser
	}
	if !isNil(o.TracerHeadPic) {
		toSerialize["tracerHeadPic"] = o.TracerHeadPic
	}
	if !isNil(o.TracerNickName) {
		toSerialize["tracerNickName"] = o.TracerNickName
	}
	if !isNil(o.TracerUserId) {
		toSerialize["tracerUserId"] = o.TracerUserId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MyTracerResult) UnmarshalJSON(bytes []byte) (err error) {
	varMyTracerResult := _MyTracerResult{}

	if err = json.Unmarshal(bytes, &varMyTracerResult); err == nil {
		*o = MyTracerResult(varMyTracerResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "accountEquity")
		delete(additionalProperties, "canRemoveTraceUser")
		delete(additionalProperties, "tracerHeadPic")
		delete(additionalProperties, "tracerNickName")
		delete(additionalProperties, "tracerUserId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMyTracerResult struct {
	value *MyTracerResult
	isSet bool
}

func (v NullableMyTracerResult) Get() *MyTracerResult {
	return v.value
}

func (v *NullableMyTracerResult) Set(val *MyTracerResult) {
	v.value = val
	v.isSet = true
}

func (v NullableMyTracerResult) IsSet() bool {
	return v.isSet
}

func (v *NullableMyTracerResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyTracerResult(val *MyTracerResult) *NullableMyTracerResult {
	return &NullableMyTracerResult{value: val, isSet: true}
}

func (v NullableMyTracerResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyTracerResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}