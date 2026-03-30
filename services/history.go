package services

import (
	clients "github.com/ismaelpereira/corinthians-mcp/clients/local"
	"github.com/ismaelpereira/corinthians-mcp/types"
)

type DataHistoryService struct {
	LocalClient *clients.LocalClient
}

func NewDataHistoryService(localClient *clients.LocalClient) *DataHistoryService {
	return &DataHistoryService{
		LocalClient: localClient,
	}
}

func (s *DataHistoryService) FetchDataHistory() (*types.HistoryFacts, error) {
	return s.LocalClient.FetchHistory()
}
