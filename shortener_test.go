package main

import (
	"strings"
	"testing"
)

func TestHello (t *testing.T) {
	var tests = []string{
		"Hello", "World", "John", "abc",
	}

	for _, input := range tests {
		if !strings.Contains(Hello(input), input) {
			t.Errorf("Hello(%q) contains %v, failed", input, input)
		}
	}
}