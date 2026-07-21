package jwt

type MapClaims map[string]any

func (m MapClaims) GetExpirationTime() (*NumericDate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m MapClaims) GetNotBefore() (*NumericDate, error) { _ = "STUB: not implemented"; return nil, nil }

func (m MapClaims) GetIssuedAt() (*NumericDate, error) { _ = "STUB: not implemented"; return nil, nil }

func (m MapClaims) GetAudience() (ClaimStrings, error) {
	_ = "STUB: not implemented"
	return *new(ClaimStrings), nil
}

func (m MapClaims) GetIssuer() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (m MapClaims) GetSubject() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (m MapClaims) parseNumericDate(key string) (*NumericDate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m MapClaims) parseClaimsString(key string) (ClaimStrings, error) {
	_ = "STUB: not implemented"
	return *new(ClaimStrings), nil
}

func (m MapClaims) parseString(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
