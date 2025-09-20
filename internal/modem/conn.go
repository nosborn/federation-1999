package modem

import (
	"context"
	"io"
	"net"

	"golang.org/x/time/rate"
)

// Modem speeds (bytes/sec, based on bits/sec ÷ 8).
// Accounts for 1990s PPP overhead: ~11% reduction from theoretical speeds.
// PPP frame overhead: 7 bytes per frame (4 header + 3 trailer).
// Byte stuffing: ~1.5% additional overhead from escaping 0x7D/0x7E.
const (
	V32_BPS       = 1070 // 9.6Kbps / 8 * 0.89 (PPP efficiency)
	V32_BIS_BPS   = 1600 // 14.4Kbps / 8 * 0.89
	V32_TERBO_BPS = 2140 // 19.2Kbps / 8 * 0.89

	V34_28K_BPS = 3200 // 28.8Kbps / 8 * 0.89
	V34_33K_BPS = 3740 // 33.6Kbps / 8 * 0.89

	V90_DOWN_BPS = 6230 // 56Kbps / 8 * 0.89 (download)
	V90_UP_BPS   = 3740 // 33.6Kbps / 8 * 0.89 (upload)
)

// NewFixedRate wraps conn with strict read/write bandwidth limits.
// No burst allowed: max throughput never exceeds configured rate.
func NewFixedRate(conn net.Conn, downloadBps, uploadBps int) net.Conn {
	return NewFixedRateWithContext(context.Background(), conn, downloadBps, uploadBps)
}

// NewFixedRateWithContext wraps conn with strict read/write bandwidth limits
// and context support.
// No burst allowed: max throughput never exceeds configured rate.
func NewFixedRateWithContext(ctx context.Context, conn net.Conn, downloadBps, uploadBps int) net.Conn {
	if ctx == nil {
		ctx = context.Background()
	}
	if downloadBps <= 0 {
		downloadBps = 1 // Minimum 1 byte/sec to avoid division by zero
	}
	if uploadBps <= 0 {
		uploadBps = 1 // Minimum 1 byte/sec to avoid division by zero
	}

	readLimiter := rate.NewLimiter(rate.Limit(downloadBps), 1)
	writeLimiter := rate.NewLimiter(rate.Limit(uploadBps), 1)

	// Drain initial token to ensure first byte is properly rate-limited.
	// Token bucket starts "initially full" so we need to consume the free
	// token.
	_ = readLimiter.Wait(ctx)
	_ = writeLimiter.Wait(ctx)

	return &rateLimitedConn{
		Conn: conn,
		reader: &rateLimitedReader{
			r:       conn,
			limiter: readLimiter,
		},
		writer: &rateLimitedWriter{
			w:       conn,
			limiter: writeLimiter,
		},
	}
}

// NewV32 returns a conn limited to 9.6Kbps up/down.
func NewV32(conn net.Conn) net.Conn {
	return NewFixedRate(conn, V32_BPS, V32_BPS)
}

// NewV32_BIS returns a conn limited to 14.4Kbps up/down.
func NewV32_BIS(conn net.Conn) net.Conn {
	return NewFixedRate(conn, V32_BIS_BPS, V32_BIS_BPS)
}

// NewV32_TERBO returns a conn limited to 19.2Kbps up/down.
func NewV32_TERBO(conn net.Conn) net.Conn {
	return NewFixedRate(conn, V32_TERBO_BPS, V32_TERBO_BPS)
}

// NewV34_28K returns a conn limited to 28.8Kbps up/down.
func NewV34_28K(conn net.Conn) net.Conn {
	return NewFixedRate(conn, V34_28K_BPS, V34_28K_BPS)
}

// NewV34_33K returns a conn limited to 33.6Kbps up/down.
func NewV34_33K(conn net.Conn) net.Conn {
	return NewFixedRate(conn, V34_33K_BPS, V34_33K_BPS)
}

// NewV90 returns a conn limited to 56Kbps down / 33.6Kbps up.
func NewV90(conn net.Conn) net.Conn {
	return NewFixedRate(conn, V90_DOWN_BPS, V90_UP_BPS)
}

type rateLimitedConn struct {
	net.Conn
	reader io.Reader
	writer io.Writer
}

func (c *rateLimitedConn) Read(b []byte) (int, error) {
	return c.reader.Read(b)
}

func (c *rateLimitedConn) Write(b []byte) (int, error) {
	return c.writer.Write(b)
}

type rateLimitedReader struct {
	r       io.Reader
	limiter *rate.Limiter
}

func (r *rateLimitedReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	total := 0
	for total < len(p) {
		// Wait for permission to read exactly 1 byte (no burst
		// allowed).
		if err := r.limiter.Wait(context.Background()); err != nil {
			return total, err
		}

		// Read exactly 1 byte.
		n, err := r.r.Read(p[total : total+1])
		total += n
		if err != nil {
			return total, err
		}
		if n == 0 {
			// No more data available.
			break
		}
	}
	return total, nil
}

type rateLimitedWriter struct {
	w       io.Writer
	limiter *rate.Limiter
}

func (w *rateLimitedWriter) Write(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	total := 0
	for total < len(p) {
		// Wait for permission to write exactly 1 byte (no burst
		// allowed).
		if err := w.limiter.Wait(context.Background()); err != nil {
			return total, err
		}

		// Write exactly 1 byte.
		n, err := w.w.Write(p[total : total+1])
		total += n
		if err != nil {
			return total, err
		}
		if n == 0 {
			// This shouldn't happen for writes, but handle it.
			break
		}
	}
	return total, nil
}
