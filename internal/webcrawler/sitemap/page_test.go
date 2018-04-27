package sitemap

import (
	"testing"
	"reflect"
)

func TestAddSubPages_Page(t *testing.T) {
	url := "url"
	page1 := "page1"
	page2 := "page2"
	page3 := "page3"
	page := Page{Url:url, SubPages:[]string{page1, page2}}
	page.AddSubPages([]string{page2, page3})

	expected := Page{Url:url, SubPages:[]string{page1, page2, page3}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}

func TestNewPage(t *testing.T) {
	url := "url"
	page := NewPage(url)

	expected := &Page{Url:url, SubPages:[]string{}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}