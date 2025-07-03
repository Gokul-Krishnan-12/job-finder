package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Gokul-Krishnan-12/job-finder/models"
	"github.com/Gokul-Krishnan-12/job-finder/scraper"
)

func jobListHandler(w http.ResponseWriter, r *http.Request) {
    // Parse query parameters
    source := r.URL.Query().Get("source")
    location := r.URL.Query().Get("location")
    tag := r.URL.Query().Get("tag")

    jobs := scraper.ScrapeAllJobs()

    // Filter
    var filtered []models.Job
    for _, job := range jobs {
        if source != "" && job.Source != source {
            continue
        }
        if location != "" && !strings.Contains(strings.ToLower(job.Location), strings.ToLower(location)) {
            continue
        }
        if tag != "" {
            match := false
            for _, t := range job.Tags {
                if strings.EqualFold(t, tag) {
                    match = true
                    break
                }
            }
            if !match {
                continue
            }
        }
        filtered = append(filtered, job)
    }

    json.NewEncoder(w).Encode(filtered)
}


func jobHandler(w http.ResponseWriter , r *http.Request){

	jobs,err := scraper.ScrapeRemoteOK()

	if err != nil {
        http.Error(w, "Failed to scrape jobs", http.StatusInternalServerError)
        return
    }

	json.NewEncoder(w).Encode(jobs)
}

func jobApiHandler(w http.ResponseWriter , r *http.Request){

	jobs,err := scraper.ScrapeAPIRemoteOK()

	if err != nil {
        http.Error(w, "Failed to scrape jobs", http.StatusInternalServerError)
        return
    }

	json.NewEncoder(w).Encode(jobs)
}
