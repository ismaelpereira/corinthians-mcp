package clients

import (
	"encoding/json"
	"io"
	"os"

	"github.com/ismaelpereira/corinthians-mcp/types"
)

type LocalClient struct {
	FilePath string
}

func NewLocalClient(filePath string) *LocalClient {
	return &LocalClient{
		FilePath: filePath,
	}
}

func (c *LocalClient) FetchHistory() (*types.HistoryFacts, error) {
	jsonFile, err := os.Open(c.FilePath)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	fileData, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var facts types.HistoryFacts
	err = json.Unmarshal(fileData, &facts)
	if err != nil {
		return nil, err
	}

	return &facts, nil
}
