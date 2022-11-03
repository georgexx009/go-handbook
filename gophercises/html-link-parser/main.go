package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {
	filePath := "./example/link/ex1.html"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	var links []Link

	var f func(node *html.Node)
	f = func(node *html.Node) {
		var link Link

		if node.Type == html.ElementNode && node.Data == "a" {
			for _, attr := range node.Attr {
				if attr.Key == "href" {
					// fmt.Println(attr.Val)
					link.Href = attr.Val
					break
				}
			}

			if node.FirstChild.Type == html.TextNode {
				// fmt.Println(node.Data)
				link.Text = node.Data
			}
		}

		if link.Text != "" && link.Href != "" {
			links = append(links, link)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	fmt.Println("------------------------------")
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	docStr, _ := html.Parse(strings.NewReader(s))
	f(docStr)

	fmt.Println(links)
}
