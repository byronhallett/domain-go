package domain

const (
	Unrecognised              Feature = "Unrecognised"
	AirConditioning           Feature = "AirConditioning"
	BuiltInWardrobes          Feature = "BuiltInWardrobes"
	CableOrSatellite          Feature = "CableOrSatellite"
	Ensuite                   Feature = "Ensuite"
	Floorboards               Feature = "Floorboards"
	Gas                       Feature = "Gas"
	InternalLaundry           Feature = "InternalLaundry"
	PetsAllowed               Feature = "PetsAllowed"
	SecureParking             Feature = "SecureParking"
	SwimmingPool              Feature = "SwimmingPool"
	Furnished                 Feature = "Furnished"
	GroundFloor               Feature = "GroundFloor"
	WaterViews                Feature = "WaterViews"
	NorthFacing               Feature = "NorthFacing"
	CityViews                 Feature = "CityViews"
	IndoorSpa                 Feature = "IndoorSpa"
	Gym                       Feature = "Gym"
	AlarmSystem               Feature = "AlarmSystem"
	Intercom                  Feature = "Intercom"
	BroadbandInternetAccess   Feature = "BroadbandInternetAccess"
	Bath                      Feature = "Bath"
	Fireplace                 Feature = "Fireplace"
	SeparateDiningRoom        Feature = "SeparateDiningRoom"
	Heating                   Feature = "Heating"
	Dishwasher                Feature = "Dishwasher"
	Study                     Feature = "Study"
	TennisCourt               Feature = "TennisCourt"
	Shed                      Feature = "Shed"
	FullyFenced               Feature = "FullyFenced"
	BalconyDeck               Feature = "BalconyDeck"
	GardenCourtyard           Feature = "GardenCourtyard"
	OutdoorSpa                Feature = "OutdoorSpa"
	DoubleGlazedWindows       Feature = "DoubleGlazedWindows"
	EnergyEfficientAppliances Feature = "EnergyEfficientAppliances"
	WaterEfficientAppliances  Feature = "WaterEfficientAppliances"
	WallCeilingInsulation     Feature = "WallCeilingInsulation"
	RainwaterStorageTank      Feature = "RainwaterStorageTank"
	GreywaterSystem           Feature = "GreywaterSystem"
	WaterEfficientFixtures    Feature = "WaterEfficientFixtures"
	SolarHotWater             Feature = "SolarHotWater"
	SolarPanels               Feature = "SolarPanels"
)

const (
	Sale     ListingType = "Sale"
	Rent     ListingType = "Rent"
	Share    ListingType = "Share"
	Sold     ListingType = "Sold"
	NewHomes ListingType = "NewHomes"
)

const (
	ACT State = "ACT"
	NSW State = "NSW"
	QLD State = "QLD"
	VIC State = "VIC"
	SA  State = "SA"
	WA  State = "WA"
	NT  State = "NT"
	TAS State = "TAS"
)

const (
	Default        SortKey = "Default"
	Suburb         SortKey = "Suburb"
	Price          SortKey = "Price"
	DateUpdated    SortKey = "DateUpdated"
	InspectionTime SortKey = "InspectionTime"
	Proximity      SortKey = "Proximity"
	SoldDate       SortKey = "SoldDate"
)

const (
	Ascending  Direction = "Ascending"
	Descending Direction = "Descending"
)

const (
	Private Source = "Private"
	Agency  Source = "Agency"
	Apm     Source = "Apm"
)

const (
	Image MediaCategory = "Image"
)

const (
	Unknown            PropertyType = "Unknown"
	AcreageSemiRural   PropertyType = "AcreageSemiRural"
	ApartmentUnitFlat  PropertyType = "ApartmentUnitFlat"
	Aquaculture        PropertyType = "Aquaculture"
	BlockOfUnits       PropertyType = "BlockOfUnits"
	CarSpace           PropertyType = "CarSpace"
	DairyFarming       PropertyType = "DairyFarming"
	DevelopmentSite    PropertyType = "DevelopmentSite"
	Duplex             PropertyType = "Duplex"
	Farm               PropertyType = "Farm"
	FishingForestry    PropertyType = "FishingForestry"
	NewHomeDesigns     PropertyType = "NewHomeDesigns"
	House              PropertyType = "House"
	NewHouseLand       PropertyType = "NewHouseLand"
	IrrigationServices PropertyType = "IrrigationServices"
	NewLand            PropertyType = "NewLand"
	Livestock          PropertyType = "Livestock"
	NewApartments      PropertyType = "NewApartments"
	Penthouse          PropertyType = "Penthouse"
	RetirementVillage  PropertyType = "RetirementVillage"
	Rural              PropertyType = "Rural"
	SemiDetached       PropertyType = "SemiDetached"
	SpecialistFarm     PropertyType = "SpecialistFarm"
	Studio             PropertyType = "Studio"
	Terrace            PropertyType = "Terrace"
	Townhouse          PropertyType = "Townhouse"
	VacantLand         PropertyType = "VacantLand"
	Villa              PropertyType = "Villa"
	Cropping           PropertyType = "Cropping"
	Viticulture        PropertyType = "Viticulture"
	MixedFarming       PropertyType = "MixedFarming"
	Grazing            PropertyType = "Grazing"
	Horticulture       PropertyType = "Horticulture"
	Equine             PropertyType = "Equine"
	Farmlet            PropertyType = "Farmlet"
	Orchard            PropertyType = "Orchard"
	RuralLifestyle     PropertyType = "RuralLifestyle"
)

const (
	NotStated           SaleMethod = "NotStated"
	SoldByAuction       SaleMethod = "SoldByAuction"
	SoldByPrivateTreaty SaleMethod = "SoldByPrivateTreaty"
	Withdrawn           SaleMethod = "Withdrawn"
	SoldPriorToAuction  SaleMethod = "SoldPriorToAuction"
)

const (
	Standard    PromoLevel = "Standard"
	StandardPP  PromoLevel = "StandardPP"
	Elite       PromoLevel = "Elite"
	ElitePP     PromoLevel = "ElitePP"
	PremiumPlus PromoLevel = "PremiumPlus"
)

const (
	PropertyListing ResultType = "PropertyListing"
	Project         ResultType = "Project"
)
