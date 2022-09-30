package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type BigFloat struct{ big.Float }

func FromBigFloat(b *big.Float) *BigFloat {
	return &BigFloat{*b}
}
func ToBigFloat(b *BigFloat) *big.Float {
	if b == nil {
		return big.NewFloat(0)
	}
	return &b.Float
}

func NewFloat(defaultValue ...float64) *BigFloat {
	if len(defaultValue) == 0 {
		return &BigFloat{*big.NewFloat(0)}
	}
	return &BigFloat{*big.NewFloat(defaultValue[0])}
}
func SetFloat(defaultValue ...*big.Float) *BigFloat {
	if len(defaultValue) == 0 {
		return &BigFloat{*big.NewFloat(0)}
	}
	return &BigFloat{*safeBigFloat(defaultValue[0])}
}
func (b *BigFloat) Set(s *big.Float) *BigFloat {
	b.Float.Set(safeBigFloat(s))
	return b
}
func (b *BigFloat) SetString(s string) *BigFloat {
	if s == "" || s == "\"\"" {
		b.SetInt64(0)
		return b
	}
	b.Float.SetString(s)
	return b
}
func (b *BigFloat) Add(x *BigFloat, y *BigFloat) *BigFloat {
	xAsBigFloat := ToBigFloat(x)
	yAsBigFloat := ToBigFloat(y)
	b.Float.Add(xAsBigFloat, yAsBigFloat)
	return b
}
func (b *BigFloat) Sub(x *BigFloat, y *BigFloat) *BigFloat {
	xAsBigFloat := ToBigFloat(x)
	yAsBigFloat := ToBigFloat(y)
	b.Float.Sub(xAsBigFloat, yAsBigFloat)
	return b
}
func (b *BigFloat) Mul(x *BigFloat, y *BigFloat) *BigFloat {
	xAsBigFloat := ToBigFloat(x)
	yAsBigFloat := ToBigFloat(y)
	b.Float.Mul(xAsBigFloat, yAsBigFloat)
	return b
}
func (b *BigFloat) Quo(x *BigFloat, y *BigFloat) *BigFloat {
	xAsBigFloat := ToBigFloat(x)
	yAsBigFloat := ToBigFloat(y)
	if y.IsZero() {
		b.SetUint64(0)
		return b
	}
	b.Float.Quo(xAsBigFloat, yAsBigFloat)
	return b
}

func (b *BigFloat) IsZero() bool {
	return b.Cmp(big.NewFloat(0)) == 0
}
func (b *BigFloat) Safe(s *BigFloat, defaultValue ...*BigFloat) *BigFloat {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}

func (b BigFloat) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.String())
}

func (b *BigFloat) UnmarshalJSON(p []byte) error {
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

func safeBigFloat(s *big.Float, defaultValue ...*big.Float) *big.Float {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}
