package cors

import "net/http"

// CORS is a middleware that adds CORS headers to the response.
func CORS(options *cors) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodOptions && r.Header.Get("Access-Control-Request-Method") != "" {
				options.preflight(w, r)
				if options.optionsPassthrough {
					h.ServeHTTP(w, r)
				} else {
					w.WriteHeader(options.optionsSuccessStatus)
				}
			} else {
				options.request(w, r)
				h.ServeHTTP(w, r)
			}
		})
	}
}
