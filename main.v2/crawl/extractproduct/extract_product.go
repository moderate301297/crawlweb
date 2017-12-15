package extractproduct

import (
	"strings"
	"github.com/buger/jsonparser"
	"github.com/PuerkitoBio/goquery"
	"fmt"
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

func ExtractProduct(url string) {
	var body string
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
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
		}
	})
	data := []byte(body)
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		b, _, _, _ := jsonparser.Get(value, "productPageUrl")
		url := OptimizeUrl(string(b))
		savedata.SaveLink(url)
		fmt.Println(url)
	}, "preso", "items")
}
