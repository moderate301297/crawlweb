package crawl

import (
	"./extractall"
	"strings"
	"./extractphoto"
)
// access url "cp" of web walmart.com
func Crawl(urlLinkText []string, urlLink map[string]string) (urlLinkNew map[string]string) {
	for _, url := range urlLinkText {
		check := strings.Index(url, "/cp/") != -1
		if check {
			urlLink = extractall.ExtractAll(url, urlLink)			
		}
	}
	// retrun map
	return urlLink					
}
// access url of web photo 
func CrawlPhoto(urlLinkText []string, urlLink map[string]string) (urlLinkNew map[string]string){
	for _, url := range urlLinkText {	
		urlLink = extractphoto.ExtractPhoto(url, urlLink)			
	}
	// return map
	return urlLink						
}