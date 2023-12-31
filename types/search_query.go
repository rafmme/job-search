package types

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/opensaucerer/goaxios"
	search "github.com/serpapi/google-search-results-golang"
)

type SearchQueryData struct {
	JobSites       []string `json:"jobSites"`
	JobTitles      []string `json:"jobTitles"`
	Ignore         []string `json:"ignore"`
	Include        []string `json:"include"`
	From           int      `json:"from"`
	Mode           string   `json:"mode"`
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

func (sqData *SearchQueryData) Execute(mode string) *SearchResult {
	if strings.ToLower(mode) == "g-cse" {
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
			var result SearchResult
			searchUrl := fmt.Sprintf(
				"https://www.googleapis.com/customsearch/v1?cx=%v&key=%v&q=%v",
				os.Getenv("CX"),
				os.Getenv("CSE_KEY"),
				url.QueryEscape(searchQueryString),
			)

			resp, err := http.Get(searchUrl)
			if err != nil {
				log.Printf("err: %v", err)
				return new(SearchResult)
			}

			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)

			if err := json.Unmarshal(body, &result); err != nil {
				log.Printf("err: %v", err)
				return new(SearchResult)
			}

			if err != nil {
				log.Printf("err: %v", err)
				return new(SearchResult)
			}

			response := &result
			response.SearchQuery = searchQueryString
			return response
		}

		response := d.(*SearchResult)
		response.SearchQuery = searchQueryString
		return response
	}

	searchQueryString := sqData.GetJobSearchQuery()
	parameter := map[string]string{
		"q":             searchQueryString,
		"hl":            "en",
		"gl":            "us",
		"google_domain": "google.com",
		"num":           "100",
		"output":        "json",
	}

	search := search.NewGoogleSearch(parameter, os.Getenv("SERAPI_KEY"))
	result, err := search.GetJSON()

	if err != nil {
		log.Printf("err: %v", err)
		return new(SearchResult)
	}

	jobsList := []FormatedJob{}

	for _, data := range result["organic_results"].([]interface{}) {
		jobData := data.(map[string]interface{})
		jobUrl := strings.ToLower(jobData["link"].(string))

		if strings.Contains(jobUrl, "linkedin.com") &&
			!strings.Contains(jobUrl, "linkedin.com/jobs") ||
			strings.Contains(jobUrl, "indeed.com") &&
				!strings.Contains(jobUrl, "indeed.com/viewjob") {
			continue

		}

		jobsList = append(jobsList, FormatedJob{
			Url:         jobUrl,
			Description: jobData["snippet"].(string),
			Title:       jobData["title"].(string),
		})
	}

	searchResult := new(SearchResult)
	searchResult.SearchQuery = searchQueryString
	searchResult.Jobs = jobsList

	return searchResult
}
