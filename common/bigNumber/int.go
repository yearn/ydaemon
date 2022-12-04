package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

type Int struct{ big.Int }

var Zero = NewInt(0)

/**************************************************************************************************
** Function to convert back a bigNumber.Int to a big.Int
**************************************************************************************************/
func ToInt(b *Int) *big.Int {
	if b == nil {
		return big.NewInt(0)
	}
	return &b.Int
}

/**************************************************************************************************
** Functions to initialize a bigNumber.Int from:
** - a big.Int
** - a int64
** - a uint64
**
** The safeInt function is used to avoid nil pointer.
**************************************************************************************************/
func NewInt(defaultValue ...int64) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*big.NewInt(defaultValue[0])}
}
func NewUint64(defaultValue ...uint64) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*big.NewInt(0).SetUint64(defaultValue[0])}
}
func SetInt(defaultValue ...*big.Int) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*safeInt(defaultValue[0])}
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

/**************************************************************************************************
** Functions to assign a valid value to an existing bigNumber.Int.
** - Clone from another bigNumber.Int (cf @dev)
** - Set from a big.Int
** - SetString from a string
** - SetUint64 from a uint64
** - Safe from another bigNumber.Int with a default value
**
** @dev
** Because we are working with pointers, we may need to clone the bigNumber value to avoid changing
** the original value.
**************************************************************************************************/
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
func (b *Int) SetUint64(s uint64) *Int {
	b.Int.SetUint64(s)
	return b
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

/**************************************************************************************************
** Override of some basic arithmetic functions from the Big package. The behavior is NOT the same,
** but it is more convenient for our use case.
**
** For most of the functions below, the second parameter is optional, unlike the Big package. If it
** is provided, only the first from this variadic parameter is used (y[0]) and then the behavior is
** the same as the Big package: b = x + y[0].
** If the second parameter is not provided, the function is applied to the current value of the
** bigNumber.Int: b = b + x.
**
** Supported overriden functions:
** - Add
** - Sub
** - Mul
** - Div
**************************************************************************************************/
func (b *Int) Add(x *Int, y ...*Int) *Int {
	xAsInt := ToInt(x)
	if len(y) > 0 {
		yAsInt := ToInt(y[0])
		b.Int.Add(xAsInt, yAsInt)
	} else {
		bAsInt := ToInt(b)
		b.Int.Add(bAsInt, xAsInt)
	}
	return b
}
func (b *Int) Sub(x *Int, y ...*Int) *Int {
	xAsInt := ToInt(x)
	if len(y) > 0 {
		yAsInt := ToInt(y[0])
		b.Int.Sub(xAsInt, yAsInt)
	} else {
		bAsInt := ToInt(b)
		b.Int.Sub(bAsInt, xAsInt)
	}
	return b
}
func (b *Int) Mul(x *Int, y ...*Int) *Int {
	xAsInt := ToInt(x)

	if len(y) > 0 {
		yAsInt := ToInt(y[0])
		b.Int.Mul(xAsInt, yAsInt)
	} else {
		bAsInt := ToInt(b)
		b.Int.Mul(bAsInt, xAsInt)
	}
	return b
}
func (b *Int) Div(x *Int, y ...*Int) *Int {
	xAsInt := ToInt(x)
	if len(y) > 0 {
		yAsInt := ToInt(y[0])
		if y[0].IsZero() {
			b.SetUint64(0)
			return b
		}
		b.Int.Div(xAsInt, yAsInt)
	} else {
		bAsInt := ToInt(b)
		if x.IsZero() {
			b.SetUint64(0)
			return b
		}
		b.Int.Div(bAsInt, xAsInt)
	}
	return b
}
func (b *Int) Exp(x *Int, y *Int, z *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	zAsInt := ToInt(z)
	b.Int.Exp(xAsInt, yAsInt, zAsInt)
	return b
}

/**************************************************************************************************
** Helpers functions that can be used to convert back a bigNumber.Int to a string or a uint64.
**************************************************************************************************/
func (b *Int) Uint64() uint64 {
	return ToInt(b).Uint64()
}
func (b *Int) String() string {
	return ToInt(b).String()
}

/**************************************************************************************************
** Theses functions are helpers to make it easier to work with the big.Int.Cmp function. This
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
func (b *Int) IsZero() bool {
	bAsInt := ToInt(b)
	return bAsInt.Cmp(big.NewInt(0)) == 0
}
func (b *Int) Gt(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) > 0
}
func (b *Int) Gte(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) >= 0
}
func (b *Int) Lt(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) < 0
}
func (b *Int) Lte(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) <= 0
}
func (b *Int) Eq(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) == 0
}
func (b *Int) Not(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) != 0
}

/**************************************************************************************************
** By default, the Big package marshals the big.Int to a number. This is not very convenient for
** most of the frontend libraries, which will be rekt by the number too big to be supported and
** expecting a string.
** Theses functions override the default behavior to marshal the big.Int to a string.
**************************************************************************************************/
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
