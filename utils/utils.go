package utils

import (
	"crypto/sha256"
	"fmt"
)

const Salt = "238de76823486e7f3d3024e3009a8c3d8b7657a844d73d35e361127c0b26cfd2"

func Ptr[T any](t T) *T {
	return &t
}

func Hash(value string) string {
	sum := sha256.Sum256([]byte(value + Salt))
	return fmt.Sprintf("%x", sum)
}
