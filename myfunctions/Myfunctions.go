package myfunctions

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func main() {
}

//generates a random number
func RandomNumberGenerator() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

//generates a random string
func GenerateRandomString(max int64) string {
	b := make([]byte, max)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//checks if the args are correct
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
