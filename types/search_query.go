package types

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/opensaucerer/goaxios"
)

type SearchQueryData struct {
	JobSites       []string `json:"jobSites"`
	JobTitles      []string `json:"jobTitles"`
	Ignore         []string `json:"ignore"`
	Include        []string `json:"include"`
	From           int      `json:"from"`
	jobSearchQuery string
}

func (sqData *SearchQueryData) GetJobSearchQuery() string {
	return sqData.jobSearchQuery
}

// Create Job Search Query String
func (sqData *SearchQueryData) CreateJobSearchQuery() *SearchQueryData {
	from := fmt.Sprintf("after:%v", time.Now().Format(time.DateOnly))
	sites := "site:lever.co | site:greenhouse.io | site:app.dover.com | site:jobs.ashbyhq.com "
	titles := "engineer | developer | remote"
	ignores := "-\"front end\" -senior -lead -\"full stack\" -react -staff "
	includes := "nodejs "

	if len(sqData.JobSites) > 0 {
		sites = ""
		for index, site := range sqData.JobSites {
			if len(sqData.JobSites)-1 == index {
				sites += fmt.Sprintf("site:%s ", site)
			} else {
				sites += fmt.Sprintf("site:%s | ", site)
			}
		}
	}

	if len(sqData.JobTitles) > 0 {
		titles = ""
		for index, title := range sqData.JobTitles {
			if len(sqData.JobTitles)-1 == index {
				titles += title
			} else {
				titles += fmt.Sprintf("%s | ", title)
			}
		}
	}

	if len(sqData.Ignore) > 0 {
		ignores = ""
		for _, ignore := range sqData.Ignore {
			if strings.Contains(ignore, " ") {
				ignores += fmt.Sprintf("-\"%s\" ", ignore)
			} else {
				ignores += fmt.Sprintf("-%s ", ignore)
			}
		}
	}

	if len(sqData.Include) > 0 {
		includes = ""
		for _, include := range sqData.Include {
			if strings.Contains(include, " ") {
				includes += fmt.Sprintf("\"%s\" ", include)
			} else {
				includes += fmt.Sprintf("%s ", include)
			}
		}
	}

	if sqData.From > 0 {
		from = fmt.Sprintf("after:%v",
			time.Now().Add(-(time.Hour * 24 * time.Duration(sqData.From))).Format(time.DateOnly))
	}

	sqData.jobSearchQuery = fmt.Sprintf("%s(%s) %s%s%s", sites, titles, ignores, includes, from)
	return sqData
}

func (sqData *SearchQueryData) Execute() *SearchResult {
	searchQueryString := sqData.GetJobSearchQuery()
	a := goaxios.GoAxios{
		Url:    "https://www.googleapis.com/customsearch/v1",
		Method: "GET",
		Query: map[string]string{
			"key": os.Getenv("CSE_KEY"),
			"cx":  os.Getenv("CX"),
			"q":   url.QueryEscape(searchQueryString),
		},
		ResponseStruct: &SearchResult{},
	}

	_, _, d, err := a.RunRest()
	if err != nil {
		log.Printf("err: %v", err)
	}

	response := d.(*SearchResult)
	response.SearchQuery = searchQueryString

	return response
}
