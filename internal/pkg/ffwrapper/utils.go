package ffwrapper

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// charset is the set of characters used to generate the random string
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// RandomString generates a random string of the specified length
func GenerateRandomString(length int) string {
	rand.New(rand.NewSource(time.Now().UnixNano())) // seed the random number generator with the current time
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // select a random character from the charset
	}
	return string(result)
}

func GetDirContents(dirPath string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, fmt.Errorf("")
	}
	return entries, nil
}
