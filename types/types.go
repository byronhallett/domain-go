package types

type SearchParameters struct {
	ListingType       ListingType       `json:"listingType,omitempty"`
	PropertyTypes     string            `json:"propertyTypes,omitempty"`
	PropertyFeatures  string            `json:"propertyFeatures,omitempty"`
	ListingAttributes string            `json:"listingAttributes,omitempty"`
	MinBedrooms       float64           `json:"minBedrooms,omitempty"`
	MaxBedrooms       float64           `json:"maxBedrooms,omitempty"`
	MinBathrooms      float64           `json:"minBathrooms,omitempty"`
	MaxBathrooms      float64           `json:"maxBathrooms,omitempty"`
	MinCarspaces      int               `json:"minCarspaces,omitempty"`
	MaxCarspaces      int               `json:"maxCarspaces,omitempty"`
	MinPrice          int               `json:"minPrice,omitempty"`
	MaxPrice          int               `json:"maxPrice,omitempty"`
	MinLandArea       int               `json:"minLandArea,omitempty"`
	MaxLandArea       int               `json:"maxLandArea,omitempty"`
	AdvertiserIds     []int             `json:"advertiserIds,omitempty"`
	AdIDs             []int             `json:"adIds,omitempty"`
	ExcludeAdIds      []int             `json:"excludeAdIds,omitempty"`
	Locations         *[]SearchLocation `json:"locations,omitempty"`
	LocationTerms     string            `json:"locationTerms,omitempty"`
	Keywords          []string          `json:"keywords,omitempty"`
	InspectionFrom    string            `json:"inspectionFrom,omitempty"`
	InspectionTo      string            `json:"inspectionTo,omitempty"`
	AuctionFrom       string            `json:"auctionFrom,omitempty"`
	AuctionTo         string            `json:"auctionTo,omitempty"`
	Sort              *NullableSortBy   `json:"sort,omitempty"`
	Page              int               `json:"page,omitempty"`
	PageSize          int               `json:"pageSize,omitempty"`
	GeoWindow         *GeoWindow        `json:"geoWindow,omitempty"`
}

type SearchLocation struct {
	State                     string `json:"state,omitempty"`
	Region                    string `json:"region,omitempty"`
	Area                      string `json:"area,omitempty"`
	Suburb                    string `json:"suburb,omitempty"`
	PostCode                  string `json:"postCode,omitempty"`
	IncludeSurroundingSuburbs bool   `json:"includeSurroundingSuburbs,omitempty"`
}

type NullableSortBy struct {
	SortKey     SortKey   `json:"sortKey,omitempty"`
	Direction   Direction `json:"direction,omitempty"`
	ProximityTo GeoPoint  `json:"proximityTo,omitempty"`
}

type GeoWindow struct {
	Box     Box     `json:"box,omitempty"`
	Circle  Circle  `json:"circle,omitempty"`
	Polygon Polygon `json:"polygon,omitempty"`
}
type GeoPoint struct {
	Lat float64 `json:"lat,omitempty"`
	Lon float64 `json:"lon,omitempty"`
}
type Box struct {
	TopLeft     GeoPoint `json:"topLeft,omitempty"`
	BottomRight GeoPoint `json:"bottomRight,omitempty"`
}
type Circle struct {
	Center         GeoPoint `json:"center,omitempty"`
	RadiusInMeters int      `json:"radiusInMeters,omitempty"`
}
type Polygon struct {
	Points []GeoPoint `json:"points,omitempty"`
}

type Contact struct {
	Name     string `json:"name,omitempty"`
	PhotoUrl string `json:"photoUrl,omitempty"`
}
type Advertiser struct {
	Type               Source    `json:"type,omitempty"`
	ID                 int       `json:"id,omitempty"`
	Name               string    `json:"name,omitempty"`
	LogoURL            string    `json:"logoURL,omitempty"`
	PreferredColourHex string    `json:"preferredColourHex,omitempty"`
	BannerUrl          string    `json:"bannerUrl,omitempty"`
	Contacts           []Contact `json:"contacts,omitempty"`
}

type Media struct {
	Category MediaCategory `json:"category,omitempty"`
	Url      string        `json:"url,omitempty"`
}

type PriceDetails struct {
	Price        int    `json:"price,omitempty"`
	PriceFrom    int    `json:"priceFrom,omitempty"`
	PriceTo      int    `json:"priceTo,omitempty"`
	DisplayPrice string `json:"displayPrice,omitempty"`
}

