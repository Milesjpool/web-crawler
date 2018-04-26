package sitemap

import (
	"testing"
	"reflect"
)

func TestToJson(t *testing.T) {
	url := "url"
	page := NewPage(url)
	actual := page.ToJson()
	expected := "{\"Url\":\"" + url + "\",\"SubPages\":[]}"

	if actual != expected {
		t.Error("Expected: ", expected, ", but was: ", actual);
	}
}

func TestNewPage(t *testing.T) {
	url := "url"
	page := NewPage(url)

	expected := Page{Url:url, SubPages:[]Page{}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}

func TestAddSubPage(t *testing.T) {
	url1 := "url1"
	url2 := "url2"
	page := NewPage(url1)

	page.AddSubPage(NewPage(url2))

	expected := Page{Url:url1, SubPages:[]Page{NewPage("url2")}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}

func TestSameAs(t *testing.T) {
	url1 := "url1"
	url2 := "url2"
	page := NewPage(url1)
	page.AddSubPage(NewPage(url2))

	samePage := NewPage(url1)
	differentPage := NewPage(url2)

	if ! page.SameAs(samePage) {
		t.Error("Expected: ", page, ", to be the same as: ", samePage);
	}

	if page.SameAs(differentPage) {
		t.Error("Expected: ", page, ", to not be the same as: ", differentPage);
	}
}