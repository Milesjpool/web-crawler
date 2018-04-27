package sitemap

import "net/url"

type Page struct {
	Url      string
	SubPages []string
}

func NewPage(url string) *Page {
	return &Page{Url: url, SubPages: []string{}}
}

func (p *Page) AddSubPages (newPages []string) {
	for _, np := range newPages {
		if !alreadyExists(p.SubPages, np) {
			p.SubPages = append(p.SubPages, np)
		}
	}
}

func (page *Page) GetUrl() *url.URL {
	parsed, _ := url.Parse(page.Url)
	return parsed
}

func alreadyExists(existingPages []string, newPage string) bool {
	for _, ep := range existingPages {
		if newPage == ep {
			return true
		}
	}
	return false
}