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

// Roadwork struct for Roadwork
type Roadwork struct {
	Extent *string `json:"extent,omitempty"`
	Identifier *string `json:"identifier,omitempty"`
	RouteRecommendation []map[string]interface{} `json:"routeRecommendation,omitempty"`
	Coordinate *Coordinate `json:"coordinate,omitempty"`
	Footer []string `json:"footer,omitempty"`
	Icon *string `json:"icon,omitempty"`
	IsBlocked *string `json:"isBlocked,omitempty"`
	Description []string `json:"description,omitempty"`
	Title *string `json:"title,omitempty"`
	Point *string `json:"point,omitempty"`
	DisplayType *DisplayType `json:"display_type,omitempty"`
	LorryParkingFeatureIcons []LorryParkingFeatureIcon `json:"lorryParkingFeatureIcons,omitempty"`
	Future *bool `json:"future,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	StartTimestamp *Time `json:"startTimestamp,omitempty"`
}

// NewRoadwork instantiates a new Roadwork object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoadwork() *Roadwork {
	this := Roadwork{}
	return &this
}

// NewRoadworkWithDefaults instantiates a new Roadwork object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoadworkWithDefaults() *Roadwork {
	this := Roadwork{}
	return &this
}

// GetExtent returns the Extent field value if set, zero value otherwise.
func (o *Roadwork) GetExtent() string {
	if o == nil || isNil(o.Extent) {
		var ret string
		return ret
	}
	return *o.Extent
}

// GetExtentOk returns a tuple with the Extent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetExtentOk() (*string, bool) {
	if o == nil || isNil(o.Extent) {
    return nil, false
	}
	return o.Extent, true
}

// HasExtent returns a boolean if a field has been set.
func (o *Roadwork) HasExtent() bool {
	if o != nil && !isNil(o.Extent) {
		return true
	}

	return false
}

// SetExtent gets a reference to the given string and assigns it to the Extent field.
func (o *Roadwork) SetExtent(v string) {
	o.Extent = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Roadwork) GetIdentifier() string {
	if o == nil || isNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.Identifier) {
    return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Roadwork) HasIdentifier() bool {
	if o != nil && !isNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Roadwork) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetRouteRecommendation returns the RouteRecommendation field value if set, zero value otherwise.
func (o *Roadwork) GetRouteRecommendation() []map[string]interface{} {
	if o == nil || isNil(o.RouteRecommendation) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RouteRecommendation
}

// GetRouteRecommendationOk returns a tuple with the RouteRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetRouteRecommendationOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.RouteRecommendation) {
    return nil, false
	}
	return o.RouteRecommendation, true
}

// HasRouteRecommendation returns a boolean if a field has been set.
func (o *Roadwork) HasRouteRecommendation() bool {
	if o != nil && !isNil(o.RouteRecommendation) {
		return true
	}

	return false
}

// SetRouteRecommendation gets a reference to the given []map[string]interface{} and assigns it to the RouteRecommendation field.
func (o *Roadwork) SetRouteRecommendation(v []map[string]interface{}) {
	o.RouteRecommendation = v
}

// GetCoordinate returns the Coordinate field value if set, zero value otherwise.
func (o *Roadwork) GetCoordinate() Coordinate {
	if o == nil || isNil(o.Coordinate) {
		var ret Coordinate
		return ret
	}
	return *o.Coordinate
}

// GetCoordinateOk returns a tuple with the Coordinate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetCoordinateOk() (*Coordinate, bool) {
	if o == nil || isNil(o.Coordinate) {
    return nil, false
	}
	return o.Coordinate, true
}

// HasCoordinate returns a boolean if a field has been set.
func (o *Roadwork) HasCoordinate() bool {
	if o != nil && !isNil(o.Coordinate) {
		return true
	}

	return false
}

// SetCoordinate gets a reference to the given Coordinate and assigns it to the Coordinate field.
func (o *Roadwork) SetCoordinate(v Coordinate) {
	o.Coordinate = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *Roadwork) GetFooter() []string {
	if o == nil || isNil(o.Footer) {
		var ret []string
		return ret
	}
	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetFooterOk() ([]string, bool) {
	if o == nil || isNil(o.Footer) {
    return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *Roadwork) HasFooter() bool {
	if o != nil && !isNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given []string and assigns it to the Footer field.
func (o *Roadwork) SetFooter(v []string) {
	o.Footer = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Roadwork) GetIcon() string {
	if o == nil || isNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetIconOk() (*string, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Roadwork) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Roadwork) SetIcon(v string) {
	o.Icon = &v
}

// GetIsBlocked returns the IsBlocked field value if set, zero value otherwise.
func (o *Roadwork) GetIsBlocked() string {
	if o == nil || isNil(o.IsBlocked) {
		var ret string
		return ret
	}
	return *o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetIsBlockedOk() (*string, bool) {
	if o == nil || isNil(o.IsBlocked) {
    return nil, false
	}
	return o.IsBlocked, true
}

// HasIsBlocked returns a boolean if a field has been set.
func (o *Roadwork) HasIsBlocked() bool {
	if o != nil && !isNil(o.IsBlocked) {
		return true
	}

	return false
}

// SetIsBlocked gets a reference to the given string and assigns it to the IsBlocked field.
func (o *Roadwork) SetIsBlocked(v string) {
	o.IsBlocked = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Roadwork) GetDescription() []string {
	if o == nil || isNil(o.Description) {
		var ret []string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetDescriptionOk() ([]string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Roadwork) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []string and assigns it to the Description field.
func (o *Roadwork) SetDescription(v []string) {
	o.Description = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Roadwork) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Roadwork) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Roadwork) SetTitle(v string) {
	o.Title = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *Roadwork) GetPoint() string {
	if o == nil || isNil(o.Point) {
		var ret string
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetPointOk() (*string, bool) {
	if o == nil || isNil(o.Point) {
    return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *Roadwork) HasPoint() bool {
	if o != nil && !isNil(o.Point) {
		return true
	}

	return false
}

// SetPoint gets a reference to the given string and assigns it to the Point field.
func (o *Roadwork) SetPoint(v string) {
	o.Point = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *Roadwork) GetDisplayType() DisplayType {
	if o == nil || isNil(o.DisplayType) {
		var ret DisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetDisplayTypeOk() (*DisplayType, bool) {
	if o == nil || isNil(o.DisplayType) {
    return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *Roadwork) HasDisplayType() bool {
	if o != nil && !isNil(o.DisplayType) {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given DisplayType and assigns it to the DisplayType field.
func (o *Roadwork) SetDisplayType(v DisplayType) {
	o.DisplayType = &v
}

// GetLorryParkingFeatureIcons returns the LorryParkingFeatureIcons field value if set, zero value otherwise.
func (o *Roadwork) GetLorryParkingFeatureIcons() []LorryParkingFeatureIcon {
	if o == nil || isNil(o.LorryParkingFeatureIcons) {
		var ret []LorryParkingFeatureIcon
		return ret
	}
	return o.LorryParkingFeatureIcons
}

// GetLorryParkingFeatureIconsOk returns a tuple with the LorryParkingFeatureIcons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetLorryParkingFeatureIconsOk() ([]LorryParkingFeatureIcon, bool) {
	if o == nil || isNil(o.LorryParkingFeatureIcons) {
    return nil, false
	}
	return o.LorryParkingFeatureIcons, true
}

// HasLorryParkingFeatureIcons returns a boolean if a field has been set.
func (o *Roadwork) HasLorryParkingFeatureIcons() bool {
	if o != nil && !isNil(o.LorryParkingFeatureIcons) {
		return true
	}

	return false
}

// SetLorryParkingFeatureIcons gets a reference to the given []LorryParkingFeatureIcon and assigns it to the LorryParkingFeatureIcons field.
func (o *Roadwork) SetLorryParkingFeatureIcons(v []LorryParkingFeatureIcon) {
	o.LorryParkingFeatureIcons = v
}

// GetFuture returns the Future field value if set, zero value otherwise.
func (o *Roadwork) GetFuture() bool {
	if o == nil || isNil(o.Future) {
		var ret bool
		return ret
	}
	return *o.Future
}

// GetFutureOk returns a tuple with the Future field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetFutureOk() (*bool, bool) {
	if o == nil || isNil(o.Future) {
    return nil, false
	}
	return o.Future, true
}

// HasFuture returns a boolean if a field has been set.
func (o *Roadwork) HasFuture() bool {
	if o != nil && !isNil(o.Future) {
		return true
	}

	return false
}

// SetFuture gets a reference to the given bool and assigns it to the Future field.
func (o *Roadwork) SetFuture(v bool) {
	o.Future = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *Roadwork) GetSubtitle() string {
	if o == nil || isNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetSubtitleOk() (*string, bool) {
	if o == nil || isNil(o.Subtitle) {
    return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *Roadwork) HasSubtitle() bool {
	if o != nil && !isNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *Roadwork) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetStartTimestamp returns the StartTimestamp field value if set, zero value otherwise.
func (o *Roadwork) GetStartTimestamp() Time {
	if o == nil || isNil(o.StartTimestamp) {
		var ret Time
		return ret
	}
	return *o.StartTimestamp
}

// GetStartTimestampOk returns a tuple with the StartTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Roadwork) GetStartTimestampOk() (*Time, bool) {
	if o == nil || isNil(o.StartTimestamp) {
    return nil, false
	}
	return o.StartTimestamp, true
}

// HasStartTimestamp returns a boolean if a field has been set.
func (o *Roadwork) HasStartTimestamp() bool {
	if o != nil && !isNil(o.StartTimestamp) {
		return true
	}

	return false
}

// SetStartTimestamp gets a reference to the given time.Time and assigns it to the StartTimestamp field.
func (o *Roadwork) SetStartTimestamp(v Time) {
	o.StartTimestamp = &v
}

func (o Roadwork) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Extent) {
		toSerialize["extent"] = o.Extent
	}
	if !isNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !isNil(o.RouteRecommendation) {
		toSerialize["routeRecommendation"] = o.RouteRecommendation
	}
	if !isNil(o.Coordinate) {
		toSerialize["coordinate"] = o.Coordinate
	}
	if !isNil(o.Footer) {
		toSerialize["footer"] = o.Footer
	}
	if !isNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	if !isNil(o.IsBlocked) {
		toSerialize["isBlocked"] = o.IsBlocked
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Point) {
		toSerialize["point"] = o.Point
	}
	if !isNil(o.DisplayType) {
		toSerialize["display_type"] = o.DisplayType
	}
	if !isNil(o.LorryParkingFeatureIcons) {
		toSerialize["lorryParkingFeatureIcons"] = o.LorryParkingFeatureIcons
	}
	if !isNil(o.Future) {
		toSerialize["future"] = o.Future
	}
	if !isNil(o.Subtitle) {
		toSerialize["subtitle"] = o.Subtitle
	}
	if !isNil(o.StartTimestamp) {
		toSerialize["startTimestamp"] = o.StartTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableRoadwork struct {
	value *Roadwork
	isSet bool
}

func (v NullableRoadwork) Get() *Roadwork {
	return v.value
}

func (v *NullableRoadwork) Set(val *Roadwork) {
	v.value = val
	v.isSet = true
}

func (v NullableRoadwork) IsSet() bool {
	return v.isSet
}

func (v *NullableRoadwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoadwork(val *Roadwork) *NullableRoadwork {
	return &NullableRoadwork{value: val, isSet: true}
}

func (v NullableRoadwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoadwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

