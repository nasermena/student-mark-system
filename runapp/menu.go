package runapp

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"student-mark-system/students"
	"student-mark-system/colors"
)

type Menu struct {
	Store *students.StudentManager
}

func NewMenu(store *students.StudentManager) *Menu {
	return &Menu{Store: store}
}

func (m *Menu) Show() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(strings.Repeat("-", 40) + "\n")
		fmt.Print("Main Menu (Choose a number):\n")
		fmt.Println("1- Students list")
		fmt.Println("2- Import students from file")
		fmt.Println("3- Add student")
		fmt.Println("4- Search student")
		fmt.Println("5- Delete student")
		fmt.Println("6- Edit student mark")
		fmt.Println("7- Summary report")
		fmt.Println("8- Grade overview")
		fmt.Println("9- Save and export students to CSV")
		fmt.Println("10- Save and export students to JSON")
		fmt.Println("11- Exit")
		fmt.Print("Enter choice: ")

		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			m.Store.ShowAll()
		case "2":
			fmt.Print("Enter File Name: ")
			scanner.Scan()
			filename := strings.TrimSpace(scanner.Text())
			m.Store.ImportFromFile(filename)
		case "3":
			m.Store.AddInteractive()
		case "4":
			m.Store.SearchInteractive()
		case "5":
			m.Store.DeleteInteractive()
		case "6":
			m.Store.EditMarkInteractive()
		case "7":
			m.Store.PrintSummary()
		case "8":
			m.Store.GradeOverview()
		case "9":
			if len(m.Store.Students) == 0 {
				fmt.Printf("❌%sNo data to export%s\n", colors.Red, colors.Reset)
			} else {
				fmt.Print("Enter filename to export as CSV: ")
				scanner.Scan()
				filename := strings.TrimSpace(scanner.Text())
				m.Store.ExportToCSV(filename)
			}
		case "10":
			if len(m.Store.Students) == 0 {
				fmt.Printf("❌%s No data to export.%s\n", colors.Red, colors.Reset)
			} else {
				fmt.Print("Enter filename to export as JSON: ")
				scanner.Scan()
				filename := strings.TrimSpace(scanner.Text())
				m.Store.ExportToJSON(filename)
			}
		case "11":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Printf("%sInvalid option. Try again.%s\n", colors.Red, colors.Reset)
		}
	}
}
