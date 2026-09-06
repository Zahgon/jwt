package jwt

import "time"

type ParserOption func(*Parser)

func WithValidMethods(methods []string) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithJSONNumber() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithoutClaimsValidation() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithLeeway(leeway time.Duration) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithTimeFunc(f func() time.Time) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithIssuedAt() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithExpirationRequired() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithNotBeforeRequired() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithAudience(aud ...string) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithAllAudiences(aud ...string) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func WithIssuer(iss string) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithSubject(sub string) ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithPaddingAllowed() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }

func WithStrictDecoding() ParserOption { _ = "STUB: not implemented"; return *new(ParserOption) }
