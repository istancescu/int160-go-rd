# Int160 – Fixed-Length 160-Bit Integer Library for Go

`Int160` is a lightweight Go package that implements a 160-bit (20-byte) unsigned integer with utility methods for bitwise operations, comparisons, and conversions. It’s useful for systems such as DHTs, cryptographic identifiers, or compact fixed-size hashes.

---

## 📦 Features

- Fixed-size 160-bit unsigned integer (`Int160`)
- Conversion from hex strings and byte slices
- Bitwise XOR and distance calculation
- Equality and lexicographic comparison
- Zero-value check and deep cloning
- Bit-level manipulation (`SetBit`)
- Common prefix length calculation
- Hexadecimal string representation

---

## 🛠️ Installation

To add this package to your project:

```bash
go get github.com/istancescu/int160-go-rd
