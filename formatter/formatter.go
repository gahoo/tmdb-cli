package formatter

import (
	"encoding/json"
	"fmt"
)

// OutputResult formats any data into JSON, Markdown, or NFO format
func OutputResult(data interface{}, format string, itemType string) error {
	switch format {
	case "json":
		return printJSON(data)
	case "markdown":
		return printMarkdown(data, itemType)
	case "table":
		return printTable(data, itemType)
	case "nfo":
		return printNFO(data, itemType)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func printJSON(data interface{}) error {
	out, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

// Very basic Markdown format for arbitrary maps/structs or SearchResultPage

