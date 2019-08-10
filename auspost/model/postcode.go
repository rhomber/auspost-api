package model

import "encoding/json"

type PostcodeSearchResult struct {
	RawLocalitiesWrapper struct {
		RawLocalities json.RawMessage `json:"localities"`
	}
	Localities Localities `json:"-"`
}

func (r *PostcodeSearchResult) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &r.RawLocalitiesWrapper); err != nil {
		return err
	}

	// Empty results come back as a string?!?
	if r.RawLocalitiesWrapper.RawLocalities[0] != '"' {
		return json.Unmarshal(r.RawLocalitiesWrapper.RawLocalities, &r.Localities)
	}

	return nil
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

	// Single result comes back not as an array?!?
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
