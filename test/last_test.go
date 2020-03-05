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
