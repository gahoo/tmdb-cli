package api

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	client := NewClient("dummy_token", "")
	if client.Token != "dummy_token" {
		t.Errorf("Expected token to be dummy_token, got %s", client.Token)
	}
	if client.HTTPClient == nil {
		t.Error("HTTP client should not be nil")
	}
}

func TestNewRequest(t *testing.T) {
	client := NewClient("dummy_token", "")
	req, err := client.newRequest("GET", "/test")
	if err != nil {
		t.Errorf("did not expect error, got %v", err)
	}

	if req.Method != "GET" {
		t.Errorf("Expected GET method, got %s", req.Method)
	}
	if req.Header.Get("Authorization") != "Bearer dummy_token" {
		t.Errorf("Expected Bearer dummy_token auth header, got %s", req.Header.Get("Authorization"))
	}
}
