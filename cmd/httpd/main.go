package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pires/go-proxyproto"
)

var docRoot string

func main() {
	log.SetPrefix(fmt.Sprintf("%s[%d]: ", filepath.Base(os.Args[0]), os.Getpid()))
	log.SetFlags(log.Lmsgprefix)

	docRoot = os.Getenv("DOCUMENT_ROOT")
	if docRoot == "" {
		docRoot = "/var/www/htdocs"
	}

	staticHandler := http.FileServer(http.Dir(docRoot))

	http.Handle("/", staticHandler)

	// Form handlers ("CGI" style)
	http.HandleFunc("/account/", handleAccountRoutes)
	http.HandleFunc("/account/signup", handleSignup)
	http.HandleFunc("/login", handleLogin)

	http.HandleFunc("/account/change-password", handleChangePassword)

	// WebSocket proxy handler
	http.HandleFunc("/play", handleWebSocketProxy)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}

// handleAccountRoutes routes /account/* requests to appropriate handlers
func handleAccountRoutes(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/account/":
		handleLogin(w, r) // FIXME
	case "/account/signup":
		handleSignup(w, r)
	default:
		http.NotFound(w, r)
	}
}

// handleLogin processes account login form submissions
func handleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := strings.TrimSpace(r.FormValue("username"))
	_ = r.FormValue("password") // TODO: Use for authentication

	// TODO: Validate input
	// TODO: Authenticate against accounts table
	// TODO: Set session cookie or redirect

	log.Printf("Login attempt: username=%q", username)

	// Placeholder response
	fmt.Fprintf(w, `<html><body>
		<h1>Account Login</h1>
		<p>Login processing not yet implemented</p>
		<p>Username: %s</p>
		<p><a href="/">Back to home</a></p>
	</body></html>`, html.EscapeString(username))
}

// handleSignup handles account signup - GET shows form, POST processes it
func handleSignup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// Process signup.html as template
		templatePath := filepath.Join(docRoot, "account", "signup.html")

		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			log.Printf("Error parsing template %s: %v", templatePath, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Execute template with empty data for GET requests
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Error executing template %s: %v", templatePath, err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := strings.TrimSpace(r.FormValue("username"))
	_ = r.FormValue("password") // TODO: Use for hashing
	email := strings.TrimSpace(r.FormValue("email"))

	// TODO: Validate input
	// TODO: Check if username already exists
	// TODO: Use auth.PasswordHash() to hash password
	// TODO: INSERT INTO accounts table
	// TODO: Handle success/error responses

	log.Printf("Signup attempt: username=%q, email=%q", username, email)

	// Placeholder response
	fmt.Fprintf(w, `<html><body>
		<h1>Account Signup</h1>
		<p>Signup processing not yet implemented</p>
		<p>Username: %s</p>
		<p>Email: %s</p>
		<p><a href="/">Back to home</a></p>
	</body></html>`, html.EscapeString(username), html.EscapeString(email))
}

// handleChangePassword processes password change form submissions
func handleChangePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	username := strings.TrimSpace(r.FormValue("username"))
	_ = r.FormValue("old_password") // TODO: Use for authentication
	_ = r.FormValue("new_password") // TODO: Use for hashing

	// TODO: Validate input
	// TODO: Authenticate current password
	// TODO: Hash new password with auth.PasswordHash()
	// TODO: UPDATE accounts table

	log.Printf("Password change attempt: username=%q", username)

	// Placeholder response
	fmt.Fprintf(w, `<html><body>
		<h1>Change Password</h1>
		<p>Password change processing not yet implemented</p>
		<p>Username: %s</p>
		<p><a href="/">Back to home</a></p>
	</body></html>`, html.EscapeString(username))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

// handleWebSocketProxy upgrades HTTP connections to WebSocket and proxies to localhost:23
func handleWebSocketProxy(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

	// Get client IP address from Fly-Client-IP header
	clientIP := r.Header.Get("Fly-Client-IP")
	var clientPort int
	if clientIP == "" {
		// Fall back to RemoteAddr
		if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
			clientIP = ip
			// clientPort = strconv.Atoi(port) -- TODO
		} else {
			clientIP = r.RemoteAddr
		}
	}
	log.Printf("WebSocket connection from %q", clientIP)

	// Connect to localhost:23 with proxy protocol
	telnetConn, err := net.Dial("tcp", "localhost:23")
	if err != nil {
		log.Printf("Failed to connect to localhost:23: %v", err)
		return
	}
	defer telnetConn.Close()

	// Send proxy protocol header
	header := &proxyproto.Header{
		Version:           1,
		Command:           proxyproto.PROXY,
		TransportProtocol: proxyproto.TCPv4,
		SourceAddr: &net.TCPAddr{
			IP:   net.ParseIP(clientIP),
			Port: clientPort,
		},
		DestinationAddr: &net.TCPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 23,
		},
	}
	if _, err := header.WriteTo(telnetConn); err != nil {
		log.Printf("Failed to write proxy protocol header: %v", err)
		return
	}

	// Start bidirectional proxy
	done := make(chan struct{})
	var once sync.Once

	// WebSocket -> Telnet
	go func() {
		defer once.Do(func() { close(done) })
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("WebSocket read error: %v", err)
				}
				return
			}
			if _, err := telnetConn.Write(message); err != nil {
				log.Printf("Telnet write error: %v", err)
				return
			}
		}
	}()

	// Telnet -> WebSocket
	go func() {
		defer once.Do(func() { close(done) })
		buffer := make([]byte, 1024)
		for {
			n, err := telnetConn.Read(buffer)
			if err != nil {
				if err != io.EOF {
					log.Printf("Telnet read error: %v", err)
				}
				return
			}
			if err := conn.WriteMessage(websocket.BinaryMessage, buffer[:n]); err != nil {
				log.Printf("WebSocket write error: %v", err)
				return
			}
		}
	}()

	// Wait for either direction to close
	<-done
	log.Printf("WebSocket proxy connection closed for %q", clientIP)
}
