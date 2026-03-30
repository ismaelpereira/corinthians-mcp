package tools

import (
	"encoding/json"
	"fmt"
	"os"

	football "github.com/ismaelpereira/corinthians-mcp/clients/football"
	historyData "github.com/ismaelpereira/corinthians-mcp/clients/local"
	rss "github.com/ismaelpereira/corinthians-mcp/clients/rss"
	"github.com/ismaelpereira/corinthians-mcp/config"
	"github.com/ismaelpereira/corinthians-mcp/services"
	mcp "github.com/metoro-io/mcp-golang"
)

type GetNewsParams struct {
	Limit int `json:"limit,omitempty" jsonschema:"description=Number of news items to return,default=10"`
}

type EmptyParams struct{}

func GenerateTools(config *config.Config, server *mcp.Server) {
	// define clients
	rssClient := rss.NewRSSClient(config.MeuTimaoFeedURL, config.CentralDoTimaoFeedURL)
	footballClient := football.NewFootballClient(config.MatchesAPIURL, config.MatchesAPIKey)
	historyDataClient := historyData.NewLocalClient(config.HistoryFilePath)
	// define services

	newsService := services.NewNewsService(rssClient)
	teamService := services.NewTeamService(footballClient)
	dataHistoryService := services.NewDataHistoryService(historyDataClient)

	server.RegisterTool(
		"get_last_news",
		"Fetches the latest news about Corinthians from Meu Timão and Central do Timão. Default return 10 last news, input a number to send more",
		func(params GetNewsParams) (*mcp.ToolResponse, error) {
			fmt.Fprintln(os.Stderr, "tool called")
			news, err := newsService.FetchNews(params.Limit)
			if err != nil {
				return nil, err
			}

			jsonBytes, err := json.Marshal(news)
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResponse(mcp.NewTextContent(string(jsonBytes))), nil
		},
	)

	server.RegisterTool(
		"get_team_info",
		"Fetches information about the Corinthians team, including players and coach details.",
		func(_ EmptyParams) (*mcp.ToolResponse, error) {
			teamInfo, err := teamService.FetchTeamInfo()
			if err != nil {
				return nil, err
			}

			jsonBytes, err := json.Marshal(teamInfo)
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResponse(mcp.NewTextContent(string(jsonBytes))), nil
		})

	server.RegisterTool(
		"get_matches_info",
		"Fetches information about upcoming and past matches.",
		func(params football.MatchParams) (*mcp.ToolResponse, error) {
			matches, err := teamService.FetchMatches(params)
			if err != nil {
				return nil, err
			}

			jsonBytes, err := json.Marshal(matches)
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResponse(mcp.NewTextContent(string(jsonBytes))), nil
		})

	server.RegisterTool(
		"get_data_history",
		"Fetches historical data about the Corinthians team. Such as idol players, historical facts, stadium info, titles",
		func(_ EmptyParams) (*mcp.ToolResponse, error) {
			historyFacts, err := dataHistoryService.FetchDataHistory()
			if err != nil {
				return nil, err
			}

			jsonBytes, err := json.Marshal(historyFacts)
			if err != nil {
				return nil, err
			}
			return mcp.NewToolResponse(mcp.NewTextContent(string(jsonBytes))), nil
		})

}
