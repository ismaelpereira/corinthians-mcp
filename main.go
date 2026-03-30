package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ismaelpereira/corinthians-mcp/config"
	"github.com/ismaelpereira/corinthians-mcp/server"
	"github.com/ismaelpereira/corinthians-mcp/tools"
)

func main() {
	config := config.LoadConfig()

	app := server.NewMCPServer(config)

	tools.GenerateTools(config, app.Server)

	err := app.Server.Serve()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(os.Stderr, "MCP server started")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case <-ctx.Done():
	case <-app.Done:
	}

}
