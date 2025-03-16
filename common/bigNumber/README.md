# BigNumber Package

## Overview

The `bigNumber` package provides a robust wrapper around Go's standard library `big.Int` and `big.Float` types for arbitrary-precision arithmetic operations. This package is specifically designed for blockchain applications where precise handling of large numbers is critical, such as token amounts, prices, and exchange rates.

The package offers two primary types:

- `Int`: A wrapper around `big.Int` for arbitrary-precision integer operations
- `Float`: A wrapper around `big.Float` for arbitrary-precision floating-point operations

## Features

Both the `Int` and `Float` types provide several enhancements over the standard library types:

- **Method Chaining**: Fluent API for combining multiple operations
- **Simplified Arithmetic**: Dual-mode operation for flexible usage patterns
- **Nil-Safety**: Built-in handling of nil pointers to prevent panics
- **JSON Serialization**: Customized JSON marshaling/unmarshaling that produces string representations for compatibility with JavaScript
- **Comparison Helpers**: Intuitive methods for value comparison
- **CSV Compatibility**: Support for CSV serialization (Float type)

## Usage Examples

### Working with Integers

```go
// Create new integer values
amount := bigNumber.NewInt(1000000000000000000) // 1 ETH in wei
balance := bigNumber.NewInt()                   // Initialized to 0

// Method chaining for calculations
result := bigNumber.NewInt().Mul(amount, bigNumber.NewInt(2)).Div(bigNumber.NewInt(4))
// Equivalent to: (amount * 2) / 4

// Comparison helpers
if amount.Gt(balance) {
    // amount is greater than balance
}

// Handling nil-safety
var maybeNilInt *bigNumber.Int
safeInt := bigNumber.NewInt().Safe(maybeNilInt) // Returns default (0) if nil

// JSON serialization (outputs as string to prevent overflow in JavaScript)
jsonData, _ := json.Marshal(amount) // Will output "1000000000000000000" as string
```

### Working with Floating Point Numbers

```go
// Create new float values
price := bigNumber.NewFloat(1234.56789)
ratio := bigNumber.NewFloat().SetString("0.001234567890123456789") // High precision

// Setting decimal precision
price.SetMode(128) // Set precision to 128 bits

// Arithmetic operations
weeklyYield := bigNumber.NewFloat()
weeklyYield.Mul(price, ratio) // Direct operation on existing values

// Converting between Int and Float
priceAsInt := price.Int()
intAsFloat := bigNumber.NewFloat().SetInt(bigNumber.NewInt(1000))

// Comparing values
if price.Gte(bigNumber.NewFloat(1000)) {
    // price is greater than or equal to 1000
}
```

## Design Principles

The `bigNumber` package follows these design principles:

1. **Immutability**: Original values are preserved unless explicitly modified
2. **Safety First**: Automatic handling of edge cases like nil pointers and division by zero
3. **Consistency**: Method names and behaviors match across Int and Float types
4. **Usability**: Fluent API with method chaining for readable code
5. **Compatibility**: String-based serialization for JavaScript interoperability

## Common Operations

### Creation and Initialization

```go
// Create from constants
n1 := bigNumber.NewInt(42)
f1 := bigNumber.NewFloat(3.14159)

// Create from strings (useful for values exceeding native type limits)
n2 := bigNumber.NewInt().SetString("115792089237316195423570985008687907853269984665640564039457584007913129639935")
f2 := bigNumber.NewFloat().SetString("0.0000000000000000000000000000000000000000000000000000000001")

// Create from other big number types
n3 := bigNumber.SetInt(new(big.Int).SetInt64(100))
f3 := bigNumber.SetFloat(new(big.Float).SetFloat64(100.5))
```

### Arithmetic Operations

Both `Int` and `Float` types support the following operations:

```go
// Single parameter mode (operate on self)
value.Add(x)      // value = value + x
value.Sub(x)      // value = value - x
value.Mul(x)      // value = value * x
value.Div(x)      // value = value / x (safely handles division by zero)

// Two parameter mode (assign result of operation)
value.Add(x, y)   // value = x + y
value.Sub(x, y)   // value = x - y
value.Mul(x, y)   // value = x * y
value.Div(x, y)   // value = x / y (safely handles division by zero)
```

### Comparison

```go
a.IsZero()        // true if a == 0
a.Gt(b)           // true if a > b
a.Gte(b)          // true if a >= b
a.Lt(b)           // true if a < b
a.Lte(b)          // true if a <= b
a.Eq(b)           // true if a == b
a.Not(b)          // true if a != b
```

## Best Practices

1. **Use Factory Methods**: Prefer factory methods like `NewInt()` and `NewFloat()` instead of direct struct initialization to ensure proper setup.

2. **Method Chaining**: Take advantage of method chaining for concise calculations:

    ```go
    result := bigNumber.NewInt().Mul(bigNumber.NewInt(10), bigNumber.NewInt(20)).Add(bigNumber.NewInt(5))
    ```

3. **Handle Division Carefully**: While the package handles division by zero safely, consider checking divisors manually for more explicit error handling.

4. **Set Appropriate Precision**: For Float calculations, set an appropriate precision using `SetMode(precision)` based on your requirements.

5. **Check Conversion Limits**: When converting between types (especially to built-in Go types), be aware of potential overflow.

## Integration with Blockchain Applications

This package is especially useful for blockchain applications where:

- Token amounts exceed JavaScript's safe integer limits
- Precise decimal calculations are needed for financial operations
- Large prime numbers are required for cryptographic operations
- Non-decimal (hex) encodings are commonly used

By using the `bigNumber` package, you can safely handle blockchain numeric values without precision loss or overflow errors.
