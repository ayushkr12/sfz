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

func CreateFolderIfNotExists(folderPath string) error {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err := os.MkdirAll(folderPath, 0755) // 0755 is the permission mode for the directory (rwxr-xr-x)
		if err != nil {
			return fmt.Errorf("failed to create folder: %v", err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check if folder exists: %v", err)
	}
	return nil
}
