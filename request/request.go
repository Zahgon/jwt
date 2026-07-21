package request

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func ParseFromRequest(req *http.Request, extractor Extractor, keyFunc jwt.Keyfunc, options ...ParseFromRequestOption) (token *jwt.Token, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseFromRequestWithClaims(req *http.Request, extractor Extractor, claims jwt.Claims, keyFunc jwt.Keyfunc) (token *jwt.Token, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fromRequestParser struct {
	req       *http.Request
	extractor Extractor
	claims    jwt.Claims
	parser    *jwt.Parser
}

type ParseFromRequestOption func(*fromRequestParser)

func WithClaims(claims jwt.Claims) ParseFromRequestOption {
	_ = "STUB: not implemented"
	return *new(ParseFromRequestOption)
}

func WithParser(parser *jwt.Parser) ParseFromRequestOption {
	_ = "STUB: not implemented"
	return *new(ParseFromRequestOption)
}