type PropertyDetails struct {
	State              State          `json:"state,omitempty"`
	Features           []Feature      `json:"features,omitempty"`
	PropertyType       PropertyType   `json:"propertyType,omitempty"`
	AllPropertyTypes   []PropertyType `json:"allPropertyTypes,omitempty"`
	Bathrooms          float64        `json:"bathrooms,omitempty"`
	Bedrooms           float64        `json:"bedrooms,omitempty"`
	Carspaces          int            `json:"carspaces,omitempty"`
	UnitNumber         string         `json:"unitNumber,omitempty"`
	StreetNumber       string         `json:"streetNumber,omitempty"`
	Street             string         `json:"street,omitempty"`
	Area               string         `json:"area,omitempty"`
	Region             string         `json:"region,omitempty"`
	Suburb             string         `json:"suburb,omitempty"`
	Postcode           string         `json:"postcode,omitempty"`
	DisplayableAddress string         `json:"displayableAddress,omitempty"`
	Latitude           float64        `json:"latitude,omitempty"`
	Longitude          float64        `json:"longitude,omitempty"`
	LandArea           float64        `json:"landArea,omitempty"`
	BuildingArea       float64        `json:"buildingArea,omitempty"`
	OnlyShowProperties []string       `json:"onlyShowProperties,omitempty"`
	DisplayAddressType string         `json:"displayAddressType,omitempty"`
	IsRural            bool           `json:"isRural,omitempty"`
	IsNew              bool           `json:"isNew,omitempty"`
}

type AuctionSchedule struct {
	Time            string `json:"time,omitempty"`
	AuctionLocation string `json:"auctionLocation,omitempty"`
}

type Times struct {
	OpeningTime string `json:"openingTime,omitempty"`
	ClosingTime string `json:"closingTime,omitempty"`
}

type InspectionSchedule struct {
	ByAppointment bool    `json:"byAppointment,omitempty"`
	Recurring     bool    `json:"recurring,omitempty"`
	Times         []Times `json:"times,omitempty"`
}
type SoldData struct {
	Source     Source     `json:"source,omitempty"`
	SaleMethod SaleMethod `json:"saleMethod,omitempty"`
	SoldDate   string     `json:"soldDate,omitempty"`
	SoldPrice  int        `json:"soldPrice,omitempty"`
}

type ListingData struct {
	PromoType          PromoLevel         `json:"promoType,omitempty"`
	ListingType        ListingType        `json:"listingType,omitempty"`
	ID                 int                `json:"id,omitempty"`
	ProjectID          int                `json:"projectID,omitempty"`
	Advertiser         Advertiser         `json:"advertiser,omitempty"`
	PriceDetails       PriceDetails       `json:"priceDetails,omitempty"`
	Media              []Media            `json:"media,omitempty"`
	PropertyDetails    PropertyDetails    `json:"propertyDetails,omitempty"`
	Headline           string             `json:"headline,omitempty"`
	SummaryDescription string             `json:"summaryDescription,omitempty"`
	HasFloorplan       bool               `json:"hasFloorplan,omitempty"`
	HasVideo           bool               `json:"hasVideo,omitempty"`
	Labels             []string           `json:"labels,omitempty"`
	AuctionSchedule    AuctionSchedule    `json:"auctionSchedule,omitempty"`
	InspectionSchedule InspectionSchedule `json:"inspectionSchedule,omitempty"`
	SoldData           SoldData           `json:"soldData,omitempty"`
	ListingSlug        string             `json:"listingSlug,omitempty"`
	UpdatedSince       string             `json:"updatedSince,omitempty"`
}

type ProjectData struct {
	PromoLevel         PromoLevel `json:"promoLevel,omitempty"`
	State              State      `json:"state,omitempty"`
	ID                 int        `json:"id,omitempty"`
	Name               string     `json:"name,omitempty"`
	BannerUrl          string     `json:"bannerUrl,omitempty"`
	PreferredColorHex  string     `json:"preferredColorHex,omitempty"`
	LogoUrl            string     `json:"logoUrl,omitempty"`
	Labels             []string   `json:"labels,omitempty"`
	DisplayableAddress string     `json:"displayableAddress,omitempty"`
	Suburb             string     `json:"suburb,omitempty"`
	Features           []Feature  `json:"features,omitempty"`
	Media              []Media    `json:"media,omitempty"`
	ProjectSlug        string     `json:"projectSlug,omitempty"`
}

type SearchResult struct {
	Type     ResultType    `json:"type,omitempty"`
	Listing  ListingData   `json:"listing,omitempty"`
	Listings []ListingData `json:"listings,omitempty"`
	Project  ProjectData   `json:"project,omitempty"`
}
type Feature string
type ListingType string
type State string
type SortKey string
type Direction string
type Source string
type MediaCategory string
type PropertyType string
type SaleMethod string
type PromoLevel string
type ResultType string
