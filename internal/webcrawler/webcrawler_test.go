package webcrawler

import (
	"testing"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
	"reflect"
	"net/url"
)

const page1 = "page1"
const page2 = "page2"
const page3 = "page3"
const page4 = "page4"

func TestCrawl(t *testing.T) {
	args := &CrawlerArgs{entryPoint: page1}
	subPageProvider := StubbedSubPageProvider{}

	page := WebCrawler{SubPageProvider: subPageProvider}.Crawl(args)

	p1 := sitemap.NewPage(page1)
	p1.AddSubPage(page2)
	p1.AddSubPage(page3)
	expectedPage := p1


	if ! reflect.DeepEqual(page, expectedPage) {
		t.Error("Expected: ", expectedPage, ", but was: ", page);
	}
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

