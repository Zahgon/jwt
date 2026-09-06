package jwt

import (
	"sync"
)

var signingMethods = map[string]func() SigningMethod{}
var signingMethodLock = new(sync.RWMutex)

type SigningMethod interface {
	Verify(signingString string, sig []byte, key any) error
	Sign(signingString string, key any) ([]byte, error)
	Alg() string
}

func RegisterSigningMethod(alg string, f func() SigningMethod) { _ = "STUB: not implemented"; return }

func GetSigningMethod(alg string) (method SigningMethod) {
	_ = "STUB: not implemented"
	return *new(SigningMethod)
}

func GetAlgorithms() (algs []string) { _ = "STUB: not implemented"; return nil }
