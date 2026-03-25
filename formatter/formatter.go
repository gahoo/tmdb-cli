package formatter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// OutputResult formats any data into JSON, Markdown, or NFO format
func OutputResult(w io.Writer, data interface{}, format string, itemType string, fields string) error {
	if fields != "" {
		filteredData, err := FilterData(data, fields)
		if err != nil {
			return err
		}
		data = filteredData
		
		switch format {
		case "json":
			outJSON, err := json.MarshalIndent(data, "", "  ")
			if err != nil {
				return err
			}
			fmt.Fprintln(w, string(outJSON))
			return nil
		case "markdown":
			return printDynamicMarkdown(w, data)
		case "table":
			return printDynamicTable(w, data)
		case "nfo":
			return printDynamicNFO(w, data)
		}
	}

	switch format {
	case "json":
		return printJSON(w, data)
	case "markdown":
		return printMarkdown(w, data, itemType)
	case "table":
		return printTable(w, data, itemType)
	case "nfo":
		return printNFO(w, data, itemType)
	default:
		return fmt.Errorf("unsupported format: %s", format)
	}
}

func printJSON(w io.Writer, data interface{}) error {
	outJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	fmt.Fprintln(w, string(outJSON))
	return nil
}

// OutputResultToFileOrStdout handles file creation if outputFile is set, then delegates to OutputResult
func OutputResultToFileOrStdout(outputFile string, data interface{}, format string, itemType string, fields string) error {
	var out io.Writer = os.Stdout
	var file *os.File

	if outputFile != "" {
		var errFile error
		file, errFile = os.Create(outputFile)
		if errFile != nil {
			return fmt.Errorf("error creating output file: %v", errFile)
		}
		defer file.Close()
		out = file
	}

	err := OutputResult(out, data, format, itemType, fields)
	if err != nil {
		return err
	}

	if file != nil {
		fmt.Printf("Results exported to %s\n", outputFile)
	}
	return nil
}

// Very basic Markdown format for arbitrary maps/structs or SearchResultPage

