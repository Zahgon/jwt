package request

import (
	"errors"
	"net/http"
)

var (
	ErrNoTokenInRequest = errors.New("no token present in request")
)

type Extractor interface {
	ExtractToken(*http.Request) (string, error)
}

type HeaderExtractor []string

func (e HeaderExtractor) ExtractToken(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type ArgumentExtractor []string

func (e ArgumentExtractor) ExtractToken(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type MultiExtractor []Extractor

func (e MultiExtractor) ExtractToken(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type PostExtractionFilter struct {
	Extractor
	Filter func(string) (string, error)
}

func (e *PostExtractionFilter) ExtractToken(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type BearerExtractor struct{}

func (e BearerExtractor) ExtractToken(req *http.Request) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
