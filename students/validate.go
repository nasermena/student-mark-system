package students

import (
	"fmt"
	"strings"
	"regexp"
)

func ValidateMark(input int) (int, error) {
	if input >= 1 && input <= 100 {

		return input, nil
	}
	return 0, fmt.Errorf("invalid mark: %d", input)
}

func ValidateName(input string) bool {
	trimmed := strings.TrimSpace(input)
	reg := regexp.MustCompile(`^[A-Za-z]+(['-]?[A-Za-z]+)*( [A-Za-z]+(['-]?[A-Za-z]+)*)*$`)
	return reg.MatchString(trimmed)
}
