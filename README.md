# int160

A 160-bit unsigned integer type for Go, stored as 20 big-endian bytes. Meant for Kademlia/BitTorrent-style DHT node IDs and SHA-1 digests, where you mostly need XOR distance and comparison.

## Install

```
go get github.com/istancescu/int160-go-rd
```

Needs Go 1.24+. No dependencies.

## Usage

The import path doesn't match the package name, so alias it.

```go
package main

import (
	"errors"
	"fmt"
	"log"

	int160 "github.com/istancescu/int160-go-rd"
)

func main() {
	a, err := int160.NewInt160FromHex("0000000000000000000000000000000000000001")
	if err != nil {
		log.Fatal(err)
	}
	b, err := int160.NewInt160FromHex("8000000000000000000000000000000000000000")
	if err != nil {
		log.Fatal(err)
	}

	d, err := int160.Distance(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)                    // 8000000000000000000000000000000000000001
	fmt.Println(a.Less(b))            // true
	fmt.Println(a.CommonPrefixLen(b)) // 0

	// wrong length
	_, err = int160.NewInt160FromHex("abcd")
	fmt.Println(errors.Is(err, int160.ErrInvalidLength)) // true
}
```

## API

```go
type Int160 struct{ Val [20]byte } // Val[0] is the most significant byte
```

| Function | Notes |
| --- | --- |
| `NewInt160FromHex(s string) (*Int160, error)` | `s` must be exactly 40 hex chars |
| `NewInt160FromBytes(b []byte) (*Int160, error)` | `b` must be exactly 20 bytes |
| `Distance(a, b *Int160) (*Int160, error)` | XOR distance, `ErrNilInput` if either is nil |

| Method | Notes |
| --- | --- |
| `Xor(other) *Int160` | bitwise XOR |
| `Equals(other) bool` | |
| `Less(other) bool` | unsigned compare |
| `Bytes() [20]byte` | copy of the value |
| `String() string` | lowercase hex, 40 chars |
| `IsZero() bool` | |
| `Clone() *Int160` | |
| `SetBit(val bool, pos uint8) error` | pos 0 is the most significant bit, error if pos >= 160 |
| `CommonPrefixLen(o) uint8` | leading bits in common, 160 if equal |

Errors: `ErrInvalidLength`, `ErrNilInput`. Both work with `errors.Is`.

## License

MIT, see [LICENSE](LICENSE).
