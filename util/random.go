package util

import (
	"fmt"
	"math/rand"
	"strings"
	// "time"
)

func init() {
	// rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(5)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(10))
}

func RandomRole() string {
	role := []string{"Man", "Women", "Other"}
	n := len(role)
	return role[rand.Intn(n)]
}
