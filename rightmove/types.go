package rightmove

import "time"

// Property represents a RightMove property listing.
type Property struct {
	ID                   int    `json:"id"`
	Bedrooms             int    `json:"bedrooms"`
	NumberOfImages       int    `json:"numberOfImages"`
	NumberOfFloorplans   int    `json:"numberOfFloorplans"`
	NumberOfVirtualTours int    `json:"numberOfVirtualTours"`
	Summary              string `json:"summary"`
	DisplayAddress       string `json:"displayAddress"`
	CountryCode          string `json:"countryCode"`
	Location             struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	PropertySubType string `json:"propertySubType"`
	ListingUpdate   struct {
		ListingUpdateReason string    `json:"listingUpdateReason"`
		ListingUpdateDate   time.Time `json:"listingUpdateDate"`
	} `json:"listingUpdate"`
	PremiumListing   bool `json:"premiumListing"`
	FeaturedProperty bool `json:"featuredProperty"`
	Price            struct {
		Amount        int    `json:"amount"`
		Frequency     string `json:"frequency"`
		CurrencyCode  string `json:"currencyCode"`
		DisplayPrices []struct {
			DisplayPrice          string `json:"displayPrice"`
			DisplayPriceQualifier string `json:"displayPriceQualifier"`
		} `json:"displayPrices"`
	} `json:"price"`
	Customer struct {
		BranchID              int         `json:"branchId"`
		BrandPlusLogoURI      string      `json:"brandPlusLogoURI"`
		ContactTelephone      string      `json:"contactTelephone"`
		BranchDisplayName     string      `json:"branchDisplayName"`
		BranchName            string      `json:"branchName"`
		BrandTradingName      string      `json:"brandTradingName"`
		BranchLandingPageURL  string      `json:"branchLandingPageUrl"`
		Development           bool        `json:"development"`
		ShowReducedProperties bool        `json:"showReducedProperties"`
		Commercial            bool        `json:"commercial"`
		ShowOnMap             bool        `json:"showOnMap"`
		EnhancedListing       bool        `json:"enhancedListing"`
		DevelopmentContent    interface{} `json:"developmentContent"`
		BrandPlusLogoURL      string      `json:"brandPlusLogoUrl"`
	} `json:"customer"`
	Distance        float64 `json:"distance"`
	TransactionType string  `json:"transactionType"`
	ProductLabel    struct {
		ProductLabelText string `json:"productLabelText"`
		SpotlightLabel   bool   `json:"spotlightLabel"`
	} `json:"productLabel"`
	Commercial                  bool      `json:"commercial"`
	Development                 bool      `json:"development"`
	Residential                 bool      `json:"residential"`
	Students                    bool      `json:"students"`
	Auction                     bool      `json:"auction"`
	FeesApply                   bool      `json:"feesApply"`
	FeesApplyText               string    `json:"feesApplyText"`
	DisplaySize                 string    `json:"displaySize"`
	ShowOnMap                   bool      `json:"showOnMap"`
	PropertyURL                 string    `json:"propertyUrl"`
	ContactURL                  string    `json:"contactUrl"`
	Channel                     string    `json:"channel"`
	FirstVisibleDate            time.Time `json:"firstVisibleDate"`
	Keywords                    []string  `json:"keywords"`
	KeywordMatchType            string    `json:"keywordMatchType"`
	Saved                       bool      `json:"saved"`
	Hidden                      bool      `json:"hidden"`
	OnlineViewingsAvailable     bool      `json:"onlineViewingsAvailable"`
	Heading                     string    `json:"heading"`
	EnhancedListing             bool      `json:"enhancedListing"`
	PropertyTypeFullDescription string    `json:"propertyTypeFullDescription"`
	PropertyImages              struct {
		Images []struct {
			SrcURL string `json:"srcUrl"`
			URL    string `json:"url"`
		} `json:"images"`
		MainImageSrc    string `json:"mainImageSrc"`
		MainMapImageSrc string `json:"mainMapImageSrc"`
	} `json:"propertyImages"`
	DisplayStatus       string `json:"displayStatus"`
	FormattedBranchName string `json:"formattedBranchName"`
	AddedOrReduced      string `json:"addedOrReduced"`
	IsRecent            bool   `json:"isRecent"`
	FormattedDistance   string `json:"formattedDistance"`
	HasBrandPlus        bool   `json:"hasBrandPlus"`
}

// jsonModel is a representation of the data includes in a rightmove property search page
type jsonModel struct {
	Properties  []Property `json:"properties"`
	ResultCount string     `json:"resultCount"`
}
