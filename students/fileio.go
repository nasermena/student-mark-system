package students

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"student-mark-system/colors"
)

func (s *StudentManager) ImportFromFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("❌ %sFailed to open file: %v%s\n", colors.Red, err, colors.Reset)
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
		if len(parts) != 2 || parts[0] == "" || !ValidateName(parts[0]) {
			fmt.Printf("⚠️ %sLine %d ignored (invalid format): %s%s\n", colors.Yellow, lineNumber, line, colors.Reset)
			skipped++
			lineNumber++
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])
		mark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("⚠️ %sLine %d ignored (invalid mark): %s%s\n", colors.Yellow, lineNumber, markStr, colors.Reset)
			skipped++
			lineNumber++
			continue
		}

		validMark, err := ValidateMark(mark)
		if err != nil {
			fmt.Printf("⚠️ %sLine %d ignored (mark out of range): %d%s\n", colors.Yellow, lineNumber, mark, colors.Reset)
			skipped++
			lineNumber++
			continue
		}

		if _, exists := s.Students[name]; exists {
			fmt.Printf("⚠️ %s%s already exists. Skipping.%s\n", colors.Yellow, name, colors.Reset)
			skipped++
		} else {
			s.Students[name] = Student{Name: name, Mark: validMark}
			added++
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ %sError reading file: %v%s\n", colors.Red, err, colors.Reset)
		return
	}

	fmt.Printf("✅ %sImport complete. %d added, %d skipped.%s\n", colors.Green, added, skipped, colors.Reset)
}

func (s *StudentManager) ExportToCSV(filename string) {
	fullPath := fmt.Sprintf("data/%s.csv", filename)
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Printf("%s❌ Could not create file: %v%s\n", colors.Red, err, colors.Reset)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Mark"})
	for _, student := range s.Students {
		record := []string{student.Name, fmt.Sprintf("%d", student.Mark)}
		if err := writer.Write(record); err != nil {
			fmt.Printf("%s❌ Failed to write record: %v%s\n", colors.Red, err, colors.Reset)
		}
	}

	fmt.Printf("%s✅ Exported %d students to CSV file: %s%s\n", colors.Green, len(s.Students), filename, colors.Reset)
}

func (s *StudentManager) ExportToJSON(filename string) {
	fullPath := fmt.Sprintf("data/%s.json", filename) // انتبه: الامتداد يجب أن يكون .json
	file, err := os.Create(fullPath)
	if err != nil {
		fmt.Printf("%s❌ Could not create file: %v%s\n", colors.Red, err, colors.Reset)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	exportData := make(map[string]int)
	for name, student := range s.Students {
		exportData[name] = student.Mark
	}

	if err := encoder.Encode(exportData); err != nil {
		fmt.Printf("%s❌ Failed to encode JSON: %v%s\n", colors.Red, err, colors.Reset)
		return
	}

	fmt.Printf("%s✅ Exported %d students to JSON file: %s%s\n", colors.Green, len(s.Students), filename, colors.Reset)
}
