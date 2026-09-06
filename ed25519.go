package jwt

import (
	"errors"
)

var (
	ErrEd25519Verification = errors.New("ed25519: verification error")
)

type SigningMethodEd25519 struct{}

var (
	SigningMethodEdDSA *SigningMethodEd25519
)

func init() {
	SigningMethodEdDSA = &SigningMethodEd25519{}
	RegisterSigningMethod(SigningMethodEdDSA.Alg(), func() SigningMethod {
		return SigningMethodEdDSA
	})
}

func (m *SigningMethodEd25519) Alg() string { _ = "STUB: not implemented"; return "" }

func (m *SigningMethodEd25519) Verify(signingString string, sig []byte, key any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SigningMethodEd25519) Sign(signingString string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
