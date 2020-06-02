/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// TimeseriesWidgetRequestMetadata Define an expression alias.
type TimeseriesWidgetRequestMetadata struct {
	// Expression alias.
	AliasName *string `json:"alias_name,omitempty"`
	// Expression name.
	Expression string `json:"expression"`
}

// NewTimeseriesWidgetRequestMetadata instantiates a new TimeseriesWidgetRequestMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeseriesWidgetRequestMetadata(expression string) *TimeseriesWidgetRequestMetadata {
	this := TimeseriesWidgetRequestMetadata{}
	this.Expression = expression
	return &this
}

// NewTimeseriesWidgetRequestMetadataWithDefaults instantiates a new TimeseriesWidgetRequestMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeseriesWidgetRequestMetadataWithDefaults() *TimeseriesWidgetRequestMetadata {
	this := TimeseriesWidgetRequestMetadata{}
	return &this
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *TimeseriesWidgetRequestMetadata) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequestMetadata) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *TimeseriesWidgetRequestMetadata) HasAliasName() bool {
	if o != nil && o.AliasName != nil {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *TimeseriesWidgetRequestMetadata) SetAliasName(v string) {
	o.AliasName = &v
}

// GetExpression returns the Expression field value
func (o *TimeseriesWidgetRequestMetadata) GetExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value
// and a boolean to check if the value has been set.
func (o *TimeseriesWidgetRequestMetadata) GetExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expression, true
}

// SetExpression sets field value
func (o *TimeseriesWidgetRequestMetadata) SetExpression(v string) {
	o.Expression = v
}

func (o TimeseriesWidgetRequestMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	if true {
		toSerialize["expression"] = o.Expression
	}
	return json.Marshal(toSerialize)
}

type NullableTimeseriesWidgetRequestMetadata struct {
	value *TimeseriesWidgetRequestMetadata
	isSet bool
}

func (v NullableTimeseriesWidgetRequestMetadata) Get() *TimeseriesWidgetRequestMetadata {
	return v.value
}

func (v *NullableTimeseriesWidgetRequestMetadata) Set(val *TimeseriesWidgetRequestMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeseriesWidgetRequestMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeseriesWidgetRequestMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeseriesWidgetRequestMetadata(val *TimeseriesWidgetRequestMetadata) *NullableTimeseriesWidgetRequestMetadata {
	return &NullableTimeseriesWidgetRequestMetadata{value: val, isSet: true}
}

func (v NullableTimeseriesWidgetRequestMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeseriesWidgetRequestMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}