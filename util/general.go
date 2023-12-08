package util

import "fmt"

func IsType(v any, expectedType string) bool {
	return fmt.Sprintf("%T", v) == expectedType
}
