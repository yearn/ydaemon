package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAPIAfterDependencyUpgrade(t *testing.T) {
	router := NewRouter()
	for _, tc := range []struct {
		name   string
		method string
		path   string
		body   string
		status int
	}{
		{"health", http.MethodGet, "/health", "", http.StatusOK},
		{"empty price query", http.MethodPost, "/prices/some", `{"addresses":""}`, http.StatusOK},
		{"unknown JSON field", http.MethodPost, "/prices/some", `{"addresses":"","unexpected":true}`, http.StatusBadRequest},
		{"malformed JSON", http.MethodPost, "/prices/some", `{"addresses":`, http.StatusBadRequest},
		{"CORS preflight", http.MethodOptions, "/prices/some", "", http.StatusNoContent},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(tc.method, tc.path, strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Origin", "https://app.yearn.fi")
			if tc.method == http.MethodOptions {
				req.Header.Set("Access-Control-Request-Method", http.MethodPost)
				req.Header.Set("Access-Control-Request-Headers", "content-type")
			}
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)
			if response.Code != tc.status {
				t.Fatalf("status = %d, want %d: %s", response.Code, tc.status, response.Body.String())
			}
			if origin := response.Header().Get("Access-Control-Allow-Origin"); origin != "*" {
				t.Errorf("public API CORS origin = %q, want *", origin)
			}
			if response.Header().Get("Access-Control-Allow-Credentials") == "true" {
				t.Error("public API must not allow cross-origin credentials")
			}
			if tc.method != http.MethodOptions && !json.Valid(response.Body.Bytes()) {
				t.Errorf("response is not JSON: %s", response.Body.String())
			}
		})
	}
}
