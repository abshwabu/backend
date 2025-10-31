package main

import (
	"fmt"
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	s = strings.ToLower(s)

	re := regexp.MustCompile(`[^a-z0-9]+`)
	cleaned := re.ReplaceAllString(s, "")

	n := len(cleaned)
	for i := 0; i < n/2; i++ {
		if cleaned[i] != cleaned[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	tests := []string{
		"Racecar",
		"A man, a plan, a canal: Panama",
		"Hello, World!",
		"12321",
		"No lemon, no melon",
	}

	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, IsPalindrome(t))
	}
}

