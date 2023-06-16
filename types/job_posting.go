package types

type Job struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DatePosted  string `json:"datePosted"`
	Url         string `json:"url"`
}
