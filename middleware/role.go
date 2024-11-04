package middleware

import (
	"day-21/model"
	"encoding/json"
	"net/http"
)

func RoleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userRole := r.Header.Get("role")
		if userRole != "user" {
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
