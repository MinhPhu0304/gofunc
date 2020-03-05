package main

import (
	"testing"

	"github.com/MinhPhu0304/gofunc"
)

func TestFirst(t *testing.T) {
	array := [3]int{1, 2, 3}

	if gofunc.First(array) != array[0] {
		t.Errorf("Test failed")
	}
}
