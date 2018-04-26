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

	expected := Page{Url:url, SubPages:[]string{}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}

func TestAddSubPage(t *testing.T) {
	url1 := "url1"
	url2 := "url2"
	page := NewPage(url1)

	page.AddSubPage(url2)

	expected := Page{Url:url1, SubPages:[]string{"url2"}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}