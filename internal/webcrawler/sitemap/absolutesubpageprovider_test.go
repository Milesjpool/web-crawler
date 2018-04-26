package sitemap

import (
	"testing"
	"reflect"
	"net/url"
)


func TestGetSubPages_AbsoluteSubPageProvider(t *testing.T) {
	rootPage, _ := url.Parse("http://page1.com/")
	pageProvider := StubSubPageProvider_AbsoluteSubPageProvider{}
	subPageProvider := AbsoluteSubPageProvider{SubPageProvider: pageProvider}
	subPages := subPageProvider.GetSubPages(rootPage)

	expectedPages := []string{
		"http://page1.com/page1",
		"http://page2.com/page2",
	}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubSubPageProvider_AbsoluteSubPageProvider struct {
}

func (s StubSubPageProvider_AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []string {
	return []string{"/page1", "http://page2.com/page2"}
}

