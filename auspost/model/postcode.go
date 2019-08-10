package model

import "encoding/json"

type PostcodeSearchResult struct {
	Localities Localities `json:"localities"`
}

type Localities struct {
	RawLocalityWrapper struct {
		RawLocality json.RawMessage `json:"locality"`
	}
	Locality []Locality `json:"-"`
}

func (l *Localities) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &l.RawLocalityWrapper); err != nil {
		return err
	}
	if l.RawLocalityWrapper.RawLocality[0] == '[' {
		return json.Unmarshal(l.RawLocalityWrapper.RawLocality, &l.Locality)
	}

	var loc Locality
	if err := json.Unmarshal(l.RawLocalityWrapper.RawLocality, &loc); err != nil {
		return err
	}

	l.Locality = []Locality{loc}

	return nil
}

type Locality struct {
	Id        int64   `json:"id"`
	Category  string  `json:"category"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Location  string  `json:"location"`
	Postcode  int     `json:"postcode"`
	State     string  `json:"state"`
}
