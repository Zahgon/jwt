package jwt

import (
	"time"
)

var TimePrecision = time.Second

var MarshalSingleStringAsArray = true

type NumericDate struct {
	time.Time
}

func NewNumericDate(t time.Time) *NumericDate { _ = "STUB: not implemented"; return nil }

func newNumericDateFromSeconds(f float64) *NumericDate { _ = "STUB: not implemented"; return nil }

func (date NumericDate) MarshalJSON() (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (date *NumericDate) UnmarshalJSON(b []byte) (err error) { _ = "STUB: not implemented"; return nil }

type ClaimStrings []string

func (s *ClaimStrings) UnmarshalJSON(data []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (s ClaimStrings) MarshalJSON() (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
