package jwt

import (
	"crypto"
)

type Keyfunc func(*Token) (any, error)

type VerificationKey interface {
	crypto.PublicKey | []uint8
}

type VerificationKeySet struct {
	Keys []VerificationKey
}

type Token struct {
	Raw       string
	Method    SigningMethod
	Header    map[string]any
	Claims    Claims
	Signature []byte
	Valid     bool
}

func New(method SigningMethod, opts ...TokenOption) *Token { _ = "STUB: not implemented"; return nil }

func NewWithClaims(method SigningMethod, claims Claims, opts ...TokenOption) *Token {
	_ = "STUB: not implemented"
	return nil
}

func (t *Token) SignedString(key any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (t *Token) SigningString() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (*Token) EncodeSegment(seg []byte) string { _ = "STUB: not implemented"; return "" }
