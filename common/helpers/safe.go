package helpers

// SafeString will return s if s is a valid string, otherwise it will use
// the defaultValue
func SafeString(s string, defaultValue string) string {
	if s == "" || s == "\"\"" {
		return defaultValue
	}
	return s
}
