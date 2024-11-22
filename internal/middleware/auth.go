package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"github.com/livghit/go-auth-spinup/internal/api"
)

// What is this middleware suppoused to do ?
// This auth middleware is suppoused to check if the
// request has the "Bearer": token Header
// if so the protected route can be accessed
// else it returns a 401

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jwt := r.Header.Get("Authorization")
		slog.Info("The jwt recived is -> %s", jwt)

		// look if there is a jwt token .
		if jwt == "" {
			authFailed(w)
			return
		}
		// if there is a jwt token look if its the form Bearer: token
		valid := validateJWT(jwt)

		if valid == false {
			authFailed(w)
			return
		}

		slog.Info("Checking auth...")
		next.ServeHTTP(w, r)
	})
}

func authFailed(w http.ResponseWriter) {
	data := map[string]interface{}{
		"error": "Authorization failed",
	}
	slog.Info("Auth failed")
	w.WriteHeader(http.StatusUnauthorized)
	api.WriteJSON(data, w)
}

func validateJWT(jwt string) bool {
	if jwt != "" {
		// assure token has prefix of Bearer
		token := strings.Split(jwt, " ")[1]
		slog.Info(fmt.Sprintf("token is -> %s", token))

		// implement logic for token validation

		return true
	}
	return false
}
