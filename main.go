package main

import (
	"fmt"
	"testing"
)

// ReverseString принимает строку и возвращает её реверс.
func ReverseString(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун для поддержки Unicode
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// main — точка входа программы.
func main() {
	fmt.Println("Reversed 'hello':", ReverseString("hello"))
	fmt.Println("Reversed 'Привет':", ReverseString("Привет"))
}

// Тесты для ReverseString.
func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"Привет", "тевирП"},
		{"", ""},
		{"a", "a"},
		{"12345", "54321"},
	}

	for _, test := range tests {
		result := ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
