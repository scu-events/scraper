package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type DateTime struct {
	DateTime string `json:"datetime"`
	TimeZone string `json:"timeZone"`
}

type Event struct {
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Location    string   `json:"location"`
	Start       DateTime `json:"start"`
	End         DateTime `json:"end"`
}

func Crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var data map[string][]Event
	json.Unmarshal(body, &data)

	fmt.Printf("\nData \n")
	for _, item := range data["items"] {
		item.Summary = strings.TrimSpace(item.Summary);
		item.Description = strings.TrimSpace(item.Description);		

		//fmt.Printf("Summary:\n")
		fmt.Println(item.Summary)
		fmt.Printf("Description:\n")
		fmt.Println(item.Description)
		fmt.Printf("Start Time:\n")
		fmt.Printf(item.Start.DateTime)
	}
	return
}

func main() {

	urls := []string{
					"https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234", 
					"https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=AIzaSyBNlYH01_9Hc5S1J9vuFmu2nUqBZJNAXxs" }

	for _, url := range urls {
		Crawl(url)
	}
}
