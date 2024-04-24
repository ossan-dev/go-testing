package utils_test

import (
	"testing"

	"github.com/ossan-dev/gotesting/internal/utils"
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
