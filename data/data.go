package data

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"student-mark-system/students"
	"encoding/csv"
	"encoding/json"
	
)

// ANSI Colors
const (
	Green = "\033[32m"
	Red   = "\033[31m"
	Reset = "\033[0m"
	Yellow = "\033[33m"

)

func ImportStudentsFromFile(filename string, studentMarks map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("❌ %sFailed to open file: %v%s\n", Red, err, Reset)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	added := 0
	skipped := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			lineNumber++
			continue
		}

		parts := strings.Split(line, ",")

		if len(parts) != 2 || parts[0] == "" || !students.ValidateName(parts[0]) {
			fmt.Printf("⚠️ %sLine %d ignored (invalid format): %s%s\n", Yellow, lineNumber, line, Reset)
			skipped++
			lineNumber++
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])
		mark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("⚠️ %sLine %d ignored (invalid mark): %s%s\n", Yellow, lineNumber, markStr, Reset)
			skipped++
			lineNumber++
			continue
		}

		validMark, err := students.ValidateMark(mark)
		if err != nil {
			fmt.Printf("⚠️ %sLine %d ignored (mark out of range): %d%s\n", Yellow, lineNumber, mark, Reset)
			skipped++
			lineNumber++
			continue
		}

		if _, exists := studentMarks[name]; exists {
			fmt.Printf("⚠️ %s%s already exists. Skipping.%s\n", Yellow, name, Reset)
			skipped++
		} else {
			studentMarks[name] = validMark
			added++
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ %sError reading file: %v%s\n", err, Red, Reset)
		return
	}

	fmt.Printf("✅ %sImport complete. %d added, %d skipped.%s\n", Green, added, skipped, Reset)
}

// ExportToCSV writes the student marks to a CSV file
func ExportToCSV(filename string, studentMarks map[string]int) {


	fullPath := fmt.Sprintf("data/%s.csv", filename)
	file, err := os.Create(fullPath)

	if err != nil {
		fmt.Printf("%s❌ Could not create file: %v%s\n", Red, err, Reset)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Name", "Mark"})

	for name, mark := range studentMarks {
		record := []string{name, fmt.Sprintf("%d", mark)}
		if err := writer.Write(record); err != nil {
			fmt.Printf("%s❌ Failed to write record: %v%s\n", Red, err, Reset)
		}
	}

	fmt.Printf("%s✅ Exported %d students to CSV file: %s%s\n", Green, len(studentMarks), filename, Reset)
}

// ExportToJSON writes the student marks to a JSON file
func ExportToJSON(filename string, studentMarks map[string]int) {
	fullPath := fmt.Sprintf("data/%s.csv", filename)
	file, err := os.Create(fullPath)

	if err != nil {
		fmt.Printf("%s❌ Could not create file: %v%s\n", Red, err, Reset)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(studentMarks); err != nil {
		fmt.Printf("%s❌ Failed to encode JSON: %v%s\n", Red, err, Reset)
		return
	}

	fmt.Printf("%s✅ Exported %d students to JSON file: %s%s\n", Green, len(studentMarks), filename, Reset)
}
