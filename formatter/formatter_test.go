package formatter

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestExtractYear(t *testing.T) {
	if got := extractYear("2021-09-08"); got != "2021" {
		t.Errorf("Expected 2021, got %s", got)
	}
	if got := extractYear("NA"); got != "N/A" {
		t.Errorf("Expected N/A, got %s", got)
	}
}

// captureOutput intercepts stdout for testing formatters
func captureOutput(f func() error) (string, error) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String(), err
}

func TestOutputResultJSON(t *testing.T) {
	data := map[string]string{"title": "Test Movie"}
	
	output, err := captureOutput(func() error {
		return OutputResult(data, "json", "movie")
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !bytes.Contains([]byte(output), []byte(`"title": "Test Movie"`)) {
		t.Errorf("JSON output did not contain expected content. Got: %s", output)
	}
}
