package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	fmt.Println(isValid("pass"))     // false
	fmt.Println(isValid("password")) // false
	fmt.Println(isValid("Password")) // false
	fmt.Println(isValid("P@ssword")) // false
	fmt.Println(isValid("P@ssw0rd")) // true
}

func isValid(s string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	re := regexp.MustCompile("[a-z]")
	re1 := regexp.MustCompile("[A-Z]")
	re2 := regexp.MustCompile("[0-9]")

	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {

		if re.MatchString(string(char)) {
			hasUpper = true
			// fmt.Println("hi")
		} else if re1.MatchString(string(char)) {
			hasLower = true
		} else if re2.MatchString(string(char)) {
			hasNumber = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
