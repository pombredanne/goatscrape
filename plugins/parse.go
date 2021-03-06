package plugins

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// ExtractAllLinks simply extracts all the <a href="<urL>"> </a> in a page
// provided that they aren't silly like a hash. It also expands relative
// links automagically.
func ExtractAllLinks(resp *http.Response) []string {
	var links []string

	tempURL := *resp.Request.URL
	tempURL.Fragment = ""
	tempURL.Path = ""
	tempURL.RawQuery = ""
	uri := tempURL.String()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		return links
	}

	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")

		if exists && href != "" {

			if href[0] == '#' {
				return
			}
			if strings.HasPrefix(href, "//") {
				href = "http:" + href
			}
			if strings.HasPrefix(href, "/") {
				href = uri + href
			}
			links = append(links, href)
		}
	})

	return links
}
