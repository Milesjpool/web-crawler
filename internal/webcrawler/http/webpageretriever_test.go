package http

import (
	"testing"
	"io/ioutil"
	"net/http"
)

func TestRetrievePage(t *testing.T) {
	client := StubHttpClient{}
	pageRetriever := WebPageRetriever{HttpClient: client}
	data, err := pageRetriever.RetrievePage("my-website")

	if err != nil {
		t.Error("Failed to retrieve page", err);
	}

	content, err := ioutil.ReadAll(data)

	if err != nil {
		t.Error("Failed to retrieve page", err);
	}
	if len(content) != 0 {
		t.Error("Expected content: ", content, " to be empty.");
	}
}

type StubHttpClient struct {
}

func (StubHttpClient) Get(string) (*http.Response, error) {
	return &http.Response{StatusCode:404}, nil
}
