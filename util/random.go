package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var builder strings.Builder
	lenAlpha := len(alphabet)

	for i:=0; i < n; i++ {
		c := alphabet[rand.Intn(lenAlpha)]
		builder.WriteByte(c)
	}
	return builder.String()
}

func RandomOwnerName() string {
	return RandomString(6)
}

func RandomAmount() int64 {
	return RandomInt(0, 2000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "KES"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
