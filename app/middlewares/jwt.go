package middlewares

import (
	"encoding/json"
	"go-commerce/app/utils"
	"net/http"
	"strings"
)

func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var header = r.Header.Get("x-access-token")

		json.NewEncoder(w).Encode(r)
		header = strings.TrimSpace(header)
		if header == "" {
			utils.JsonResponse(w, "Access token required!", 203)
			return;
		}

		next.ServeHTTP(w, r)
	})
}
