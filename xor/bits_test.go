package xor

import "testing"

func TestBits(t *testing.T) {
	cases := []struct {
		a, b   int
		expect int
	}{
		{a: 0b0000, b: 0b0000, expect: 0b0000},
		{a: 0b0101, b: 0b0011, expect: 0b0110},
		{a: 0b1111, b: 0b0000, expect: 0b1111},
		{a: 0b1010, b: 0b0101, expect: 0b1111},
		{a: 0b1111, b: 0b1111, expect: 0b0000},
	}

	for _, c := range cases {
		result := bits(c.a, c.b)
		if result != c.expect {
			t.Errorf("bits(%b, %b) = %b; want %b", c.a, c.b, result, c.expect)
		}
	}
}
