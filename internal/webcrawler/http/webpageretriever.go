package http

import (
	"io"
	"io/ioutil"
	"bytes"
	"net/http"
)

type WebPageRetriever struct {
	HttpClient HttpClient
}

func (r WebPageRetriever) RetrieveData(url string) (io.Reader, error) {
	response, err := r.HttpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return bytes.NewBuffer([]byte{}), nil
	}

	body, err := readBody(response)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func readBody(response *http.Response) (io.Reader, error) {
	data, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	return bytes.NewBuffer(data), err
}
