package etc

import (
	"fmt"
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().UnixNano()) }

// RandomString generates a random alphanumeric string of length n. This
// function uses a random number generator that is not cryptographically
// secure.
func RandomString(n int) string {
	var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	buf := make([]byte, n)
	for i := 0; i < len(buf); i++ {
		buf[i] = alpha[rand.Intn(len(alpha)-1)]
	}
	return string(buf)
}

func CodeString(n int) string {
	var alpha = "ABCDEFGHIJKLMNPQRSTWXYZ"
	buf := make([]byte, n)
	for i := 0; i < len(buf); i++ {
		buf[i] = alpha[rand.Intn(len(alpha)-1)]
	}
	return string(buf)
}

func RandomEmail(n, m int) string {
	return fmt.Sprintf("%s@%s.com", RandomString(n), RandomString(m))
}
