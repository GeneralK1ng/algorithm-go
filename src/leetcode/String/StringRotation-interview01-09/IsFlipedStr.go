package StringRotation_interview01_09

import "strings"

func isFlippedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1s1 := s1 + s1
	return strings.Contains(s1s1, s2)
}
