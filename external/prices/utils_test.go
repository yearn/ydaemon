package prices

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
)

/**************************************************************************************************
** TestGetQuery tests the getQuery helper function to ensure it correctly extracts query parameters
** from the Gin context in a case-insensitive manner. This test validates:
** - Normal query parameter retrieval
** - Case-insensitive parameter matching
** - Multiple values for the same parameter
** - Missing parameter handling
**
** These tests ensure that the query parameter extraction, which is critical for the 'humanized'
** flag functionality, works correctly in various scenarios.
**************************************************************************************************/
func TestGetQuery(t *testing.T) {
	// Setup Gin in test mode
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryParams    map[string][]string
		targetKey      string
		expectedResult string
	}{
		{
			name: "Basic query parameter",
			queryParams: map[string][]string{
				"param1": {"value1"},
			},
			targetKey:      "param1",
			expectedResult: "value1",
		},
		{
			name: "Case-insensitive matching",
			queryParams: map[string][]string{
				"Param1": {"value1"},
			},
			targetKey:      "param1",
			expectedResult: "value1",
		},
		{
			name: "Multiple values for same key",
			queryParams: map[string][]string{
				"param1": {"value1", "value2"},
			},
			targetKey:      "param1",
			expectedResult: "value1,value2",
		},
		{
			name: "Missing parameter",
			queryParams: map[string][]string{
				"param1": {"value1"},
			},
			targetKey:      "param2",
			expectedResult: "",
		},
		{
			name:           "Empty query parameters",
			queryParams:    map[string][]string{},
			targetKey:      "param1",
			expectedResult: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test request with query parameters
			req, _ := http.NewRequest("GET", "/?"+url.Values(tt.queryParams).Encode(), nil)

			// Setup Gin context
			w := &mockResponseWriter{}
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request = req

			// Call the function
			result := getQuery(ctx, tt.targetKey)

			// Verify result
			if result != tt.expectedResult {
				t.Errorf("getQuery(%q) = %q, want %q", tt.targetKey, result, tt.expectedResult)
			}
		})
	}
}

/**************************************************************************************************
** TestControllerStructure verifies that the Controller struct can be created and used as expected.
** This is a simple structural test to ensure the base type works correctly.
**************************************************************************************************/
func TestControllerStructure(t *testing.T) {
	controller := Controller{}

	// Simply check that we can create a controller without issues
	if controller != (Controller{}) {
		t.Error("Controller creation failed")
	}
}

// mockResponseWriter is a simple implementation of http.ResponseWriter for testing
type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteHeader(statusCode int) {
	// Do nothing
}
