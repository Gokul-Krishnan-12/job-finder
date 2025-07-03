package main

import (
	"log"
	"net/http"
)

func main()  {
	 http.HandleFunc("/", jobListHandler)
    http.HandleFunc("/jobs", jobHandler)
	http.HandleFunc("/api/jobs",jobApiHandler)

	log.Println("Server started at port :9000")
    log.Fatal(http.ListenAndServe(":9000",nil))
}