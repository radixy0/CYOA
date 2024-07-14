package main

import (
	"flag"
	"fmt"
	"net/http"
	link "sitemap"
)

func getLinksFromUrl(url string) ([]string, error) {
	//request html
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	urls, err := link.Parse(response.Body)
	if err != nil {
		return nil, err
	}
	result := []string{}

	for _, url := range urls {
		result = append(result, url.Href)
	}

	return result, nil
}

func main() {
	//get url from param
	flag.Parse()
	args := flag.Args()
	//validate args?
	url := args[0]
	//get links from url
	urlSet := map[string]bool{}
	urlsToCheck := []string{url}

	for {
		if len(urlsToCheck) == 0 {
			break
		}
		newUrls := []string{url}
		for _, url := range urlsToCheck {
			links, err := getLinksFromUrl(url)
			if err != nil {
				fmt.Println("oopsie", err)
				continue
			}
			for _, link := range links {
				_, ok := urlSet[link]
				if ok {
					continue
				}
				urlSet[link] = true
				newUrls = append(newUrls, link)
			}
		}
		urlsToCheck = newUrls
	}

	//parse to xml
	fmt.Println(urlSet)
}
