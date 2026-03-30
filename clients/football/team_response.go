package clients

const (
	TeamCorinthians = "1779"
)

type TeamResponse struct {
	ID                  int           `json:"id"`
	Name                string        `json:"name"`
	ShortName           string        `json:"shortName"`
	TLA                 string        `json:"tla"`
	Logo                string        `json:"crestUrl"`
	Venue               string        `json:"venue"`
	Address             string        `json:"address"`
	Website             string        `json:"website"`
	Founded             int           `json:"founded"`
	ClubColors          string        `json:"clubColors"`
	Area                Area          `json:"area"`
	RunningCompetitions []Competition `json:"runningCompetitions"`
	Coach               Coach         `json:"coach"`
	Squad               []Player      `json:"squad"`
	LastUpdated         string        `json:"lastUpdated"`
}

type Area struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
	Flag string `json:"flag"`
}

type Competition struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Type   string `json:"type"`
	Emblem string `json:"emblemUrl"`
}

type Coach struct {
	ID          int      `json:"id"`
	FirstName   string   `json:"firstName"`
	LastName    string   `json:"lastName"`
	Name        string   `json:"name"`
	DateOfBirth string   `json:"dateOfBirth"`
	Nationality string   `json:"nationality"`
	Contract    Contract `json:"contract"`
}

type Contract struct {
	Start string `json:"start"`
	Until string `json:"until"`
}

type Player struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Position    string `json:"position"`
	DateOfBirth string `json:"dateOfBirth"`
	Nationality string `json:"nationality"`
}
