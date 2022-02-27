package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	var SetApplicationJSON = func(wr http.ResponseWriter, req *http.Request) {
		wr.Header().Set("Content-type", "application/json")
		next.ServeHTTP(wr, req)
	}
	return http.HandlerFunc(SetApplicationJSON)
}
