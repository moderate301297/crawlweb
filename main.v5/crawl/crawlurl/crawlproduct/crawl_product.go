package crawlproduct

import (
	"strings"
	"github.com/buger/jsonparser"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"../../../savedata"
)

func QueryBody(url string) {
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
		fmt.Println(string(title), string(dataPath), string(imageUrl))
		savedata.SaveData(string(title), string(dataPath), string(imageUrl))
	}, "items")
}