# Webcam

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
**Operator** | Pointer to **string** |  | [optional] 
**Imageurl** | Pointer to **string** |  | [optional] 
**Linkurl** | Pointer to **string** |  | [optional] 

## Methods

### NewWebcam

`func NewWebcam() *Webcam`

NewWebcam instantiates a new Webcam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebcamWithDefaults

`func NewWebcamWithDefaults() *Webcam`

NewWebcamWithDefaults instantiates a new Webcam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtent

`func (o *Webcam) GetExtent() string`

GetExtent returns the Extent field if non-nil, zero value otherwise.

### GetExtentOk

`func (o *Webcam) GetExtentOk() (*string, bool)`

GetExtentOk returns a tuple with the Extent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtent

`func (o *Webcam) SetExtent(v string)`

SetExtent sets Extent field to given value.

### HasExtent

`func (o *Webcam) HasExtent() bool`

HasExtent returns a boolean if a field has been set.

### GetIdentifier

`func (o *Webcam) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *Webcam) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *Webcam) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *Webcam) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### GetRouteRecommendation

`func (o *Webcam) GetRouteRecommendation() []map[string]interface{}`

GetRouteRecommendation returns the RouteRecommendation field if non-nil, zero value otherwise.

### GetRouteRecommendationOk

`func (o *Webcam) GetRouteRecommendationOk() (*[]map[string]interface{}, bool)`

GetRouteRecommendationOk returns a tuple with the RouteRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouteRecommendation

`func (o *Webcam) SetRouteRecommendation(v []map[string]interface{})`

SetRouteRecommendation sets RouteRecommendation field to given value.

### HasRouteRecommendation

`func (o *Webcam) HasRouteRecommendation() bool`

HasRouteRecommendation returns a boolean if a field has been set.

### GetCoordinate

`func (o *Webcam) GetCoordinate() Coordinate`

GetCoordinate returns the Coordinate field if non-nil, zero value otherwise.

### GetCoordinateOk

`func (o *Webcam) GetCoordinateOk() (*Coordinate, bool)`

GetCoordinateOk returns a tuple with the Coordinate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinate

`func (o *Webcam) SetCoordinate(v Coordinate)`

SetCoordinate sets Coordinate field to given value.

### HasCoordinate

`func (o *Webcam) HasCoordinate() bool`

HasCoordinate returns a boolean if a field has been set.

### GetFooter

`func (o *Webcam) GetFooter() []string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *Webcam) GetFooterOk() (*[]string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *Webcam) SetFooter(v []string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *Webcam) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetIcon

`func (o *Webcam) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Webcam) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Webcam) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Webcam) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetIsBlocked

`func (o *Webcam) GetIsBlocked() string`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *Webcam) GetIsBlockedOk() (*string, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *Webcam) SetIsBlocked(v string)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *Webcam) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetDescription

`func (o *Webcam) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Webcam) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Webcam) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Webcam) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *Webcam) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Webcam) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Webcam) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Webcam) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPoint

`func (o *Webcam) GetPoint() string`

GetPoint returns the Point field if non-nil, zero value otherwise.

### GetPointOk

`func (o *Webcam) GetPointOk() (*string, bool)`

GetPointOk returns a tuple with the Point field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoint

`func (o *Webcam) SetPoint(v string)`

SetPoint sets Point field to given value.

### HasPoint

`func (o *Webcam) HasPoint() bool`

HasPoint returns a boolean if a field has been set.

### GetDisplayType

`func (o *Webcam) GetDisplayType() DisplayType`

GetDisplayType returns the DisplayType field if non-nil, zero value otherwise.

### GetDisplayTypeOk

`func (o *Webcam) GetDisplayTypeOk() (*DisplayType, bool)`

GetDisplayTypeOk returns a tuple with the DisplayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayType

`func (o *Webcam) SetDisplayType(v DisplayType)`

SetDisplayType sets DisplayType field to given value.

### HasDisplayType

`func (o *Webcam) HasDisplayType() bool`

HasDisplayType returns a boolean if a field has been set.

### GetLorryParkingFeatureIcons

`func (o *Webcam) GetLorryParkingFeatureIcons() []LorryParkingFeatureIcon`

GetLorryParkingFeatureIcons returns the LorryParkingFeatureIcons field if non-nil, zero value otherwise.

### GetLorryParkingFeatureIconsOk

`func (o *Webcam) GetLorryParkingFeatureIconsOk() (*[]LorryParkingFeatureIcon, bool)`

GetLorryParkingFeatureIconsOk returns a tuple with the LorryParkingFeatureIcons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLorryParkingFeatureIcons

`func (o *Webcam) SetLorryParkingFeatureIcons(v []LorryParkingFeatureIcon)`

SetLorryParkingFeatureIcons sets LorryParkingFeatureIcons field to given value.

### HasLorryParkingFeatureIcons

`func (o *Webcam) HasLorryParkingFeatureIcons() bool`

HasLorryParkingFeatureIcons returns a boolean if a field has been set.

### GetFuture

`func (o *Webcam) GetFuture() bool`

GetFuture returns the Future field if non-nil, zero value otherwise.

### GetFutureOk

`func (o *Webcam) GetFutureOk() (*bool, bool)`

GetFutureOk returns a tuple with the Future field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFuture

`func (o *Webcam) SetFuture(v bool)`

SetFuture sets Future field to given value.

### HasFuture

`func (o *Webcam) HasFuture() bool`

HasFuture returns a boolean if a field has been set.

### GetSubtitle

`func (o *Webcam) GetSubtitle() string`

GetSubtitle returns the Subtitle field if non-nil, zero value otherwise.

### GetSubtitleOk

`func (o *Webcam) GetSubtitleOk() (*string, bool)`

GetSubtitleOk returns a tuple with the Subtitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtitle

`func (o *Webcam) SetSubtitle(v string)`

SetSubtitle sets Subtitle field to given value.

### HasSubtitle

`func (o *Webcam) HasSubtitle() bool`

HasSubtitle returns a boolean if a field has been set.

### GetOperator

`func (o *Webcam) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Webcam) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Webcam) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Webcam) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetImageurl

`func (o *Webcam) GetImageurl() string`

GetImageurl returns the Imageurl field if non-nil, zero value otherwise.

### GetImageurlOk

`func (o *Webcam) GetImageurlOk() (*string, bool)`

GetImageurlOk returns a tuple with the Imageurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageurl

`func (o *Webcam) SetImageurl(v string)`

SetImageurl sets Imageurl field to given value.

### HasImageurl

`func (o *Webcam) HasImageurl() bool`

HasImageurl returns a boolean if a field has been set.

### GetLinkurl

`func (o *Webcam) GetLinkurl() string`

GetLinkurl returns the Linkurl field if non-nil, zero value otherwise.

### GetLinkurlOk

`func (o *Webcam) GetLinkurlOk() (*string, bool)`

GetLinkurlOk returns a tuple with the Linkurl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkurl

`func (o *Webcam) SetLinkurl(v string)`

SetLinkurl sets Linkurl field to given value.

### HasLinkurl

`func (o *Webcam) HasLinkurl() bool`

HasLinkurl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


