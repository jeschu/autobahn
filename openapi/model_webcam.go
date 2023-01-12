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

// Webcam struct for Webcam
type Webcam struct {
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
	Operator *string `json:"operator,omitempty"`
	Imageurl *string `json:"imageurl,omitempty"`
	Linkurl *string `json:"linkurl,omitempty"`
}

// NewWebcam instantiates a new Webcam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebcam() *Webcam {
	this := Webcam{}
	return &this
}

// NewWebcamWithDefaults instantiates a new Webcam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebcamWithDefaults() *Webcam {
	this := Webcam{}
	return &this
}

// GetExtent returns the Extent field value if set, zero value otherwise.
func (o *Webcam) GetExtent() string {
	if o == nil || isNil(o.Extent) {
		var ret string
		return ret
	}
	return *o.Extent
}

// GetExtentOk returns a tuple with the Extent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetExtentOk() (*string, bool) {
	if o == nil || isNil(o.Extent) {
    return nil, false
	}
	return o.Extent, true
}

// HasExtent returns a boolean if a field has been set.
func (o *Webcam) HasExtent() bool {
	if o != nil && !isNil(o.Extent) {
		return true
	}

	return false
}

// SetExtent gets a reference to the given string and assigns it to the Extent field.
func (o *Webcam) SetExtent(v string) {
	o.Extent = &v
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *Webcam) GetIdentifier() string {
	if o == nil || isNil(o.Identifier) {
		var ret string
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetIdentifierOk() (*string, bool) {
	if o == nil || isNil(o.Identifier) {
    return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *Webcam) HasIdentifier() bool {
	if o != nil && !isNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given string and assigns it to the Identifier field.
func (o *Webcam) SetIdentifier(v string) {
	o.Identifier = &v
}

// GetRouteRecommendation returns the RouteRecommendation field value if set, zero value otherwise.
func (o *Webcam) GetRouteRecommendation() []map[string]interface{} {
	if o == nil || isNil(o.RouteRecommendation) {
		var ret []map[string]interface{}
		return ret
	}
	return o.RouteRecommendation
}

// GetRouteRecommendationOk returns a tuple with the RouteRecommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetRouteRecommendationOk() ([]map[string]interface{}, bool) {
	if o == nil || isNil(o.RouteRecommendation) {
    return nil, false
	}
	return o.RouteRecommendation, true
}

// HasRouteRecommendation returns a boolean if a field has been set.
func (o *Webcam) HasRouteRecommendation() bool {
	if o != nil && !isNil(o.RouteRecommendation) {
		return true
	}

	return false
}

// SetRouteRecommendation gets a reference to the given []map[string]interface{} and assigns it to the RouteRecommendation field.
func (o *Webcam) SetRouteRecommendation(v []map[string]interface{}) {
	o.RouteRecommendation = v
}

// GetCoordinate returns the Coordinate field value if set, zero value otherwise.
func (o *Webcam) GetCoordinate() Coordinate {
	if o == nil || isNil(o.Coordinate) {
		var ret Coordinate
		return ret
	}
	return *o.Coordinate
}

// GetCoordinateOk returns a tuple with the Coordinate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetCoordinateOk() (*Coordinate, bool) {
	if o == nil || isNil(o.Coordinate) {
    return nil, false
	}
	return o.Coordinate, true
}

// HasCoordinate returns a boolean if a field has been set.
func (o *Webcam) HasCoordinate() bool {
	if o != nil && !isNil(o.Coordinate) {
		return true
	}

	return false
}

// SetCoordinate gets a reference to the given Coordinate and assigns it to the Coordinate field.
func (o *Webcam) SetCoordinate(v Coordinate) {
	o.Coordinate = &v
}

// GetFooter returns the Footer field value if set, zero value otherwise.
func (o *Webcam) GetFooter() []string {
	if o == nil || isNil(o.Footer) {
		var ret []string
		return ret
	}
	return o.Footer
}

// GetFooterOk returns a tuple with the Footer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetFooterOk() ([]string, bool) {
	if o == nil || isNil(o.Footer) {
    return nil, false
	}
	return o.Footer, true
}

// HasFooter returns a boolean if a field has been set.
func (o *Webcam) HasFooter() bool {
	if o != nil && !isNil(o.Footer) {
		return true
	}

	return false
}

// SetFooter gets a reference to the given []string and assigns it to the Footer field.
func (o *Webcam) SetFooter(v []string) {
	o.Footer = v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Webcam) GetIcon() string {
	if o == nil || isNil(o.Icon) {
		var ret string
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetIconOk() (*string, bool) {
	if o == nil || isNil(o.Icon) {
    return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Webcam) HasIcon() bool {
	if o != nil && !isNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given string and assigns it to the Icon field.
func (o *Webcam) SetIcon(v string) {
	o.Icon = &v
}

// GetIsBlocked returns the IsBlocked field value if set, zero value otherwise.
func (o *Webcam) GetIsBlocked() string {
	if o == nil || isNil(o.IsBlocked) {
		var ret string
		return ret
	}
	return *o.IsBlocked
}

// GetIsBlockedOk returns a tuple with the IsBlocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetIsBlockedOk() (*string, bool) {
	if o == nil || isNil(o.IsBlocked) {
    return nil, false
	}
	return o.IsBlocked, true
}

// HasIsBlocked returns a boolean if a field has been set.
func (o *Webcam) HasIsBlocked() bool {
	if o != nil && !isNil(o.IsBlocked) {
		return true
	}

	return false
}

// SetIsBlocked gets a reference to the given string and assigns it to the IsBlocked field.
func (o *Webcam) SetIsBlocked(v string) {
	o.IsBlocked = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Webcam) GetDescription() []string {
	if o == nil || isNil(o.Description) {
		var ret []string
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetDescriptionOk() ([]string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Webcam) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []string and assigns it to the Description field.
func (o *Webcam) SetDescription(v []string) {
	o.Description = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Webcam) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Webcam) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Webcam) SetTitle(v string) {
	o.Title = &v
}

// GetPoint returns the Point field value if set, zero value otherwise.
func (o *Webcam) GetPoint() string {
	if o == nil || isNil(o.Point) {
		var ret string
		return ret
	}
	return *o.Point
}

// GetPointOk returns a tuple with the Point field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetPointOk() (*string, bool) {
	if o == nil || isNil(o.Point) {
    return nil, false
	}
	return o.Point, true
}

// HasPoint returns a boolean if a field has been set.
func (o *Webcam) HasPoint() bool {
	if o != nil && !isNil(o.Point) {
		return true
	}

	return false
}

// SetPoint gets a reference to the given string and assigns it to the Point field.
func (o *Webcam) SetPoint(v string) {
	o.Point = &v
}

// GetDisplayType returns the DisplayType field value if set, zero value otherwise.
func (o *Webcam) GetDisplayType() DisplayType {
	if o == nil || isNil(o.DisplayType) {
		var ret DisplayType
		return ret
	}
	return *o.DisplayType
}

// GetDisplayTypeOk returns a tuple with the DisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetDisplayTypeOk() (*DisplayType, bool) {
	if o == nil || isNil(o.DisplayType) {
    return nil, false
	}
	return o.DisplayType, true
}

// HasDisplayType returns a boolean if a field has been set.
func (o *Webcam) HasDisplayType() bool {
	if o != nil && !isNil(o.DisplayType) {
		return true
	}

	return false
}

// SetDisplayType gets a reference to the given DisplayType and assigns it to the DisplayType field.
func (o *Webcam) SetDisplayType(v DisplayType) {
	o.DisplayType = &v
}

// GetLorryParkingFeatureIcons returns the LorryParkingFeatureIcons field value if set, zero value otherwise.
func (o *Webcam) GetLorryParkingFeatureIcons() []LorryParkingFeatureIcon {
	if o == nil || isNil(o.LorryParkingFeatureIcons) {
		var ret []LorryParkingFeatureIcon
		return ret
	}
	return o.LorryParkingFeatureIcons
}

// GetLorryParkingFeatureIconsOk returns a tuple with the LorryParkingFeatureIcons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetLorryParkingFeatureIconsOk() ([]LorryParkingFeatureIcon, bool) {
	if o == nil || isNil(o.LorryParkingFeatureIcons) {
    return nil, false
	}
	return o.LorryParkingFeatureIcons, true
}

// HasLorryParkingFeatureIcons returns a boolean if a field has been set.
func (o *Webcam) HasLorryParkingFeatureIcons() bool {
	if o != nil && !isNil(o.LorryParkingFeatureIcons) {
		return true
	}

	return false
}

// SetLorryParkingFeatureIcons gets a reference to the given []LorryParkingFeatureIcon and assigns it to the LorryParkingFeatureIcons field.
func (o *Webcam) SetLorryParkingFeatureIcons(v []LorryParkingFeatureIcon) {
	o.LorryParkingFeatureIcons = v
}

// GetFuture returns the Future field value if set, zero value otherwise.
func (o *Webcam) GetFuture() bool {
	if o == nil || isNil(o.Future) {
		var ret bool
		return ret
	}
	return *o.Future
}

// GetFutureOk returns a tuple with the Future field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetFutureOk() (*bool, bool) {
	if o == nil || isNil(o.Future) {
    return nil, false
	}
	return o.Future, true
}

