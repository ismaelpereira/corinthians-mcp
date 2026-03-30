package clients

import (
	"time"

	"github.com/k3a/html2text"
	"github.com/mmcdole/gofeed"
)

type RSSClient struct {
	MeuTimaoFeedURL       string
	CentralDoTimaoFeedURL string
	parser                *gofeed.Parser
}

func NewRSSClient(meuTimaoFeedURL string, centralDoTimaoFeedURL string) *RSSClient {
	return &RSSClient{
		MeuTimaoFeedURL:       meuTimaoFeedURL,
		CentralDoTimaoFeedURL: centralDoTimaoFeedURL,
		parser:                gofeed.NewParser(),
	}
}

func (c *RSSClient) FetchMeuTimaoNews() ([]NewsItem, error) {
	return c.fetchNews(c.MeuTimaoFeedURL)
}

func (c *RSSClient) FetchCentralDoTimaoNews() ([]NewsItem, error) {
	return c.fetchNews(c.CentralDoTimaoFeedURL)
}

func (c *RSSClient) fetchNews(feedURL string) ([]NewsItem, error) {
	feed, err := c.parser.ParseURL(feedURL)
	if err != nil {
		return nil, err
	}

	var news []NewsItem
	for _, item := range feed.Items {
		plainContent := html2text.HTML2Text(item.Content)
		publishedAt, err := time.Parse(time.RFC1123Z, item.Published)
		if err != nil {
			return nil, err
		}
		news = append(news, NewsItem{
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			Content:     plainContent,
			Published:   publishedAt,
		})
	}
	return news, nil
}
