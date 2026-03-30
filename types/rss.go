package types

import (
	"time"

	clients "github.com/ismaelpereira/corinthians-mcp/clients/rss"
)

type New struct {
	Title       string    `json:"title"`
	Link        string    `json:"link"`
	Description string    `json:"description"`
	Content     string    `json:"content"`
	PublishedAt time.Time `json:"published_at"`
}

func MapToDomainNews(newsItems []clients.NewsItem) []New {
	var news []New
	for _, newsItem := range newsItems {
		news = append(news, mapToDomainItem(newsItem))
	}
	return news
}

func mapToDomainItem(newsItem clients.NewsItem) New {
	return New{
		Title:       newsItem.Title,
		Link:        newsItem.Link,
		Description: newsItem.Description,
		Content:     newsItem.Content,
		PublishedAt: newsItem.Published,
	}
}
