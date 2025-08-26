package modem

import (
	"io"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNoBurstBehavior(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV34_33K(client) // 4200 Bps = ~4.2 bytes/ms

	go func() {
		// Send 40 bytes total (4 bytes * 10 times).
		for range 10 {
			server.Write([]byte("abcd")) // nolint:errcheck
			time.Sleep(10 * time.Millisecond)
		}
	}()

	start := time.Now()
	total := 0
	buf := make([]byte, 4)
	for range 10 {
		n, err := limited.Read(buf)
		require.NoError(t, err)
		total += n
	}
	elapsed := time.Since(start)

	assert.Equal(t, 40, total)

	// At 4200 bytes/sec, 40 bytes should take at least 40/4200 = ~9.5ms.
	// But we rate limit byte-by-byte, so it should take closer to 40ms
	// (1ms per byte at 4200 Bps = 4.2 bytes/ms means ~0.24ms per byte).
	minExpectedTime := time.Duration(40*1000/4200) * time.Millisecond // ~9.5ms
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "rate limiting not working - too fast")
}

func TestV32SymmetricTiming(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV32(client)
	payload := make([]byte, V32_BPS) // 1200 bytes = 1 sec

	// Write test.
	go func() {
		buf := make([]byte, len(payload))
		n, err := io.ReadFull(server, buf)
		assert.NoError(t, err)
		assert.Equal(t, payload, buf)
		assert.Equal(t, len(payload), n)
	}()

	start := time.Now()
	n, err := limited.Write(payload)
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)

	// At 1200 bytes/sec, 1200 bytes should take ~1 second. Allow 5%
	// tolerance for system scheduling/timing variations.
	minExpectedTime := 950 * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "upload too fast - rate limiting failed")
}

func TestV32_BISSymmetricTiming(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV32_BIS(client)
	payload := make([]byte, V32_BIS_BPS) // 1800 bytes = 1 sec

	// Write test
	go func() {
		buf := make([]byte, len(payload))
		n, err := io.ReadFull(server, buf)
		assert.NoError(t, err)
		assert.Equal(t, payload, buf)
		assert.Equal(t, len(payload), n)
	}()

	start := time.Now()
	n, err := limited.Write(payload)
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)

	// At 1800 bytes/sec, 1800 bytes should take ~1 second. Allow 5%
	// tolerance for system scheduling/timing variations.
	minExpectedTime := 950 * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "upload too fast - rate limiting failed")
}

func TestV32_TERBOSymmetricTiming(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV32_TERBO(client)
	payload := make([]byte, V32_TERBO_BPS) // 2400 bytes = 1 sec

	// Write test.
	go func() {
		buf := make([]byte, len(payload))
		n, err := io.ReadFull(server, buf)
		assert.NoError(t, err)
		assert.Equal(t, payload, buf)
		assert.Equal(t, len(payload), n)
	}()

	start := time.Now()
	n, err := limited.Write(payload)
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)

	// At 2400 bytes/sec, 2400 bytes should take ~1 second. Allow 5%
	// tolerance for system scheduling/timing variations.
	minExpectedTime := 950 * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "upload too fast - rate limiting failed")
}

func TestV34_28KSymmetricTiming(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV34_28K(client)
	payload := make([]byte, V34_28K_BPS) // 3600 bytes = 1 sec

	// Write test.
	go func() {
		buf := make([]byte, len(payload))
		n, err := io.ReadFull(server, buf)
		assert.NoError(t, err)
		assert.Equal(t, payload, buf)
		assert.Equal(t, len(payload), n)
	}()

	start := time.Now()
	n, err := limited.Write(payload)
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)

	// At 3600 bytes/sec, 3600 bytes should take ~1 second. Allow 5%
	// tolerance for system scheduling/timing variations.
	minExpectedTime := 950 * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "upload too fast - rate limiting failed")

	// Read test.
	go func() {
		_, err := server.Write(payload)
		assert.NoError(t, err)
	}()

	start = time.Now()
	buf := make([]byte, len(payload))
	n, err = io.ReadFull(limited, buf)
	elapsed = time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "download too fast - rate limiting failed")
}

func TestV34_33KUploadTiming(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV34_33K(client)

	payload := make([]byte, V34_33K_BPS) // 4200 bytes = 1 sec upload
	go func() {
		buf := make([]byte, len(payload))
		n, err := io.ReadFull(server, buf)
		assert.NoError(t, err)
		assert.Equal(t, payload, buf)
		assert.Equal(t, len(payload), n)
	}()

	start := time.Now()
	n, err := limited.Write(payload)
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)

	// At 4200 bytes/sec, 4200 bytes should take ~1 second. Allow 5%
	// tolerance for system scheduling/timing variations.
	minExpectedTime := 950 * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "upload too fast - rate limiting failed")
}

func TestV90DownloadTiming(t *testing.T) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV90(client)

	payload := make([]byte, V90_DOWN_BPS) // 7000 bytes = 1 second of downstream
	go func() {
		_, err := server.Write(payload)
		assert.NoError(t, err)
	}()

	start := time.Now()
	buf := make([]byte, len(payload))
	n, err := io.ReadFull(limited, buf)
	elapsed := time.Since(start)

	require.NoError(t, err)
	assert.Equal(t, len(payload), n)

	// At 7000 bytes/sec, 7000 bytes should take ~1 second. Allow 5%
	// tolerance for system scheduling/timing variations.
	minExpectedTime := 950 * time.Millisecond
	assert.GreaterOrEqual(t, elapsed, minExpectedTime, "download too fast - rate limiting failed")
}
