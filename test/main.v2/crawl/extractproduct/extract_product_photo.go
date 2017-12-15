package extractproduct

import (
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"fmt"
	"../../savedata"
)

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
func OptimizeUrlPhoto(href string) string{
	if strings.Index(href, "http") == -1 {
		if string(href[0]) == "/" && string(href[1]) == "/" {
			href = "https:" + href
		} else {
			href = "https://photos3.walmart.com" + href
		}
	}	
	return href
}
// find url product of web photo
func ExtractProductPhoto(url string, urlLink map[string]string) (urlLinkNew map[string]string) {
	var body []byte
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		savedata.SaveUrlError(url)
		return urlLink
	} else {
        body,_ = ioutil.ReadAll(response.Body)
        defer response.Body.Close()      
    }
	doc, _ := html.Parse(strings.NewReader(string(body)))
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
			href := GetHref(n)
			check := strings.Index(href, "blankets?") != -1
			if check {
				href := OptimizeUrlPhoto(href)
				if urlLink[href] != href {
					urlLink[href] = href
					savedata.SaveLink(href)					
				} 
			}
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
	f(doc)
	// return map
	return urlLink
}
