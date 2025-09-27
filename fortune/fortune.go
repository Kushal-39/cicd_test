package fortune

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

// Load fortunes from file
func LoadFortunes(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	fortunes := strings.Split(string(data), "\n%\n")
	return fortunes, nil
}

// Pick a random fortune
func RandomFortune(fortunes []string) string {
	rand.Seed(time.Now().UnixNano())
	return fortunes[rand.Intn(len(fortunes))]
}
