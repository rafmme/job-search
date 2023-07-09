package api

import (
	"fmt"
	"net/http"

	"github.com/rafmme/job-search/util"
)

func Home(w http.ResponseWriter, r *http.Request) {
	resultString, _ := util.GetMyJobs("ser-api")

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(
		w,
		resultString,
	)
}
