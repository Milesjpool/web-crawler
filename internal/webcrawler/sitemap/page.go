package sitemap

import (
	"encoding/json"
	"net/url"
)

type Page struct {
	Url      string
	SubPages []string
}

func NewPage(url string) Page {
	return Page{Url: url, SubPages: []string{}}
}

func (page *Page) AddSubPage(subpage string) {
	page.SubPages = append(page.SubPages, subpage)
}

func (page *Page) ToJson() string {
	jsonResult, _ := json.Marshal(page)
	return string(jsonResult)
}
func (page *Page) GetUrl() *url.URL {
	parsed, _ := url.Parse(page.Url)
	return parsed
}
