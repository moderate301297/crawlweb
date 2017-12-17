package crawlproductphoto

import (
	"golang.org/x/net/html"
	"strconv"
)
// find rel in token
func GetRel(t *html.Node) (href string) {
	for _, a := range t.Attr {
		if a.Key == "rel" {
			href = a.Val
		}
	}
	return
}
// find next page
func NextPage(href string, i int64) string{
	href = href + "&_page=" + strconv.FormatInt(i, 10)
	return href
}