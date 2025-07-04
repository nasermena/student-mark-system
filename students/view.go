package students
import (
	"fmt"
	"sort"
	"strings"
	"student-mark-system/colors"
	"math"
)

func ShowStudents(studentMarks map[string]int){
	if len(studentMarks) == 0{
		fmt.Printf("%s❌ Students list is empty.%s\n", colors.Red, colors.Reset)

		}else{
			keys := []string{}
			for k := range studentMarks {
				keys = append(keys, k)
			}
			
			sort.Strings(keys)
			fmt.Println(strings.Repeat("=", 40))
			fmt.Printf("%s🎓 Students list:%s\n", colors.Cyan, colors.Reset)
			for _, name:= range keys{
				fmt.Printf("👨‍🎓 %-12s : %d\n", name, studentMarks[name])
			}
		}
		fmt.Println(strings.Repeat("=", 40))
	}

func PrintSummary(studentMarks map[string]int) {
	if len(studentMarks) == 0 {
		fmt.Printf("❌ %sNo students to summarize.%s", colors.Red, colors.Reset)
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
	fmt.Printf("%s📊 Students Summary Report%s\n", colors.Cyan, colors.Reset)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("👥 Total students: %d\n", len(studentMarks))
	fmt.Printf("🎯 Highest mark: %d by %v\n", maxMark, topStudents)
	fmt.Printf("❗ Lowest mark: %d by %v\n", minMark, lowStudents)
	fmt.Printf("📈 Average mark: %.2f\n", avg)
	fmt.Printf("✅ Passed: %d student(s)\n", passCount)
	fmt.Printf("❌ Failed: %d student(s)\n", failCount)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("%s📊 Advanced Statistics:%s\n", colors.Cyan, colors.Reset)
	ComputeAdvancedStats(studentMarks)
	fmt.Println(strings.Repeat("=", 40))

}


func GradeOverview(studentMarks map[string]int) {
	if len(studentMarks) == 0 {
		fmt.Printf("%s❌ No student marks available.%s\n", colors.Red, colors.Reset)
		return
	}
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf("📄%s Grade Overview:%s\n", colors.Cyan, colors.Reset)
	fmt.Println(strings.Repeat("=", 40))

	distribution := map[string]int{
		"A+ (90-100) 🎖️": 0,
		"A (80-89) 🥇":    0,
		"B (70-79) 🥈":    0,
		"C (60-69) 🥉":    0,
		"D (50-59) ✅":    0,
		"F (0-49) ❌":     0,
	}

	fmt.Printf("📄%s Grade Classification per Student:%s\n", colors.Cyan, colors.Reset)
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
	fmt.Printf("%s📊 Grade Distribution Summary:%s\n", colors.Cyan, colors.Reset)
	fmt.Println(strings.Repeat("-", 40))

	orderedGrades := []string{
		"A+ (90-100) 🎖️",
		"A (80-89) 🥇",
		"B (70-79) 🥈",
		"C (60-69) 🥉",
		"D (50-59) ✅",
		"F (0-49) ❌",
	}

	for _, grade := range orderedGrades {
		count := distribution[grade]
		fmt.Printf("%-20s : %d student(s)\n", grade, count)
	}
	fmt.Println(strings.Repeat("=", 40))
}

// ComputeAdvancedStats calculates and displays advanced statistical metrics
func ComputeAdvancedStats(studentMarks map[string]int) {
	if len(studentMarks) == 0 {
		fmt.Printf("%s❌ No data available for statistics.%s\n", colors.Red, colors.Reset)
		return
	}

	var marks []int
	for _, mark := range studentMarks {
		marks = append(marks, mark)
	}

	sort.Ints(marks)

	// Median
	var median float64
	n := len(marks)
	if n%2 == 0 {
		median = float64(marks[n/2-1]+marks[n/2]) / 2.0
	} else {
		median = float64(marks[n/2])
	}

	// Mean
	var sum int
	for _, mark := range marks {
		sum += mark
	}
	mean := float64(sum) / float64(n)

	// Standard Deviation
	var variance float64
	for _, mark := range marks {
		diff := float64(mark) - mean
		variance += diff * diff
	}
	stdDev := math.Sqrt(variance / float64(n))

	// Range
	valueRange := marks[n-1] - marks[0]

	// Unique sorted marks
	uniqueMap := make(map[int]bool)
	for _, m := range marks {
		uniqueMap[m] = true
	}
	var uniqueMarks []int
	for m := range uniqueMap {
		uniqueMarks = append(uniqueMarks, m)
	}
	sort.Ints(uniqueMarks)

	fmt.Printf("📊 Median: %.2f\n", median)
	fmt.Printf("📏 Standard Deviation: %.2f\n", stdDev)
	fmt.Printf("📉 Range: %d\n", valueRange)
	fmt.Printf("🔢 Unique Marks: %v\n", uniqueMarks)
}
