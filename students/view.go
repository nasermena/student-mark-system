package students
import (
	"fmt"
	"sort"
	"strings"
)

func ShowStudents(studentMarks map[string]int){
	if len(studentMarks) == 0{
		fmt.Println("❌ Students list is empty.")

		}else{
			keys := []string{}
			for k := range studentMarks {
				keys = append(keys, k)
			}
			
			sort.Strings(keys)
			fmt.Println("🎓 The students list:")
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
	fmt.Printf("📊 Students Summary Report\n")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("👥 Total students: %d\n", len(studentMarks))
	fmt.Printf("🎯 Highest mark: %d by %v\n", maxMark, topStudents)
	fmt.Printf("❗ Lowest mark: %d by %v\n", minMark, lowStudents)
	fmt.Printf("📈 Average mark: %.2f\n", avg)
	fmt.Printf("✅ Passed: %d student(s)\n", passCount)
	fmt.Printf("❌ Failed: %d student(s)\n", failCount)
	fmt.Println(strings.Repeat("=", 40))
}
