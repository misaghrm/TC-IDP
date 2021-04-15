package utils

import (
	"github.com/mavihq/persian"
	"regexp"
	"strings"
)

func CleanUpPhone(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return value
	}
	value = DigitsToLatin(value)
	re := regexp.MustCompile(`^(?P<prefix>0098{1}|\+?98{1}|0{1})?(?P<phone>9{1}[0-9]{9})$`)
	if re.MatchString(value) {
		value = strings.TrimPrefix(value, "0098")
		value = strings.TrimPrefix(value, "+98")
		value = strings.TrimPrefix(value, "98")
		value = strings.TrimPrefix(value, "0")
		return value
	} else {
		return ""
	}
}

func DigitsToLatin(value string) string {
	return persian.ToEnglishDigits(value)
}

func NormalizePersian(value string) string {
	value = strings.ReplaceAll(value, "\u1610", "\u1740")
	value = strings.ReplaceAll(value, "\u1603", "\u1705")
	return value
}
