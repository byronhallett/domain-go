package domain

type SearchParameters struct {
	ListingType       ListingType      `json:"listingType"`
	PropertyTypes     string           `json:"propertyTypes"`
	PropertyFeatures  string           `json:"propertyFeatures"`
	ListingAttributes string           `json:"listingAttributes"`
	MinBedrooms       float64          `json:"minBedrooms"`
	MaxBedrooms       float64          `json:"maxBedrooms"`
	MinBathrooms      float64          `json:"minBathrooms"`
	MaxBathrooms      float64          `json:"maxBathrooms"`
	MinCarspaces      int              `json:"minCarspaces"`
	MaxCarspaces      int              `json:"maxCarspaces"`
	MinPrice          int              `json:"minPrice"`
	MaxPrice          int              `json:"maxPrice"`
	MinLandArea       int              `json:"minLandArea"`
	MaxLandArea       int              `json:"maxLandArea"`
	AdvertiserIds     []int            `json:"advertiserIds"`
	AdIDs             []int            `json:"adIds"`
	ExcludeAdIds      []int            `json:"excludeAdIds"`
	Locations         []SearchLocation `json:"locations"`
	LocationTerms     string           `json:"locationTerms"`
	Keywords          []string         `json:"keywords"`
	InspectionFrom    string           `json:"inspectionFrom"`
	InspectionTo      string           `json:"inspectionTo"`
	AuctionFrom       string           `json:"auctionFrom"`
	AuctionTo         string           `json:"auctionTo"`
	Sort              NullableSortBy   `json:"sort"`
	Page              int              `json:"page"`
	PageSize          int              `json:"pageSize"`
	GeoWindow         GeoWindow        `json:"geoWindow"`
}

type SearchLocation struct {
	State                     string `json:"state"`
	Region                    string `json:"region"`
	Area                      string `json:"area"`
	Suburb                    string `json:"suburb"`
	PostCode                  string `json:"postCode"`
	IncludeSurroundingSuburbs bool   `json:"includeSurroundingSuburbs"`
}

type NullableSortBy struct {
	SortKey     SortKey   `json:"sortKey"`
	Direction   Direction `json:"direction"`
	ProximityTo GeoPoint  `json:"proximityTo"`
}

type GeoWindow struct {
	Box     Box     `json:"box"`
	Circle  Circle  `json:"circle"`
	Polygon Polygon `json:"polygon"`
}
type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}
type Box struct {
	TopLeft     GeoPoint `json:"topLeft"`
	BottomRight GeoPoint `json:"bottomRight"`
}
type Circle struct {
	Center         GeoPoint `json:"center"`
	RadiusInMeters int      `json:"radiusInMeters"`
}
type Polygon struct {
	Points []GeoPoint `json:"points"`
}

type Contact struct {
	Name     string `json:"name"`
	PhotoUrl string `json:"photoUrl"`
}
type Advertiser struct {
	Type               Source    `json:"type"`
	ID                 int       `json:"id"`
	Name               string    `json:"name"`
	LogoURL            string    `json:"logoURL"`
	PreferredColourHex string    `json:"preferredColourHex"`
	BannerUrl          string    `json:"bannerUrl"`
	Contacts           []Contact `json:"contacts"`
}

type Media struct {
	Category MediaCategory `json:"category"`
	Url      string        `json:"url"`
}

type PriceDetails struct {
	Price        int    `json:"price"`
	PriceFrom    int    `json:"priceFrom"`
	PriceTo      int    `json:"priceTo"`
	DisplayPrice string `json:"displayPrice"`
}

type PropertyDetails struct {
	State              State          `json:"state"`
	Features           []Feature      `json:"features"`
	PropertyType       PropertyType   `json:"propertyType"`
	AllPropertyTypes   []PropertyType `json:"allPropertyTypes"`
	Bathrooms          float64        `json:"bathrooms"`
	Bedrooms           float64        `json:"bedrooms"`
	Carspaces          int            `json:"carspaces"`
	UnitNumber         string         `json:"unitNumber"`
	StreetNumber       string         `json:"streetNumber"`
	Street             string         `json:"street"`
	Area               string         `json:"area"`
	Region             string         `json:"region"`
	Suburb             string         `json:"suburb"`
	Postcode           string         `json:"postcode"`
	DisplayableAddress string         `json:"displayableAddress"`
	Latitude           float64        `json:"latitude"`
	Longitude          float64        `json:"longitude"`
	LandArea           float64        `json:"landArea"`
	BuildingArea       float64        `json:"buildingArea"`
	OnlyShowProperties []string       `json:"onlyShowProperties"`
	DisplayAddressType string         `json:"displayAddressType"`
	IsRural            bool           `json:"isRural"`
	IsNew              bool           `json:"isNew"`
}

type AuctionSchedule struct {
	Time            string `json:"time"`
	AuctionLocation string `json:"auctionLocation"`
}

type Times struct {
	OpeningTime string `json:"openingTime"`
	ClosingTime string `json:"closingTime"`
}

type InspectionSchedule struct {
	ByAppointment bool  `json:"byAppointment"`
	Recurring     bool  `json:"recurring"`
	Times         Times `json:"times"`
}
type SoldData struct {
	Source     Source     `json:"source"`
	SaleMethod SaleMethod `json:"saleMethod"`
	SoldDate   string     `json:"soldDate"`
	SoldPrice  int        `json:"soldPrice"`
}

type ListingData struct {
	PromoType          PromoLevel         `json:"promoType"`
	ListingType        ListingType        `json:"listingType"`
	ID                 int                `json:"id"`
	ProjectID          int                `json:"projectID"`
	Advertiser         Advertiser         `json:"advertiser"`
	PriceDetails       PriceDetails       `json:"priceDetails"`
	Media              []Media            `json:"media"`
	PropertyDetails    PropertyDetails    `json:"propertyDetails"`
	Headline           string             `json:"headline"`
	SummaryDescription string             `json:"summaryDescription"`
	HasFloorplan       bool               `json:"hasFloorplan"`
	HasVideo           bool               `json:"hasVideo"`
	Labels             []string           `json:"labels"`
	AuctionSchedule    AuctionSchedule    `json:"auctionSchedule"`
	InspectionSchedule InspectionSchedule `json:"inspectionSchedule"`
	SoldData           SoldData           `json:"soldData"`
	ListingSlug        string             `json:"listingSlug"`
	UpdatedSince       string             `json:"updatedSince"`
}

type ProjectData struct {
	PromoLevel         PromoLevel `json:"promoLevel"`
	State              State      `json:"state"`
	ID                 int        `json:"id"`
	Name               string     `json:"name"`
	BannerUrl          string     `json:"bannerUrl"`
	PreferredColorHex  string     `json:"preferredColorHex"`
	LogoUrl            string     `json:"logoUrl"`
	Labels             []string   `json:"labels"`
	DisplayableAddress string     `json:"displayableAddress"`
	Suburb             string     `json:"suburb"`
	Features           []Feature  `json:"features"`
	Media              []Media    `json:"media"`
	ProjectSlug        string     `json:"projectSlug"`
}

type SearchResult struct {
	Type     ResultType    `json:"type"`
	Listing  ListingData   `json:"listing"`
	Listings []ListingData `json:"listings"`
	Project  ProjectData   `json:"project"`
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
