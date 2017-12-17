package crawlurl

import (
	"fmt"
	"golang.org/x/net/html"
	"github.com/PuerkitoBio/goquery"
)

// find rel in token
func GetRel(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "rel" {
			href = a.Val
		}
	}
	return
}

// get href in token
func GetHref(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}
	return
}

// find next page
func NextPage(href string) (url string) {
	doc, err := goquery.NewDocument(href)
	if err != nil {
		fmt.Println("Error: ", err)
		return url
	}
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		for _, k := range s.Nodes {
			if GetRel(k) == "next" {
				url = GetHref(k)
			}
		}
	})
	return url
}