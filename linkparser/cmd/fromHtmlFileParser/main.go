package main

import (
	"fmt"
	"myhtmlparser"
	"os"
)

var htmlContent string

func init() {
	byteContent, err := os.ReadFile("ex1.html")
	if err != nil {
		panic(err)
	}
	htmlContent = string(byteContent)
}

func main() {
	links, err := myhtmlparser.ParseHtml(htmlContent)
	if err != nil {
		panic(err)
	}
	fmt.Println(links)
}
