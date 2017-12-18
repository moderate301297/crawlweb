package crawlphoto

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"./crawlproductphoto"
	"../../savedata"
)

// find customer riview
func Check(href string) string {
	var hrefNew string
	index := strings.Index(href, "#")
	if index != -1 {
		for i := 0; i < index; i++ {
			hrefNew = hrefNew + string(href[i])
		}
		return hrefNew
	} else {
		return href
	}
}
// find href in token
func GetHref(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}
	return
}
// optimize url
func OptimizeHref(href string) string{
	if strings.Index(href, "http") == -1 {
		if string(href[0]) == "/" && string(href[1]) == "/" {
			href = "https:" + href
		}
		href = "https://photos3.walmart.com" + href
	}	
	return href
}
// find url have "/category/"
func CrawlPhoto(url string) {
	urlMap := make(map[string]bool)
	var body []byte
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		savedata.SaveUrlError(url)
		return
	} else {
        body,_ = ioutil.ReadAll(response.Body)
        defer response.Body.Close()      
    }
	doc, _ := html.Parse(strings.NewReader(string(body)))
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
			href := GetHref(n)
			check1 := strings.Index(href, "/about/") != -1 
			if check1 {
				href := Check(href)
				urlNew := OptimizeHref(href)
				_, checkUrl := urlMap[urlNew] 
				if !checkUrl {
					urlMap[urlNew] = true
					urlMap = crawlproductphoto.CrawlProductPhoto(urlNew, urlMap)
				}	
			}
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
	f(doc)
}
