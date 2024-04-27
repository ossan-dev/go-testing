package utils_test

import (
	"testing"

	"github.com/ossan-dev/gotesting/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	// Arrange
	input := "Ninja"
	expected := "ajniN"
	// Act
	actual := utils.ReverseString(input)
	// Assert
	if actual != expected {
		t.Errorf("want: %q, got %q", expected, actual)
	}
}

func FuzzReverseString(f *testing.F) {
	testCases := []string{"Ninja", "Hello, World!", ""}
	for _, tc := range testCases {
		f.Add(tc) // add tests to seed corpus
	}
	f.Fuzz(func(t *testing.T, input string) { // fuzz target func
		reversed := utils.ReverseString(input)
		reversedTwice := utils.ReverseString(reversed)
		assert.Equal(t, input, reversedTwice)
	})
}
