package myhtmlparser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseHtml(input string) ([]Link, error) {
	result := []Link{}
	f := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, key := range n.Attr {
				if key.Key == "href" {
					found := Link{
						Href: key.Val,
						Text: n.Data,
					}
					result = append(result, found)
				}
			}
		}
	}

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}
	for n := doc.FirstChild; n != nil; n = n.NextSibling {
		f(n)
	}
	return result, nil
}
