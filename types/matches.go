package types

import (
	clients "github.com/ismaelpereira/corinthians-mcp/clients/football"
)

type Matches struct {
	Matches   []Match   `json:"matches"`
	ResultSet ResultSet `json:"result_set"`
}

type Match struct {
	Competition string `json:"competition"`
	Date        string `json:"date"`
	HomeTeam    string `json:"home_team"`
	AwayTeam    string `json:"away_team"`
	Status      string `json:"status"`
	Winner      string `json:"winner"`
	Score       Score  `json:"score"`
}

type ResultSet struct {
	ScheduledAmount   int    `json:"scheduled_amount"`
	Competitions      string `json:"competitions"`
	FirstGamePLayedAt string `json:"first_game_played_at"`
	LastGamePLayedAt  string `json:"last_game_played_at"`
	PlayedTotal       int    `json:"played_total"`
	WinsTotal         int    `json:"wins_total"`
	DrawsTotal        int    `json:"draws_total"`
	LossesTotal       int    `json:"losses_total"`
}

type Score struct {
	Duration string    `json:"duration"`
	FullTime MatchTime `json:"fullTime"`
	HalfTime MatchTime `json:"halfTime"`
}

type MatchTime struct {
	Home int `json:"home"`
	Away int `json:"away"`
}

func MapToDomainMatches(matches *clients.MatchesResponse) *Matches {
	return &Matches{
		Matches:   mapToDomainMatch(matches.Matches),
		ResultSet: mapToDomainResultSet(matches.ResultSet),
	}
}

func mapToDomainMatch(matches []clients.Match) []Match {
	var domainMatches []Match
	for _, match := range matches {
		domainMatches = append(domainMatches, Match{
			Competition: match.Competition.Name,
			Date:        match.UtcDate,
			HomeTeam:    match.HomeTeam.Name,
			AwayTeam:    match.AwayTeam.Name,
			Status:      match.Status,
			Winner:      match.Score.Winner,
			Score:       mapToDomainScore(match.Score),
		})
	}
	return domainMatches
}

func mapToDomainScore(score clients.Score) Score {
	return Score{
		Duration: score.Duration,
		FullTime: mapToDomainMatchTime(score.FullTime),
		HalfTime: mapToDomainMatchTime(score.HalfTime),
	}
}

func mapToDomainMatchTime(matchTime clients.MatchTime) MatchTime {
	return MatchTime{
		Home: matchTime.Home,
		Away: matchTime.Away,
	}
}

func mapToDomainResultSet(resultSet clients.ResultSet) ResultSet {
	return ResultSet{
		ScheduledAmount:   resultSet.Count,
		Competitions:      resultSet.Competitios,
		FirstGamePLayedAt: resultSet.First,
		LastGamePLayedAt:  resultSet.Last,
		PlayedTotal:       resultSet.Played,
		WinsTotal:         resultSet.Wins,
		DrawsTotal:        resultSet.Draws,
		LossesTotal:       resultSet.Losses,
	}
}
