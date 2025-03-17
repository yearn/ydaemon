package bigNumber

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
)

/**************************************************************************************************
** The Int type provides a wrapper around Go's standard library big.Int to enhance usability
** when working with arbitrary-precision integers in blockchain applications. This wrapper
** adds method chaining, improved error handling, and better JSON serialization.
**
** This file contains implementations for working with integers of unlimited precision,
** which is essential for representing token amounts and other large numeric values common
** in blockchain environments.
**************************************************************************************************/

/**************************************************************************************************
** Int is a wrapper type around Go's big.Int that provides enhanced functionality:
** - Method chaining for fluent API usage
** - Simplified arithmetic operations
** - Safe handling of nil pointers
** - JSON serialization to string format for easy interchange with web applications
** - Comparison helper methods for intuitive value checking
**************************************************************************************************/
type Int struct{ big.Int }

/**************************************************************************************************
** Zero is a pre-initialized bigNumber.Int with a value of 0.
** This provides a convenient constant for comparison and initialization.
**************************************************************************************************/
var Zero = NewInt(0)

/**************************************************************************************************
** ToInt converts a bigNumber.Int back to a standard library big.Int.
** This function provides a bridge to the standard library when needed.
** If the input is nil, a zero-initialized big.Int is returned.
**
** @param b The bigNumber.Int to convert
** @return *big.Int The equivalent standard library big.Int
**************************************************************************************************/
func ToInt(b *Int) *big.Int {
	if b == nil {
		return big.NewInt(0)
	}
	return &b.Int
}

/**************************************************************************************************
** NewInt creates a new bigNumber.Int initialized with the provided value.
** If no value is provided, it initializes to zero.
**
** @param defaultValue Optional initial int64 value (defaults to 0 if omitted)
** @return *Int A new bigNumber.Int initialized to the specified value
**************************************************************************************************/
func NewInt(defaultValue ...int64) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*big.NewInt(defaultValue[0])}
}

/**************************************************************************************************
** NewUint64 creates a new bigNumber.Int initialized with the provided uint64 value.
** If no value is provided, it initializes to zero.
**
** @param defaultValue Optional initial uint64 value (defaults to 0 if omitted)
** @return *Int A new bigNumber.Int initialized to the specified value
**************************************************************************************************/
func NewUint64(defaultValue ...uint64) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*big.NewInt(0).SetUint64(defaultValue[0])}
}

/**************************************************************************************************
** SetInt creates a new bigNumber.Int from an existing big.Int.
** If no value is provided, or if the provided value is nil, it initializes to zero.
**
** @param defaultValue Optional existing big.Int to use as the initial value
** @return *Int A new bigNumber.Int initialized with the value of the provided big.Int
**************************************************************************************************/
func SetInt(defaultValue ...*big.Int) *Int {
	if len(defaultValue) == 0 {
		return &Int{*big.NewInt(0)}
	}
	return &Int{*safeInt(defaultValue[0])}
}

/**************************************************************************************************
** safeInt provides a nil-safe way to access a standard library big.Int.
** If the provided big.Int is nil, a default value is returned instead.
**
** @param s The big.Int to check for nil
** @param defaultValue Optional default value to return if s is nil
** @return *big.Int The original big.Int if not nil, otherwise the default value
**************************************************************************************************/
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
** Clone copies the value from another bigNumber.Int into the current instance.
** This is useful for working with pointers without modifying the original value.
** If the source is nil, the current value is left unchanged.
**
** @param s The source bigNumber.Int to copy from
** @return *Int The current bigNumber.Int with the updated value (for method chaining)
**************************************************************************************************/
func (b *Int) Clone(s *Int) *Int {
	if s == nil {
		return b
	}
	b.Int.Set(&s.Int)
	return b
}

/**************************************************************************************************
** Set assigns a value from a standard library big.Int to the current bigNumber.Int.
** This function safely handles nil inputs and returns the modified object for chaining.
**
** @param s The big.Int to copy the value from
** @return *Int The current bigNumber.Int with the updated value (for method chaining)
**************************************************************************************************/
func (b *Int) Set(s *big.Int) *Int {
	b.Int.Set(safeInt(s))
	return b
}

/**************************************************************************************************
** SetString parses a string representation of a number and assigns it to the current Int.
** This function handles empty strings and quoted empty strings by setting the value to zero.
**
** @param s The string representation of a number
** @return *Int The current bigNumber.Int with the updated value (for method chaining)
**************************************************************************************************/
func (b *Int) SetString(s string) *Int {
	if s == "" || s == "\"\"" {
		b.SetInt64(0)
		return b
	}
	b.Int.SetString(s, 10)
	return b
}

/**************************************************************************************************
** SetUint64 assigns a uint64 value to the current bigNumber.Int.
**
** @param s The uint64 value to assign
** @return *Int The current bigNumber.Int with the updated value (for method chaining)
**************************************************************************************************/
func (b *Int) SetUint64(s uint64) *Int {
	b.Int.SetUint64(s)
	return b
}

/**************************************************************************************************
** Safe provides a nil-safe way to access a bigNumber.Int.
** If the provided Int is nil, a default value is returned instead.
**
** @param s The Int to check for nil
** @param defaultValue Optional default value to return if s is nil
** @return *Int The original Int if not nil, otherwise the default value
**************************************************************************************************/
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
** Add calculates the sum of bigNumber.Int values and stores the result in the receiver.
** This function has two modes of operation:
** 1. If only one parameter is provided: b = b + x
** 2. If two parameters are provided: b = x + y[0]
**
** This dual behavior allows for more flexible and concise arithmetic operations.
**
** @param x The first operand
** @param y Optional second operand
** @return *Int The result of the addition (for method chaining)
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

