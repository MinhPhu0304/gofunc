package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/go-func"
)

func Test3Elements(t *testing.T) {
	array := []int{1, 2, 3}
	expected := 6
	actual := gofunc.Sum(array)
	if actual != expected { 
		t.Errorf("Test failed: Sum function does not give expected result")
	}
}

func Test10Elements(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 55
	actual := gofunc.Sum(array)
	if actual != expected { 
		t.Errorf("Test failed: Sum function does not give expected result")
	}
}
