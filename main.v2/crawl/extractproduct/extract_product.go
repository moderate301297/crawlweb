package extractproduct

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
	"../../savedata"
)

// optimize url
func OptimizeUrl(value string) (url string) {
	if (strings.Index(value, "http") == 0){
		url = value
		return url
	} else {
		url = "https://www.walmart.com" + value
		return url
	}
}
// find url product in text of token script
func ExtractProduct(url string, urlLink map[string]string) (urlLinkNew map[string]string) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error: ", err)
		// save url error to my sql
		savedata.SaveUrlError(url)
		return urlLink
	}
	doc.Find("head script").Each(func(i int, s *goquery.Selection) {
		var (
			band string
			index int
		)
		band = s.Text()
		check := strings.Index(band, "__WML_REDUX_INITIAL_STATE__") != -1
		if check {
			for {
				var value string
				index = strings.Index(band, "productPageUrl")
				if index == -1 {
					break
				}
				for i := 0; string(band[index + 17 + i + 1]) != ","; i++ {
					value = value + string(band[index + 17 + i])
				}
				for i := 0; i < index + 17 ; i ++ {
					band = strings.Replace(band, string(band[i]), " ", 1)
				}
				url := OptimizeUrl(value)
				if (urlLink[url] != url){
					urlLink[url] = url
					savedata.SaveLink(url)
				}				
			}
		}
	})
	// return map
	return urlLink
}
