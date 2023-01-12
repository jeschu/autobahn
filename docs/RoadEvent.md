# RoadEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extent** | Pointer to **string** |  | [optional] 
**Identifier** | Pointer to **string** |  | [optional] 
**RouteRecommendation** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Coordinate** | Pointer to [**Coordinate**](Coordinate.md) |  | [optional] 
**Footer** | Pointer to **[]string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**IsBlocked** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **[]string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Point** | Pointer to **string** |  | [optional] 
**DisplayType** | Pointer to [**DisplayType**](DisplayType.md) |  | [optional] 
**LorryParkingFeatureIcons** | Pointer to [**[]LorryParkingFeatureIcon**](LorryParkingFeatureIcon.md) |  | [optional] 
**Future** | Pointer to **bool** |  | [optional] 
**Subtitle** | Pointer to **string** |  | [optional] 
**StartTimestamp** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRoadEvent

`func NewRoadEvent() *RoadEvent`

NewRoadEvent instantiates a new RoadEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadEventWithDefaults

`func NewRoadEventWithDefaults() *RoadEvent`

NewRoadEventWithDefaults instantiates a new RoadEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtent

`func (o *RoadEvent) GetExtent() string`

GetExtent returns the Extent field if non-nil, zero value otherwise.

### GetExtentOk

`func (o *RoadEvent) GetExtentOk() (*string, bool)`

GetExtentOk returns a tuple with the Extent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtent

`func (o *RoadEvent) SetExtent(v string)`

SetExtent sets Extent field to given value.

### HasExtent

`func (o *RoadEvent) HasExtent() bool`

HasExtent returns a boolean if a field has been set.

### GetIdentifier

`func (o *RoadEvent) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *RoadEvent) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *RoadEvent) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *RoadEvent) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRouteRecommendation

`func (o *RoadEvent) GetRouteRecommendation() []map[string]interface{}`

GetRouteRecommendation returns the RouteRecommendation field if non-nil, zero value otherwise.

### GetRouteRecommendationOk

`func (o *RoadEvent) GetRouteRecommendationOk() (*[]map[string]interface{}, bool)`

GetRouteRecommendationOk returns a tuple with the RouteRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteRecommendation

`func (o *RoadEvent) SetRouteRecommendation(v []map[string]interface{})`

SetRouteRecommendation sets RouteRecommendation field to given value.

### HasRouteRecommendation

`func (o *RoadEvent) HasRouteRecommendation() bool`

HasRouteRecommendation returns a boolean if a field has been set.

### GetCoordinate

`func (o *RoadEvent) GetCoordinate() Coordinate`

GetCoordinate returns the Coordinate field if non-nil, zero value otherwise.

### GetCoordinateOk

`func (o *RoadEvent) GetCoordinateOk() (*Coordinate, bool)`

GetCoordinateOk returns a tuple with the Coordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate

`func (o *RoadEvent) SetCoordinate(v Coordinate)`

SetCoordinate sets Coordinate field to given value.

### HasCoordinate

`func (o *RoadEvent) HasCoordinate() bool`

HasCoordinate returns a boolean if a field has been set.

### GetFooter

`func (o *RoadEvent) GetFooter() []string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *RoadEvent) GetFooterOk() (*[]string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *RoadEvent) SetFooter(v []string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *RoadEvent) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetIcon

`func (o *RoadEvent) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *RoadEvent) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *RoadEvent) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *RoadEvent) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsBlocked

`func (o *RoadEvent) GetIsBlocked() string`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *RoadEvent) GetIsBlockedOk() (*string, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *RoadEvent) SetIsBlocked(v string)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *RoadEvent) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetDescription

`func (o *RoadEvent) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoadEvent) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoadEvent) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoadEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *RoadEvent) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RoadEvent) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RoadEvent) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RoadEvent) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPoint

`func (o *RoadEvent) GetPoint() string`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *RoadEvent) GetPointOk() (*string, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *RoadEvent) SetPoint(v string)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *RoadEvent) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetDisplayType

`func (o *RoadEvent) GetDisplayType() DisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *RoadEvent) GetDisplayTypeOk() (*DisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *RoadEvent) SetDisplayType(v DisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *RoadEvent) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetLorryParkingFeatureIcons

`func (o *RoadEvent) GetLorryParkingFeatureIcons() []LorryParkingFeatureIcon`

GetLorryParkingFeatureIcons returns the LorryParkingFeatureIcons field if non-nil, zero value otherwise.

### GetLorryParkingFeatureIconsOk

`func (o *RoadEvent) GetLorryParkingFeatureIconsOk() (*[]LorryParkingFeatureIcon, bool)`

GetLorryParkingFeatureIconsOk returns a tuple with the LorryParkingFeatureIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLorryParkingFeatureIcons

`func (o *RoadEvent) SetLorryParkingFeatureIcons(v []LorryParkingFeatureIcon)`

SetLorryParkingFeatureIcons sets LorryParkingFeatureIcons field to given value.

### HasLorryParkingFeatureIcons

`func (o *RoadEvent) HasLorryParkingFeatureIcons() bool`

HasLorryParkingFeatureIcons returns a boolean if a field has been set.

### GetFuture

`func (o *RoadEvent) GetFuture() bool`

GetFuture returns the Future field if non-nil, zero value otherwise.

### GetFutureOk

`func (o *RoadEvent) GetFutureOk() (*bool, bool)`

GetFutureOk returns a tuple with the Future field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuture

`func (o *RoadEvent) SetFuture(v bool)`

SetFuture sets Future field to given value.

### HasFuture

`func (o *RoadEvent) HasFuture() bool`

HasFuture returns a boolean if a field has been set.

### GetSubtitle

`func (o *RoadEvent) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *RoadEvent) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *RoadEvent) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *RoadEvent) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *RoadEvent) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *RoadEvent) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *RoadEvent) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *RoadEvent) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


