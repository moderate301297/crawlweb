package crawl

import (
	"fmt"
	"./extractall"
	"strings"
	"./extractphoto"
	"time"
)
// access url "cp" of web walmart.com
func Crawl(urlLinkText []string, urlLink map[string]string) {
	for _, url := range urlLinkText {
		check := strings.Index(url, "/cp/") != -1
		if check {
			fmt.Println("AAAAAAAA", time.Now())
			urlLink = extractall.ExtractAll(url, urlLink)			
		}
	}
}
// access url of web photo 
func CrawlPhoto(urlLinkText []string, urlLink map[string]string) {
	for _, url := range urlLinkText {	
		urlLink = extractphoto.ExtractPhoto(url, urlLink)			
	}		
}