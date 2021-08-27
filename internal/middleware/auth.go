package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/danisbagus/mini-pos-app/internal/core/port"
	"github.com/danisbagus/mini-pos-app/pkg/errs"
)

type AuthMiddleware struct {
	Repo port.IAuthRepo
}

func (rc AuthMiddleware) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader != "" {
				token := getTokenFromHeader(authHeader)
				err := rc.Repo.Verify(token)
				if err != nil {
					writeResponse(w, err.Code, err.AsMessage())
				} else {
					next.ServeHTTP(w, r)
				}

			} else {
				err := errs.NewAuthorizationError("Missing token!")
				writeResponse(w, err.Code, err.AsMessage())
			}
		})
	}
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
