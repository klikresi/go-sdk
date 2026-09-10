package klikresi

import "time"

// TrackingInfo is the full tracking information for a shipment.
type TrackingInfo struct {
	// Status is the normalized delivery status.
	Status DeliveryStatus `json:"status"`
	// Origin is the sender address.
	Origin Address `json:"origin"`
	// Destination is the recipient address.
	Destination Address `json:"destination"`
	// Histories lists the tracking events, newest first.
	Histories []History `json:"histories"`
}

// Address identifies a sender or a recipient.
type Address struct {
	ContactName string `json:"contact_name"`
	Address     string `json:"address"`
}

// History is a single tracking event.
type History struct {
	Status  DeliveryStatus `json:"status"`
	Message string         `json:"message"`
	Date    time.Time      `json:"date"`
}

// RateResult contains shipping rates between an origin and a destination.
type RateResult struct {
	Origin      LocationRef `json:"origin"`
	Destination LocationRef `json:"destination"`
	Pricing     []Pricing   `json:"pricing"`
}

// LocationRef identifies a location by id and name.
type LocationRef struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Pricing is a single shipping rate offer from a courier.
type Pricing struct {
	Type        string  `json:"type"`
	CourierCode string  `json:"courier_code"`
	CourierName string  `json:"courier_name"`
	Service     string  `json:"service"`
	Price       float64 `json:"price"`
	Duration    string  `json:"duration"`
}

// LocationInfo is a district-level location returned by the location
// search.
type LocationInfo struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	District string `json:"district"`
	City     string `json:"city"`
	Province string `json:"province"`
}

// Province is an Indonesian province.
type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// City is an Indonesian city or regency.
type City struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// District is an Indonesian district.
type District struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// LocationPage is one page of location search results.
type LocationPage struct {
	Data       []LocationInfo `json:"data"`
	NextCursor string         `json:"next_cursor"`
}

// ProvincePage is one page of provinces.
type ProvincePage struct {
	Data       []Province `json:"data"`
	NextCursor string     `json:"next_cursor"`
}

// CityPage is one page of cities.
type CityPage struct {
	Data       []City `json:"data"`
	NextCursor string `json:"next_cursor"`
}

// DistrictPage is one page of districts.
type DistrictPage struct {
	Data       []District `json:"data"`
	NextCursor string     `json:"next_cursor"`
}
