package jwt

type RegisteredClaims struct {
	Issuer string `json:"iss,omitempty"`

	Subject string `json:"sub,omitempty"`

	Audience ClaimStrings `json:"aud,omitempty"`

	ExpiresAt *NumericDate `json:"exp,omitempty"`

	NotBefore *NumericDate `json:"nbf,omitempty"`

	IssuedAt *NumericDate `json:"iat,omitempty"`

	ID string `json:"jti,omitempty"`
}

func (c RegisteredClaims) GetExpirationTime() (*NumericDate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c RegisteredClaims) GetNotBefore() (*NumericDate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c RegisteredClaims) GetIssuedAt() (*NumericDate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c RegisteredClaims) GetAudience() (ClaimStrings, error) {
	_ = "STUB: not implemented"
	return *new(ClaimStrings), nil
}

func (c RegisteredClaims) GetIssuer() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (c RegisteredClaims) GetSubject() (string, error) { _ = "STUB: not implemented"; return "", nil }
