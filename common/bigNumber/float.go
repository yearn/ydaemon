package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
)

type Float struct{ big.Float }

/**************************************************************************************************
** Function to convert back a bigNumber.Float to a big.Float
**************************************************************************************************/
func ToFloat(b *Float) *big.Float {
	if b == nil {
		return big.NewFloat(0)
	}
	return &b.Float
}

/**************************************************************************************************
** Functions to initialize a bigNumber.Float from:
** - a big.Float
** - a float64
**************************************************************************************************/
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

/**************************************************************************************************
** Functions to assign a valid value to an existing bigNumber.Float.
** - Clone from another bigNumber.Float (cf @dev)
** - Set from a big.Float
** - SetString from a string
** - SetInt64 from a int64
** - SetUint64 from a uint64
** - SetInt from a bigNumber.Int
**
** @dev
** Because we are working with pointers, we may need to clone the bigNumber value to avoid changing
** the original value.
**************************************************************************************************/
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
func (b *Float) SetInt64(s int64) *Float {
	b.Float.SetInt64(s)
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

/**************************************************************************************************
** Replicate of some basic arithmetic functions from the Big package.
**************************************************************************************************/
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

/**************************************************************************************************
** Helpers functions that can be used to convert back a bigNumber.Float to a string or a BigInt.
**************************************************************************************************/
func (b *Float) Int() *Int {
	i, _ := b.Float.Int(nil)
	return SetInt(i)
}
func (b *Float) String() string {
	if b == nil {
		return `0`
	}
	return b.Text('f', -1)
}

/**************************************************************************************************
** Theses functions are helpers to make it easier to work with the big.Float.Cmp function. This
** function returns an int, which is not very convenient to use and to remember which value means
** what.
**
** The functions below are just wrappers to make it easier to use with:
** - IsZero (returns true if the current value b is equal to 0)
** - Gt (returns true if the current value b is greater than the provided value x)
** - Gte (returns true if the current value b is greater than or equal to the provided value x)
** - Lt (returns true if the current value b is less than the provided value x)
** - Lte (returns true if the current value b is less than or equal to the provided value x)
** - Eq (returns true if the current value b is equal to the provided value x)
** - Not (returns true if the current value b is not equal to the provided value x)
**************************************************************************************************/
func (b *Float) IsZero() bool {
	return b.Cmp(big.NewFloat(0)) == 0
}
func (b *Float) Gt(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) > 0
}
func (b *Float) Gte(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) >= 0
}
func (b *Float) Lt(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) < 0
}
func (b *Float) Lte(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) <= 0
}
func (b *Float) Eq(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) == 0
}
func (b *Float) Not(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) != 0
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

func safeFloat(s *big.Float, defaultValue ...*big.Float) *big.Float {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}

/**************************************************************************************************
** By default, the Big package marshals the big.Float to a number. This is not very convenient for
** most of the frontend libraries, which will be rekt by the number too big to be supported and
** expecting a string.
** Theses functions override the default behavior to marshal the big.Float to a string.
**************************************************************************************************/
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

// Convert the internal date as CSV string
func (b Float) MarshalCSV() (string, error) {
	return b.Text('f', -1), nil
}
