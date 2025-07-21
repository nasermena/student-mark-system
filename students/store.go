package students

type Student struct {
	Name string
	Mark int
}

type StudentManager struct {
	Students map[string]Student
}

func NewStudentManager() *StudentManager {
	return &StudentManager{
		Students: make(map[string]Student),
	}
}
