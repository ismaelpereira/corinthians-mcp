# Corinthians MCP Server

A Model Context Protocol (MCP) server that provides information about the Sport Club Corinthians Paulista football team, including news, team details, match schedules, and historical data.

## Features

- **Latest News**: Fetch recent news from Meu Timão and Central do Timão RSS feeds.
- **Team Information**: Get details about the current Corinthians team, including players and coach.
- **Match Information**: Retrieve upcoming and past match details.
- **Historical Data**: Access historical facts, idol players, stadium info, and titles.

## APIs Used

This project integrates with the following APIs and data sources:

- **Meu Timão RSS Feed**: `https://www.meutimao.com.br/rss` - For fetching latest news articles.
- **Central do Timão RSS Feed**: `https://www.centraldotimao.com.br/feed/` - Additional news source.
- **Football Data API**: `https://api.football-data.org` - For team information and match data. Requires an API key.
- **Local Data**: `history.json` - Static file containing historical data about the team.

## Project Structure

```
corinthians-mcp/
├── main.go                 # Entry point of the application
├── go.mod                  # Go module file
├── go.sum                  # Go dependencies checksum
├── .env                    # Environment variables (not committed)
├── history.json            # Local historical data
├── mcp-example.json        # Example MCP server configuration
├── prompt.json             # MCP prompt configuration
├── clients/                # API client implementations
│   ├── football/           # Football Data API client
│   ├── local/              # Local data client
│   └── rss/                # RSS feed client
├── config/                 # Configuration management
├── server/                 # MCP server setup
├── services/               # Business logic services
├── tools/                  # MCP tool definitions
└── types/                  # Data type definitions
```

## Prerequisites

- Go 1.25.7 or later
- API key for Football Data API (sign up at https://www.football-data.org/)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/ismaelpereira/corinthians-mcp.git
   cd corinthians-mcp
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory with the following variables:

   ```
   MEU_TIMAO_FEED_URL=https://www.meutimao.com.br/rss
   CENTRAL_DO_TIMAO_FEED_URL=https://www.centraldotimao.com.br/feed/
   MATCHES_API_URL=https://api.football-data.org
   MATCHES_API_KEY=your_football_data_api_key_here
   HISTORY_FILE_PATH=./history.json
   ```

   Replace `your_football_data_api_key_here` with your actual API key from Football Data.

## Running the MCP Server

1. Build the server:

   ```bash
   go build -o corinthians-mcp main.go
   ```

2. Run the server:
   ```bash
   ./corinthians-mcp
   ```

The server will start and communicate via stdin/stdout using the MCP protocol.

## Configuration in VS Code

To use this MCP server in VS Code, add the following configuration to your VS Code settings (replace paths as needed):

```json
{
   "mcpServers": {
      "corinthians-mcp": {
         "command": "/path/to/binary",
         "args": [],
         "env": {
            "MEU_TIMAO_FEED_URL": "https://www.meutimao.com.br/rss",
            "CENTRAL_DO_TIMAO_FEED_URL": "https://www.centraldotimao.com.br/feed/",
            "MATCHES_API_KEY": "your_api_key",
            "MATCHES_API_URL": "https://api.football-data.org",
            "HISTORY_FILE_PATH": "/path/to/history.json"
         }
      }
   }
}
```

## Available Tools

The MCP server provides the following tools:

- `get_last_news`: Fetches the latest news (default 10 items, configurable limit)
- `get_team_info`: Retrieves current team information
- `get_matches_info`: Gets match information with optional filters
- `get_data_history`: Fetches historical team data

## Development

To contribute or modify the server:

1. Make changes to the code
2. Run tests (if any)
3. Build and test the server locally
4. Submit a pull request

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Future Features

- Show more Players statistics and Profilling tool
- Check transfermarket to see the possible players coming and going
- Publish this as a Gitlab Application and have this Public
- Add the data from the different modalities (Basketball, Feminine Soccer, Futsal, Sub-20 etc.)