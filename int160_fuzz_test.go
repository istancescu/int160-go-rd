package main

import (
	"testing"
)

func FuzzNewInt160FromHex(f *testing.F) {
	f.Add("0000000000000000000000000000000000000001")   // Valid input (40 chars)
	f.Add("aabbccddeeff112233445566778899aabbccddeeff") // Valid input (40 chars)

	f.Fuzz(func(t *testing.T, hexStr string) {
		if !isValidHex(hexStr) || len(hexStr) != 40 {
			return
		}
		if len(hexStr) != 40 {
			return
		}

		result, err := NewInt160FromHex(hexStr)

		if err != nil {
			t.Errorf("Unexpected error for input %s: %v", hexStr, err)
		}

		if result == nil {
			t.Errorf("Expected result to be non-nil for input %s", hexStr)
		}
	})
}

func isValidHex(s string) bool {
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
			return false
		}
	}
	return true
}
