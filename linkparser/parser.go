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
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n == nil {
			return
		}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, key := range n.Attr {
				if key.Key == "href" {
					found := Link{
						Href: key.Val,
						//TODO use dfs to get text
						Text: n.Val,
					}
					result = append(result, found)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		panic(err)
	}
	f(doc)
	return result, nil
}
