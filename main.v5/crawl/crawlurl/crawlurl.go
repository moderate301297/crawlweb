package crawlurl

import (
	"strings"
	"github.com/buger/jsonparser"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"../../savedata"
	"sync"
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
		var wg sync.WaitGroup
		for {
			wg.Add(1)
			go func (url string) {
				defer wg.Done()
				fmt.Println(url)
				doc, err := goquery.NewDocument(url)
				if err != nil {
					fmt.Println("Error: ", err)
					savedata.SaveUrlError(url)
					return
				}
				var body string
				doc.Find("head script").Each(func(i int, s *goquery.Selection) {
					var band string
					band = s.Text()
					check := strings.Index(band, "__WML_REDUX_INITIAL_STATE__") != -1
					if check {
						index := strings.Index(band, "{")
						for i := 0; i < index ; i ++ {
							band = strings.Replace(band, string(band[i]), " ", 1)
						}
						body = band
						savedata.SaveDataBody(body)	
					}
				})
				data := []byte(body)
				dataProduct,_,_,_ := jsonparser.Get(data, "preso")
				dataPath,_,_,_ := jsonparser.Get(dataProduct, "adContext","categoryPathName")
				jsonparser.ArrayEach(dataProduct, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
					title, _, _, _ := jsonparser.Get(value, "title")
					imageUrl, _, _, _ := jsonparser.Get(value, "imageUrl")
					url, _, _, _ := jsonparser.Get(value, "productPageUrl")
					savedata.SaveData(string(title), string(dataPath), string(imageUrl), string(url))
				}, "items")
			} (url)
			// next page
			urlNext := NextPage(url)
			if urlNext == "" {
				wg.Wait()
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
				urlMap = CrawlUrl(urlNew, urlMap)
			}
		})
	}
	return urlMap
}
