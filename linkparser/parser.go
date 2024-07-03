package myhtmlparser

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func text(n *html.Node) string {
	if n.Type == html.CommentNode {
		return ""
	}
	result := ""
	if n.Type == html.ElementNode {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result += text(c)
		}
	} else {
		result += n.Data
	}
	return strings.Trim(result, " ")
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
						Text: text(n),
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
