package model

type PostcodeSearchResult struct {
	Localities Localities `json:"localities"`
}

type Localities struct {
	Locality []Locality `json:"locality"`
}

type Locality struct {
	Id        int64   `json:"id"`
	Category  string  `json:"category"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Location  string  `json:"location"`
	Postcode  string  `json:"postcode"`
	State     string  `json:"state"`
}
