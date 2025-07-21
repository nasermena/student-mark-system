package main

import (
	"fmt"
	"student-mark-system/colors"
	"student-mark-system/runapp"
)
func main() {
    fmt.Printf("%sWelcome to Student System!%s\n", colors.BgMagenta, colors.Reset)
    runapp.MainMenu()
}
