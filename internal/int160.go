package int160

import (
	hex2 "encoding/hex"
	"errors"
	"fmt"
	"github.com/istancescu/int160-go-rd/internal/logger"
)

type Int160 struct {
	Val [20]byte
}

type IInt160 interface {
	Bytes() [20]byte
	Xor(other Int160) *Int160
	Equals(other *Int160) bool
	String() string
	IsZero() bool
	Clone() Int160
}

func (i *Int160) Xor(other Int160) *Int160 {
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
		logger.LogError(err.Error())
		return nil, err
	}

	byte20, err := hex2.DecodeString(hex)

	if err != nil {
		logErr := fmt.Errorf("failure decoding hex string %s: %w", hex, err)
		logger.LogError(logErr.Error())
		return nil, logErr
	}

	var result Int160
	copy(result.Val[:], byte20)

	return &result, nil
}

func NewInt160FromBytes(bytes []byte) (*Int160, error) {
	if len(bytes) != 20 {
		logger.LogError("Failure on conversion from bytes, length not 20")
		return nil, errors.New("conversion failure")
	}
	var x Int160
	copy(x.Val[:], bytes)
	return &x, nil
}
