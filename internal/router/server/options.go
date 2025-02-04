package server

import (
	"net/http"
	"strings"
)

func optionsHandler(methods []string, extraAllowedHeaders []string) http.HandlerFunc {
	methodString := strings.Join(methods, ",") + ",OPTIONS"
	accessControlAllowHeaders := "Content-Type,Cache-Control,Authorization,X-Amz-Date,X-Api-Key,X-Amz-Security-Token,X-Requested-With"
	if len(extraAllowedHeaders) > 0 {
		accessControlAllowHeaders += "," + strings.Join(extraAllowedHeaders, ",")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			origin = "*"
		}
		w.Header().Add("Access-Control-Allow-Headers", accessControlAllowHeaders)
		w.Header().Add("Access-Control-Allow-Methods", methodString)
		w.Header().Add("Access-Control-Allow-Origin", origin)
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Max-Age", "600")
		w.WriteHeader(http.StatusOK)
	})
}
