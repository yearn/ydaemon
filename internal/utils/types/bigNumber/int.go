package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type Int struct{ big.Int }

func FromInt(b *big.Int) *Int {
	return &Int{*b}
}
func ToInt(b *Int) *big.Int {
	if b == nil {
		return big.NewInt(0)
	}
	return &b.Int
}
func NewInt(defaultValue ...int64) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*big.NewInt(defaultValue[0])}
}
func SetInt(defaultValue ...*big.Int) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*safeInt(defaultValue[0])}
}
func (b *Int) Clone(s *Int) *Int {
	if s == nil {
		return b
	}
	b.Int.Set(&s.Int)
	return b
}
func (b *Int) Set(s *big.Int) *Int {
	b.Int.Set(safeInt(s))
	return b
}
func (b *Int) SetString(s string) *Int {
	if s == "" || s == "\"\"" {
		b.SetInt64(0)
		return b
	}
	b.Int.SetString(s, 10)
	return b
}
func (b *Int) Add(x *Int, y *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	b.Int.Add(xAsInt, yAsInt)
	return b
}
func (b *Int) Sub(x *Int, y *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	b.Int.Sub(xAsInt, yAsInt)
	return b
}
func (b *Int) Mul(x *Int, y *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	b.Int.Mul(xAsInt, yAsInt)
	return b
}
func (b *Int) Div(x *Int, y *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	if y.IsZero() {
		b.SetUint64(0)
		return b
	}
	b.Int.Div(xAsInt, yAsInt)
	return b
}
func (b *Int) Exp(x *Int, y *Int, z *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	zAsInt := ToInt(z)
	b.Int.Exp(xAsInt, yAsInt, zAsInt)
	return b
}
func (b *Int) Uint64() uint64 {
	return ToInt(b).Uint64()
}
func (b *Int) String() string {
	return ToInt(b).String()
}

func (b *Int) IsZero() bool {
	return b.Cmp(big.NewInt(0)) == 0
}
func (b *Int) Safe(s *Int, defaultValue ...*Int) *Int {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return NewInt(0)
		}
		return defaultValue[0]
	}
	return s
}

func (b Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Int) UnmarshalJSON(p []byte) error {
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

func safeInt(s *big.Int, defaultValue ...*big.Int) *big.Int {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewInt(0)
		}
		return defaultValue[0]
	}
	return s
}
