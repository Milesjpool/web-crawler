package sitemap

import (
	"net/url"
	"encoding/json"
)

type PageListing struct {
	Domain     string
	EntryPoint string
	Pages      []*Page
}

func NewPageListing(domain string, entryPoint string) PageListing {
	listing := PageListing{Domain: domain, EntryPoint: entryPoint, Pages: []*Page{}}
	listing.Put(entryPoint)
	return listing
}
func (listing *PageListing) GetEntryPoint() *url.URL {
	parsed, _ := url.Parse(listing.EntryPoint)
	return parsed
}

func (listing *PageListing) ToJson() string {
	jsonResult, _ := json.Marshal(listing)
	return string(jsonResult)
}
func (listing *PageListing) Put(newPage string) (added bool) {
	for _, p := range listing.Pages {
		if p.Url == newPage {
			return false
		}
	}
	listing.Pages = append(listing.Pages, NewPage(newPage))
	return true
}
func (listing *PageListing) AddSubPages(page string, subPages []string) {
	for _, p := range listing.Pages {
		if p.Url == page {
			p.AddSubPages(subPages)
		}
	}
}


