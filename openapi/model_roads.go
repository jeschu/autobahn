/*
Autobahn App API

Was passiert auf Deutschlands Bundesstraßen? API für aktuelle Verwaltungsdaten zu Baustellen, Staus und Ladestationen. Außerdem Zugang zu Verkehrsüberwachungskameras und vielen weiteren Datensätzen. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Roads struct for Roads
type Roads struct {
	Roads []string `json:"roads,omitempty"`
}

// NewRoads instantiates a new Roads object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoads() *Roads {
	this := Roads{}
	return &this
}

// NewRoadsWithDefaults instantiates a new Roads object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoadsWithDefaults() *Roads {
	this := Roads{}
	return &this
}

// GetRoads returns the Roads field value if set, zero value otherwise.
func (o *Roads) GetRoads() []string {
	if o == nil || isNil(o.Roads) {
		var ret []string
		return ret
	}
	return o.Roads
}

// GetRoadsOk returns a tuple with the Roads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roads) GetRoadsOk() ([]string, bool) {
	if o == nil || isNil(o.Roads) {
    return nil, false
	}
	return o.Roads, true
}

// HasRoads returns a boolean if a field has been set.
func (o *Roads) HasRoads() bool {
	if o != nil && !isNil(o.Roads) {
		return true
	}

	return false
}

// SetRoads gets a reference to the given []string and assigns it to the Roads field.
func (o *Roads) SetRoads(v []string) {
	o.Roads = v
}

func (o Roads) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Roads) {
		toSerialize["roads"] = o.Roads
	}
	return json.Marshal(toSerialize)
}

type NullableRoads struct {
	value *Roads
	isSet bool
}

func (v NullableRoads) Get() *Roads {
	return v.value
}

func (v *NullableRoads) Set(val *Roads) {
	v.value = val
	v.isSet = true
}

func (v NullableRoads) IsSet() bool {
	return v.isSet
}

func (v *NullableRoads) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoads(val *Roads) *NullableRoads {
	return &NullableRoads{value: val, isSet: true}
}

func (v NullableRoads) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoads) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


