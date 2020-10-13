package utils

import (
	"testing"
)

func TestAdd(t *testing.T) {
	expected := 6
	actual := Add(4, 0, 2)

	if actual != expected {
		t.Errorf("Addition was incorrect. Expected %d, Actual: %d", expected, actual)
	}
}
