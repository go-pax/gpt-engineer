package io

import (
	"crypto/rand"
	"fmt"
)

func randString(n int) string {
	const alphanumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanumeric[b%byte(len(alphanumeric))]
	}
	return string(bytes)
}

func GenerateRandomFilename() string {
	return fmt.Sprintf("filename-unknown-%s.txt", randString(4))
}
