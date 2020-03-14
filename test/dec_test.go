package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/gofunc"
)

func TestDec(t *testing.T) {
	input := 20
	expected := 19
	actual := gofunc.Dec(input)
	if expected != actual {
		t.Errorf("Error the actual result is not equals expected")
	}
}
