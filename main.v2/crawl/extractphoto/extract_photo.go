package extractphoto

import (
	"fmt"
	"strings"
	"golang.org/x/net/html"
	"net/http"
	"io/ioutil"
	"../extractproduct"
	"../../savedata"
)

// find same urls
func Check1(href string) string{
	var hrefNew string
	index1 := strings.Index(href, "?")
	for i := 0; i < index1; i++ {
		hrefNew = hrefNew + string(href[i])
	}
	return hrefNew
}
// find customer riview
func Check(href string) string {
	var hrefNew string
	index := strings.Index(href, "#")
	for i := 0; i < index; i++ {
		hrefNew = hrefNew + string(href[i])
	}
	return hrefNew
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
func ExtractPhoto(url string, urlLink map[string]string) (urlLinkNew map[string]string) {

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
			check1 := strings.Index(href, "/category/") != -1 
			if check1 {
				href := OptimizeHref(href)
				if urlLink[href] != href {
					var i int64
					for i = 0; i < 25; i ++ {
						urlLink[href] = href
						fmt.Println(href)
						// find url product
						urlLink = extractproduct.ExtractProductPhoto(href, urlLink)
						// next page
						d := ScaleLinkText(href, i)
						if d == "" || urlLink[d] == d {
							break
						}
						href = d
					}
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
