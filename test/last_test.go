package test

import  (
	"testing"

	"github.com/MinhPhu0304/gofunc"
)

func TestLast (t *testing.T) {
	input:= [3] int {1,2,3}

	result:= gofunc.Last(input)

	if result != input[len(input) -1] {
		t.Errorf("Failed Test ")
	}
}