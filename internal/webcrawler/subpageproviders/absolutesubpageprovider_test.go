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
	url3, _ := url.Parse("http://page3.com/page3")
	url4, _ := url.Parse("http://page4.com/page4")
	expectedPages := []url.URL{*url1,*url2,*url3,*url4}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubSubPageProvider_AbsoluteSubPageProvider struct {
}

func (s StubSubPageProvider_AbsoluteSubPageProvider) GetSubPages(pageUrl *url.URL) []url.URL {
	url1, _ := url.Parse("/page1")
	url2, _ := url.Parse("http://page2.com/page2")
	url3, _ := url.Parse("http://page3.com/page3#abc?a=1")
	url4, _ := url.Parse("http://page4.com/page4?z=9#xyz")
	return []url.URL{*url1,*url2,*url3,*url4}
}

