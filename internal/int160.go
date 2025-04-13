package int160

import (
	hex2 "encoding/hex"
	"errors"
	"fmt"
	"github.com/istancescu/int160-go-rd/internal/logger"
)

type int160 struct {
	val [20]byte
}

type Int160 interface {
	Bytes() [20]byte
	Xor(other Int160) Int160
	Equals(other Int160) bool
	String() string
	IsZero() bool
	Clone() Int160
}

// Xor xor table
// a 	b 	a^b
// 1 	0 	1
// 1 	1 	0
// 0 	0 	0
// 0 	1 	0
func (i *int160) Xor(other Int160) Int160 {
	o := other.(*int160)
	var out int160
	for j := 0; j < 20; j++ {
		out.val[j] = i.val[j] ^ o.val[j]
	}
	return &out
}

func (i *int160) Equals(other Int160) bool {
	o := other.(*int160)

	for j := 0; j < 20; j++ {
		if i.val[j] != o.val[j] {
			return false
		}
	}
	return true
}

func (i *int160) Bytes() [20]byte {
	return i.val
}

func (i *int160) String() string {
	return fmt.Sprintf("%x", i.val)
}

func (i *int160) IsZero() bool {
	for _, v := range i.val {
		if v != 0 {
			return false
		}
	}
	return true
}

func (i *int160) Clone() Int160 {
	var cloned int160

	for j := 0; j < 20; j++ {
		cloned.val[j] = i.val[j]
	}

	return &cloned
}

func NewInt160FromHex(hex string) (Int160, error) {
	if len(hex) != 40 {
		logger.LogError(fmt.Sprintf("Invalid hexadecimal string: %s", hex))
		return nil, errors.New("invalid hexadecimal string")
	}

	byte20 := make([]byte, 20)
	byte20, err := hex2.DecodeString(hex)

	if err != nil {
		logger.LogError(fmt.Sprintf("Failure on conversion from %s: ", hex))
		return nil, errors.New("conversion failure")
	}

	var result int160
	copy(result.val[:], byte20)

	return &result, nil
}

func NewInt160FromBytes(bytes []byte) (Int160, error) {
	if len(bytes) != 20 {
		logger.LogError("Failure on conversion from bytes, length not 20")
		return nil, errors.New("conversion failure")
	}
	var x int160
	copy(x.val[:], bytes)
	return &x, nil
}
