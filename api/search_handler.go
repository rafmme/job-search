package api

import (
	"net/http"

	"github.com/opensaucerer/barf"
	"github.com/rafmme/job-search/types"
)

func SearchJobs(w http.ResponseWriter, r *http.Request) {
	var data types.SearchQueryData

	err := barf.Request(r).Body().Format(&data)

	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "Invalid request body",
		})
		return
	}

	result := data.CreateJobSearchQuery().Execute("g-cse").FormatJobList()

	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Message: "Job Search Results",
		Data:    result,
	})
}
