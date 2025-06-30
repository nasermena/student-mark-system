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
			fmt.Printf("%s🎓 Students list:%s\n", Cyan, Reset)
			for _, name:= range keys{
				fmt.Printf("👨‍🎓 %s = %d\n", name, studentMarks[name])
			}
		}
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


func GradeDistribution(studentMarks map[string]int) {
	if len(studentMarks) == 0 {
		fmt.Println("❌ No student marks available.")
		return
	}

	distribution := map[string]int{
		"A (90-100)": 0,
		"B (80-89)":  0,
		"C (70-79)":  0,
		"D (60-69)":  0,
		"F (<60)":    0,
	}

	for _, mark := range studentMarks {
		switch {
		case mark >= 90:
			distribution["A (90-100)"]++
		case mark >= 80:
			distribution["B (80-89)"]++
		case mark >= 70:
			distribution["C (70-79)"]++
		case mark >= 60:
			distribution["D (60-69)"]++
		default:
			distribution["F (<60)"]++
		}
	}

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("%s📊 Grade Distribution:%s\n", Cyan, Reset)
	for grade, count := range distribution {
		fmt.Printf("%-12s: %d student(s)\n", grade, count)
	}
	fmt.Println(strings.Repeat("=", 40))
}
