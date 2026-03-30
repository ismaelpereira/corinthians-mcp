package services

import (
	"fmt"

	clients "github.com/ismaelpereira/corinthians-mcp/clients/football"
	"github.com/ismaelpereira/corinthians-mcp/types"
)

type TeamService struct {
	FootballClient *clients.FootballClient
}

func NewTeamService(footballClient *clients.FootballClient) *TeamService {
	return &TeamService{
		FootballClient: footballClient,
	}
}

func (s *TeamService) FetchTeamInfo() (*types.Team, error) {
	teamInfo, err := s.FootballClient.FetchTeamInfo()
	if err != nil {
		return nil, err
	}
	return types.MapToDomainTeam(teamInfo), nil

}

func (s *TeamService) FetchMatches(params clients.MatchParams) (*types.Matches, error) {
	if params.Status != "" && !params.Status.IsValid() {
		return nil, fmt.Errorf("invalid param status: %s", params.Status)
	}
	if params.Venue != "" && !params.Venue.IsValid() {
		return nil, fmt.Errorf("invalid param venue: %s", params.Venue)
	}

	matches, err := s.FootballClient.FetchMatches(params)
	if err != nil {
		return nil, err
	}
	return types.MapToDomainMatches(matches), nil

}
