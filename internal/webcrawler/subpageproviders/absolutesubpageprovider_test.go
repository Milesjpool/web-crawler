package subpageproviders

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

	url1, _ := url.Parse("http://page1.com/page1")
	url2, _ := url.Parse("http://page2.com/page2")
	expectedPages := []url.URL{*url1,*url2}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubSubPageProvider_AbsoluteSubPageProvider struct {
}

func (s StubSubPageProvider_AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	url1, _ := url.Parse("/page1")
	url2, _ := url.Parse("http://page2.com/page2")
	return []url.URL{*url1, *url2}
}

