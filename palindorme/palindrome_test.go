package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"Racecar", true},
		{"A man, a plan, a canal: Panama", true},
		{"No lemon, no melon", true},
		{"Hello", false},
		{"12321", true},
		{"Was it a car or a cat I saw?", true},
	}

	for _, c := range cases {
		result := IsPalindrome(c.input)
		if result != c.expected {
			t.Errorf("IsPalindrome(%q) = %v; want %v", c.input, result, c.expected)
		}
	}
}

