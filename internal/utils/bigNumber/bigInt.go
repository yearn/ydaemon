package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type BigInt struct{ big.Int }

func FromBigInt(b *big.Int) *BigInt {
	return &BigInt{*b}
}
func ToBigInt(b *BigInt) *big.Int {
	if b == nil {
		return big.NewInt(0)
	}
	return &b.Int
}
func NewInt(defaultValue ...int64) *BigInt {
	if len(defaultValue) == 0 {
		return &BigInt{*big.NewInt(0)}
	}
	return &BigInt{*big.NewInt(defaultValue[0])}
}
func SetInt(defaultValue ...*big.Int) *BigInt {
	if len(defaultValue) == 0 {
		return &BigInt{*big.NewInt(0)}
	}
	return &BigInt{*safeBigInt(defaultValue[0])}
}
func (b *BigInt) Set(s *big.Int) *BigInt {
	b.Int.Set(safeBigInt(s))
	return b
}
func (b *BigInt) SetString(s string) *BigInt {
	if s == "" || s == "\"\"" {
		b.SetInt64(0)
		return b
	}
	b.Int.SetString(s, 10)
	return b
}
func (b *BigInt) Add(x *BigInt, y *BigInt) *BigInt {
	xAsBigInt := ToBigInt(x)
	yAsBigInt := ToBigInt(y)
	b.Int.Add(xAsBigInt, yAsBigInt)
	return b
}
func (b *BigInt) Sub(x *BigInt, y *BigInt) *BigInt {
	xAsBigInt := ToBigInt(x)
	yAsBigInt := ToBigInt(y)
	b.Int.Sub(xAsBigInt, yAsBigInt)
	return b
}
func (b *BigInt) Mul(x *BigInt, y *BigInt) *BigInt {
	xAsBigInt := ToBigInt(x)
	yAsBigInt := ToBigInt(y)
	b.Int.Mul(xAsBigInt, yAsBigInt)
	return b
}
func (b *BigInt) Div(x *BigInt, y *BigInt) *BigInt {
	xAsBigInt := ToBigInt(x)
	yAsBigInt := ToBigInt(y)
	if y.IsZero() {
		b.SetUint64(0)
		return b
	}
	b.Int.Div(xAsBigInt, yAsBigInt)
	return b
}
func (b *BigInt) Uint64() uint64 {
	return ToBigInt(b).Uint64()
}

func (b *BigInt) IsZero() bool {
	return b.Int.Cmp(big.NewInt(0)) == 0
}
func (b *BigInt) Safe(s *BigInt, defaultValue ...*BigInt) *BigInt {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return NewInt(0)
		}
		return defaultValue[0]
	}
	return s
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BigInt) UnmarshalJSON(p []byte) error {
	s := strings.Trim(string(p), "\"")
	if s == "null" {
		return nil
	}
	var z big.Int
	_, ok := z.SetString(s, 10)
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Int = z
	return nil
}

func safeBigInt(s *big.Int, defaultValue ...*big.Int) *big.Int {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewInt(0)
		}
		return defaultValue[0]
	}
	return s
}
