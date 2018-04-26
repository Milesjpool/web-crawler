package http

import (
	"io"
	"strings"
)

type WebPageRetriever struct {

}

func (r WebPageRetriever) RetrievePage(url string) io.Reader {
	return strings.NewReader("")
}
