package main

import (
    "fmt"
    "github.com/renquinn/rss/rss"
)

func main() {
    link := "http://xkcd.com/rss.xml"

    headlines, err := rss.Get(link)
    if err != nil {
        panic(err)
    }

    for _, headline := range headlines {
        fmt.Println(headline.Source, headline.Title, headline.Link, headline.Since)
    }
}
