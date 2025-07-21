package runapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"student-mark-system/students"
    "student-mark-system/colors"
)

func MainMenu(){
    
    store := students.NewStudentManager()

    for {
        scanner := bufio.NewScanner(os.Stdin)        
        fmt.Print(strings.Repeat("-", 40) + "\n")
        fmt.Print("Main Menu (Choose a number):\n1- Students list\n2- Import students from file\n3- Add student\n4- Search student\n5- Delete student\n6- Edit student mark\n7- Summary report\n8- Grade overviewn\n9- Save and export students to CSV\n10- Save and export students to JSON\n11- Exit\nEnter choice: ")
        scanner.Scan()
        option := scanner.Text() 
        switch option {
        case "1":			
           store.ShowAll()
        case "2":
            scanner := bufio.NewScanner(os.Stdin)
            fmt.Print("Enter File Name: ")
            scanner.Scan()
            filename := strings.TrimSpace(scanner.Text())
            store.ImportFromFile(filename)
        case "3":
            store.AddInteractive()
        case "4":
            store.SearchInteractive()
        case "5":
            store.DeleteInteractive()
        case "6":
            store.EditMarkInteractive()
        case "7":
            store.PrintSummary()
        case "8":
            store.GradeOverview()
        case "9":
            if len(store.Students) == 0 {
                fmt.Printf("❌%sNo data to export%s\n", colors.Red, colors.Reset)
            }else {
                fmt.Print("Enter filename to export as CSV: ")
                scanner.Scan()
                filename := strings.TrimSpace(scanner.Text())
                store.ExportToCSV(filename)
                return
            }
        case "10":
            if len(store.Students) == 0 {
                fmt.Printf("❌%s No data to export.%s\n", colors.Red, colors.Reset)
            }else {
                fmt.Print("Enter filename to export as JSON: ")
                scanner.Scan()
                filename := strings.TrimSpace(scanner.Text())
                store.ExportToJSON(filename)
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