package main

import (
	"regexp"
)

func IsPalindrome(s string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]`)
	strCleaned := re.ReplaceAllString(s, "")
	last := len(strCleaned) - 1
	for i := 0; i < len(strCleaned); i++ {
		if !(strCleaned[i] == strCleaned[last]) {
			return false
		}
		last -= 1
	}
	return true
}
