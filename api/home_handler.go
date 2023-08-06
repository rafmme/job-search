package api

import (
	"net/http"

	"github.com/opensaucerer/barf"
)

func Home(w http.ResponseWriter, r *http.Request) {
	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Message: "Server Up!",
		Data:    nil,
	})
}
