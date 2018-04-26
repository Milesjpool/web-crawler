package sitemap

import "encoding/json"

type Page struct {
	Url      string
	SubPages []Page
}

func NewPage(url string) Page {
	return Page{Url: url, SubPages: []Page{}}
}

func (page *Page) AddSubPage(url string) {
	page.SubPages = append(page.SubPages, NewPage(url))
}

func (page *Page) SameAs(otherPage Page) bool {
	return page.Url == otherPage.Url
}

func (page *Page) ToJson() string {
	jsonResult, _ := json.Marshal(page)
	return string(jsonResult)
}
