package crawlproduct

import(
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
	"../../savedata"
	"strings"
)

// get src of token
func GetSrc(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "src" {
			href = a.Val
		}
	}
	return
}
// get title, link, link image
func CrawlProduct(url string) (title string, link string, linkImage string){
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("Error: ", err)
		savedata.SaveUrlError(url)
		return title, link, linkImage
	}
	fmt.Print("Title: ")
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		tmp := s.Text()
		index := strings.Index(tmp, "Walmart.com")
		for i := 0; i < index -2; i ++ {
			title = title + string(tmp[i])
		}
		fmt.Println(title)
	})
	fmt.Print("Link: ")
	doc.Find("ol.breadcrumb-list li").Each(func(i int, s *goquery.Selection) {
		link = link + s.Find("a").Text() + " > "
	})
	fmt.Print(link)
	fmt.Println()
	fmt.Print("Link Image: ")
	doc.Find("div.prod-HeroImage-container img").Each(func(i int, s *goquery.Selection) {
		for _, k := range s.Nodes {
			linkImage = GetSrc(k)
			fmt.Print(linkImage)
		}
	})
	fmt.Println()
	return title, link, linkImage
}