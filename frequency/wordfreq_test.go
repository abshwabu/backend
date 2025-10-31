package main

import (
	"reflect"
	"testing"
)

func TestWordFrequency(t *testing.T) {
	text := "Hello, hello! This is a test. A test, this is."
	expected := map[string]int{
		"hello": 2,
		"this":  2,
		"is":    2,
		"a":     2,
		"test":  2,
	}

	result := WordFrequency(text)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

