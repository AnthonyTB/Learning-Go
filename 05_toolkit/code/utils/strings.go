package utils

import (
	"strings"
)

// MakeExcited returns a string in all uppercase with an appended !
func MakeExcited(value string) string {
	return strings.ToUpper(value) + "!"
}
