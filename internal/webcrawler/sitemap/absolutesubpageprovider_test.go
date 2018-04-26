package sitemap

import (
	"testing"
	"reflect"
)


func TestGetSubPages_AbsoluteSubPageProvider(t *testing.T) {
	pageProvider := StubSubPageProvider{}
	subPageProvider := AbsoluteSubPageProvider{SubPageProvider: pageProvider}
	subPages := subPageProvider.GetSubPages("http://page1.com/")

	expectedPages := []string{
		"http://page1.com/page1",
		"http://page2.com/page2",
	}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubSubPageProvider struct {
}

func (s StubSubPageProvider) GetSubPages(url string) []string {
	return []string{"/page1", "http://page2.com/page2"}
}

