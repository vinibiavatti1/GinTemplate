package utils

import (
	"crypto/sha256"
	"fmt"
)

func Ptr[T any](t T) *T {
	return &t
}

func Hash(value string) string {
	sum := sha256.Sum256([]byte(value))
	return fmt.Sprintf("%x", sum)
}
