package int160

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

func Test_newInt160FromHex(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name string
		args args
		want *Int160
	}{
		{
			"Should transform this Hex String into 160 bytes",
			args{
				"2f00000000000000000000000000000000000000",
			},
			&Int160{
				Val: [20]byte{
					0x2f, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
					0x00, 0x00, 0x00, 0x00,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newInt160FromHexString(tt.args.hex))
		})
	}
}

func Test_newInt160FromHexSha(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *Int160
	}{
		{
			name:  "Should convert SHA-1 hex to int160 correctly",
			input: "e0c09862faafc7d2b315b5f8c14f9f38e2a3ac8b",
			want: &Int160{
				Val: [20]byte{
					0xe0, 0xc0, 0x98, 0x62, 0xfa, 0xaf, 0xc7, 0xd2,
					0xb3, 0x15, 0xb5, 0xf8, 0xc1, 0x4f, 0x9f, 0x38,
					0xe2, 0xa3, 0xac, 0x8b,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, newInt160FromHexString(tt.input))
		})
	}
}

func TestXor(t *testing.T) {
	type args struct {
		a *Int160
		b *Int160
	}
	tests := []struct {
		name string
		args args
		want *Int160
	}{
		{
			"Should perform XOR correctly",
			args{
				a: &Int160{
					Val: [20]byte{
						0x01, 0x02, 0x03, 0x04, 0x05,
						0x06, 0x07, 0x08, 0x09, 0x0A,
						0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
						0x10, 0x11, 0x12, 0x13, 0x14,
					},
				},
				b: &Int160{
					Val: [20]byte{
						0xFF, 0xEE, 0xDD, 0xCC, 0xBB,
						0xAA, 0x99, 0x88, 0x77, 0x66,
						0x55, 0x44, 0x33, 0x22, 0x11,
						0x00, 0xFF, 0xEE, 0xDD, 0xCC,
					},
				},
			},
			&Int160{
				Val: [20]byte{
					0xFE, 0xEC, 0xDE, 0xC8, 0xBE,
					0xAC, 0x9E, 0x80, 0x7E, 0x6C,
					0x5E, 0x48, 0x3E, 0x2C, 0x1E,
					0x10, 0xEE, 0xFC, 0xCE, 0xD8,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Xor(*tt.args.a, *tt.args.b))
		})
	}
}
