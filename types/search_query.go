package types

type SearchQueryData struct {
	JobSites  []string `json:"jobSites"`
	JobTitles []string `json:"jobTitles"`
	Ignore    []string `json:"ignore"`
	Include   []string `json:"include"`
	From      string   `json:"from"`
}
