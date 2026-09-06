package jwt

var SigningMethodNone *signingMethodNone

const UnsafeAllowNoneSignatureType unsafeNoneMagicConstant = "none signing method allowed"

var NoneSignatureTypeDisallowedError error

type signingMethodNone struct{}
type unsafeNoneMagicConstant string

func init() {
	SigningMethodNone = &signingMethodNone{}
	NoneSignatureTypeDisallowedError = newError("'none' signature type is not allowed", ErrTokenUnverifiable)

	RegisterSigningMethod(SigningMethodNone.Alg(), func() SigningMethod {
		return SigningMethodNone
	})
}

func (m *signingMethodNone) Alg() string { _ = "STUB: not implemented"; return "" }

func (m *signingMethodNone) Verify(signingString string, sig []byte, key any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (m *signingMethodNone) Sign(signingString string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
