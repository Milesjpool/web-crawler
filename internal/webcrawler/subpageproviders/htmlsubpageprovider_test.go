package subpageproviders

import (
	"testing"
	"reflect"
	"io"
	"strings"
	"net/url"
)

func TestGetSubPages_HtmlSubPageProvider(t *testing.T) {
	pageUrl, _ := url.Parse("http://page1.com/url")
	pageRetriever := StubPageRetriever{}
	subPageProvider := HtmlSubPageProvider{pageRetriever}
	subPages := subPageProvider.GetSubPages(pageUrl)

	url1, _ := url.Parse("/page1")
	url2, _ := url.Parse("http://page2.com/page2")
	expectedPages := []url.URL{*url1,*url2}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubPageRetriever struct {
}

func (s StubPageRetriever) RetrieveData(url string) (io.Reader, error) {
	return strings.NewReader("<html><body>"+
		"<a href=\"/page1\">link1</a>"+
		"<a href=\"http://page2.com/page2\">link2</a>"+
	"</body></html>"), nil
}

