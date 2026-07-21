package jwt

import (
	"crypto"
	"crypto/rsa"
)

type SigningMethodRSAPSS struct {
	*SigningMethodRSA
	Options *rsa.PSSOptions

	VerifyOptions *rsa.PSSOptions
}

var (
	SigningMethodPS256 *SigningMethodRSAPSS
	SigningMethodPS384 *SigningMethodRSAPSS
	SigningMethodPS512 *SigningMethodRSAPSS
)

func init() {

	SigningMethodPS256 = &SigningMethodRSAPSS{
		SigningMethodRSA: &SigningMethodRSA{
			Name: "PS256",
			Hash: crypto.SHA256,
		},
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
		VerifyOptions: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		},
	}
	RegisterSigningMethod(SigningMethodPS256.Alg(), func() SigningMethod {
		return SigningMethodPS256
	})

	SigningMethodPS384 = &SigningMethodRSAPSS{
		SigningMethodRSA: &SigningMethodRSA{
			Name: "PS384",
			Hash: crypto.SHA384,
		},
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
		VerifyOptions: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		},
	}
	RegisterSigningMethod(SigningMethodPS384.Alg(), func() SigningMethod {
		return SigningMethodPS384
	})

	SigningMethodPS512 = &SigningMethodRSAPSS{
		SigningMethodRSA: &SigningMethodRSA{
			Name: "PS512",
			Hash: crypto.SHA512,
		},
		Options: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthEqualsHash,
		},
		VerifyOptions: &rsa.PSSOptions{
			SaltLength: rsa.PSSSaltLengthAuto,
		},
	}
	RegisterSigningMethod(SigningMethodPS512.Alg(), func() SigningMethod {
		return SigningMethodPS512
	})
}

func (m *SigningMethodRSAPSS) Verify(signingString string, sig []byte, key any) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SigningMethodRSAPSS) Sign(signingString string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
