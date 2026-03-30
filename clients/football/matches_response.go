package clients

type MatchParams struct {
	Status       MatchStatus
	Venue        MatchVenue
	Limit        int
	DateFrom     string
	DateTo       string
	Competitions string
}

type MatchStatus string

type MatchVenue string

const (
	StatusScheduled MatchStatus = "SCHEDULED"
	StatusLive      MatchStatus = "LIVE"
	StatusInPlay    MatchStatus = "IN_PLAY"
	StatusPaused    MatchStatus = "PAUSED"
	StatusFinished  MatchStatus = "FINISHED"
	StatusPostponed MatchStatus = "POSTPONED"
	StatusSuspended MatchStatus = "SUSPENDED"
	StatusCancelled MatchStatus = "CANCELLED"
)

const (
	VenueHome MatchVenue = "HOME"
	VenueAway MatchVenue = "AWAY"
)

func (s MatchStatus) String() string {
	return string(s)
}

func (s MatchStatus) IsValid() bool {
	switch s {
	case
		StatusScheduled,
		StatusLive,
		StatusInPlay,
		StatusPaused,
		StatusFinished,
		StatusPostponed,
		StatusSuspended,
		StatusCancelled:
		return true
	default:
		return false
	}
}
func (v MatchVenue) String() string {
	return string(v)
}

func (v MatchVenue) IsValid() bool {
	switch v {
	case
		VenueHome,
		VenueAway:
		return true
	default:
		return false
	}
}

type MatchesResponse struct {
	Filters   Filters   `json:"filters"`
	ResultSet ResultSet `json:"resultSet"`
	Matches   []Match   `json:"matches"`
}

type Filters struct {
	Status       []string `json:"status"`
	Venue        string   `json:"venue"`
	Limit        int      `json:"limit"`
	DateFrom     string   `json:"dateFrom"`
	DateTo       string   `json:"dateTo"`
	Competitions string   `json:"competitions"`
}

type ResultSet struct {
	Count       int    `json:"count"`
	Competitios string `json:"competitions"`
	First       string `json:"first"`
	Last        string `json:"last"`
	Played      int    `json:"played"`
	Wins        int    `json:"wins"`
	Draws       int    `json:"draws"`
	Losses      int    `json:"losses"`
}

type Match struct {
	Area        Area        `json:"area"`
	Competition Competition `json:"competition"`
	ID          int         `json:"id"`
	UtcDate     string      `json:"utcDate"`
	Status      string      `json:"status"`
	Matchday    int         `json:"matchday"`
	Stage       string      `json:"stage"`
	Group       string      `json:"group"`
	LastUpdated string      `json:"lastUpdated"`
	HomeTeam    TeamSummary `json:"homeTeam"`
	AwayTeam    TeamSummary `json:"awayTeam"`
	Score       Score       `json:"score"`
	Referees    []Referee   `json:"referees"`
}

type TeamSummary struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
	TLA       string `json:"tla"`
	Crest     string `json:"crest"`
}

type Score struct {
	Winner   string    `json:"winner"`
	Duration string    `json:"duration"`
	FullTime MatchTime `json:"fullTime"`
	HalfTime MatchTime `json:"halfTime"`
}

type MatchTime struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

type Referee struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Nationality string `json:"nationality"`
}
