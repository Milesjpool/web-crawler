package http

import "net/http"

type HttpClient interface {
	Get(string) (*http.Response, error)
}