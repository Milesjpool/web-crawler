package http

import (
	"testing"
	"strings"
	"io/ioutil"
	"net/http"
)

func TestRetrievePage_Integration(t *testing.T) {
	pageRetriever := WebPageRetriever{HttpClient:http.DefaultClient}
	data, err := pageRetriever.RetrievePage("https://monzo.com/")

	if err != nil {
		t.Error("Failed to retrieve page", err);
	}

	content, err := ioutil.ReadAll(data)
	expected := "<!DOCTYPE html>"

	if err != nil {
		t.Error("Failed to retrieve page", err);
	}
	if !strings.Contains(string(content), expected) {
		t.Error("Expected: ", expected, ", but was: ", content);
	}
}