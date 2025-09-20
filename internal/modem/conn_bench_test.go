package modem

import (
	"net"
	"testing"
)

func BenchmarkV32_Write(b *testing.B) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV32(client)
	data := []byte("Hello, World!")

	go func() {
		buf := make([]byte, 1024)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	b.ResetTimer()
	for b.Loop() {
		limited.Write(data) // nolint:errcheck
	}
}

func BenchmarkV32_BIS_Write(b *testing.B) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV32_BIS(client)
	data := []byte("Hello, World!")

	go func() {
		buf := make([]byte, 1024)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	b.ResetTimer()
	for b.Loop() {
		limited.Write(data) // nolint:errcheck
	}
}

func BenchmarkV32_TERBO_Write(b *testing.B) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV32_TERBO(client)
	data := []byte("Hello, World!")

	go func() {
		buf := make([]byte, 1024)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	b.ResetTimer()
	for b.Loop() {
		limited.Write(data) // nolint:errcheck
	}
}

func BenchmarkV34_28K_Read(b *testing.B) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV34_28K(client)
	data := []byte("Hello, World!")

	go func() {
		for {
			_, err := server.Write(data)
			if err != nil {
				return
			}
		}
	}()

	buf := make([]byte, len(data))
	b.ResetTimer()
	for b.Loop() {
		limited.Read(buf) // nolint:errcheck
	}
}

func BenchmarkV34_28K_Write(b *testing.B) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV34_28K(client)
	data := []byte("Hello, World!")

	go func() {
		buf := make([]byte, 1024)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	b.ResetTimer()
	for b.Loop() {
		limited.Write(data) // nolint:errcheck
	}
}

func BenchmarkV90_Write(b *testing.B) {
	client, server := net.Pipe()
	defer client.Close() // nolint:errcheck
	defer server.Close() // nolint:errcheck

	limited := NewV90(client)
	data := []byte("Hello, World!")

	go func() {
		buf := make([]byte, 1024)
		for {
			_, err := server.Read(buf)
			if err != nil {
				return
			}
		}
	}()

	b.ResetTimer()
	for b.Loop() {
		limited.Write(data) // nolint:errcheck
	}
}
