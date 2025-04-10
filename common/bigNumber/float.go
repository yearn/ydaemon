package bigNumber

import (
	"encoding/json"
	"fmt"
	"math"
	"math/big"
)

/**************************************************************************************************
** The bigNumber package provides a wrapper around Go's standard library big.Int and big.Float
** types to enhance usability when working with arbitrary-precision arithmetic. These wrappers
** add method chaining, improved error handling, and JSON serialization optimized for blockchain
** applications.
**
** This file contains the Float type, which extends big.Float with blockchain-friendly features.
**************************************************************************************************/

/**************************************************************************************************
** Float is a wrapper type around Go's big.Float that provides enhanced functionality:
** - Method chaining for fluent API usage
** - Simplified arithmetic operations
** - Safe handling of nil pointers
** - JSON serialization to string format for easy interchange with web applications
** - Comparison helper methods for intuitive value checking
**************************************************************************************************/
type Float struct{ big.Float }

/**************************************************************************************************
** ToFloat converts a bigNumber.Float back to a standard library big.Float.
** This function provides a bridge to the standard library when needed.
** If the input is nil, a zero-initialized big.Float is returned.
**
** @param b The bigNumber.Float to convert
** @return *big.Float The equivalent standard library big.Float
**************************************************************************************************/
func ToFloat(b *Float) *big.Float {
	if b == nil {
		return big.NewFloat(0)
	}
	return &b.Float
}

/**************************************************************************************************
** NewFloat creates a new bigNumber.Float initialized with the provided value.
** If no value is provided, it initializes to zero.
**
** @param defaultValue Optional initial float64 value (defaults to 0 if omitted)
** @return *Float A new bigNumber.Float initialized to the specified value
**************************************************************************************************/
func NewFloat(defaultValue ...float64) *Float {
	if len(defaultValue) == 0 {
		return &Float{*big.NewFloat(0)}
	}
	return &Float{*big.NewFloat(defaultValue[0])}
}

/**************************************************************************************************
** SetFloat creates a new bigNumber.Float from an existing big.Float.
** If no value is provided, or if the provided value is nil, it initializes to zero.
**
** @param defaultValue Optional existing big.Float to use as the initial value
** @return *Float A new bigNumber.Float initialized with the value of the provided big.Float
**************************************************************************************************/
func SetFloat(defaultValue ...*big.Float) *Float {
	if len(defaultValue) == 0 {
		return &Float{*big.NewFloat(0)}
	}
	return &Float{*safeFloat(defaultValue[0])}
}

/**************************************************************************************************
** Clone copies the value from another bigNumber.Float into the current instance.
** This is useful for working with pointers without modifying the original value.
** If the source is nil, the current value is left unchanged.
**
** @param s The source bigNumber.Float to copy from
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) Clone(s *Float) *Float {
	if s == nil {
		return b
	}
	b.Float.Set(&s.Float)
	return b
}

/**************************************************************************************************
** Set assigns a value from a standard library big.Float to the current bigNumber.Float.
** This function safely handles nil inputs and returns the modified object for chaining.
**
** @param s The big.Float to copy the value from
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) Set(s *big.Float) *Float {
	b.Float.Set(safeFloat(s))
	return b
}

/**************************************************************************************************
** SetString parses a string representation of a number and assigns it to the current Float.
** This function handles empty strings and quoted empty strings by setting the value to zero.
**
** @param s The string representation of a number
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) SetString(s string) *Float {
	if s == "" || s == "\"\"" {
		b.SetInt64(0)
		return b
	}
	b.Float.SetString(s)
	return b
}

/**************************************************************************************************
** SetInt64 assigns an int64 value to the current bigNumber.Float.
**
** @param s The int64 value to assign
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) SetInt64(s int64) *Float {
	b.Float.SetInt64(s)
	return b
}

/**************************************************************************************************
** SetFloat64 assigns a float64 value to the current bigNumber.Float.
**
** @param s The float64 value to assign
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) SetFloat64(s float64) *Float {
	b.Float.SetFloat64(s)
	return b
}

/**************************************************************************************************
** SetInt converts a bigNumber.Int to a bigNumber.Float and assigns the value.
**
** @param s The bigNumber.Int to convert and assign
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) SetInt(s *Int) *Float {
	b.Float.SetInt(ToInt(s))
	return b
}

/**************************************************************************************************
** SetUint64 assigns a uint64 value to the current bigNumber.Float.
**
** @param s The uint64 value to assign
** @return *Float The current bigNumber.Float with the updated value (for method chaining)
**************************************************************************************************/
func (b *Float) SetUint64(s uint64) *Float {
	b.Float.SetUint64(s)
	return b
}

