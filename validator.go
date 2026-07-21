package jwt

import (
	"time"
)

type ClaimsValidator interface {
	Claims
	Validate() error
}

type Validator struct {
	leeway time.Duration

	timeFunc func() time.Time

	requireExp bool

	requireNbf bool

	verifyIat bool

	expectedAud []string

	expectAllAud bool

	expectedIss string

	expectedSub string
}

func NewValidator(opts ...ParserOption) *Validator { _ = "STUB: not implemented"; return nil }

func (v *Validator) Validate(claims Claims) error { _ = "STUB: not implemented"; return nil }

func (v *Validator) verifyExpiresAt(claims Claims, cmp time.Time, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) verifyIssuedAt(claims Claims, cmp time.Time, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) verifyNotBefore(claims Claims, cmp time.Time, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) verifyAudience(claims Claims, cmp []string, expectAllAud bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) verifyIssuer(claims Claims, cmp string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (v *Validator) verifySubject(claims Claims, cmp string, required bool) error {
	_ = "STUB: not implemented"
	return nil
}

func errorIfFalse(value bool, err error) error { _ = "STUB: not implemented"; return nil }

func errorIfRequired(required bool, claim string) error { _ = "STUB: not implemented"; return nil }
