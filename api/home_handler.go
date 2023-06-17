package api

import (
	"net/http"

	"github.com/opensaucerer/barf"
)

func Home(w http.ResponseWriter, r *http.Request) {
	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Data:    nil,
		Message: "Job Search API Home",
	})
}
