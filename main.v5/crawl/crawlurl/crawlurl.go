package crawlurl

import (
	"time"
	"strings"
	"github.com/buger/jsonparser"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"./crawlproduct"
	"../../savedata"
)

func OptimizeUrl(value string) (url string) {
	if (strings.Index(value, "http") == 0){
		url = value
		return url
	} else {
		url = "https://www.walmart.com" + value
		return url
	}
}

func CrawlUrl(url string, urlMap map[string]bool) (urlMapNew map[string]bool){
	if url == "" {
		return urlMap
	}
	check := strings.Index(url, "/browse/") != -1
	if check {
		for {
			go crawlproduct.QueryBody(url)
			time.Sleep(time.Millisecond * 1)
			// next page
			urlNext := NextPage(url)
			if urlNext == "" {
				break
			}
			urlMap[urlNext] = true
			url = urlNext
		}
	} else {
		var body string
		doc, err := goquery.NewDocument(url)
		if err != nil {
			fmt.Println("Error: ", err)
			savedata.SaveUrlError(url)
			return
		}
		doc.Find("head script").Each(func(i int, s *goquery.Selection) {
			var band string
			band = s.Text()
			check := strings.Index(band, "__WML_REDUX_INITIAL_STATE__") != -1
			if check {
				index := strings.Index(band, "= {")
				for i := 0; i < index +1 ; i ++ {
					band = strings.Replace(band, string(band[i]), " ", 1)
				}
				for i := 0; i < len(band) - 1 ; i ++ {
					body = body + string(band[i])
				}				
			}
		})
		data := []byte(body)
		shopCategory,_,_,_ := jsonparser.Get(data, "presoData", "modules", "left", "[0]", "data")
		jsonparser.ArrayEach(shopCategory, func(value1 []byte, dataType jsonparser.ValueType, offset int, err error) {
			value, _, _, _ := jsonparser.Get(value1, "url")
			urlNew := OptimizeUrl(string(value))
			_, checkUrl := urlMap[urlNew]
			if !checkUrl {
				urlMap[urlNew] = true
				CrawlUrl(urlNew, urlMap)
			}
		})
	}
	return urlMap
}
