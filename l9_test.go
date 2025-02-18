package base10quant_test

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/ndx-technologies/base10quant"
)

func ExampleL9() {
	var v base10quant.L9
	v.UnmarshalText([]byte("123456789"))
	fmt.Println(v)
	// Output: 123456789
}

func TestL9(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		tests := []struct {
			s string
			v uint32
		}{
			{s: "999999999", v: 999_999_999},
			{s: "123456789", v: 123_456_789},
		}
		for _, tc := range tests {
			v := base10quant.L9FromUint32(tc.v)
			if tc.s != v.String() {
				t.Errorf("expected %s, got %s", tc.s, v.String())
			}
			if tc.v != v.UInt32() {
				t.Errorf("expected %d, got %d", tc.v, v.UInt32())
			}

			x, err := base10quant.L9FromString(tc.s)
			if err != nil {
				t.Error(err)
			}
			if v != x {
				t.Errorf("expected %v, got %v", v, x)
			}
		}
	})

	t.Run("error", func(t *testing.T) {
		tests := []string{
			"0123456789",
			"a",
			"",
			"1234a6789",
			"1234 6789",
		}
		for _, tc := range tests {
			_, err := base10quant.L9FromString(tc)
			if err == nil {
				t.Errorf("expected error")
			}
		}
	})
}

func FuzzL9(f *testing.F) {
	f.Add(uint32(999_999_999))
	f.Add(uint32(1))
	f.Add(uint32(0))

	f.Fuzz(func(t *testing.T, v uint32) {
		v = v % 999_999_999

		q := base10quant.L9FromUint32(v)
		if v != q.UInt32() {
			t.Errorf("expected %d, got %d", v, q.UInt32())
		}

		g, err := base10quant.L9FromString(q.String())
		if err != nil {
			t.Error(err)
		}
		if q != g {
			t.Errorf("expected %v, got %v", q, g)
		}

		b, err := q.MarshalText()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		var f base10quant.L9
		if !f.IsEmpty() {
			t.Errorf("expected empty")
		}
		if err := (f).UnmarshalText(b); err != nil {
			t.Error(err)
		}
		if f != q {
			t.Errorf("expected %v, got %v", q, f)
		}
	})
}

func BenchmarkL9(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		x := rand.Uint32()
		v := base10quant.L9FromUint32(x)

		for b.Loop() {
			v.String()
		}
	})

	b.Run("from_string", func(b *testing.B) {
		x := rand.Uint32()
		v := base10quant.L9FromUint32(x)
		s := v.String()

		for b.Loop() {
			base10quant.L9FromString(s)
		}
	})
}
