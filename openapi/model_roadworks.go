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

// Roadworks struct for Roadworks
type Roadworks struct {
	Roadworks []Roadwork `json:"roadworks,omitempty"`
}

// NewRoadworks instantiates a new Roadworks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoadworks() *Roadworks {
	this := Roadworks{}
	return &this
}

// NewRoadworksWithDefaults instantiates a new Roadworks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoadworksWithDefaults() *Roadworks {
	this := Roadworks{}
	return &this
}

// GetRoadworks returns the Roadworks field value if set, zero value otherwise.
func (o *Roadworks) GetRoadworks() []Roadwork {
	if o == nil || isNil(o.Roadworks) {
		var ret []Roadwork
		return ret
	}
	return o.Roadworks
}

// GetRoadworksOk returns a tuple with the Roadworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadworks) GetRoadworksOk() ([]Roadwork, bool) {
	if o == nil || isNil(o.Roadworks) {
    return nil, false
	}
	return o.Roadworks, true
}

// HasRoadworks returns a boolean if a field has been set.
func (o *Roadworks) HasRoadworks() bool {
	if o != nil && !isNil(o.Roadworks) {
		return true
	}

	return false
}

// SetRoadworks gets a reference to the given []Roadwork and assigns it to the Roadworks field.
func (o *Roadworks) SetRoadworks(v []Roadwork) {
	o.Roadworks = v
}

func (o Roadworks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Roadworks) {
		toSerialize["roadworks"] = o.Roadworks
	}
	return json.Marshal(toSerialize)
}

type NullableRoadworks struct {
	value *Roadworks
	isSet bool
}

func (v NullableRoadworks) Get() *Roadworks {
	return v.value
}

func (v *NullableRoadworks) Set(val *Roadworks) {
	v.value = val
	v.isSet = true
}

func (v NullableRoadworks) IsSet() bool {
	return v.isSet
}

func (v *NullableRoadworks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoadworks(val *Roadworks) *NullableRoadworks {
	return &NullableRoadworks{value: val, isSet: true}
}

func (v NullableRoadworks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoadworks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


