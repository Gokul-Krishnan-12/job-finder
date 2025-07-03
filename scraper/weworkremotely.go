package scraper

import (
	"net/http"
	"strings"

	"github.com/Gokul-Krishnan-12/job-finder/models"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeWeWorkRemotely()([]models.Job,error)  {
	var jobs []models.Job

	res, err := http.Get("https://weworkremotely.com/remote-jobs")
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.StatusCode != 200 {
        return nil, err
    }

	 doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
        return nil, err
    }

	 doc.Find(".jobs li.feature").Each(func(i int, s *goquery.Selection) {
        title := strings.TrimSpace(s.Find("span.title").Text())
        company := strings.TrimSpace(s.Find("span.company").Text())
        location := strings.TrimSpace(s.Find("span.region").Text())
        href, _ := s.Find("a").Attr("href")

        if title != "" && href != "" {
            job := models.Job{
                Title:    title,
                Company:  company,
                Location: location,
                URL:      "https://weworkremotely.com" + href,
                Source:   "weworkremotely",
            }
            jobs = append(jobs, job)
        }
    })
	return jobs, nil
}