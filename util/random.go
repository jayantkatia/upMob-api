package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrtsuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInteger(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
func RandomDeviceName() string {
	return RandomString(6)
}

func RandomPrice() int32 {
	return int32(RandomInteger(1000, 1000000))
}
func RandomDate() string {
	return fmt.Sprint(RandomInteger(1, 31))
}
func RandomYear() string {
	return fmt.Sprint(RandomInteger(21, 23))
}
func RandomSpecScore() int32 {
	return int32(RandomInteger(0, 100))
}