/**************************************************************************************************
** Sub calculates the difference between bigNumber.Int values and stores the result in the receiver.
** This function has two modes of operation:
** 1. If only one parameter is provided: b = b - x
** 2. If two parameters are provided: b = x - y[0]
**
** @param x The first operand (minuend if y is provided, or subtrahend if y is omitted)
** @param y Optional second operand (subtrahend)
** @return *Int The result of the subtraction (for method chaining)
**************************************************************************************************/
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

/**************************************************************************************************
** Mul calculates the product of bigNumber.Int values and stores the result in the receiver.
** This function has two modes of operation:
** 1. If only one parameter is provided: b = b * x
** 2. If two parameters are provided: b = x * y[0]
**
** @param x The first operand
** @param y Optional second operand
** @return *Int The result of the multiplication (for method chaining)
**************************************************************************************************/
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

/**************************************************************************************************
** Div calculates the quotient of bigNumber.Int values and stores the result in the receiver.
** This function has two modes of operation:
** 1. If only one parameter is provided: b = b / x (with checks for division by zero)
** 2. If two parameters are provided: b = x / y[0] (with checks for division by zero)
**
** @param x The first operand (dividend if y is provided, or divisor if y is omitted)
** @param y Optional second operand (divisor)
** @return *Int The result of the division, or zero if dividing by zero (for method chaining)
**************************************************************************************************/
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

/**************************************************************************************************
** Exp calculates x^y mod |z| (modular exponentiation) and stores the result in the receiver.
** This is particularly useful for cryptographic operations.
**
** @param x The base value
** @param y The exponent value
** @param z The modulus value
** @return *Int The result of the modular exponentiation (for method chaining)
**************************************************************************************************/
func (b *Int) Exp(x *Int, y *Int, z *Int) *Int {
	xAsInt := ToInt(x)
	yAsInt := ToInt(y)
	zAsInt := ToInt(z)
	b.Int.Exp(xAsInt, yAsInt, zAsInt)
	return b
}

/**************************************************************************************************
** Uint64 converts the bigNumber.Int to a uint64 value.
** This is useful when standard Go integer types are required.
** Note that this may overflow if the value is larger than can be represented by uint64.
**
** @return uint64 The value of the bigNumber.Int as a uint64
**************************************************************************************************/
func (b *Int) Uint64() uint64 {
	return ToInt(b).Uint64()
}

/**************************************************************************************************
** String converts the bigNumber.Int to its string representation.
**
** @return string The string representation of the Int value
**************************************************************************************************/
func (b *Int) String() string {
	return ToInt(b).String()
}

/**************************************************************************************************
** IsZero checks if the current bigNumber.Int value is equal to zero.
**
** @return bool True if the value is zero, false otherwise
**************************************************************************************************/
func (b *Int) IsZero() bool {
	bAsInt := ToInt(b)
	return bAsInt.Cmp(big.NewInt(0)) == 0
}

/**************************************************************************************************
** Gt checks if the current bigNumber.Int is greater than the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is greater than x, false otherwise
**************************************************************************************************/
func (b *Int) Gt(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) > 0
}

/**************************************************************************************************
** Gte checks if the current bigNumber.Int is greater than or equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is greater than or equal to x, false otherwise
**************************************************************************************************/
func (b *Int) Gte(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) >= 0
}

/**************************************************************************************************
** Lt checks if the current bigNumber.Int is less than the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is less than x, false otherwise
**************************************************************************************************/
func (b *Int) Lt(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) < 0
}

/**************************************************************************************************
** Lte checks if the current bigNumber.Int is less than or equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is less than or equal to x, false otherwise
**************************************************************************************************/
func (b *Int) Lte(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) <= 0
}

/**************************************************************************************************
** Eq checks if the current bigNumber.Int is equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is equal to x, false otherwise
**************************************************************************************************/
func (b *Int) Eq(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) == 0
}

/**************************************************************************************************
** Not checks if the current bigNumber.Int is not equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is not equal to x, false otherwise
**************************************************************************************************/
func (b *Int) Not(x *Int) bool {
	bAsInt := ToInt(b)
	xAsInt := ToInt(x)
	return bAsInt.Cmp(xAsInt) != 0
}

/**************************************************************************************************
** MarshalJSON customizes the JSON serialization of bigNumber.Int.
** While the standard big.Int marshals to a number, this implementation marshals to a string,
** which is more compatible with JavaScript and frontend libraries that cannot handle very
** large integer values.
**
** @return []byte The JSON representation of the Int as a string
** @return error Any error encountered during marshaling
**************************************************************************************************/
func (b *Int) MarshalJSON() ([]byte, error) {
	if b == nil {
		return json.Marshal(big.NewInt(0).String())
	}
	return json.Marshal(b.String())
}

/**************************************************************************************************
** UnmarshalJSON customizes the JSON deserialization of bigNumber.Int.
** This allows parsing numeric JSON values into a bigNumber.Int, handling string-encoded
** large integers that are common in blockchain applications.
**
** @param p The JSON data to unmarshal
** @return error Any error encountered during unmarshaling
**************************************************************************************************/
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