// HasFuture returns a boolean if a field has been set.
func (o *Webcam) HasFuture() bool {
	if o != nil && !isNil(o.Future) {
		return true
	}

	return false
}

// SetFuture gets a reference to the given bool and assigns it to the Future field.
func (o *Webcam) SetFuture(v bool) {
	o.Future = &v
}

// GetSubtitle returns the Subtitle field value if set, zero value otherwise.
func (o *Webcam) GetSubtitle() string {
	if o == nil || isNil(o.Subtitle) {
		var ret string
		return ret
	}
	return *o.Subtitle
}

// GetSubtitleOk returns a tuple with the Subtitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetSubtitleOk() (*string, bool) {
	if o == nil || isNil(o.Subtitle) {
    return nil, false
	}
	return o.Subtitle, true
}

// HasSubtitle returns a boolean if a field has been set.
func (o *Webcam) HasSubtitle() bool {
	if o != nil && !isNil(o.Subtitle) {
		return true
	}

	return false
}

// SetSubtitle gets a reference to the given string and assigns it to the Subtitle field.
func (o *Webcam) SetSubtitle(v string) {
	o.Subtitle = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *Webcam) GetOperator() string {
	if o == nil || isNil(o.Operator) {
		var ret string
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetOperatorOk() (*string, bool) {
	if o == nil || isNil(o.Operator) {
    return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *Webcam) HasOperator() bool {
	if o != nil && !isNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given string and assigns it to the Operator field.
func (o *Webcam) SetOperator(v string) {
	o.Operator = &v
}

// GetImageurl returns the Imageurl field value if set, zero value otherwise.
func (o *Webcam) GetImageurl() string {
	if o == nil || isNil(o.Imageurl) {
		var ret string
		return ret
	}
	return *o.Imageurl
}

// GetImageurlOk returns a tuple with the Imageurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetImageurlOk() (*string, bool) {
	if o == nil || isNil(o.Imageurl) {
    return nil, false
	}
	return o.Imageurl, true
}

// HasImageurl returns a boolean if a field has been set.
func (o *Webcam) HasImageurl() bool {
	if o != nil && !isNil(o.Imageurl) {
		return true
	}

	return false
}

// SetImageurl gets a reference to the given string and assigns it to the Imageurl field.
func (o *Webcam) SetImageurl(v string) {
	o.Imageurl = &v
}

// GetLinkurl returns the Linkurl field value if set, zero value otherwise.
func (o *Webcam) GetLinkurl() string {
	if o == nil || isNil(o.Linkurl) {
		var ret string
		return ret
	}
	return *o.Linkurl
}

// GetLinkurlOk returns a tuple with the Linkurl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Webcam) GetLinkurlOk() (*string, bool) {
	if o == nil || isNil(o.Linkurl) {
    return nil, false
	}
	return o.Linkurl, true
}

// HasLinkurl returns a boolean if a field has been set.
func (o *Webcam) HasLinkurl() bool {
	if o != nil && !isNil(o.Linkurl) {
		return true
	}

	return false
}

// SetLinkurl gets a reference to the given string and assigns it to the Linkurl field.
func (o *Webcam) SetLinkurl(v string) {
	o.Linkurl = &v
}

func (o Webcam) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !isNil(o.Imageurl) {
		toSerialize["imageurl"] = o.Imageurl
	}
	if !isNil(o.Linkurl) {
		toSerialize["linkurl"] = o.Linkurl
	}
	return json.Marshal(toSerialize)
}

type NullableWebcam struct {
	value *Webcam
	isSet bool
}

func (v NullableWebcam) Get() *Webcam {
	return v.value
}

func (v *NullableWebcam) Set(val *Webcam) {
	v.value = val
	v.isSet = true
}

func (v NullableWebcam) IsSet() bool {
	return v.isSet
}

func (v *NullableWebcam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebcam(val *Webcam) *NullableWebcam {
	return &NullableWebcam{value: val, isSet: true}
}

func (v NullableWebcam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebcam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


