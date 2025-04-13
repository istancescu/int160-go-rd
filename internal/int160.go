package int160

import (
	hex2 "encoding/hex"
	"math/big"
)

type Int160 struct {
	Val [20]byte
}

// Xor xor table
// a 	b 	a^b
// 1 	0 	1
// 1 	1 	0
// 0 	0 	0
// 0 	1 	0
func Xor(a, b Int160) *Int160 {
	var result Int160

	for i := 0; i < 20; i++ {
		result.Val[i] = a.Val[i] ^ b.Val[i]
	}

	return &result
}

func newInt160FromHexString(hex string) *Int160 {
	byte20 := make([]byte, 20)
	byte20, err := hex2.DecodeString(hex)

	if err != nil || len(byte20) != 20 {
		return nil
	}

	var x Int160

	copy(x.Val[:], byte20)
	return &x
}

// TODO: implement this
func newInt160FromBigInt(i *big.Int) *Int160 {
	var x big.Int

	//rep of int 160
	x.Exp(i, big.NewInt(2), big.NewInt(0))

	return nil
}
