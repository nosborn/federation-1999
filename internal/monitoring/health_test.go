package monitoring

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HealthHandler(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	// Check required fields
	require.Equal(t, "ok", response["status"])
	require.NotNil(t, response["timestamp"])
	require.NotNil(t, response["uptime"])

	// Check uptime is reasonable (should be a duration string)
	uptime, ok := response["uptime"].(string)
	require.True(t, ok, "uptime should be a string")
	require.Contains(t, uptime, "s", "uptime should contain seconds")
}

func TestHealthResponseFormat(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	beforeTime := time.Now()
	HealthHandler(w, req)
	afterTime := time.Now()

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	require.NoError(t, err)

	// Parse timestamp
	timestampStr, ok := response["timestamp"].(string)
	require.True(t, ok, "timestamp should be a string")

	timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
	require.NoError(t, err, "timestamp should be valid RFC3339")

	// Timestamp should be between before and after
	require.True(t, timestamp.After(beforeTime) || timestamp.Equal(beforeTime))
	require.True(t, timestamp.Before(afterTime) || timestamp.Equal(afterTime))
}
