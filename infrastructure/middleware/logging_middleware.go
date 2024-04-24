package middleware

import (
	"github.com/filmons/go-ingnerd-backend/pkg/utils"
	"net/http"
)

// LoggingMiddleware logs the incoming HTTP request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		utils.Logger.Printf("%s %s %s", r.Method, r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
