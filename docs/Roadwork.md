# Roadwork

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

### NewRoadwork

`func NewRoadwork() *Roadwork`

NewRoadwork instantiates a new Roadwork object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoadworkWithDefaults

`func NewRoadworkWithDefaults() *Roadwork`

NewRoadworkWithDefaults instantiates a new Roadwork object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtent

`func (o *Roadwork) GetExtent() string`

GetExtent returns the Extent field if non-nil, zero value otherwise.

### GetExtentOk

`func (o *Roadwork) GetExtentOk() (*string, bool)`

GetExtentOk returns a tuple with the Extent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtent

`func (o *Roadwork) SetExtent(v string)`

SetExtent sets Extent field to given value.

### HasExtent

`func (o *Roadwork) HasExtent() bool`

HasExtent returns a boolean if a field has been set.

### GetIdentifier

`func (o *Roadwork) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Roadwork) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Roadwork) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Roadwork) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRouteRecommendation

`func (o *Roadwork) GetRouteRecommendation() []map[string]interface{}`

GetRouteRecommendation returns the RouteRecommendation field if non-nil, zero value otherwise.

### GetRouteRecommendationOk

`func (o *Roadwork) GetRouteRecommendationOk() (*[]map[string]interface{}, bool)`

GetRouteRecommendationOk returns a tuple with the RouteRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteRecommendation

`func (o *Roadwork) SetRouteRecommendation(v []map[string]interface{})`

SetRouteRecommendation sets RouteRecommendation field to given value.

### HasRouteRecommendation

`func (o *Roadwork) HasRouteRecommendation() bool`

HasRouteRecommendation returns a boolean if a field has been set.

### GetCoordinate

`func (o *Roadwork) GetCoordinate() Coordinate`

GetCoordinate returns the Coordinate field if non-nil, zero value otherwise.

### GetCoordinateOk

`func (o *Roadwork) GetCoordinateOk() (*Coordinate, bool)`

GetCoordinateOk returns a tuple with the Coordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate

`func (o *Roadwork) SetCoordinate(v Coordinate)`

SetCoordinate sets Coordinate field to given value.

### HasCoordinate

`func (o *Roadwork) HasCoordinate() bool`

HasCoordinate returns a boolean if a field has been set.

### GetFooter

`func (o *Roadwork) GetFooter() []string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *Roadwork) GetFooterOk() (*[]string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *Roadwork) SetFooter(v []string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *Roadwork) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetIcon

`func (o *Roadwork) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Roadwork) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Roadwork) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Roadwork) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsBlocked

`func (o *Roadwork) GetIsBlocked() string`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *Roadwork) GetIsBlockedOk() (*string, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *Roadwork) SetIsBlocked(v string)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *Roadwork) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetDescription

`func (o *Roadwork) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Roadwork) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Roadwork) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Roadwork) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *Roadwork) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Roadwork) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Roadwork) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Roadwork) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPoint

`func (o *Roadwork) GetPoint() string`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *Roadwork) GetPointOk() (*string, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *Roadwork) SetPoint(v string)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *Roadwork) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetDisplayType

`func (o *Roadwork) GetDisplayType() DisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Roadwork) GetDisplayTypeOk() (*DisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Roadwork) SetDisplayType(v DisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *Roadwork) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetLorryParkingFeatureIcons

`func (o *Roadwork) GetLorryParkingFeatureIcons() []LorryParkingFeatureIcon`

GetLorryParkingFeatureIcons returns the LorryParkingFeatureIcons field if non-nil, zero value otherwise.

### GetLorryParkingFeatureIconsOk

`func (o *Roadwork) GetLorryParkingFeatureIconsOk() (*[]LorryParkingFeatureIcon, bool)`

GetLorryParkingFeatureIconsOk returns a tuple with the LorryParkingFeatureIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLorryParkingFeatureIcons

`func (o *Roadwork) SetLorryParkingFeatureIcons(v []LorryParkingFeatureIcon)`

SetLorryParkingFeatureIcons sets LorryParkingFeatureIcons field to given value.

### HasLorryParkingFeatureIcons

`func (o *Roadwork) HasLorryParkingFeatureIcons() bool`

HasLorryParkingFeatureIcons returns a boolean if a field has been set.

### GetFuture

`func (o *Roadwork) GetFuture() bool`

GetFuture returns the Future field if non-nil, zero value otherwise.

### GetFutureOk

`func (o *Roadwork) GetFutureOk() (*bool, bool)`

GetFutureOk returns a tuple with the Future field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuture

`func (o *Roadwork) SetFuture(v bool)`

SetFuture sets Future field to given value.

### HasFuture

`func (o *Roadwork) HasFuture() bool`

HasFuture returns a boolean if a field has been set.

### GetSubtitle

`func (o *Roadwork) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *Roadwork) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *Roadwork) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *Roadwork) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *Roadwork) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Roadwork) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Roadwork) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *Roadwork) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


