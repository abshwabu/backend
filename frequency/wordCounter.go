
package main

import (
	"fmt"
	"regexp"
	"strings"
)

func WordFrequency(text string) map[string]int {
	text = strings.ToLower(text)

	re := regexp.MustCompile(`[^\w\s]`)
	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)

	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func main() {
	text := "Hello, hello! This is a test. A test, this is."
	result := WordFrequency(text)
	fmt.Println(result)
}
