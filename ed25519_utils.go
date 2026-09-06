package jwt

import (
	"crypto"
	"errors"
)

var (
	ErrNotEdPrivateKey = errors.New("key is not a valid Ed25519 private key")
	ErrNotEdPublicKey  = errors.New("key is not a valid Ed25519 public key")
)

func ParseEdPrivateKeyFromPEM(key []byte) (crypto.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey), nil
}

func ParseEdPublicKeyFromPEM(key []byte) (crypto.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey), nil
}
