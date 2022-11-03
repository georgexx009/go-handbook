package main

import (
	"fmt"
	urlshortener "github.com/georgexx009/handbook/gophercises/url-shortener/urlShortener"
	"log"
	"net/http"
)

func main() {
	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := urlshortener.MapHandler(pathsToUrls, mux)

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := urlshortener.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":8080", yamlHandler); err != nil {
		log.Println(err.Error())
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "Hello world!"); err != nil {
		log.Println(err.Error())
	}
}
