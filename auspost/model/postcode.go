package model

import (
	"encoding/json"
	"fmt"
)

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
	Locality []LocalityRaw `json:"-"`
}

func (l *Localities) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &l.RawLocalityWrapper); err != nil {
		return err
	}

	// Single result comes back not as an array?!?
	if l.RawLocalityWrapper.RawLocality[0] == '[' {
		return json.Unmarshal(l.RawLocalityWrapper.RawLocality, &l.Locality)
	}

	var loc LocalityRaw
	if err := json.Unmarshal(l.RawLocalityWrapper.RawLocality, &loc); err != nil {
		return err
	}

	l.Locality = []LocalityRaw{loc}

	return nil
}

type LocalityRaw struct {
	Id        int64       `json:"id"`
	Category  string      `json:"category"`
	Latitude  float64     `json:"latitude"`
	Longitude float64     `json:"longitude"`
	Location  string      `json:"location"`
	Postcode  interface{} `json:"postcode"`
	State     string      `json:"state"`
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

func LocalityRawToLocality(locality LocalityRaw) Locality {
	res := Locality{
		Id:        locality.Id,
		Category:  locality.Category,
		Latitude:  locality.Latitude,
		Longitude: locality.Longitude,
		Location:  locality.Location,
		State:     locality.State,
	}

	if pc, ok := locality.Postcode.(string); ok {
		res.Postcode = pc
	} else {
		res.Postcode = fmt.Sprintf("%v", locality.Postcode)
	}

	return res
}
