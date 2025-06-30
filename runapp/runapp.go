package runapp

import (
    "student-mark-system/students"
    "fmt"
    "strings"
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
        fmt.Print(strings.Repeat("-", 40) + "\n")
        fmt.Print("Main Menu (Choose a number):\n1- Show students list\n2- Add student\n3- Search student\n4- Delete student\n5- Edit student mark\n6- Show summary report\n7- Show grade distribution\n8- Exit\nEnter choice: ")
        var option string
        fmt.Scan(&option)
        switch option {
        case "1":			
           students.ShowStudents(studentMarks)
        case "2":
            students.AddStudent(studentMarks)
        case "3":
           students.SearchStudent(studentMarks)
        case "4":
            students.DeleteStudent(studentMarks)
        case "5":
            students.EditStudentMark(studentMarks)
        case "6":
            students.PrintSummary(studentMarks)
        case "7":
            students.GradeDistribution(studentMarks)
        case "8":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Printf("%sInvalid option. Try again.%s\n", Red, Reset)
        }
    }
}