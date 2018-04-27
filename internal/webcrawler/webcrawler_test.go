package webcrawler

import (
	"testing"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
	"reflect"
	"net/url"
)

const myDomain = "my.domain"
const page1, page2, page3, page4 = "page1", "page2", "page3", "page4"

func TestCrawl(t *testing.T) {
	args := StubCrawlerArgs{domain: myDomain, entryPoint: page1}
	subPageProvider := StubbedSubPageProvider{}

	pageListing := WebCrawler{SubPageProvider: subPageProvider}.Crawl(args)

	expectedPageListing := sitemap.NewPageListing(myDomain, page1)
	expectedPageListing.Pages[0].SubPages = []string{page2, page3}


	if ! reflect.DeepEqual(pageListing, expectedPageListing) {
		t.Error("Expected: ", expectedPageListing, ", but was: ", pageListing);
	}
}

type StubCrawlerArgs struct {
	domain string
	entryPoint string
}

func (s StubCrawlerArgs) GetDomain() string  {
	return s.domain
}

func (s StubCrawlerArgs) GetEntryPoint() string  {
	return s.entryPoint
}

type StubbedSubPageProvider struct {
}

func (s StubbedSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	url1, _ := url.Parse(page1)
	url2, _ := url.Parse(page2)
	url3, _ := url.Parse(page3)
	url4, _ := url.Parse(page4)

	if pageUrl.String() == page1{return []url.URL{*url2, *url3}}
	if pageUrl.String() == page2{return []url.URL{*url3, *url4}}
	if pageUrl.String() == page3{return []url.URL{*url1, *url4}}
	if pageUrl.String() == page4{return []url.URL{*url1, *url2}}
	return []url.URL{}
}

