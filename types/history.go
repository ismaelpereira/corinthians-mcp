package types

import (
	"encoding/json"
	"fmt"
)

type HistoryFacts struct {
	HistorySummary string        `json:"history_summary"`
	References     []string      `json:"references"`
	Stadium        Stadium       `json:"stadium"`
	Titles         []Title       `json:"titles"`
	Idols          []Idol        `json:"idols"`
	HistoryFacts   []HistoryFact `json:"history_facts"`
}

type StringSlice []string

func (s *StringSlice) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		*s = nil
		return nil
	}

	switch b[0] {
	case '"':
		var single string
		if err := json.Unmarshal(b, &single); err != nil {
			return err
		}
		*s = []string{single}
		return nil
	case '[':
		var arr []string
		if err := json.Unmarshal(b, &arr); err != nil {
			return err
		}
		*s = arr
		return nil
	default:
		return fmt.Errorf("expected string or []string, got: %s", string(b))
	}
}

type Stadium struct {
	Name                     string   `json:"name"`
	OtherNames               []string `json:"other_names"`
	Location                 string   `json:"location"`
	OfficialInaugurationDate string   `json:"official_inauguration_date"`
	Notes                    string   `json:"notes"`
	References               []string `json:"references"`
}

type Title struct {
	Championship string   `json:"championship"`
	Count        int      `json:"count"`
	Years        []string `json:"years"`
	References   []string `json:"references"`
}

type Idol struct {
	Name         string   `json:"name"`
	Position     string   `json:"position"`
	YearsActive  []string `json:"years_active"`
	Contribution string   `json:"contribution"`
	KnowAs       StringSlice `json:"know_as"`
	References   []string `json:"references"`
}

type HistoryFact struct {
	Year         string   `json:"year"`
	Event        string   `json:"event"`
	NameForEvent string   `json:"name_for_event"`
	References   []string `json:"references"`
}
