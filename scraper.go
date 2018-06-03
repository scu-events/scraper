package main

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"
	//"log"
	"net/http"
	"strings"

	"github.com/mmcdole/gofeed"

//	"golang.org/x/net/context"
//  "golang.org/x/oauth2"
 // "golang.org/x/oauth2/google"
    "google.golang.org/api/calendar/v3"
)

type DateTime struct {
	DateTime string `json:"datetime"`
	TimeZone string `json:"timeZone"`
}

func CrawlGoogleCal(url string, calendarId string) {
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

	// Throw event into database


	// printing out info
	fmt.Printf("\n O \n")
	for _, item := range data["items"] {
		item.Summary = strings.TrimSpace(item.Summary);
		item.Description = strings.TrimSpace(item.Description);	


		fmt.Println(item.Summary)
		fmt.Printf("Start Time: ")
		fmt.Println(item.Start.DateTime)
		fmt.Printf("Description:  ")
		fmt.Println(item.Description)
		fmt.Printf("id:  ")
		fmt.Println(item.Id)
		fmt.Printf("time:  ")
		fmt.Println(item.Id)

		fmt.Printf("\n")

	//	event, err = srv.Events.Insert(calendarId, item).Do()
	//	if err != nil {
  	//		log.Fatalf("Unable to create event. %v\n", err)
	//	}
		
	//	fmt.Printf("Event created: %s\n", item.HtmlLink)

	}
	return
}

func CrawlOfficialCal(url string) {

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	fmt.Println(feed.Title)
	fmt.Println()

	for _, element := range feed.Items {
		fmt.Println(element.Title)

		fmt.Printf("	")
		fmt.Printf(element.Description)
	}

	return
}

func main() {

	calendarId := "info.scuevents@gmail.com"

	//GOOGLE CAL URLS	
	//acm club: http://acm.engr.scu.edu/events
	var url_google_acm = "https://www.googleapis.com/calendar/v3/calendars/santaclara.acm@gmail.com/events?key=AIzaSyCnRyFyPuJ9WSeu602Q7CE13TsxWVNbw10&timeMin=2018-02-24T00:00:00Z&timeMax=2030-04-09T00:00:00Z&singmaxResults=9999&_=1520708172234"
	
	//offical calendar of CSO's and RSO's: https://www.scu.edu/csi/calendar/
//	var url_google_clubs = "https://clients6.google.com/calendar/v3/calendars/csl@scu.edu/events?calendarId=csl@scu.edu&singleEvents=true&timeZone=America/Los_Angeles&maxAttendees=1&maxResults=250&sanitizeHtml=true&timeMin=2018-02-26T00:00:00-08:00&timeMax=2018-04-02T00:00:00-08:00&key=AIzaSyBNlYH01_9Hc5S1J9vuFmu2nUqBZJNAXxs"


//	var officialUrl = "https://lwcal.scu.edu/live/rss/events/exclude_group/Math%20Tutors/exclude_tag/private/header/All%20Events"

	//CrawlGoogleCal(url_google_clubs)
	CrawlGoogleCal(url_google_acm, calendarId)

//	CrawlOfficialCal(officialUrl)

	t := time.Now()
	fmt.Println(t.Format(time.RFC850))

}

//Using natural language processing to determine if an event has free food
func freeFood(description string) bool{
	foodWords := [7]string{"food", "refreshments", "dinner", "lunch", "breakfast", "snacks", "drinks"}

	for _, element := range foodWords {
		if(strings.Contains(description, element)){
			return true
		}
	}

	return false
}