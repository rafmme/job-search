package types

import (
	"fmt"
	"strings"
)

type Job struct {
	Title   string `json:"title"`
	Snippet string `json:"snippet"`
	Link    string `json:"link"`
	Pagemap struct {
		Metatags []map[string]string `json:"metatags"`
	} `json:"pagemap"`
}

type FormatedJob struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	Location    string `json:"location"`
}

type SearchResult struct {
	Items             []Job             `json:"items"`
	Jobs              []FormatedJob     `json:"jobs"`
	SearchInformation map[string]string `json:"searchInformation"`
	Url               map[string]string `json:"url"`
	SearchQuery       string            `json:"searchQuery"`
}

func (s *SearchResult) FormatJobList() *SearchResult {
	jobs := []FormatedJob{}

	for _, job := range s.Items {
		jobUrl := strings.ToLower(job.Link)

		if !strings.Contains(jobUrl, "linkedin.com") ||
			strings.Contains(jobUrl, "linkedin.com/jobs") {
			jobs = append(jobs, FormatedJob{
				Title:       job.Title,
				Description: job.Snippet,
				Url:         jobUrl,
				Location:    job.Pagemap.Metatags[0]["og:description"],
			})
		}
	}

	s.Jobs = jobs
	s.Items = nil
	return s
}

func createHTML(jobData FormatedJob) string {
	return fmt.Sprintf(`<div></div><div>
	</div><div>
    <p><strong>Job Title:</strong>  %s </p>
	<p><strong>Job Desc:</strong>  %s </p>
	<p><strong>Job Link:</strong>  <a href="%s">%s</a> </p>
    <p><strong>Location:</strong>  %s </p>
  </div>`, jobData.Title, jobData.Description, jobData.Url, jobData.Url, jobData.Location)
}

func createTGMessage(jobData FormatedJob) string {
	return strings.ReplaceAll(fmt.Sprintf(
		`%s\n%s\n%s\n\n`,
		jobData.Title, jobData.Description, jobData.Url,
	), `\n`, "\n")
}

func (s *SearchResult) CreateResultString() (emailHTMLString, tgMessageString string) {
	emailHTMLString = fmt.Sprintf(`<div><h3>%s</h3>`, strings.ToUpper(s.SearchQuery))

	for _, jobData := range s.Jobs {
		emailHTMLString += createHTML(jobData)

		// Implement sending many results with Ser API here
		tgMessageString += createTGMessage(jobData)
	}

	return emailHTMLString + `</div>`, tgMessageString
}
