package util

import (
	"testing"
	"reflect"
	"net/url"
)

func TestAsStrings(t *testing.T) {
	u1,_ := url.Parse("url1")
	u2,_ := url.Parse("url2")
	urls := []url.URL{*u1, *u2}

	actual := AsStrings(urls)
	expected := []string{"url1","url2"}

	if ! reflect.DeepEqual(actual, expected) {
		t.Error("Expected: ", expected, ", but was: ", actual);
	}
}