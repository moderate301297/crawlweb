package crawlproduct

import (
	"golang.org/x/net/html"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"../../../../savedata"
)

func OptimizeHref(href string) string{
	if strings.Index(href, "http") == -1 {
		if string(href[0]) == "/" && string(href[1]) == "/" {
			href = "https:" + href
		}
		href = "https://photos3.walmart.com" + href
	}	
	return href
}

func GetSrc(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "src" {
			href = a.Val
		}
	}
	return
}

func QueryBody(url string) {
	var title, link, linkImage string
	var linkSlice []string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error: ", err)
		savedata.SaveUrlError(url)
		return
	}
	doc.Find("div.product-page").Each(func(i int, s *goquery.Selection) {
		title = s.Find("h1").Text()
	})
	doc.Find("div.product-page ol.breadcrumb li").Each(func(i int, s *goquery.Selection) {
		linkSlice = append(linkSlice, s.Find("a").Text())
	})
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		linkSlice = append(linkSlice, s.Text())
	})
	for _, k := range linkSlice {
		if k != "" {
			link = link + k + "/"
		}
	}
	doc.Find("div.product-image.undefined img").Each(func(i int, s *goquery.Selection) {
		for _, k := range s.Nodes {
			href := GetSrc(k)
			linkImage = OptimizeHref(href)
		}
	})
	savedata.SaveData(title, link, linkImage)
}