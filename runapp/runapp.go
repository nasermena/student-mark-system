package runapp

import (
	"student-mark-system/students"
)

func MainMenu() {
	store := students.NewStudentManager()
	menu := NewMenu(store)
	menu.Show()
}
