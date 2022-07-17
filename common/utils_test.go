package common

import (
	"regexp"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	str1 := GenerateRandomString(8, "1234567890")

	if len(str1) != 8 {
		t.Error("Result length unequal to passed argument")
	}

	matched, _ := regexp.MatchString(`^(\d+){8}$`, str1)

	if !matched {
		t.Error("Result contains characters out of the allowed range")
	}

	str2 := GenerateRandomString(100, "")

	if len(str2) != 100 {
		t.Error("Result length unequal to passed argument")
	}

	str3 := GenerateRandomString(0, "")

	if str3 != "" {
		t.Error("Result is not empty")
	}
}
