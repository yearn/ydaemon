package helpers

import "strconv"

/** 🔵 - Yearn *************************************************************************************
** SafeString will return s if s is a valid string, otherwise it will use
** the defaultValue
**************************************************************************************************/
func SafeString(s string, defaultValue string) string {
	if s == "" || s == "\"\"" {
		return defaultValue
	}
	return s
}

/** 🔵 - Yearn *************************************************************************************
** SafeStringToUint64 will return n if n is a valid uint64, otherwise it will use
** the defaultValue
**************************************************************************************************/
func SafeStringToUint64(n string, defaultValue uint64) uint64 {
	if n == "" || n == "\"\"" {
		return defaultValue
	}
	result, err := strconv.ParseUint(n, 10, 64)
	if err != nil {
		return defaultValue
	}
	return result
}
