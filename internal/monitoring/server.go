package monitoring

import (
	"log"
	"net/http"
	"time"
)

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", Handler())
	mux.HandleFunc("/health", HealthHandler)

	newServer := func(addr string) *http.Server {
		return &http.Server{
			Addr:              addr,
			Handler:           mux,
			IdleTimeout:       30 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      5 * time.Second,
		}
	}

	go func() {
		// #nosec G102 -- Binding to all IPv4 addresses is required for Fly.io health checks
		if err := newServer(port).ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("monitoring server error:", err)
		}
	}()
}
