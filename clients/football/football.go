package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FootballClient struct {
	BaseURL    string
	APIKey     string
	TeamID     string
	HTTPClient *http.Client
}

func NewFootballClient(baseURL, apiKey string) *FootballClient {
	return &FootballClient{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		TeamID:     TeamCorinthians,
		HTTPClient: &http.Client{},
	}
}

func (c *FootballClient) FetchTeamInfo() (*TeamResponse, error) {
	url := fmt.Sprintf("%s/v4/teams/%s", c.BaseURL, c.TeamID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Auth-Token", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var team TeamResponse
	err = json.Unmarshal(body, &team)
	if err != nil {
		return nil, err
	}

	return &team, nil
}

func (c *FootballClient) FetchMatches(params MatchParams) (*MatchesResponse, error) {
	url := fmt.Sprintf("%s/v4/teams/%s/matches", c.BaseURL, c.TeamID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	if params.Status != "" && params.Status.IsValid() {
		q.Add("status", params.Status.String())
	}

	if params.Venue != "" && params.Venue.IsValid() {
		q.Add("venue", params.Venue.String())
	}

	if params.Limit > 0 {
		q.Add("limit", fmt.Sprintf("%d", params.Limit))
	}

	if params.DateFrom != "" {
		q.Add("dateFrom", params.DateFrom)
	}

	if params.DateTo != "" {
		q.Add("dateTo", params.DateTo)
	}

	if params.Competitions != "" {
		q.Add("competition", params.Competitions)
	}

	req.URL.RawQuery = q.Encode()

	req.Header.Set("X-Auth-Token", c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var matches MatchesResponse
	err = json.Unmarshal(body, &matches)
	if err != nil {
		return nil, err
	}

	return &matches, nil
}
