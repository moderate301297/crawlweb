package crawl

import (
	"strings"
	"./crawlurl"
	"./crawlphoto"
)
// access shops category of website walmart.com
func Crawl(urlSlice []string, urlMap map[string]bool) {	
	for _, url := range urlSlice {
		check := strings.Index(url, "photos3") != -1 
		if check {
			crawlphoto.CrawlPhoto(url)
		} else {
			urlMap = crawlurl.CrawlUrl(url, urlMap)
		}
	}			
}