package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/go-func"
)

func TestLast(t *testing.T) {
	input := [3]int{1, 2, 3}

	result := gofunc.Last(input)

	if result != input[len(input)-1] {
		t.Errorf("Failed Test ")
	}
}

func TestLastWithEmptyIntegerArray(t *testing.T) {
	var input [5]int
	result := gofunc.Last(input)

	if result != 0 {
		t.Errorf("Failed Test ")
	}
}

func TestLastWithEmptyStringArray(t *testing.T) {
	var input [5]string
	result := gofunc.Last(input)

	if result != "" {
		t.Errorf("Failed Test ")
	}
}

func TestLastWithWrongDataType(t *testing.T) {
	input := 3

	result := gofunc.Last(input)

	if result != nil {
		t.Errorf("Failed Test ")
	}
}
