package types

import clients "github.com/ismaelpereira/corinthians-mcp/clients/football"

type Team struct {
	Name                string        `json:"name"`
	ShortName           string        `json:"short_name"`
	TLA                 string        `json:"tla"`
	Logo                string        `json:"logo"`
	Country             string        `json:"country"`
	Founded             int           `json:"founded"`
	Address             string        `json:"address"`
	Website             string        `json:"website"`
	ClubColors          string        `json:"club_colors"`
	RunningCompetitions []Competition `json:"running_competitions"`
	Coach               Coach         `json:"coach"`
	Squad               []Player      `json:"squad"`
	LastUpdated         string        `json:"last_updated"`
}

type Competition struct {
	Name   string `json:"name"`
	Code   string `json:"code"`
	Emblem string `json:"emblem"`
}

type Coach struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	Name              string `json:"name"`
	DateOfBirth       string `json:"date_of_birth"`
	Nationality       string `json:"nationality"`
	ContractStartedAt string `json:"contract_started_at"`
	ContractEndAt     string `json:"contract_end_at"`
}

type Player struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	DateOfBirth string `json:"date_of_birth"`
	Nationality string `json:"nationality"`
}

func MapToDomainTeam(teamResponse *clients.TeamResponse) *Team {
	return &Team{
		Name:                teamResponse.Name,
		ShortName:           teamResponse.ShortName,
		TLA:                 teamResponse.TLA,
		Logo:                teamResponse.Logo,
		Country:             teamResponse.Area.Name,
		Founded:             teamResponse.Founded,
		Address:             teamResponse.Address,
		Website:             teamResponse.Website,
		ClubColors:          teamResponse.ClubColors,
		RunningCompetitions: mapToDomainCompetitions(teamResponse.RunningCompetitions),
		Coach:               mapToDomainCoach(teamResponse.Coach),
		Squad:               mapToDomainSquad(teamResponse.Squad),
		LastUpdated:         teamResponse.LastUpdated,
	}
}

func mapToDomainCompetitions(competitions []clients.Competition) []Competition {
	var domainCompetitions []Competition
	for _, competition := range competitions {
		domainCompetitions = append(domainCompetitions, Competition{
			Name:   competition.Name,
			Code:   competition.Code,
			Emblem: competition.Emblem,
		})
	}
	return domainCompetitions
}

func mapToDomainCoach(coach clients.Coach) Coach {
	return Coach{
		FirstName:         coach.FirstName,
		LastName:          coach.LastName,
		Name:              coach.Name,
		DateOfBirth:       coach.DateOfBirth,
		Nationality:       coach.Nationality,
		ContractStartedAt: coach.Contract.Start,
		ContractEndAt:     coach.Contract.Until,
	}
}

func mapToDomainSquad(squad []clients.Player) []Player {
	var domainSquad []Player
	for _, player := range squad {
		domainSquad = append(domainSquad, Player{
			Name:        player.Name,
			Position:    player.Position,
			DateOfBirth: player.DateOfBirth,
			Nationality: player.Nationality,
		})
	}
	return domainSquad
}
