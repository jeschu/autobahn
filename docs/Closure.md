# Closure

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

### NewClosure

`func NewClosure() *Closure`

NewClosure instantiates a new Closure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClosureWithDefaults

`func NewClosureWithDefaults() *Closure`

NewClosureWithDefaults instantiates a new Closure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtent

`func (o *Closure) GetExtent() string`

GetExtent returns the Extent field if non-nil, zero value otherwise.

### GetExtentOk

`func (o *Closure) GetExtentOk() (*string, bool)`

GetExtentOk returns a tuple with the Extent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtent

`func (o *Closure) SetExtent(v string)`

SetExtent sets Extent field to given value.

### HasExtent

`func (o *Closure) HasExtent() bool`

HasExtent returns a boolean if a field has been set.

### GetIdentifier

`func (o *Closure) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Closure) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Closure) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Closure) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRouteRecommendation

`func (o *Closure) GetRouteRecommendation() []map[string]interface{}`

GetRouteRecommendation returns the RouteRecommendation field if non-nil, zero value otherwise.

### GetRouteRecommendationOk

`func (o *Closure) GetRouteRecommendationOk() (*[]map[string]interface{}, bool)`

GetRouteRecommendationOk returns a tuple with the RouteRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteRecommendation

`func (o *Closure) SetRouteRecommendation(v []map[string]interface{})`

SetRouteRecommendation sets RouteRecommendation field to given value.

### HasRouteRecommendation

`func (o *Closure) HasRouteRecommendation() bool`

HasRouteRecommendation returns a boolean if a field has been set.

### GetCoordinate

`func (o *Closure) GetCoordinate() Coordinate`

GetCoordinate returns the Coordinate field if non-nil, zero value otherwise.

### GetCoordinateOk

`func (o *Closure) GetCoordinateOk() (*Coordinate, bool)`

GetCoordinateOk returns a tuple with the Coordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate

`func (o *Closure) SetCoordinate(v Coordinate)`

SetCoordinate sets Coordinate field to given value.

### HasCoordinate

`func (o *Closure) HasCoordinate() bool`

HasCoordinate returns a boolean if a field has been set.

### GetFooter

`func (o *Closure) GetFooter() []string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *Closure) GetFooterOk() (*[]string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *Closure) SetFooter(v []string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *Closure) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetIcon

`func (o *Closure) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Closure) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Closure) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Closure) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsBlocked

`func (o *Closure) GetIsBlocked() string`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *Closure) GetIsBlockedOk() (*string, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *Closure) SetIsBlocked(v string)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *Closure) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetDescription

`func (o *Closure) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Closure) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Closure) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Closure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *Closure) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Closure) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Closure) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Closure) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPoint

`func (o *Closure) GetPoint() string`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *Closure) GetPointOk() (*string, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *Closure) SetPoint(v string)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *Closure) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetDisplayType

`func (o *Closure) GetDisplayType() DisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Closure) GetDisplayTypeOk() (*DisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Closure) SetDisplayType(v DisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *Closure) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetLorryParkingFeatureIcons

`func (o *Closure) GetLorryParkingFeatureIcons() []LorryParkingFeatureIcon`

GetLorryParkingFeatureIcons returns the LorryParkingFeatureIcons field if non-nil, zero value otherwise.

### GetLorryParkingFeatureIconsOk

`func (o *Closure) GetLorryParkingFeatureIconsOk() (*[]LorryParkingFeatureIcon, bool)`

GetLorryParkingFeatureIconsOk returns a tuple with the LorryParkingFeatureIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLorryParkingFeatureIcons

`func (o *Closure) SetLorryParkingFeatureIcons(v []LorryParkingFeatureIcon)`

SetLorryParkingFeatureIcons sets LorryParkingFeatureIcons field to given value.

### HasLorryParkingFeatureIcons

`func (o *Closure) HasLorryParkingFeatureIcons() bool`

HasLorryParkingFeatureIcons returns a boolean if a field has been set.

### GetFuture

`func (o *Closure) GetFuture() bool`

GetFuture returns the Future field if non-nil, zero value otherwise.

### GetFutureOk

`func (o *Closure) GetFutureOk() (*bool, bool)`

GetFutureOk returns a tuple with the Future field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuture

`func (o *Closure) SetFuture(v bool)`

SetFuture sets Future field to given value.

### HasFuture

`func (o *Closure) HasFuture() bool`

HasFuture returns a boolean if a field has been set.

### GetSubtitle

`func (o *Closure) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *Closure) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *Closure) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *Closure) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetStartTimestamp

`func (o *Closure) GetStartTimestamp() time.Time`

GetStartTimestamp returns the StartTimestamp field if non-nil, zero value otherwise.

### GetStartTimestampOk

`func (o *Closure) GetStartTimestampOk() (*time.Time, bool)`

GetStartTimestampOk returns a tuple with the StartTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTimestamp

`func (o *Closure) SetStartTimestamp(v time.Time)`

SetStartTimestamp sets StartTimestamp field to given value.

### HasStartTimestamp

`func (o *Closure) HasStartTimestamp() bool`

HasStartTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


