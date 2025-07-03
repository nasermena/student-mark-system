package runapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"student-mark-system/data"
	"student-mark-system/students"
    "student-mark-system/colors"
)

func MainMenu(){

    studentMarks := make(map[string]int)

    for {
        scanner := bufio.NewScanner(os.Stdin)        
        fmt.Print(strings.Repeat("-", 40) + "\n")
        fmt.Print("Main Menu (Choose a number):\n1- Students list\n2- Import students from file\n3- Add student\n4- Search student\n5- Delete student\n6- Edit student mark\n7- Summary report\n8- Grade overviewn\n9- Save and export students to CSV\n10- Save and export students to JSON\n11- Exit\nEnter choice: ")
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
	        data.ImportStudentsFromFile(filename, studentMarks)
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
            if len(studentMarks) == 0 {
                fmt.Printf("❌%sNo data to export%s\n", colors.Red, colors.Reset)
            }else {
                fmt.Print("Enter filename to export as CSV: ")
                scanner.Scan()
                filename := strings.TrimSpace(scanner.Text())
                data.ExportToCSV(filename, studentMarks)
                return
            }
        case "10":
            if len(studentMarks) == 0 {
                fmt.Printf("❌%s No data to export.%s\n", colors.Red, colors.Reset)
            }else {
                fmt.Print("Enter filename to export as JSON: ")
                scanner.Scan()
                filename := strings.TrimSpace(scanner.Text())
                data.ExportToJSON(filename, studentMarks)
                return
            }
        case "11":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Printf("%sInvalid option. Try again.%s\n", colors.Red, colors.Reset)
        }
    }
}