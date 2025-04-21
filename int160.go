package int160

import (
	hex2 "encoding/hex"
	"errors"
	"fmt"
)

type Int160 struct {
	Val [20]byte
}

func (i *Int160) Xor(other *Int160) *Int160 {
	var out Int160
	for j := 0; j < 20; j++ {
		out.Val[j] = i.Val[j] ^ other.Val[j]
	}
	return &out
}

func (i *Int160) Equals(other *Int160) bool {
	if other == nil {
		return false
	}
	for j := 0; j < 20; j++ {
		if i.Val[j] != other.Val[j] {
			return false
		}
	}
	return true
}

func (i *Int160) Bytes() [20]byte {
	return i.Val
}

func (i *Int160) String() string {
	return fmt.Sprintf("%x", i.Val)
}

func (i *Int160) hexString() string {
	return hex2.EncodeToString(i.Val[:])
}

func (i *Int160) IsZero() bool {
	for _, v := range i.Val {
		if v != 0 {
			return false
		}
	}
	return true
}

func (i *Int160) Clone() *Int160 {
	var cloned Int160

	for j := 0; j < 20; j++ {
		cloned.Val[j] = i.Val[j]
	}

	return &cloned
}

func NewInt160FromHex(hex string) (*Int160, error) {
	if len(hex) != 40 {
		err := fmt.Errorf("invalid hexadecimal string length: got %d, want 40", len(hex))
		LogError(err.Error())
		return nil, err
	}

	byte20, err := hex2.DecodeString(hex)

	if err != nil {
		logErr := fmt.Errorf("failure decoding hex string %s: %w", hex, err)
		LogError(logErr.Error())
		return nil, logErr
	}

	var result Int160
	copy(result.Val[:], byte20)

	return &result, nil
}

func NewInt160FromBytes(bytes []byte) (*Int160, error) {
	if len(bytes) != 20 {
		LogError("Failure on conversion from bytes, length not 20")
		return nil, errors.New("conversion failure")
	}
	var x Int160
	copy(x.Val[:], bytes)
	return &x, nil
}

// TODO: test this
func Distance(a, b *Int160) (*Int160, error) {
	if (a == nil) != (b == nil) {
		return nil, fmt.Errorf("either a, or b pointer are null\n")
	}
	return a.Xor(b), nil
}

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

func (i *Int160) SetBit(val bool, pos uint8) error {
	if pos >= 160 {
		return fmt.Errorf("can't set byte %t at pos %d \n", val, pos)
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
