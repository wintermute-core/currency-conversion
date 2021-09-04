package env

import "os"

// Env - Fetch environment variable
func Env(name string, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		return defaultValue
	}

	return v
}

// IsDefined - check if environment variable is defined
func IsDefined(name string) bool {
	v := os.Getenv(name)
	return v != ""
}
