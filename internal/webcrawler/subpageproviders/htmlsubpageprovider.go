package subpageproviders

import (
	"golang.org/x/net/html"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/http"
	"net/url"
)

const aTag, hrefTag = "a", "href"

type HtmlSubPageProvider struct {
	PageRetriever http.DataRetriever
}

func (s HtmlSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	content, _ := s.PageRetriever.RetrieveData(pageUrl.String())

	tokenizer := html.NewTokenizer(content)

	subpages := extractSubPages(tokenizer)

	return subpages;
}

func extractSubPages(tokenizer *html.Tokenizer) []url.URL {
	subPages := make([]url.URL, 0)

	for tokenType := tokenizer.Next(); tokenType != html.ErrorToken; tokenType = tokenizer.Next() {
		token := tokenizer.Token()
		if token.Data == aTag {
			href, found := getHref(token)
			if found {
				subPages = append(subPages, *href)
			}
		}
	}
	return subPages
}

func getHref(t html.Token) (href *url.URL, found bool) {
	for _, attribute := range t.Attr {
		if attribute.Key == hrefTag {
			href, _ := url.Parse(attribute.Val)
			return href, true
		}
	}
	return nil, false
}