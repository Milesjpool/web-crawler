package sitemap

import (
	"testing"
	"reflect"
	"io"
	"strings"
)

func TestGetSubPages(t *testing.T) {
	pageRetriever := StubPageRetriever{}
	subPageProvider := HtmlSubPageProvider{pageRetriever}
	subPages := subPageProvider.GetSubPages("http://page1.com/url")

	expectedPages := []string{
		"/page1",
		"http://page2.com/page2",
	}

	if ! reflect.DeepEqual(subPages, expectedPages) {
		t.Error("Expected: ", expectedPages, ", but was: ", subPages);
	}
}

type StubPageRetriever struct {
}

func (s StubPageRetriever) RetrievePage(url string) io.Reader {
	return strings.NewReader("<html><body>"+
		"<a href=\"/page1\">link1</a>"+
		"<a href=\"http://page2.com/page2\">link2</a>"+
	"</body></html>")
}

