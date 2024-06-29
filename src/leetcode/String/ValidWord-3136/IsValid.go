package ValidWord_3136

import "regexp"

func isValid(word string) bool {
	n := len(word)

	return n >= 3 && letter(word) && pronounce(word)
}

func letter(word string) bool {
	reg := `^[a-zA-Z0-9]+$`
	re := regexp.MustCompile(reg)
	return re.MatchString(word)
}

func pronounce(word string) bool {
	reg1 := `[aeiouAEIOU]`
	re1 := regexp.MustCompile(reg1)

	reg2 := `[^aeiouAEIOU1-9]`
	re2 := regexp.MustCompile(reg2)

	return re1.MatchString(word) && re2.MatchString(word)
}
