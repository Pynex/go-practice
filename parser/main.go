package main

import (
	"bufio"
	"fmt"
	"os"
	"parser/random"
	"strconv"
	"strings"
)

func main() {
	err := random.GenerateRandomDates()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating random number: %v\n", err)
		os.Exit(1)
	}

	inputFile, err := os.ReadFile("dates.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading dates from file: %v\n", err)
		os.Exit(1)
	}

	var days []int
	var months []int
	var years []int

	scanner := bufio.NewScanner(strings.NewReader(string(inputFile)))

	for scanner.Scan() {
		date := strings.TrimSpace(scanner.Text())
		parts := strings.Split(date, "/")
		if len(parts) != 3 {
			fmt.Fprintf(os.Stderr, "Invalid date format: %s\n", date)
			continue
		}

		day, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing day: %v\n", err)
			continue
		}

		month, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing month: %v\n", err)
			continue
		}

		year, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing year: %v\n", err)
			continue
		}

		days = append(days, day)
		months = append(months, month)
		years = append(years, year)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading dates from file: %v\n", err)
		os.Exit(1)
	}

	daysFile, err := os.Create("days.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating days file: %v\n", err)
		os.Exit(1)
	}
	defer daysFile.Close()

	for _, day := range days {
		daysFile.WriteString(strconv.Itoa(day) + "\n")
	}

	monthsFile, err := os.Create("months.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating months file: %v\n", err)
		os.Exit(1)
	}
	defer monthsFile.Close()

	for _, month := range months {
		monthsFile.WriteString(strconv.Itoa(month) + "\n")
	}

	yearsFile, err := os.Create("years.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating years file: %v\n", err)
		os.Exit(1)
	}
	defer yearsFile.Close()

	for i := len(years) - 1; i >= 0; i-- {
		yearsFile.WriteString(strconv.Itoa(years[i]) + "\n")
	}

	fmt.Println("Days: ", days)
	fmt.Println("Months: ", months)
	fmt.Println("Years: ", years)

	fmt.Println("Days, months and years have been saved to days.txt, months.txt and years.txt")
}
