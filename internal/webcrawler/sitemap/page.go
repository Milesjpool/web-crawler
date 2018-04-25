package sitemap

import "encoding/json"

type Page struct {
	Url      string
	SubPages []Page
}

func (p *Page) ToJson() string {
	jsonResult, _ := json.Marshal(p)
	return string(jsonResult)
}
