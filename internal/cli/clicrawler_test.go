package cli

import (
	"testing"
	"net/http"
)


func TestExecute(t *testing.T) {
	writer := new(MockWriter)
	page := "https://monzo.com/"
	args := []string{page}

	httpClient := new(MockClient)
	crawler := CliCrawler{Writer: writer, HttpClient:httpClient}
	crawler.Execute(args)

	actual := writer.written
	expected := []string{
		"{" +
			"\"Domain\":\"monzo.com\"," +
			"\"EntryPoint\":\"" + page + "\"," +
			"\"Pages\":[{\"Url\":\"" + page + "\",\"SubPages\":[]}]" +
		"}",
	}

	if len(actual) != len(expected) {
		t.Error("Incorrect output - Expected: ", expected, ", but was: ", actual)
		return
	}
	for i, v := range expected {
		if v != actual[i] {
			t.Error("Expected: ", v, ", but was: ", actual[i])
		}
	}
}

type MockWriter struct {
	written []string
}

func (mw *MockWriter) Write(b []byte) (int, error) {
	mw.written = append(mw.written, string(b[:]))
	return 1, nil
}

type MockClient struct{
}

func (MockClient) Get(string) (*http.Response, error) {
	return &http.Response{StatusCode:404}, nil
}
