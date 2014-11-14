package rss

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	// Mon Jan 2 15:04:05 -0700 MST 2006
	UPDATE_RANGE = 1
	RSS_DATE     = "Mon, 02 Jan 2006 15:04:05 -0700"
	ATOM_DATE    = "2006-01-02T15:04"
)

type Headline struct {
    Source string
    Title string
    Link string
    Since time.Duration
}

// RSS

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	XMLName  xml.Name `xml:"channel"`
	Title    string   `xml:"title"`
	ItemList []Item   `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Date        string `xml:"pubDate"`
}

// Atom

type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Entry   []*Entry `xml:"entry"`
}

type Entry struct {
	Title   string  `xml:"title"`
	Link    Link    `xml:"link"`
	Updated string  `xml:"updated"`
	Author  *Person `xml:"author"`
}

type Link struct {
	Rel  string `xml:"rel,attr"`
	Href string `xml:"href,attr"`
	Type string `xml:"type,attr"`
}

type Person struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

func getFeed(link string) []byte {
	resp, err := http.Get(link)
	if err != nil {
		log.Println("***ERROR: http.Get", err, link)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body
}

// TODO: Add picture if any to headline (e.g., xkcd)
func GetHeadlines(link string) ([]Headline, error) {
    headlines := make([]Headline, 0)
	feed := getFeed(link)
	var i RSS

	err := xml.Unmarshal(feed, &i)
	if err == nil {
        // If it succeeds, it was an RSS feed

		for _, item := range i.Channel.ItemList {
			t, err := time.Parse(RSS_DATE, item.Date)
			if err != nil {
                return headlines, err
			}

            var headline Headline
            headline.Source = i.Channel.Title
            headline.Title = item.Title
            headline.Link = item.Link
            headline.Since = time.Since(t.Local())

            headlines = append(headlines, headline)
		}
	} else {
        // If it failed, it might just be an Atom feed

		var j Feed

		err := xml.Unmarshal(feed, &j)
		if err != nil {
            // If it failed again, you either messed up, or it was an
            // unsupported feed
            return headlines, err
		}

		for _, entry := range j.Entry {
			date := entry.Updated

			// We have to truncate the date because some atom dates are
			// different.
			t, err := time.Parse(ATOM_DATE, date[:len(ATOM_DATE)])

			if err != nil {
                return headlines, err
			}

            var headline Headline
            headline.Source = j.Title
            headline.Title = entry.Title
            headline.Link = entry.Link.Href
            headline.Since = time.Since(t.Local())
            headlines = append(headlines, headline)
		}
	}

    return headlines, nil
}
