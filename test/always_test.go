package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/go-func"
)

func TestAlwaysReturnTrue(t *testing.T) {
	alwaysTrue := gofunc.Always(true)
	actual := alwaysTrue()
	expected := true

	if actual != expected {
		t.Errorf("Failed Always function unitest")
	}
}
func TestAlwaysReturnFalse(t *testing.T) {
	alwaysFalse := gofunc.Always(false)
	actual := alwaysFalse()
	expected := false

	if actual != expected {
		t.Errorf("Failed Always function unitest")
	}
}