/**************************************************************************************************
** Add calculates the sum of two bigNumber.Float values and stores the result in the receiver.
** This function safely handles nil inputs by converting them to zero.
**
** @param x The first operand
** @param y The second operand
** @return *Float The result of x + y (for method chaining)
**************************************************************************************************/
func (b *Float) Add(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	b.Float.Add(xAsFloat, yAsFloat)
	return b
}

/**************************************************************************************************
** Sub calculates the difference between two bigNumber.Float values and stores the result in the receiver.
** This function safely handles nil inputs by converting them to zero.
**
** @param x The minuend (value to subtract from)
** @param y The subtrahend (value to subtract)
** @return *Float The result of x - y (for method chaining)
**************************************************************************************************/
func (b *Float) Sub(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	b.Float.Sub(xAsFloat, yAsFloat)
	return b
}

/**************************************************************************************************
** Mul calculates the product of two bigNumber.Float values and stores the result in the receiver.
** This function safely handles nil inputs by converting them to zero.
**
** @param x The first operand
** @param y The second operand
** @return *Float The result of x * y (for method chaining)
**************************************************************************************************/
func (b *Float) Mul(x *Float, y *Float) *Float {
	xAsFloat := ToFloat(x)
	yAsFloat := ToFloat(y)
	b.Float.Mul(xAsFloat, yAsFloat)
	return b
}

/**************************************************************************************************
** Quo calculates the quotient of two bigNumber.Float values and stores the result in the receiver.
** This function safely handles nil inputs and division by zero by returning zero in those cases.
**
** @param x The dividend (numerator)
** @param y The divisor (denominator)
** @return *Float The result of x / y, or zero if y is zero (for method chaining)
**************************************************************************************************/
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

/**************************************************************************************************
** Div is an alias for Quo, calculating the quotient of two bigNumber.Float values.
** This provides a more familiar name for division operations.
**
** @param x The dividend (numerator)
** @param y The divisor (denominator)
** @return *Float The result of x / y, or zero if y is zero (for method chaining)
**************************************************************************************************/
func (b *Float) Div(x *Float, y *Float) *Float {
	return b.Quo(x, y)
}

/**************************************************************************************************
** Pow calculates x raised to the power of y and stores the result in the receiver.
** This function implements exponentiation through repeated multiplication.
**
** @param x The base
** @param y The exponent (unsigned integer)
** @return *Float The result of x^y (for method chaining)
**************************************************************************************************/
func (b *Float) Pow(x *Float, y uint64) *Float {
	result := NewFloat(1)
	for i := uint64(0); i < y; i++ {
		result.Mul(result, x)
	}
	return result
}

/**************************************************************************************************
** SetMode sets the rounding mode for operations on the current bigNumber.Float.
** This controls how results are rounded when they cannot be represented exactly.
**
** @param mode The rounding mode to use from the big package
** @return *Float The current bigNumber.Float with the updated rounding mode (for method chaining)
**************************************************************************************************/
func (b *Float) SetMode(mode big.RoundingMode) *Float {
	b.Float.SetMode(mode)
	return b
}

/**************************************************************************************************
** Int converts the bigNumber.Float to a bigNumber.Int, truncating any fractional part.
**
** @return *Int The integer part of the Float value as a bigNumber.Int
**************************************************************************************************/
func (b *Float) Int() *Int {
	i, _ := b.Float.Int(nil)
	return SetInt(i)
}

/**************************************************************************************************
** String converts the bigNumber.Float to its string representation.
** If the receiver is nil, "0" is returned to prevent nil pointer dereference.
**
** @return string The string representation of the Float value
**************************************************************************************************/
func (b *Float) String() string {
	if b == nil {
		return `0`
	}
	return b.Text('f', -1)
}

