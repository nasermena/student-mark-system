package colors

import "fmt"

func Success(msg string) {
	fmt.Println(Green + "✅ " + msg + Reset)
}

func Error(msg string) {
	fmt.Println(Red + "❌ " + msg + Reset)
}

func Warning(msg string) {
	fmt.Print(Yellow + "⚠️ " + msg + Reset)
}

func Info(msg string) {
	fmt.Println(Cyan + "ℹ️ " + msg + Reset)
}

func Highlight(msg string) {
	fmt.Println(BgMagenta + msg + Reset)
}
