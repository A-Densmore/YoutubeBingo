package handlers

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../../web/templates/index.html")
}

func GetRandomBingoItem(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	filepath := "../../list.txt"

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

	// Generate a random date between January 1, 2009 and today
	startDate := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now()
	randomDays := rand.Intn(int(endDate.Sub(startDate).Hours()/24)) + 1
	randomDate := startDate.AddDate(0, 0, randomDays)

	dateYYYYMMDD := randomDate.Format("20060102")
	dateYYYY := randomDate.Format("2006")

	// Replace date placeholders with random dates
	dateYYYYMMDD_spaced := randomDate.Format("2006 01 02")
	result := regexp.MustCompile(`YYYY MM DD`).ReplaceAllString(selectedLine, dateYYYYMMDD_spaced)
	result = regexp.MustCompile(`YYYYMMDD`).ReplaceAllString(result, dateYYYYMMDD)
	result = regexp.MustCompile(`YYYY`).ReplaceAllString(result, dateYYYY)

	// Replace XXXX with random 4-digit numbers
	result = regexp.MustCompile(`XXXX`).ReplaceAllStringFunc(result, func(match string) string {
		return fmt.Sprintf("%04d", rand.Intn(10000))
	})

	// Replace XXX (3 X's) with random 3-digit numbers
	result = regexp.MustCompile(`XXX`).ReplaceAllStringFunc(result, func(match string) string {
		return fmt.Sprintf("%03d", rand.Intn(1000))
	})

	// Replace XX (2 X's) with random 2-digit numbers
	result = regexp.MustCompile(`XX`).ReplaceAllStringFunc(result, func(match string) string {
		return fmt.Sprintf("%02d", rand.Intn(100))
	})

	fmt.Fprintln(w, result)
}
