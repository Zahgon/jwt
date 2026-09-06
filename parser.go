package jwt

const tokenDelimiter = "."

type Parser struct {
	validMethods []string

	useJSONNumber bool

	skipClaimsValidation bool

	validator *Validator

	decodeStrict bool

	decodePaddingAllowed bool
}

func NewParser(options ...ParserOption) *Parser { _ = "STUB: not implemented"; return nil }

func (p *Parser) Parse(tokenString string, keyFunc Keyfunc) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *Parser) ParseUnverified(tokenString string, claims Claims) (token *Token, parts []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func splitToken(token string) ([]string, bool) { _ = "STUB: not implemented"; return nil, false }

func (p *Parser) DecodeSegment(seg string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Parse(tokenString string, keyFunc Keyfunc, options ...ParserOption) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc, options ...ParserOption) (*Token, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
