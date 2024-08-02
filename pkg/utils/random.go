package utils

import (
	"math/rand"
	"strings"
)

const (
	chars = "0123456789abcefghijklmnopqrstuvwxyz"
)

func GenerateRandomCode(length int) string {
	var sb strings.Builder

	k := len(chars)
	for i := 0; i < length; i++ {
		sb.WriteByte(chars[rand.Intn(k)])
	}

	return strings.ToUpper(sb.String())
}
