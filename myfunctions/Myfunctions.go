package myfunctions

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func main() {
}
func RandomNumberGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func GenerateRandomString(max int64) string {
	b := make([]byte, max)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
func CheckArgs(args string) bool {
	switch args {
	case
		"search",
		"list",
		"buy",
		"delete":
		return false
	}
	return true
}
