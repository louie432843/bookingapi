package bookingapi

type Campground struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    *Location `json:"location"`
	Sites       []Site    `json:"sites"`
}

type Location struct {
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zipCode"`
	Country      string `json:"country"`
	StreetNumber string `json:"streetNumber"`
	StreetLine1  string `json:"streetLine1"`
	StreetLine2  string `json:"streetLine2"`
}
