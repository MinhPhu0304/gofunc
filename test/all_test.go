package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/gofunc"
)

func TestAllReturnTrue(t *testing.T) {
	allEqual3 := gofunc.All(3)
	input := []int{3, 3, 3}
	actual := allEqual3(input)
	expected := true

	if actual != expected {
		t.Errorf("Failed All function unitest")
	}
}

func TestAllReturnFalse(t *testing.T) {
	allNotEqual3 := gofunc.All(3)
	input := []int{3, 4, 3}
	actual := allNotEqual3(input)
	expected := false

	if actual != expected {
		t.Errorf("Failed All function unitest")
	}
}
