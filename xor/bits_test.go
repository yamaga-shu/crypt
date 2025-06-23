package xor

import "testing"

func TestBinary(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{a: 0b0000, b: 0b0000, expect: 0b0000},
		{a: 0b0101, b: 0b0011, expect: 0b0110},
		{a: 0b1111, b: 0b0000, expect: 0b1111},
		{a: 0b1010, b: 0b0101, expect: 0b1111},
		{a: 0b1111, b: 0b1111, expect: 0b0000},
	}

	for _, tt := range tests {
		result := bits(tt.a, tt.b)
		if result != tt.expect {
			t.Errorf("bits(%b, %b) = %b; want %b", tt.a, tt.b, result, tt.expect)
		}
	}
}
