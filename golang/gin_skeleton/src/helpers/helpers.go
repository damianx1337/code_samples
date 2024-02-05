package helpers

import (
	"regexp"
)

func CheckFormatSample1 (s string) bool {
	re := regexp.MustCompile("^\\d{4}-\\d{1}-\\d*$")	// format e.g., 1323-5-0459045
	return re.MatchString(s)
}

func CheckFormatSample2 (s string) bool {
	re := regexp.MustCompile("^\\w{4}$")	// format e.g., gh2s
	return re.MatchString(s)
}
