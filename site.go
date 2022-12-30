package bookingapi

type Site struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	RoomType    string   `json:"roomType"`
	Beds        []Bed    `json:"beds"`
	Capacity    int      `json:"capacity"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Amenities   []string `json:"amenities"`
	Active      bool     `json:"active"`
}

type Bed struct {
	Size    string `json:"size"`
	BedType string `json:"bedType"`
}

type Plot struct {
	SqFeet     float64 `json:"sqFeet"`
	Width      float64 `json:"width"`
	Height     float64 `json:"height"`
	GroundType string  `json:"groundType"`
}
