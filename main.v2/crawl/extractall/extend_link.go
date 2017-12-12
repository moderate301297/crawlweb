package extractall

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
// find next page
func ScaleLinkText(href string) (a string) {
	doc, err := goquery.NewDocument(href)
	if err != nil {
		fmt.Println("Error: ", err)
		return a
	}
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		for _, k := range s.Nodes {
			if GetRel(k) == "next" {
				a = GetHref(k)
			}
		}
	})
	return a
}