package http

import (
	"io"
)

type DataRetriever interface {
	RetrieveData(url string) (io.Reader, error)
}