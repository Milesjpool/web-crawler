package sitemap

import (
	"golang.org/x/net/html"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/http"
)

const aTag, hrefTag = "a", "href"

type HtmlSubPageProvider struct {
	PageRetriever http.PageRetriever
}

func (s HtmlSubPageProvider) GetSubPages(url string) []string {
	content := s.PageRetriever.RetrievePage(url)

	tokenizer := html.NewTokenizer(content)

	subpages := extractSubPages(tokenizer)

	return subpages;
}

func extractSubPages(tokenizer *html.Tokenizer) []string {
	subPages := make([]string, 0)

	for tokenType := tokenizer.Next(); tokenType != html.ErrorToken; tokenType = tokenizer.Next() {
		token := tokenizer.Token()
		if token.Data == aTag {
			href, found := getHref(token)
			if found {
				subPages = append(subPages, href)
			}
		}
	}
	return subPages
}

func getHref(t html.Token) (href string, found bool) {
	for _, attribute := range t.Attr {
		if attribute.Key == hrefTag {
			return attribute.Val, true
		}
	}
	return "", false
}