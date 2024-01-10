package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
}

type SearchResult struct {
	ResulRank   int
	ResultUrl   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{
	"Mozilla/5.0 (X11; Linux x86_64; rv:85.0) Gecko/20100101 Firefox/85.0",
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	radNum := rand.Int() % len(userAgents)
	return userAgents[radNum]
}

func buildGoogleUrls(searchTerm, languageCode, countryCode string, pages, count int) ([]string, error) {
	toScrappe := []string{}
	searchTerm = strings.Trim(searchTerm, "")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)

	if domain, found := googleDomains[countryCode]; found {
		for i := 0; i < pages; i++ {
			start := i * count
			scrapperUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0", domain, searchTerm, count, languageCode, start)
			toScrappe = append(toScrappe, scrapperUrl)
		}
	} else {
		err := fmt.Errorf("You current (%s) country is not support", countryCode)
		return nil, err
	}
	return toScrappe, nil
}

func getScrapperClient(proxyString interface{}) *http.Client {
	switch v := proxyString.(type) {
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			},
		}

	default:
		return &http.Client{}
	}
}

func scrapeClientRequest(page string, proxyString interface{}) (*http.Response, error) {
	baseClient := getScrapperClient(proxyString)

	req, _ := http.NewRequest("GET", page, nil)
	req.Header.Set("User-Agent", randomUserAgent())

	res, err := baseClient.Do(req)

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Error, Scrapper recive status code not 200, suggesting to ban")
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

func googlResultParsing(res *http.Response, result int) ([]SearchResult, error) {
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	sel := doc.Find("div.g")
	result++
	results := []SearchResult{}
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		links := strings.Trim(link, "")
		titleTag := item.Find("h3.r")
		title := titleTag.Text()
		desctag := item.Find("span.st")
		desc := desctag.Text()

		if links != "" && links != "#" && !strings.HasPrefix(links, "/") {
			res := SearchResult{
				ResulRank:   result,
				ResultUrl:   links,
				ResultTitle: title,
				ResultDesc:  desc,
			}
			results = append(results, res)
			result++
		}
	}

	return results, nil
}

func GoogleScrapper(search, languageCode, countryCode string, proxyString interface{}, pages, count, backoff int) ([]SearchResult, error) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(search, languageCode, countryCode, pages, count)
	if err != nil {
		return nil, err
	}

	for _, page := range googlePages {
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}

		data, err := googlResultParsing(res, resultCounter)
		if err != nil {
			return nil, err
		}
		resultCounter += len(data)

		for _, result := range data {
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff) * time.Second)
	}

	return results, nil
}

func main() {
	response, err := GoogleScrapper("akhil sharma", "es", "com", nil, 1, 30, 10)
	if err == nil {
		for _, res := range response {
			fmt.Println(res)
		}
	}
}
