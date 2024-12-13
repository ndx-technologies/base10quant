package base10quant

import (
	"errors"
)

var ErrInvalidFormat = errors.New("invalid format")

const MaxL9 = 999_999_999

// L9 is base10 chars encoding least significant 9 decimals of uint32.
type L9 struct{ v uint32 }

func L9FromUint32(v uint32) L9 { return L9{v: v} }

func L9FromString(s string) (L9, error) {
	var v L9
	err := (&v).UnmarshalText([]byte(s))
	return v, err
}

func (s L9) UInt32() uint32 { return s.v }

func (s L9) IsEmpty() bool { return s.v == 0 }

func (s L9) AppendBytes(b []byte) {
	v := s.v
	v, b[8] = v/10, '0'+byte(v%10)
	v, b[7] = v/10, '0'+byte(v%10)
	v, b[6] = v/10, '0'+byte(v%10)
	v, b[5] = v/10, '0'+byte(v%10)
	v, b[4] = v/10, '0'+byte(v%10)
	v, b[3] = v/10, '0'+byte(v%10)
	v, b[2] = v/10, '0'+byte(v%10)
	v, b[1] = v/10, '0'+byte(v%10)
	_, b[0] = v/10, '0'+byte(v%10)
}

func (s L9) MarshalText() ([]byte, error) {
	b := make([]byte, 9)
	s.AppendBytes(b)
	return b, nil
}

func (s *L9) UnmarshalText(b []byte) error {
	if len(b) != 9 {
		return ErrInvalidFormat
	}
	s.v = 0
	for i := 0; i < 9; i++ {
		if b[i] < '0' || b[i] > '9' {
			return ErrInvalidFormat
		}
		s.v *= 10
		s.v += uint32(b[i] - '0')
	}
	return nil
}

func (s L9) String() string {
	b, _ := s.MarshalText()
	return string(b)
}
