package crawlproduct

import (
	"strings"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"../../../../savedata"
)

// find same urls
func Check1(href string) string{
	var hrefNew string
	index := strings.Index(href, "?")
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

func CrawlProduct(url string, urlMap map[string]bool) (urlMapNew map[string]bool) {	 
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
			check := strings.Index(href, "/category/") != -1 
			if check {
				href := Check1(href)
				url := OptimizeHref(href)
				_, checkUrl := urlMap[url]
				if !checkUrl {
					urlMap[url] = true
					go QueryBody(url)
				}	
			}
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
	f(doc)
	return urlMap
}