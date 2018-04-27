package subpageproviders

import (
	"testing"
	"reflect"
	"net/url"
)


func TestGetSubPages_DomainSubPageProvider(t *testing.T) {
	rootPage, _ := url.Parse("http://page1.com/")
	pageProvider := StubSubPageProvider_DomainSubPageProvider{}
	subPageProvider := DomainSubPageProvider{SubPageProvider: pageProvider}
	subPages := subPageProvider.GetSubPages(rootPage)

	url1, _ := url.Parse("http://page1.com/page1")
	expectedPages := []url.URL{*url1}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubSubPageProvider_DomainSubPageProvider struct {
}

func (s StubSubPageProvider_DomainSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	url1, _ := url.Parse("http://page1.com/page1")
	url2, _  := url.Parse("http://page2.com/page2")
	return []url.URL{*url1, *url2}
}

