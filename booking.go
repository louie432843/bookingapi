package bookingapi

type Booking struct {
	BookingID      int      `json:"bookingID"`
	CampgroundID   int      `json:"campgroundID"`
	Sites          []int    `json:"sites"`
	NumberOfGuests int      `json:"numberOfGuests"`
	NumberOfPets   int      `json:"numberOfPets"`
	Requests       string   `json:"requests"`
	Guests         []*Guest `json:"guests"`
}

type Guest struct {
	FirstName      string `json:"firstName"`
	MiddleName     string `json:"middleName"`
	LastName       string `json:"lastName"`
	Prefix         string `json:"prefix"`
	Suffix         string `json:"suffix"`
	PrimaryContact bool   `json:"primaryContact"`
	PhoneNumber    string `json:"phoneNumber"`
	EmailAddress   string `json:"emailAddress"`
	NumberOfStays  int    `json:"numberOfStays"`
}
