package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Gokul-Krishnan-12/job-finder/models"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeRemoteOK() ([]models.Job,error)  {

	var jobs []models.Job

	res,err := http.Get("https://remoteok.com/")

	if err != nil{
		return nil,err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
        return nil, err
    }

	doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }
	doc.Find("tr[id^='job-']").Each(func(i int, s *goquery.Selection) {
        title := strings.TrimSpace(s.Find("h2").Text())
        company := strings.TrimSpace(s.Find(".companyLink").Text())
        location := s.Find(".location").Text()
        href, _ := s.Find("a.preventLink").Attr("href")
        tags := []string{}

        s.Find(".tag").Each(func(i int, tag *goquery.Selection) {
            tags = append(tags, strings.TrimSpace(tag.Text()))
        })

        if title != "" && href != "" {
            job := models.Job{
                Title:    title,
                Company:  company,
                Location: location,
                URL:      "https://remoteok.com" + href,
                Source:   "remoteok",
                Tags:     tags,
            }
            jobs = append(jobs, job)
        }
    })
	 return jobs, nil
	
}

func ScrapeAPIRemoteOK() ([]models.Job, error) {
    res, err := http.Get("https://remoteok.com/api")
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    var apiData []map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&apiData); err != nil {
        return nil, err
    }

    var jobs []models.Job
    for _, item := range apiData {
        // Skip metadata (the first item is metadata)
        if _, ok := item["position"]; !ok {
            continue
        }

        job := models.Job{
            Title:    fmt.Sprintf("%v", item["position"]),
            Company:  fmt.Sprintf("%v", item["company"]),
            Location: fmt.Sprintf("%v", item["location"]),
            URL:      fmt.Sprintf("%v", item["url"]),
            Source:   "remoteok",
        }

        // Handle tags (array of strings)
        if tags, ok := item["tags"].([]interface{}); ok {
            for _, tag := range tags {
                job.Tags = append(job.Tags, fmt.Sprintf("%v", tag))
            }
        }

        jobs = append(jobs, job)
    }

    return jobs, nil
}
