package main

import (
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	result := IsPalindrome("amanaplanacanalpanama")
	if result != true {
		t.Errorf("Result was incorrect, got: %s, want: %s.", strconv.FormatBool(result), "true")
	}
}