/**************************************************************************************************
** IsZero checks if the current bigNumber.Float value is equal to zero.
**
** @return bool True if the value is zero, false otherwise
**************************************************************************************************/
func (b *Float) IsZero() bool {
	if b == nil {
		return true
	}
	return b.Cmp(big.NewFloat(0)) == 0
}

/**************************************************************************************************
** Gt checks if the current bigNumber.Float is greater than the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is greater than x, false otherwise
**************************************************************************************************/
func (b *Float) Gt(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) > 0
}

/**************************************************************************************************
** Gte checks if the current bigNumber.Float is greater than or equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is greater than or equal to x, false otherwise
**************************************************************************************************/
func (b *Float) Gte(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) >= 0
}

/**************************************************************************************************
** Lt checks if the current bigNumber.Float is less than the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is less than x, false otherwise
**************************************************************************************************/
func (b *Float) Lt(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) < 0
}

/**************************************************************************************************
** Lte checks if the current bigNumber.Float is less than or equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is less than or equal to x, false otherwise
**************************************************************************************************/
func (b *Float) Lte(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) <= 0
}

/**************************************************************************************************
** Eq checks if the current bigNumber.Float is equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is equal to x, false otherwise
**************************************************************************************************/
func (b *Float) Eq(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) == 0
}

/**************************************************************************************************
** Not checks if the current bigNumber.Float is not equal to the provided value.
** This is a convenience wrapper around the Cmp method.
**
** @param x The value to compare against
** @return bool True if the current value is not equal to x, false otherwise
**************************************************************************************************/
func (b *Float) Not(x *Float) bool {
	bAsFloat := ToFloat(b)
	xAsFloat := ToFloat(x)
	return bAsFloat.Cmp(xAsFloat) != 0
}

/**************************************************************************************************
** Safe provides a nil-safe way to access a bigNumber.Float.
** If the provided Float is nil, a default value is returned instead.
**
** @param s The Float to check for nil
** @param defaultValue Optional default value to return if s is nil
** @return *Float The original Float if not nil, otherwise the default value
**************************************************************************************************/
func (b *Float) Safe(s *Float, defaultValue ...*Float) *Float {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}

/**************************************************************************************************
** safeFloat provides a nil-safe way to access a standard library big.Float.
** If the provided big.Float is nil, a default value is returned instead.
**
** @param s The big.Float to check for nil
** @param defaultValue Optional default value to return if s is nil
** @return *big.Float The original big.Float if not nil, otherwise the default value
**************************************************************************************************/
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
** MarshalJSON customizes the JSON serialization of bigNumber.Float.
** While the standard big.Float marshals to a number, this implementation marshals to a float64,
** which is more compatible with JavaScript and frontend libraries.
**
** @return []byte The JSON representation of the Float as a float64
** @return error Any error encountered during marshaling
**************************************************************************************************/
func (b *Float) MarshalJSON() ([]byte, error) {
	if b == nil {
		return json.Marshal(big.NewFloat(0).String())
	}
	if b.IsInf() {
		return json.Marshal(big.NewFloat(0).String())
	}
	toFloat64, _ := b.Float64()
	return json.Marshal(toFloat64)
}

/**************************************************************************************************
** UnmarshalJSON customizes the JSON deserialization of bigNumber.Float.
** This allows parsing numeric JSON values into a bigNumber.Float.
**
** @param p The JSON data to unmarshal
** @return error Any error encountered during unmarshaling
**************************************************************************************************/
func (b *Float) UnmarshalJSON(p []byte) error {
	if string(p) == "null" {
		return nil
	}
	var z big.Float
	_, ok := z.SetString(string(p))
	if !ok {
		return fmt.Errorf("not a valid big integer: %s", p)
	}
	if z.IsInf() {
		b.Float = *big.NewFloat(math.MaxFloat64)
		return nil
	}
	b.Float = z
	return nil
}

/**************************************************************************************************
** MarshalCSV provides CSV serialization for bigNumber.Float.
** This converts the Float to a string representation suitable for CSV files.
**
** @return string The string representation of the Float for CSV
** @return error Any error encountered during marshaling
**************************************************************************************************/
func (b Float) MarshalCSV() (string, error) {
	return b.Text('f', -1), nil
}
