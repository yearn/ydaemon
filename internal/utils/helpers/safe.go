package helpers

import "math/big"

// SafeBigFloat will return s if s is a valid bigFloat, otherwise it will use
// the first value of the variadic defaultValue parameter if this one is valid
func SafeBigFloat(s *big.Float, defaultValue ...*big.Float) *big.Float {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewFloat(0)
		}
		return defaultValue[0]
	}
	return s
}

// SafeBigInt will return s if s is a valid bigInt, otherwise it will use
// the first value of the variadic defaultValue parameter if this one is valid
func SafeBigInt(s *big.Int, defaultValue ...*big.Int) *big.Int {
	if s == nil {
		if len(defaultValue) == 0 || defaultValue[0] == nil {
			return big.NewInt(0)
		}
		return defaultValue[0]
	}
	return s
}

// SafeString will return s if s is a valid string, otherwise it will use
// the defaultValue
func SafeString(s string, defaultValue string) string {
	if s == "" || s == "\"\"" {
		return defaultValue
	}
	return s
}
