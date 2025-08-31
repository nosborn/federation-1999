package monitoring

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestMetricsHandler(t *testing.T) {
	// Initialize some metrics first so they appear in output
	ConnectionsTotal.WithLabelValues("test").Add(0)
	ConnectionsCurrent.WithLabelValues("test").Set(0)

	// Create test request
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	w := httptest.NewRecorder()

	// Call handler
	Handler().ServeHTTP(w, req)

	// Check response
	require.Equal(t, http.StatusOK, w.Code)
	require.Contains(t, w.Header().Get("Content-Type"), "text/plain")

	body := w.Body.String()

	// Should contain our metrics
	require.Contains(t, body, "connections_total")
	require.Contains(t, body, "connections_current")
	require.Contains(t, body, "uptime_seconds")
}

func TestConnectionMetrics(t *testing.T) {
	// Reset metrics for clean test
	ConnectionsTotal.Reset()
	ConnectionsCurrent.Reset()

	// Increment counters
	ConnectionsTotal.WithLabelValues("test").Inc()
	ConnectionsCurrent.WithLabelValues("test").Inc()

	// Get metrics
	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	w := httptest.NewRecorder()
	Handler().ServeHTTP(w, req)

	body := w.Body.String()

	// Should show our incremented values
	require.Contains(t, body, `connections_total{service="test"} 1`)
	require.Contains(t, body, `connections_current{service="test"} 1`)
}

func TestMetricsRegistration(t *testing.T) {
	// Test that metrics are properly registered
	metricFamilies, err := prometheus.DefaultGatherer.Gather()
	require.NoError(t, err)

	var foundConnections, foundUptime bool
	for _, mf := range metricFamilies {
		switch mf.GetName() {
		case "connections_total", "connections_current":
			foundConnections = true
		case "uptime_seconds":
			foundUptime = true
		}
	}

	require.True(t, foundConnections, "connection metrics should be registered")
	require.True(t, foundUptime, "uptime metric should be registered")
}
