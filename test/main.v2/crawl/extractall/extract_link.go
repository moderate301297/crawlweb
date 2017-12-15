package extractall

import (
	"time"
	"fmt"
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	// "../extractproduct"
	// "../../savedata"
)
// check urk customer review
func Check(href string) string {
	var hrefNew string
	index := strings.Index(href, "#")
	for i := 0; i < index; i++ {
		hrefNew = hrefNew + string(href[i])
	}
	return hrefNew
}
// get href in token
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
		href = "https://www.walmart.com" + href
	}	
	return href
}
// func recursive, find url "browse"
var count int

func ExtractAll(url string, urlLink map[string]string) (urlLinkNew map[string]string) {

	var body []byte
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		// savedata.SaveUrlError(url)
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
			check1 := strings.Index(href, "/browse/") != -1 
			if check1 {
				href := OptimizeHref(href)
				if urlLink[href] != href {
					for {
						urlLink[href] = href
						count ++;
						// find url of product
						// urlLink = extractproduct.ExtractProduct(href, urlLink)
						// next page 
						d := ScaleLinkText(href)
						if d == "" {
							break
						}
						href = d
					}
					fmt.Println(count, time.Now())
				}	
			}
			check2 := strings.Index(href, "/cp/") != -1
			// if url has "cp" call recursive function
			if check2 {
				href := OptimizeHref(href)
				if (urlLink[href] != href) {
					urlLink[href] = href
					ExtractAll(href, urlLink)
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
