package handlers

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/templates/index.html")
}

func GetRandomBingoItem(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	filepath := "list.txt"

	file, err := os.Open(filepath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		http.Error(w, fmt.Sprintf("Error reading file: %v", err), http.StatusInternalServerError)
		return
	}

	if len(lines) == 0 {
		http.Error(w, "File is empty", http.StatusInternalServerError)
		return
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

	fmt.Fprintln(w, result)
}
