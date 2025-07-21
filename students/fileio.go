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
		colors.Error(fmt.Sprintf("Failed to open file: %v", err))
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
			colors.Warning(fmt.Sprintf("Line %d ignored (invalid format): %s\n", lineNumber, line))
			skipped++
			lineNumber++
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])
		mark, err := strconv.Atoi(markStr)
		if err != nil {
			colors.Warning(fmt.Sprintf("Line %d ignored (invalid mark): %s\n", lineNumber, markStr))
			skipped++
			lineNumber++
			continue
		}

		validMark, err := ValidateMark(mark)
		if err != nil {
			colors.Warning(fmt.Sprintf("Line %d ignored (mark out of range): %d\n", lineNumber, mark))
			skipped++
			lineNumber++
			continue
		}

		if _, exists := s.Students[name]; exists {
			colors.Warning(fmt.Sprintf("%s already exists. Skipping.\n", name))
			skipped++
		} else {
			s.Students[name] = Student{Name: name, Mark: validMark}
			added++
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		colors.Error(fmt.Sprintf("Error reading file: %v", err))
		return
	}
	colors.Success(fmt.Sprintf("Import complete. %d added, %d skipped.", added, skipped))
}

func (s *StudentManager) ExportToCSV(filename string) {
	fullPath := fmt.Sprintf("data/%s.csv", filename)
	file, err := os.Create(fullPath)
	if err != nil {
		colors.Error(fmt.Sprintf("Could not create file: %v", err))
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Mark"})
	for _, student := range s.Students {
		record := []string{student.Name, fmt.Sprintf("%d", student.Mark)}
		if err := writer.Write(record); err != nil {
			colors.Error(fmt.Sprintf("Failed to write record: %v", err))
		}
	}

	colors.Success(fmt.Sprintf("Exported %d students to CSV file: %s", len(s.Students), filename))
}

func (s *StudentManager) ExportToJSON(filename string) {
	fullPath := fmt.Sprintf("data/%s.json", filename) // انتبه: الامتداد يجب أن يكون .json
	file, err := os.Create(fullPath)
	if err != nil {
		colors.Error(fmt.Sprintf("Could not create file: %v", err))
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
		colors.Error(fmt.Sprintf("Failed to encode JSON: %v", err))
		return
	}

	colors.Success(fmt.Sprintf("Exported %d students to JSON file: %s", len(s.Students), filename))
}
