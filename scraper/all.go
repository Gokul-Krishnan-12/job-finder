package scraper

import (
	"sync"

	"github.com/Gokul-Krishnan-12/job-finder/models"
)

func ScrapeAllJobs() []models.Job {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var allJobs []models.Job

	sources:= []func() ([]models.Job,error){
		ScrapeAPIRemoteOK,
		ScrapeWeWorkRemotely,
	}

	 wg.Add(len(sources))

	 for _, scrape := range sources {
		go func(scrapeFunc func() ([]models.Job, error)) {
			defer wg.Done()
			jobs, err := scrapeFunc()
			if err == nil {
				mu.Lock()
				allJobs = append(allJobs, jobs...)
				mu.Unlock()
			}
		}(scrape)
	}
	
	wg.Wait()
    return allJobs
}