package jwt

import (
	"crypto/ecdsa"
	"errors"
)

var (
	ErrNotECPublicKey  = errors.New("key is not a valid ECDSA public key")
	ErrNotECPrivateKey = errors.New("key is not a valid ECDSA private key")
)

func ParseECPrivateKeyFromPEM(key []byte) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseECPublicKeyFromPEM(key []byte) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
