package api

import (
	"net/http"
	"strings"

	"github.com/opensaucerer/barf"
	"github.com/rafmme/job-search/types"
)

func SearchJobs(w http.ResponseWriter, r *http.Request) {
	var data types.SearchQueryData
	var result *types.SearchResult

	err := barf.Request(r).Body().Format(&data)

	if err != nil {
		barf.Response(w).Status(http.StatusBadRequest).JSON(barf.Res{
			Status:  false,
			Data:    nil,
			Message: "Invalid request body",
		})
		return
	}

	if strings.ToLower(data.Mode) == "g-cse" {
		result = data.CreateJobSearchQuery().Execute(data.Mode).FormatJobList()
	} else {
		result = data.CreateJobSearchQuery().Execute(data.Mode)
	}

	barf.Response(w).Status(http.StatusOK).JSON(barf.Res{
		Status:  true,
		Message: "Job Search Results",
		Data:    result,
	})
}
