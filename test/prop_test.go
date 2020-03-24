package test

import (
	"testing"

	gofunc "github.com/MinhPhu0304/gofunc"
)

func TestPropReturnNilForInteger(t *testing.T) {
	actual := gofunc.Prop("1", 1)

	if actual != nil {
		t.Errorf("Failed Always function unitest")
	}
}

type DummyStruct struct {
	Name string
}

func TestPropReturnForStruct(t *testing.T) {
	test := DummyStruct{Name: "test"}
	actual := gofunc.Prop("Name", test)
	expected := "test"
	if actual != expected {
		t.Errorf("Failed Always function unitest")
	}
}
