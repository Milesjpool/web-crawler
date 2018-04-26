package http

import "io"

type PageRetriever interface {
	RetrievePage(url string) io.Reader
}
