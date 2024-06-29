package main

import (
	"cyoa"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	templ, err := template.ParseFiles("cmd/webapp/chapter.html")
	if err != nil {
		log.Fatal(err)
		return
	}

	tpl = template.Must(templ, err)
}

type HandlerOption func(h *handler)

func WithLogging() HandlerOption {
	return func(h *handler) {
		fmt.Println("logging option chosen")
	}
}

func NewHandler(s cyoa.Story, opts ...HandlerOption) http.Handler {
	h := handler{s}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s cyoa.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wront", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)

}

func main() {
	port := flag.Int("port", 3000, "port to run on")
	filename := flag.String("file", "story.json", "file containing cyoa")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(file)
	if err != nil {
		panic(err)
	}

	h := NewHandler(story, WithLogging())
	fmt.Printf("starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", *port), h))
}
