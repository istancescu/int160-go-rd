// Package int160 provides a fixed-size 160-bit unsigned integer, mostly useful
// for Kademlia-style DHT node IDs and SHA-1 digests.
package int160

import (
	"encoding/hex"
	"errors"
	"fmt"
)

// ErrInvalidLength is returned when an input does not have the required length.
var ErrInvalidLength = errors.New("int160: invalid length")

// ErrNilInput is returned when a required pointer argument is nil.
var ErrNilInput = errors.New("int160: nil input")

// Int160 is a 160-bit value stored as 20 bytes in big-endian order.
type Int160 struct {
	Val [20]byte
}

// Xor returns the bitwise XOR of i and other.
func (i *Int160) Xor(other *Int160) *Int160 {
	var out Int160
	for j := 0; j < 20; j++ {
		out.Val[j] = i.Val[j] ^ other.Val[j]
	}
	return &out
}

// Equals reports whether i and other hold the same value. It returns false if other is nil.
func (i *Int160) Equals(other *Int160) bool {
	return other != nil && i.Val == other.Val
}

// Bytes returns a copy of the underlying 20 bytes.
func (i *Int160) Bytes() [20]byte {
	return i.Val
}

// String returns the value as a 40-character lowercase hex string.
func (i *Int160) String() string {
	return hex.EncodeToString(i.Val[:])
}

// IsZero reports whether all bits of i are zero.
func (i *Int160) IsZero() bool {
	for _, v := range i.Val {
		if v != 0 {
			return false
		}
	}
	return true
}

// Clone returns an independent copy of i.
func (i *Int160) Clone() *Int160 {
	c := *i
	return &c
}

// NewInt160FromHex parses a 40-character hex string into an Int160.
func NewInt160FromHex(s string) (*Int160, error) {
	if len(s) != 40 {
		return nil, fmt.Errorf("%w: got %d hex chars, want 40", ErrInvalidLength, len(s))
	}

	byte20, err := hex.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("int160: decode hex: %w", err)
	}

	var result Int160
	copy(result.Val[:], byte20)

	return &result, nil
}

// NewInt160FromBytes creates an Int160 from exactly 20 bytes.
func NewInt160FromBytes(b []byte) (*Int160, error) {
	if len(b) != 20 {
		return nil, fmt.Errorf("%w: got %d bytes, want 20", ErrInvalidLength, len(b))
	}
	var x Int160
	copy(x.Val[:], b)
	return &x, nil
}

// Distance returns the XOR distance between a and b, or ErrNilInput if either is nil.
func Distance(a, b *Int160) (*Int160, error) {
	if a == nil || b == nil {
		return nil, ErrNilInput
	}
	return a.Xor(b), nil
}

// Less reports whether i is numerically less than other.
func (i *Int160) Less(other *Int160) bool {
	for j := 0; j < 20; j++ {
		if i.Val[j] < other.Val[j] {
			return true
		}
		if i.Val[j] > other.Val[j] {
			return false
		}
	}
	return false
}

// SetBit sets the bit at pos to val. Position 0 is the most significant bit of
// Val[0]; it returns an error if pos is not in [0,160).
func (i *Int160) SetBit(val bool, pos uint8) error {
	if pos >= 160 {
		return fmt.Errorf("int160: bit position %d out of range [0,160)", pos)
	}

	byteIndex := pos / 8
	bitIndex := 7 - (pos % 8)

	mask := byte(1 << bitIndex)

	if val {
		i.Val[byteIndex] |= mask
	} else {
		i.Val[byteIndex] &= ^mask
	}

	return nil
}

// CommonPrefixLen returns the number of leading bits i and o have in common (160 if equal).
func (i *Int160) CommonPrefixLen(o *Int160) uint8 {
	xor := i.Xor(o)

	for j := 0; j < 20; j++ {
		if xor.Val[j] == 0 {
			continue
		}
		for k := 0; k < 8; k++ {
			if xor.Val[j]&(0x80>>k) != 0 {
				return uint8(j*8 + k)
			}
		}
	}
	return 160
}
