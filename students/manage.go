package students

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
	"student-mark-system/colors"
)

func (s *StudentManager) AddInteractive() {
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
			colors.Error("Invalid format. Use: Name/Grade")
			continue
		}

		name := strings.TrimSpace(parts[0])
		markStr := strings.TrimSpace(parts[1])

		mark, err := strconv.Atoi(markStr)
		if err != nil {
			colors.Error("Invalid mark. Please enter a number.")
			continue
		}
		
		validMark, err := ValidateMark(mark)
		if err != nil {
			colors.Error("Invalid range. Enter a number between 1-100.")
			continue
		}

		if _, exists := s.Students[name]; exists {
			colors.Warning("already exists. Skipping.\n")
			continue
		}

		s.Students[name] = Student{Name: name, Mark: validMark}
		colors.Success(fmt.Sprintf("%s was added successfully.", name))
		counter++
	}
}

func (s *StudentManager) SearchInteractive(){
	if len(s.Students) == 0{
		colors.Error("Students list is empty.")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("🔍 Search for a student: ")
	scanner.Scan()
	searchName := strings.TrimSpace(scanner.Text())
	
	student, exists := s.Students[searchName]
	if exists{
		colors.Success(fmt.Sprintf("Found! %s got %d.", searchName, student.Mark))
		}else{
			colors.Error(fmt.Sprintf("%s not found", searchName))
		}
	}

func (s *StudentManager) DeleteInteractive(){
	if len(s.Students) == 0{
		colors.Error("Students list is empty.")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("🗑️ Enter a student to delete: ")
	scanner.Scan()
	searchName := strings.TrimSpace(scanner.Text())

	if _, exists := s.Students[searchName]; exists {
		colors.Warning(fmt.Sprintf("Are you sure you want to delete %s? (y/n): ", searchName))
		scanner.Scan()
		confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if confirm == "y" {
			delete(s.Students, searchName)
			colors.Success(fmt.Sprintf("%s was deleted successfully.", searchName))
		} else {
			colors.Warning("Operation canceled.\n")
		}
	} else {
		colors.Error(fmt.Sprintf("%s not found.", searchName))
	}
}

func (s *StudentManager) EditMarkInteractive(){
	if len(s.Students) == 0{
		colors.Error("Students list is empty.")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("👨‍🎓 Enter the student's name to edit the mark: ")

	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())

	studentOldMark, exists := s.Students[name]
	if !exists {
		colors.Error(fmt.Sprintf(" %s not found.", name))
		return
	}

	fmt.Printf("%s✏️  Current mark for %s is: %d%s\n", colors.Blue, name, studentOldMark.Mark, colors.Reset)

	colors.Warning("Are you sure you want to update the mark? (y/n): ")
	scanner.Scan()
	confirm := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if confirm != "y" {
		colors.Warning("Operation canceled.\n")
		return
	}

	for {
		fmt.Print("🔖 Enter the new mark (1-100): ")
		scanner.Scan()
		markStr := strings.TrimSpace(scanner.Text())

		newMark, err := strconv.Atoi(markStr)
		if err != nil {
			colors.Error("Invalid input. Please enter a number.")
			continue
		}

		validMark, err := ValidateMark(newMark)
		if err != nil {
			colors.Error("Invalid range. Enter a number between 1-100.")
			continue
		}
		
		s.Students[name] = Student{Name: name, Mark: validMark}
		colors.Success(fmt.Sprintf(" %s's mark has been updated successfully to %d.", name, validMark))
		break
	}
}
