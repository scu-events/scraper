package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	//"log"
	"net/http"
	"strings"

	//	"golang.org/x/net/context"
	//  "golang.org/x/oauth2"
	// "golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

//DateTime struct
type DateTime struct {
	DateTime string `json:"datetime"`
	TimeZone string `json:"timeZone"`
}

//CrawlGoogleCal gathers data from google calendars
func CrawlGoogleCal(url string, calendarID string) {
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

	var data map[string][]calendar.Event
	json.Unmarshal(body, &data)

	// printing out info
	fmt.Printf("\n O \n")
	for _, item := range data["items"] {
		item.Summary = strings.TrimSpace(item.Summary)
		item.Description = strings.TrimSpace(item.Description)

		fmt.Println(item.Summary)
		fmt.Printf("Start Time: ")
		fmt.Println(item.Start.DateTime)
		fmt.Printf("Description:  ")
		fmt.Println(item.Description)
		fmt.Printf("id:  ")
		fmt.Println(item.Id)
		fmt.Printf("time:  ")
		fmt.Println(item.Id) //?

		fmt.Printf("\n")

		//	event, err = srv.Events.Insert(calendarId, item).Do()
		//	if err != nil {
		//		log.Fatalf("Unable to create event. %v\n", err)
		//	}

		//	fmt.Printf("Event created: %s\n", item.HtmlLink)

	}
	return
}

func main() {

	//Host Calendar
	calendarID := "scuhackers@gmail.com"

	//Array of Google Calendars to Scrape
	var googleCals [2]string

	//ACM Club Calendar
	googleCals[0] = "https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"
	//Official CSO/RSO Calendar
	googleCals[1] = "https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=AIzaSyBNlYH01_9Hc5S1J9vuFmu2nUqBZJNAXxs"

	//Crawl through Google Calendars
	for _, item := range googleCals {
		CrawlGoogleCal(item, calendarID)
	}

	t := time.Now()
	fmt.Println(t.Format(time.RFC850))

}

//Using natural language processing to determine if an event has free food
func freeFood(description string) bool {
	foodWords := [7]string{"food", "refreshments", "dinner", "lunch", "breakfast", "snacks", "drinks"}

	for _, element := range foodWords {
		if strings.Contains(description, element) {
			return true
		}
	}

	return false
}
