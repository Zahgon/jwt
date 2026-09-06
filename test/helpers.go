package test

import (
	"crypto"
	"crypto/rsa"

	"github.com/golang-jwt/jwt/v5"
)

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	_ = "STUB: not implemented"
	return nil
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	_ = "STUB: not implemented"
	return nil
}

func MakeSampleToken(c jwt.Claims, method jwt.SigningMethod, key any) string {
	_ = "STUB: not implemented"
	return ""
}

func LoadECPrivateKeyFromDisk(location string) crypto.PrivateKey {
	_ = "STUB: not implemented"
	return *new(crypto.PrivateKey)
}

func LoadECPublicKeyFromDisk(location string) crypto.PublicKey {
	_ = "STUB: not implemented"
	return *new(crypto.PublicKey)
}
