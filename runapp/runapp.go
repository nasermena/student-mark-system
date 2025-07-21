package runapp

import (
	"fmt"
	"student-mark-system/students"
)

func MainMenu() {
	fmt.Println("Welcome to Student System!")
	store := students.NewStudentManager()
	menu := NewMenu(store)
	menu.Show()
}
