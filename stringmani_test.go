package main

import (
	"testing"
)

func TestUpperCase(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"software testing", "SOFTWARE TESTING"},
		{"SOFTWARE TESTING", "SOFTWARE TESTING"},
		{"qWeRtY", "QWERTY"},
		{"", ""},
		{"zxcvbn", "ZXCVBN"},
	}

	for _, tc := range testCases {
		result := UpperCase(tc.input)
		if result != tc.expected {
			t.Errorf("UpperCase(%q) = %q; expected %q", tc.input, result, tc.expected)
		}
	}
}