package sitemap

import (
	"encoding/json"
	"net/url"
)

type Page struct {
	Url      string
	SubPages []Page
}

func NewPage(url string) Page {
	return Page{Url: url, SubPages: []Page{}}
}

func (page *Page) AddSubPage(newPage Page) {
	page.SubPages = append(page.SubPages, newPage)
}

func (page *Page) SameAs(otherPage Page) bool {
	return page.Url == otherPage.Url
}

func (page *Page) ToJson() string {
	jsonResult, _ := json.Marshal(page)
	return string(jsonResult)
}
func (page *Page) GetUrl() *url.URL {
	parsed, _ := url.Parse(page.Url)
	return parsed
}
