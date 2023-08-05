package api

import (
	"net/http"
)

// Serve the Swagger UI HTML page
func SwaggerAPIDocHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/swagger/index.html", http.StatusFound)
}
