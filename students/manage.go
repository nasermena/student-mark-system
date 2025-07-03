package students

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
	"student-mark-system/colors"
)

func AddStudent(studentMarks map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)
	
	fmt.Printf("%s📥 Enter students as: Name/Grade (Type '0' to exit)%s\n", colors.Cyan, colors.Reset)
	counter := 1
	for {
		fmt.Printf("- Student #%d: ",counter)
		scanner.Scan()
		line := strings.TrimSpace(scanner.Text())

		if line == "0" {
			break
		}

		parts := strings.Split(line, "/")
		if len(parts) != 2 || !ValidateName(parts[0]){
			fmt.Printf("%s❌ Invalid format. Use: Name/Grade%s\n",colors.Red, colors.Reset)
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])

		mark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("%s❌ Invalid mark. Please enter a number.%s\n", colors.Red, colors.Reset)
			continue
		}
		
		validMark, err := ValidateMark(mark)
		if err != nil {
			fmt.Printf("%s🔢 Invalid range. Enter a number between 1-100.%s\n", colors.Red, colors.Reset)
			continue
		}

		if _, exists := studentMarks[name]; exists {
			fmt.Printf("%s⚠️ %s already exists. Skipping.%s\n", colors.Yellow, name, colors.Reset)
			continue
		}

		studentMarks[name] = validMark
		fmt.Printf("%s✅ %s was added successfully.%s\n", colors.Green, name, colors.Reset)
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
		fmt.Printf("%s✅ Found! %s got %d.%s\n", colors.Green, searchName, mark, colors.Reset)
		}else{
			fmt.Printf("%s❌ %s not found.%s\n", colors.Red, searchName, colors.Reset)
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
			fmt.Printf("%s✅ %s was deleted successfully.%s\n", colors.Green, searchName, colors.Reset)
		} else {
			fmt.Printf("%sℹ️ Operation canceled.%s\n", colors.Yellow, colors.Reset)
		}
	} else {
		fmt.Printf("%s❌ %s not found.%s\n", colors.Red, searchName, colors.Reset)
	}
}

func EditStudentMark(studentMarks map[string]int) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("👨‍🎓 Enter the student's name to edit the mark: ")

	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	oldMark, exists := studentMarks[name]
	if !exists {
		fmt.Printf("%s❌ %s not found.%s\n", colors.Red, name, colors.Reset)
		return
	}

	fmt.Printf("%s✏️  Current mark for %s is: %d%s\n", colors.Blue, name, oldMark, colors.Reset)

	fmt.Print("⚠️  Are you sure you want to update the mark? (y/n): ")
	scanner.Scan()
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if confirm != "y" {
		fmt.Printf("%sℹ️  Operation canceled.%s\n", colors.Yellow, colors.Reset)
		return
	}

	for {
		fmt.Print("🔖 Enter the new mark (1-100): ")
		scanner.Scan()
		markStr := strings.TrimSpace(scanner.Text())

		newMark, err := strconv.Atoi(markStr)
		if err != nil {
			fmt.Printf("%s❌ Invalid input. Please enter a number.%s\n", colors.Red, colors.Reset)
			continue
		}

		validMark, err := ValidateMark(newMark)
		if err != nil {
			fmt.Printf("%s🔢 Invalid range. Enter a number between 1-100.%s\n", colors.Red, colors.Reset)
			continue
		}

		studentMarks[name] = validMark
		fmt.Printf("%s✅ %s's mark has been updated successfully to %d.%s\n", colors.Green, name, validMark, colors.Reset)
		break
	}
}
