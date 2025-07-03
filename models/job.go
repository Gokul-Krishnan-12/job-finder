package models

type Job struct {
    Title    string   `json:"title"`
    Company  string   `json:"company"`
    Location string   `json:"location"`
    URL      string   `json:"url"`
    Source   string   `json:"source"`
    Tags     []string `json:"tags,omitempty"`
}