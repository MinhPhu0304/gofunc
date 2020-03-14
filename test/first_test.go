package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/gofunc"
)

func TestFirst(t *testing.T) {
	array := [3]int{1, 2, 3}

	if gofunc.First(array) != array[0] {
		t.Errorf("Test failed")
	}
}

func TestFirstWithEmptyArray(t *testing.T) {
	var array [5]int

	if gofunc.First(array) != array[0] {
		t.Errorf("Test failed")
	}
}

func TestNotArrayInput(t *testing.T) {
	var array int

	if gofunc.First(array) != nil {
		t.Errorf("Test failed")
	}
}
