package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Float struct{ big.Float }

func ToFloat(b *Float) *big.Float {
	if b == nil {
		return big.NewFloat(0)
	}
	return &b.Float
}
func NewFloat(defaultValue ...float64) *Float {
	if len(defaultValue) == 0 {
		return &Float{*big.NewFloat(0)}
	}
	return &Float{*big.NewFloat(defaultValue[0])}
}
func SetFloat(defaultValue ...*big.Float) *Float {
	if len(defaultValue) == 0 {
		return &Float{*big.NewFloat(0)}
	}
	return &Float{*safeFloat(defaultValue[0])}
}
func (b *Float) Clone(s *Float) *Float {
	if s == nil {
		return b
	}
	b.Float.Set(&s.Float)
	return b
}
func (b *Float) Set(s *big.Float) *Float {
	b.Float.Set(safeFloat(s))
	return b
}
func (b *Float) SetString(s string) *Float {
	if s == "" || s == "\"\"" {
		b.SetInt64(0)
		return b
	}
	b.Float.SetString(s)
	return b
}
func (b *Float) SetInt(s *Int) *Float {
	b.Float.SetInt(ToInt(s))
	return b
}
func (b *Float) SetUint64(s uint64) *Float {
	b.Float.SetUint64(s)
	return b
}
func (b *Float) Add(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	b.Float.Add(xAsFloat, yAsFloat)
	return b
}
func (b *Float) Sub(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	b.Float.Sub(xAsFloat, yAsFloat)
	return b
}
func (b *Float) Mul(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	b.Float.Mul(xAsFloat, yAsFloat)
	return b
}
func (b *Float) Quo(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	if y.IsZero() {
		b.SetUint64(0)
		return b
	}
	b.Float.Quo(xAsFloat, yAsFloat)
	return b
}
func (b *Float) Div(x *Float, y *Float) *Float {
	return b.Quo(x, y)
}
func (b *Float) Int() *Int {
	i, _ := b.Float.Int(nil)
	return SetInt(i)
}

func (b *Float) IsZero() bool {
	return b.Cmp(big.NewFloat(0)) == 0
}
func (b *Float) Safe(s *Float, defaultValue ...*Float) *Float {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}
func (b *Float) String() string {
	if b == nil {
		return `0`
	}
	return b.Text('f', -1)
}

func (b Float) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *Float) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Float
	_, ok := z.SetString(string(p))
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	b.Float = z
	return nil
}

func safeFloat(s *big.Float, defaultValue ...*big.Float) *big.Float {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}
