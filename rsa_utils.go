package jwt

import (
	"crypto/rsa"
	"errors"
)

var (
	ErrKeyMustBePEMEncoded = errors.New("invalid key: Key must be a PEM encoded PKCS1 or PKCS8 key")
	ErrNotRSAPrivateKey    = errors.New("key is not a valid RSA private key")
	ErrNotRSAPublicKey     = errors.New("key is not a valid RSA public key")
)

func ParseRSAPrivateKeyFromPEM(key []byte) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseRSAPrivateKeyFromPEMWithPassword(key []byte, password string) (*rsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ParseRSAPublicKeyFromPEM(key []byte) (*rsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
