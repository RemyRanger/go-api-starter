package middleware

import (
	"net/http"
	"time"
)

// HealthCheckMiddleware is a middleware that intercepts all requests
// and handles the /healthcheck endpoint to return a simple health status.
func HealthCheckMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the /health endpoint
		if r.URL.Path == "/healthcheck" {
			// Respond with a 200 OK and a simple JSON indicating the service is up
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"UP", "time":"` + time.Now().Format(time.RFC3339) + `"}`))
			return
		}

		// If the request is not for /healthcheck, pass it to the next handler
		next.ServeHTTP(w, r)
	})
}
