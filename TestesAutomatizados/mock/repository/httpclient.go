package repository

import (
	"io"
	"net/http"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}
