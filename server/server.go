package server

import (
	"github.com/ismaelpereira/corinthians-mcp/config"
	mcp "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

type MCPServer struct {
	Server *mcp.Server
	Done   <-chan struct{}
}

func NewMCPServer(config *config.Config) *MCPServer {
	transport := stdio.NewStdioServerTransport()
	server := mcp.NewServer(transport)

	return &MCPServer{
		Server: server,
		Done:   make(chan struct{}),
	}
}
