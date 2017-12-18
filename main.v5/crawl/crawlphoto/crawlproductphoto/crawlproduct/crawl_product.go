package crawlproduct

import (
	"strings"
	"io/ioutil"
	"net/http"
	"golang.org/x/net/html"
	"fmt"
	"../../../../savedata"
	"sync"
	"github.com/PuerkitoBio/goquery"
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

func OptimizeHref(href string) string{
	if strings.Index(href, "http") == -1 {
		if string(href[0]) == "/" && string(href[1]) == "/" {
			href = "https:" + href
		}
		href = "https://photos3.walmart.com" + href
	}	
	return href
}

func GetSrc(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "src" {
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
	var wg sync.WaitGroup
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
					wg.Add(1)
					go func (url string) {
						defer wg.Done()
						var title, link, linkImage string
						var linkSlice []string
						doc, err := goquery.NewDocument(url)
						if err != nil {
							fmt.Println("Error: ", err)
							savedata.SaveUrlError(url)
							return
						}
						doc.Find("div.product-page").Each(func(i int, s *goquery.Selection) {
							title = s.Find("h1").Text()
						})
						doc.Find("div.product-page ol.breadcrumb li").Each(func(i int, s *goquery.Selection) {
							linkSlice = append(linkSlice, s.Find("a").Text())
						})
						doc.Find("title").Each(func(i int, s *goquery.Selection) {
							linkSlice = append(linkSlice, s.Text())
						})
						for _, k := range linkSlice {
							if k != "" {
								link = link + k + "/"
							}
						}
						doc.Find("div.product-image.undefined img").Each(func(i int, s *goquery.Selection) {
							for _, k := range s.Nodes {
								href := GetSrc(k)
								linkImage = OptimizeHref(href)
							}
						})
						savedata.SaveData(title, link, linkImage, url)
					} (url)
				}	
			}
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
	f(doc)
	wg.Wait()
	return urlMap
}