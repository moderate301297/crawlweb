package main

import (
	"fmt"
	"strings"

	"./crawl"
	"github.com/PuerkitoBio/goquery"
	"github.com/buger/jsonparser"
)

func OptimizeUrl(value string) (url string) {
	if strings.Index(value, "http") == 0 {
		url = value
		return url
	} else {
		url = "https://www.walmart.com" + value
		return url
	}
}

func main() {
	var body string
	doc, err := goquery.NewDocument("https://www.walmart.com/all-departments")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	doc.Find("head script").Each(func(i int, s *goquery.Selection) {
		var band string
		band = s.Text()
		check := strings.Index(band, "_setReduxState") != -1
		if check {
			index := strings.Index(band, "= {")
			for i := 0; i < index+1; i++ {
				band = strings.Replace(band, string(band[i]), " ", 1)
			}
			for i := 0; i < len(band)-3; i++ {
				body = body + string(band[i])
			}

		}
	})
	urlMap := make(map[string]bool)
	data := []byte(body)
	shopCategory, _, _, _ := jsonparser.Get(data, "header", "quimbyData", "global_header", "headerZone3", "configs", "departments", "[0]")
	level1, _, _, _ := jsonparser.Get(shopCategory, "departments")
	var urlSlice []string
	jsonparser.ArrayEach(level1, func(value2 []byte, dataType jsonparser.ValueType, offset int, err error) {
		level2, _, _, _ := jsonparser.Get(value2, "department")
		value, _, _, _ := jsonparser.Get(level2, "clickThrough", "value")
		url := OptimizeUrl(string(value))
		_, check := urlMap[url]
		if !check {
			urlMap[url] = true
			urlSlice = append(urlSlice, url)
		}
	})
	crawl.Crawl(urlSlice, urlMap)
}
