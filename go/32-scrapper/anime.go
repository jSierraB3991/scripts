package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	res, err := http.Get("https://www3.animeflv.net/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".Title").Text()
		newChapter := s.Find(".Capi").Text()
		if newChapter != "" {
			fmt.Printf("Anime %s: Capitulo Nuevo %s\n", title, newChapter)
		}
	})
}

func main() {
	ExampleScrape()
}
