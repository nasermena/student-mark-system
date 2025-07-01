package runapp

import (
    "student-mark-system/students"
    "fmt"
    "strings"
    "bufio"
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

func MainMenu(){

    studentMarks := make(map[string]int)

    for {
        scanner := bufio.NewScanner(os.Stdin)        
        fmt.Print(strings.Repeat("-", 40) + "\n")
        fmt.Print("Main Menu (Choose a number):\n1- Students list\n2- Import students from file\n3- Add student\n4- Search student\n5- Delete student\n6- Edit student mark\n7- Summary report\n8- Grade overviewn\n9- Exit\nEnter choice: ")
        scanner.Scan()
        option := scanner.Text() 
        switch option {
        case "1":			
           students.ShowStudents(studentMarks)
        case "2":
            scanner := bufio.NewScanner(os.Stdin)
            fmt.Print("Enter File Name: ")
            scanner.Scan()
            filename := strings.TrimSpace(scanner.Text())
	        students.ImportStudentsFromFile(filename, studentMarks)
        case "3":
            students.AddStudent(studentMarks)
        case "4":
           students.SearchStudent(studentMarks)
        case "5":
            students.DeleteStudent(studentMarks)
        case "6":
            students.EditStudentMark(studentMarks)
        case "7":
            students.PrintSummary(studentMarks)
        case "8":
            students.GradeOverview(studentMarks)
        case "9":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Printf("%sInvalid option. Try again.%s\n", Red, Reset)
        }
    }
}