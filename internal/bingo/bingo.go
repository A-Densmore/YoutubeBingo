package bingo

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GetRandomItem(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if len(lines) == 0 {
		return "", nil
	}

	randomIndex := rand.Intn(len(lines))
	selectedLine := lines[randomIndex]

	// Replace X with random digits
	result := ""
	for _, char := range selectedLine {
		if char == 'X' || char == 'x' {
			result += fmt.Sprintf("%d", rand.Intn(10))
		} else {
			result += string(char)
		}
	}

	return result, nil
}
