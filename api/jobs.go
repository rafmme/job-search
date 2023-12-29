package api

import (
	"fmt"
	"net/http"

	"github.com/rafmme/job-search/util"
)

func FetchJobs(w http.ResponseWriter, r *http.Request) {
	resultString, _ := util.GetMyJobs("g-cse")

	w.Header().Add("Content-Type", "text/html")
	fmt.Fprintln(
		w,
		resultString,
	)

}
