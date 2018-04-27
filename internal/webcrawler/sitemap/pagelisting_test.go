package sitemap

import (
	"testing"
	"reflect"
)

func TestNewPageListing(t *testing.T) {
	domain := "my.domain"
	url := "url"
	actual := NewPageListing(domain, url)
	expected := PageListing{Domain:domain, EntryPoint:url, Pages:[]*Page{NewPage(url)}}

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected: ", expected, ", but was: ", actual);
	}
}

func TestPut(t *testing.T) {
	domain := "my.domain"
	url := "url"
	pageListing := NewPageListing(domain, url)

	url2 := "newPage"
	added := pageListing.Put(url2)
	expected := PageListing{Domain:domain, EntryPoint:url, Pages:[]*Page{NewPage(url),NewPage(url2)}}

	if !added {
		t.Error("Expected new element to be added.");
	}
	if !reflect.DeepEqual(pageListing, expected) {
		t.Error("Expected: ", expected, ", but was: ", pageListing);
	}
	added = pageListing.Put(url2)
	if added {
		t.Error("Expected existing element to not be added.");
	}
	if !reflect.DeepEqual(pageListing, expected) {
		t.Error("Expected: ", expected, ", but was: ", pageListing);
	}
}


func TestAddSubPages_PageListing(t *testing.T) {
	domain := "my.domain"
	page1 := "page1"
	pageListing := NewPageListing(domain, page1)
	pageListing.AddSubPages(page1, []string{"page2", "page3"})

	expectedPage1 := NewPage(page1)
	expectedPage1.SubPages = []string{"page2", "page3"}
	expected := PageListing{Domain:domain, EntryPoint:page1, Pages:[]*Page{expectedPage1}}

	if !reflect.DeepEqual(pageListing, expected) {
		t.Error("Expected: ", expected, ", but was: ", pageListing);
	}
}

func TestToJson(t *testing.T) {
	domain := "my.domain"
	url := "url"
	pageListing := NewPageListing(domain, url)
	actual := pageListing.ToJson()
	expected := "{" +
		"\"Domain\":\"" + domain + "\"," +
		"\"EntryPoint\":\"" + url + "\"," +
		"\"Pages\":[{\"Url\":\""+url+"\",\"SubPages\":[]}]}"

	if actual != expected {
		t.Error("Expected: ", expected, ", but was: ", actual);
	}
}