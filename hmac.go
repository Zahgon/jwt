package jwt

import (
	"crypto"
	"errors"
)

type SigningMethodHMAC struct {
	Name string
	Hash crypto.Hash
}

var (
	SigningMethodHS256  *SigningMethodHMAC
	SigningMethodHS384  *SigningMethodHMAC
	SigningMethodHS512  *SigningMethodHMAC
	ErrSignatureInvalid = errors.New("signature is invalid")
)

func init() {

	SigningMethodHS256 = &SigningMethodHMAC{"HS256", crypto.SHA256}
	RegisterSigningMethod(SigningMethodHS256.Alg(), func() SigningMethod {
		return SigningMethodHS256
	})

	SigningMethodHS384 = &SigningMethodHMAC{"HS384", crypto.SHA384}
	RegisterSigningMethod(SigningMethodHS384.Alg(), func() SigningMethod {
		return SigningMethodHS384
	})

	SigningMethodHS512 = &SigningMethodHMAC{"HS512", crypto.SHA512}
	RegisterSigningMethod(SigningMethodHS512.Alg(), func() SigningMethod {
		return SigningMethodHS512
	})
}

func (m *SigningMethodHMAC) Alg() string { _ = "STUB: not implemented"; return "" }

func (m *SigningMethodHMAC) Verify(signingString string, sig []byte, key any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SigningMethodHMAC) Sign(signingString string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
