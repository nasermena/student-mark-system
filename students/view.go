package students
import (
	"fmt"
	"sort"
	"strings"
)

func ShowStudents(studentMarks map[string]int){
	if len(studentMarks) == 0{
		fmt.Printf("%s❌ Students list is empty.%s\n", Red, Reset)

		}else{
			keys := []string{}
			for k := range studentMarks {
				keys = append(keys, k)
			}
			
			sort.Strings(keys)
			fmt.Println(strings.Repeat("=", 40))
			fmt.Printf("%s🎓 Students list:%s\n", Cyan, Reset)
			for _, name:= range keys{
				fmt.Printf("👨‍🎓 %-12s : %d\n", name, studentMarks[name])
			}
		}
		fmt.Println(strings.Repeat("=", 40))
	}

	func PrintSummary(studentMarks map[string]int) {
	if len(studentMarks) == 0 {
		fmt.Println("❌ No students to summarize.")
		return
	}

	var (
		total       int
		maxMark     = -1
		minMark     = 101
		topStudents []string
		lowStudents []string
		passCount   int
	)

	for name, mark := range studentMarks {
		total += mark

		if mark > maxMark {
			maxMark = mark
			topStudents = []string{name}
		} else if mark == maxMark {
			topStudents = append(topStudents, name)
		}

		if mark < minMark {
			minMark = mark
			lowStudents = []string{name}
		} else if mark == minMark {
			lowStudents = append(lowStudents, name)
		}

		if mark >= 50 {
			passCount++
		}
	}

	avg := float64(total) / float64(len(studentMarks))
	failCount := len(studentMarks) - passCount

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%s📊 Students Summary Report%s\n", Cyan, Reset)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("👥 Total students: %d\n", len(studentMarks))
	fmt.Printf("🎯 Highest mark: %d by %v\n", maxMark, topStudents)
	fmt.Printf("❗ Lowest mark: %d by %v\n", minMark, lowStudents)
	fmt.Printf("📈 Average mark: %.2f\n", avg)
	fmt.Printf("✅ Passed: %d student(s)\n", passCount)
	fmt.Printf("❌ Failed: %d student(s)\n", failCount)
	fmt.Println(strings.Repeat("=", 40))
}


func GradeOverview(studentMarks map[string]int) {
	if len(studentMarks) == 0 {
		fmt.Println("❌ No student marks available.")
		return
	}
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("📄%s Grade Overview:%s\n", Cyan, Reset)
	fmt.Println(strings.Repeat("=", 40))

	distribution := map[string]int{
		"A+ (90-100) 🎖️": 0,
		"A (80-89) 🥇":    0,
		"B (70-79) 🥈":    0,
		"C (60-69) 🥉":    0,
		"D (50-59) ✅":    0,
		"F (0-49) ❌":     0,
	}

	fmt.Printf("📄%s Grade Classification per Student:%s\n", Cyan, Reset)
	fmt.Println(strings.Repeat("-", 40))

	for name, mark := range studentMarks {
		var grade string
		switch {
		case mark >= 90:
			grade = "A+ (90-100) 🎖️"
		case mark >= 80:
			grade = "A (80-89) 🥇"
		case mark >= 70:
			grade = "B (70-79) 🥈"
		case mark >= 60:
			grade = "C (60-69) 🥉"
		case mark >= 50:
			grade = "D (50-59) ✅"
		default:
			grade = "F (0-49) ❌"
		}

		distribution[grade]++
		fmt.Printf("👨‍🎓 %-15s → %3d (%s)\n", name, mark, grade)
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%s📊 Grade Distribution Summary:%s\n", Cyan, Reset)
	fmt.Println(strings.Repeat("-", 40))

	for category, count := range distribution {
		fmt.Printf("%-20s : %d student(s)\n", category, count)
	}
	fmt.Println(strings.Repeat("=", 40))
}
