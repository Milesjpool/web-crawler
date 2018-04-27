package sitemap

import (
	"testing"
	"reflect"
)

func TestNewPage(t *testing.T) {
	url := "url"
	page := NewPage(url)

	expected := Page{Url:url, SubPages:[]string{}}

	if !reflect.DeepEqual(page, expected) {
		t.Error("Expected: ", expected, ", but was: ", page);
	}
}