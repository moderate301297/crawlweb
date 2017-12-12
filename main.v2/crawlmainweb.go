package main

import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"./crawl"
)
// optimize url have found
func OptimizeUrl(value string) (url string) {
	if (strings.Index(value, "https") == 0){
		url = value
		return url
	} else {
		url = "https://www.walmart.com" + value
		return url
	}
}
// Crawl main web find link of shops 
func CrawlMainWeb() {
	doc, err := goquery.NewDocument("https://www.walmart.com/all-departments")
	if err != nil {
		fmt.Println("Error: ", err)
	}
  
	// Find the review items (text of token script)
	doc.Find("head script").Each(func(i int, s *goquery.Selection) {
		var (
			band string
			index int
			urlLinkText []string
			urlLinkTextPhoto []string
		)
		urlLink := make(map[string]string)
	  band = s.Text()
	  // Parse text of token script, get url
	  check := strings.Index(band, "_setReduxState") == -1
	  if !check {
		for {
			var name string
			index = strings.Index(band, "name")
			if index == -1 {
				break
			}
			if string(band[index - 2]) == "," {
				for i := 0; ((string(band[index + 7 + i + 1]) != ",") || (string(band[index + 7 + i + 3]) != "l")); i++ {
					name = name + string(band[index + 7 + i])
				}
				fmt.Println("-> Name category: ", name)
				// delete text have parsed
				for i := 0; i < index + 7 ; i ++ {
					band = strings.Replace(band, string(band[i]), " ", 1)
				}
				indexNext := strings.Index(band, "name")
				for {
					var value string
					index = strings.Index(band, "linkText")
					if index >= indexNext {
						break
					}
					for i := 0; i < index + 11 ; i ++ {    
						band = strings.Replace(band, string(band[i]), " ", 1)
					}
					// get url
					index = strings.Index(band, "value")
					for i := 0; string(band[index + 8 + i + 1]) != "}"; i++ {
						value = value + string(band[index + 8 + i])
					}
					for i := 0; i < index + 8 ; i ++ {
						band = strings.Replace(band, string(band[i]), " ", 1)
					}
					// optimize url have found
					url := OptimizeUrl(value)
					// save url to a slice (urlLINkText)
					switch name {
					case "Electronics & Office": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Home, Furniture & Appliances": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Home Improvement & Patio": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Clothing, Shoes & Jewwlry": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Baby & Toddler": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Toys & Video Games": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Food, Household & Pets": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break					
					}
					case "Pharmacy, Health & Beauty": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Auto & Tires": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break
					}
					case "Photo & Personalized Shop": {
						check, _ := urlLink[url]
						if check != url {
							check1 := strings.Index(url, "/cp/") != -1
							if check1 {
								urlLink[url] = url
								urlLinkText = append(urlLinkText,url)
							}
							check2 := strings.Index(url, "/about/") != -1
							if check2 {
								urlLink[url] = url
								urlLinkTextPhoto = append(urlLinkTextPhoto,url)
							}
						}
						break
					}
					case "Sewing, Crafts & Christmas Decor": {
						check, _ := urlLink[url]
						if check != url {
							urlLink[url] = url
							urlLinkText = append(urlLinkText,url)
						}
						break		
					}
					}
				}
			} else {
				for i := 0; i < index + 7 ; i ++ {
					band = strings.Replace(band, string(band[i]), " ", 1)
				}
			}
		}
		urlLink := crawl.Crawl(urlLinkText, urlLink)
		crawl.CrawlPhoto(urlLinkTextPhoto, urlLink)
	  }	  
	})
}
// main
func main() {
	CrawlMainWeb()
}