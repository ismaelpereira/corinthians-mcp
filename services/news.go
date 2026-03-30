package services

import (
	"sort"

	clients "github.com/ismaelpereira/corinthians-mcp/clients/rss"
	"github.com/ismaelpereira/corinthians-mcp/types"
)

type NewsService struct {
	RSSClient *clients.RSSClient
}

func NewNewsService(rssClient *clients.RSSClient) *NewsService {
	return &NewsService{
		RSSClient: rssClient,
	}
}

func (s *NewsService) FetchNews(limit int) ([]types.New, error) {
	var offset int
	if limit > 0 {
		offset = limit
	} else {
		offset = 10
	}

	meuTimaoNews, err := s.RSSClient.FetchMeuTimaoNews()
	if err != nil {
		return nil, err
	}

	centralDoTimaoNews, err := s.RSSClient.FetchCentralDoTimaoNews()
	if err != nil {
		return nil, err
	}
	var news []types.New
	news = append(news, types.MapToDomainNews(meuTimaoNews)...)
	news = append(news, types.MapToDomainNews(centralDoTimaoNews)...)

	sort.Slice(news, func(i, j int) bool {
		return news[i].PublishedAt.After(news[j].PublishedAt)
	})

	return news[:offset], nil
}
