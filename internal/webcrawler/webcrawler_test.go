package webcrawler

import (
	"testing"
	"github.com/Milesjpool/web-crawler/internal/webcrawler/sitemap"
	"reflect"
)

func TestCrawl(t *testing.T) {
	args := &CrawlerArgs{entryPoint:"page1"}
	subPageProvider := StubbedSubPageProvider{}

	page := WebCrawler{SubPageProvider: subPageProvider}.Crawl(args)

	expectedPage := sitemap.Page{Url: "page1", SubPages: []sitemap.Page{
		sitemap.NewPage("page2"),
		sitemap.NewPage("page3"),
	}}

	if ! reflect.DeepEqual(page, expectedPage) {
		t.Error("Expected: ", expectedPage, ", but was: ", page);
	}
}

type StubbedSubPageProvider struct {
}

func (s StubbedSubPageProvider) GetSubPages(page string) []string {
	return []string{"page2", "page3"}
}

