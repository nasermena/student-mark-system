package students

import (
	"fmt"
)



func AddStudent(studentMarks map[string]int){
	var studentName string
	var studentIndex  int
	
	for{
		fmt.Printf("👨‍🎓 Please enter student name (%d) (enter 0 to exit): ", studentIndex +1)
		fmt.Scan(&studentName)
		
		if _, exists := studentMarks[studentName]; exists {
			fmt.Printf("⚠️ %s already exists. Please enter a new name.\n", studentName)
			continue
		}

		
		if studentName == "0"{
			break
		}
		for{
			var studentMark int

			fmt.Printf("🔖 Enter the mark for (%s) (1-100) = ", studentName)
			fmt.Scan(&studentMark)
			
			validMark, err := validateMark(studentMark)

			if err != nil {
				fmt.Println("🔢 Please enter valid input (number 1-100)")
				}else{
					studentMarks[studentName] = validMark
					fmt.Printf("✅ Done, %s was added\n", studentName)
					studentIndex++
					break
				}
			}
		}
	}

func SearchStudent(studentMarks map[string]int){

	var searchName string
	fmt.Print("🔍 Search for a student: ")
	fmt.Scan(&searchName)
	
	mark, exists := studentMarks[searchName]
	if exists{
		fmt.Printf("✅ Found!, %s got %d.\n", searchName, mark)
		}else{
			fmt.Printf("❌ %s not found.\n", searchName)
		}
	}


func DeleteStudent(studentMarks map[string]int){
	var searchName string
	fmt.Print("🗑️ Enter a student to delete: ")
	fmt.Scan(&searchName)

	_, exists := studentMarks[searchName]
	if exists{
		delete(studentMarks, searchName)
		fmt.Printf("✅ Done, %s was deleted.\n", searchName)
		}else{
			fmt.Printf("❌%s not found.\n", searchName)
		}
	}
			
func EditStudentMark(studentMarks map[string]int){
	var searchName string
	var newMark int
	
	fmt.Print("👨‍🎓 Enter the student's name to edit his mark: ")
	fmt.Scan(&searchName)
	oldMark, exists := studentMarks[searchName]

	if exists{
		for {
			fmt.Printf("🔖 The old mark = %d, renter the new mark: ", oldMark)
			fmt.Scan(&newMark)

			validMark, err := validateMark(newMark)

			if err != nil{
				fmt.Println("🔢 Please enter valid input (number 1-100)")
			}else{
				studentMarks[searchName] = validMark
				fmt.Printf("✅ %s's mark has been updated successfully.\n", searchName)
				break
			}
		}
		}else{
			fmt.Printf("❌ %s not found.\n", searchName)
		}
	}


