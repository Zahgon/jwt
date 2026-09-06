package jwt

import (
	"crypto"
	"errors"
)

var (
	ErrECDSAVerification = errors.New("crypto/ecdsa: verification error")
)

type SigningMethodECDSA struct {
	Name      string
	Hash      crypto.Hash
	KeySize   int
	CurveBits int
}

var (
	SigningMethodES256 *SigningMethodECDSA
	SigningMethodES384 *SigningMethodECDSA
	SigningMethodES512 *SigningMethodECDSA
)

func init() {

	SigningMethodES256 = &SigningMethodECDSA{"ES256", crypto.SHA256, 32, 256}
	RegisterSigningMethod(SigningMethodES256.Alg(), func() SigningMethod {
		return SigningMethodES256
	})

	SigningMethodES384 = &SigningMethodECDSA{"ES384", crypto.SHA384, 48, 384}
	RegisterSigningMethod(SigningMethodES384.Alg(), func() SigningMethod {
		return SigningMethodES384
	})

	SigningMethodES512 = &SigningMethodECDSA{"ES512", crypto.SHA512, 66, 521}
	RegisterSigningMethod(SigningMethodES512.Alg(), func() SigningMethod {
		return SigningMethodES512
	})
}

func (m *SigningMethodECDSA) Alg() string { _ = "STUB: not implemented"; return "" }

func (m *SigningMethodECDSA) Verify(signingString string, sig []byte, key any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SigningMethodECDSA) Sign(signingString string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
