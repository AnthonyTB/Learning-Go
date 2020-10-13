package utils

import (
	"testing"
)

func TestMakeExcited(t *testing.T) {
	expected := "OMG SO EXCITING!"
	actual := MakeExcited("omg so exciting")

	if expected != actual {
		t.Errorf("Sorry your string wasn't very excited. Expected: %s, Actual: %s", expected, actual)
	}
}
