package students
import (
	"fmt"
	"sort"
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