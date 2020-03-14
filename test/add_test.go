package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/gofunc"
)

func TestAddInt(t *testing.T) {
	input := 3

	result := gofunc.AddInt(input)
	if result != input {
		t.Errorf("Adding failed")
	}
}

func TestAddIntMultiParams(t *testing.T) {
	expected := 9
	result := gofunc.AddInt(1, 2, 2, 4)
	if result != expected {
		t.Errorf("Adding failed")
	}
}

func TestAddStringSingleParam(t *testing.T) {
	input := "Hello"

	result := gofunc.AddString(input)

	if result != input {
		t.Errorf("Add string function does not give expected result")
	}
}

func TestAddStringMultiParams(t *testing.T) {
	input := "Hello"
	secondInpt := "world"
	result := gofunc.AddString(input, secondInpt)
	expected := input + " " + secondInpt
	if result != expected {
		t.Errorf("Add string function does not give expected result")
	}
}
