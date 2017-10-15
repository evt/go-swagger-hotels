package file

type Results struct {
	SearchResponse SearchResponse
}

type SearchResponse struct {
	ServiceHotels []*ServiceHotel `xml:"ServiceHotel"`
}

type ServiceHotel struct {
	AvailToken    string           `xml:"availToken,attr"`
	Contract      Contract         `xml:"ContractList>Contract"`
	DirectPayment string           `xml:"DirectPayment"`
	Currency      Currency         `xml:"Currency"`
	PackageRate   string           `xml:"PackageRate"`
	HotelInfo     HotelInfo        `xml:"HotelInfo"`
	AvailableRoom []*AvailableRoom `xml:"AvailableRoom"`
}

type Currency struct {
	Name string `xml:"code,attr"`
}

type Contract struct {
	Name           string        `xml:"Name"`
	Class          ContractClass `xml:"Classification"`
	IncomingOffice Office        `xml:"IncomingOffice"`
}

type Office struct {
	Code string `xml:"code,attr"`
}

type ContractClass struct {
	Code string `xml:"code,attr"`
	Name string `xml:",chardata"`
}

type HotelInfo struct {
	Code        string      `xml:"Code"`
	Name        string      `xml:"Name"`
	Destination Destination `xml:"Destination"`
}

type Destination struct {
	Code string `xml:"code,attr"`
}

type AvailableRoom struct {
	RoomCount  int       `xml:"HotelOccupancy>RoomCount"`
	AdultCount int       `xml:"HotelOccupancy>Occupancy>AdultCount"`
	ChildCount int       `xml:"HotelOccupancy>Occupancy>ChildCount"`
	Info       HotelRoom `xml:"HotelRoom"`
}

type HotelRoom struct {
	RoomType             RoomType             `xml:"RoomType"`
	Board                RoomBoard            `xml:"Board"`
	AvailCount           int                  `xml:"availCount,attr"`
	OnRequest            string               `xml:"onRequest,attr"`
	SHRUI                string               `xml:"SHRUI,attr"`
	Price                Price                `xml:"Price"`
	Discounts            []Discount           `xml:"DiscountList>Discount"`
	CancellationPolicies []CancellationPolicy `xml:"CancellationPolicies>CancellationPolicy"`
	Taxes                Taxes                `xml:"TaxAndHotelFeeList"`
}

type Price struct {
	Amount       float64      `xml:"Amount"`
	SellingPrice SellingPrice `xml:"SellingPrice"`
	Commission   float64      `xml:"Commission"`
}

type SellingPrice struct {
	Value     float64 `xml:",chardata"`
	Mandatory string  `xml:"mandatory,attr"`
}

type CancellationPolicy struct {
	Amount   float64 `xml:"amount,attr"`
	DateFrom string  `xml:"dateFrom,attr"`
	Time     string  `xml:"time,attr"`
}

type Discount struct {
	Description string  `xml:"Description"`
	Price       float64 `xml:"Price"`
}

type RoomType struct {
	Name           string `xml:",chardata"`
	Code           string `xml:"code,attr"`
	Type           string `xml:"type,attr"`
	Characteristic string `xml:"characteristic,attr"`
}

type RoomBoard struct {
	Name string `xml:",chardata"`
	Code string `xml:"code,attr"`
}

type Taxes struct {
	TaxItem     []Tax  `xml:"Concept"`
	AllIncluded string `xml:"allIncluded,attr"`
}

type Tax struct {
	Type     string  `xml:"type,attr"`
	Included string  `xml:"included,attr"`
	Percent  string  `xml:"percent,attr"`
	Currency string  `xml:"currencyCode,attr"`
	Total    float64 `xml:"import,attr"`
}
