package int160

import (
	"bytes"
	"github.com/alecthomas/assert/v2"
	"reflect"
	"testing"
)

func Test_newInt160FromHex(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *Int160
	}{
		{
			"Should transform this Hex String into 160 bytes",
			"2f00000000000000000000000000000000000000",
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
			res, _ := NewInt160FromHex(tt.input)
			assert.Equal(t, tt.want, res)
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
			res, _ := NewInt160FromHex(tt.input)
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_newInt160FromHexSha_Too_Long(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  *Int160
	}{
		{
			name:  "Should fail to convert from SHA1 - too long",
			input: "e0c09862faafc7d2b315b5f8c14f9f38e2a3ac8b11",
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
			_, err := NewInt160FromHex(tt.input)
			assert.Error(t, err)
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
			assert.Equal(t, tt.want, tt.args.a.Xor(*tt.args.b))
		})
	}
}

func TestNewInt160FromBytes(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "valid 20-byte input",
			input:   []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
			wantErr: false,
		},
		{
			name:    "too short input",
			input:   make([]byte, 10),
			wantErr: true,
		},
		{
			name:    "too long input",
			input:   make([]byte, 21),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewInt160FromBytes(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewInt160FromBytes() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if got == nil {
					t.Fatalf("Expected non-nil Int160 on success")
				}
				gotBytes := got

				if !bytes.Equal(got.Val[:], tt.input) {
					t.Errorf("Int160 value mismatch. Got %v, want %v", gotBytes, tt.input)
				}
			}
		})
	}
}

func Test_int160_Equals(t *testing.T) {
	type fields struct {
		val [20]byte
	}
	tests := []struct {
		name   string
		fields fields
		input  *Int160
		want   bool
	}{
		{
			"should equal",
			fields{[20]byte(make([]byte, 20))},
			&Int160{
				Val: [20]byte(make([]byte, 20)),
			},
			true,
		},
		{
			"should not equal",
			fields{[20]byte(make([]byte, 20))},
			&Int160{
				Val: [20]byte{
					0xFE, 0xEC, 0xDE, 0xC8, 0xBE,
					0xAC, 0x9E, 0x80, 0x7E, 0x6C,
					0x5E, 0x48, 0x3E, 0x2C, 0x1E,
					0x10, 0xEE, 0xFC, 0xCE, 0xD8,
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Int160{
				Val: tt.fields.val,
			}
			if got := i.Equals(tt.input); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_int160_IsZero(t *testing.T) {
	type args struct {
		val [20]byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should return true, as its 0",
			args{
				[20]byte(make([]byte, 20)),
			},
			true,
		},
		{
			"should return false",
			args{
				[20]byte(bytes.Repeat([]byte{1}, 20)),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Int160{
				Val: tt.args.val,
			}
			if got := i.IsZero(); got != tt.want {
				t.Errorf("IsZero() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	original := &Int160{}
	for i := 0; i < 20; i++ {
		original.Val[i] = byte(i)
	}

	cloned := original.Clone()

	if original == cloned {
		t.Errorf("Clone() returned the same instance, expected different instance")
	}
	if !reflect.DeepEqual(original.Val, cloned.Val) {
		t.Errorf("Clone() returned incorrect values: got %v, want %v", cloned.Val, original.Val)
	}

	original.Val[0] = 255
	if original.Val[0] == cloned.Val[0] {
		t.Errorf("Clone() failed to create a deep copy, clone was modified when original was changed")
	}
}
