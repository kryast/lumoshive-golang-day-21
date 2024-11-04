package middleware

import (
	"day-21/model"
	"encoding/json"
	"net/http"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("token")
		if authHeader != "12345" {
			badResponse := model.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "Unauthorized",
				Data:       nil,
			}
			json.NewEncoder(w).Encode(badResponse)
			return
		}

		next.ServeHTTP(w, r)
	})
}
