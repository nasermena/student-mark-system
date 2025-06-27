package students

import "fmt"

func validateMark(input int) (int, error) {
	if input >= 1 && input <= 100 {
		return input, nil
	}
	return 0, fmt.Errorf("invalid mark: %d", input)
}