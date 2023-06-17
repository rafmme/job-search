package types

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
}

func (s *SearchResult) FormatJobList() *SearchResult {
	jobs := []FormatedJob{}

	for _, job := range s.Items {
		jobs = append(jobs, FormatedJob{
			Title:       job.Title,
			Description: job.Snippet,
			Url:         job.Link,
			Location:    job.Pagemap.Metatags[0]["og:description"],
		})
	}

	s.Jobs = jobs
	s.Items = nil
	return s
}
