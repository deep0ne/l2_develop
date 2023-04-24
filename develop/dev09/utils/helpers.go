package utils

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// Функция для поиска ссылок на другие файлы
func GetLinks(baseURL string, body io.Reader) ([]string, error) {
	links := make([]string, 0)
	doc, err := html.Parse(body)
	if err != nil {
		return nil, fmt.Errorf("Error parsing HTML: %v", err)
	}
	// https://stackoverflow.com/questions/29318672/parsing-list-items-from-html-with-go
	var search func(*html.Node)
	search = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && !strings.HasPrefix(attr.Val, "mailto:") {
					links = append(links, attr.Val)
				}
			}
		}
		if n.Type == html.ElementNode && n.Data == "<img" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		}

		if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "link") {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					links = append(links, attr.Val)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			search(c)
		}
	}
	search(doc)

	linkSet := make(map[string]struct{})
	for _, link := range links {
		if !strings.HasPrefix(link, "#") && !strings.HasPrefix(link, "../") {
			linkSet[link] = struct{}{}
		}
	}
	outLinks := make([]string, 0, len(linkSet))
	for link := range linkSet {
		outLinks = append(outLinks, link)
	}

	return outLinks, nil
}

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}
