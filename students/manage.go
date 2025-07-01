package students

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

// ANSI Colors
const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Reset  = "\033[0m"
)

func AddStudent(studentMarks map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("%s📥 Enter students as: Name/Grade (Type '0' to exit)%s\n", Cyan, Reset)
	counter := 1
	for {
		fmt.Printf("- Student #%d: ",counter)
		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())

		if line == "0" {
			break
		}

		parts := strings.Split(line, "/")
		if len(parts) != 2 {
			fmt.Printf("%s❌ Invalid format. Use: Name/Grade%s\n",Red, Reset)
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])

		mark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("%s❌ Invalid mark. Please enter a number.%s\n", Red, Reset)
			continue
		}
		
		validMark, err := ValidateMark(mark)
		if err != nil {
			fmt.Printf("%s🔢 Invalid range. Enter a number between 1-100.%s\n", Red, Reset)
			continue
		}

		if _, exists := studentMarks[name]; exists {
			fmt.Printf("%s⚠️ %s already exists. Skipping.%s\n", Yellow, name, Reset)
			continue
		}

		studentMarks[name] = validMark
		fmt.Printf("%s✅ %s was added successfully.%s\n", Green, name, Reset)
		counter++
	}
}

func SearchStudent(studentMarks map[string]int){
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("🔍 Search for a student: ")
	scanner.Scan()
	searchName := strings.TrimSpace(scanner.Text())
	
	mark, exists := studentMarks[searchName]
	if exists{
		fmt.Printf("%s✅ Found! %s got %d.%s\n", Green, searchName, mark, Reset)
		}else{
			fmt.Printf("%s❌ %s not found.%s\n", Red, searchName, Reset)
		}
	}


func DeleteStudent(studentMarks map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("🗑️ Enter a student to delete: ")
	scanner.Scan()
	searchName := strings.TrimSpace(scanner.Text())

	if _, exists := studentMarks[searchName]; exists {
		fmt.Printf("⚠️ Are you sure you want to delete %s? (y/n): ", searchName)
		scanner.Scan()
		confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if confirm == "y" {
			delete(studentMarks, searchName)
			fmt.Printf("%s✅ %s was deleted successfully.%s\n", Green, searchName, Reset)
		} else {
			fmt.Printf("%sℹ️ Operation canceled.%s\n", Yellow, Reset)
		}
	} else {
		fmt.Printf("%s❌ %s not found.%s\n", Red, searchName, Reset)
	}
}
			


func EditStudentMark(studentMarks map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("👨‍🎓 Enter the student's name to edit the mark: ")

	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	oldMark, exists := studentMarks[name]
	if !exists {
		fmt.Printf("%s❌ %s not found.%s\n", Red, name, Reset)
		return
	}

	fmt.Printf("%s✏️  Current mark for %s is: %d%s\n", Blue, name, oldMark, Reset)

	fmt.Print("⚠️  Are you sure you want to update the mark? (y/n): ")
	scanner.Scan()
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if confirm != "y" {
		fmt.Printf("%sℹ️  Operation canceled.%s\n", Yellow, Reset)
		return
	}

	for {
		fmt.Print("🔖 Enter the new mark (1-100): ")
		scanner.Scan()
		markStr := strings.TrimSpace(scanner.Text())

		newMark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("%s❌ Invalid input. Please enter a number.%s\n", Red, Reset)
			continue
		}

		validMark, err := ValidateMark(newMark)
		if err != nil {
			fmt.Printf("%s🔢 Invalid range. Enter a number between 1-100.%s\n", Red, Reset)
			continue
		}

		studentMarks[name] = validMark
		fmt.Printf("%s✅ %s's mark has been updated successfully to %d.%s\n", Green, name, validMark, Reset)
		break
	}
}


func ImportStudentsFromFile(filename string, studentMarks map[string]int) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("❌ Failed to open file: %v\n", err)
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
		if len(parts) != 2 {
			fmt.Printf("⚠️ Line %d ignored (invalid format): %s\n", lineNumber, line)
			skipped++
			lineNumber++
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])
		mark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("⚠️ Line %d ignored (invalid mark): %s\n", lineNumber, markStr)
			skipped++
			lineNumber++
			continue
		}

		validMark, err := ValidateMark(mark)
		if err != nil {
			fmt.Printf("⚠️ Line %d ignored (mark out of range): %d\n", lineNumber, mark)
			skipped++
			lineNumber++
			continue
		}

		if _, exists := studentMarks[name]; exists {
			fmt.Printf("⚠️ %s already exists. Skipping.\n", name)
			skipped++
		} else {
			studentMarks[name] = validMark
			added++
		}

		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("❌ Error reading file: %v\n", err)
		return
	}

	fmt.Printf("✅ Import complete. %d added, %d skipped.\n", added, skipped)
}