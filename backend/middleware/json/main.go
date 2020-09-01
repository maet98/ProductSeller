package ContentType

import "net/http"

func middleware(next http.Handler) http.Handler {
	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
