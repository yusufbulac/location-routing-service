package validation

import (
	"regexp"
)

var hexColorRegex = regexp.MustCompile(`^#(?:[0-9a-fA-F]{3}){1,2}$`)

func IsHexColor(s string) bool {
	return hexColorRegex.MatchString(s)
}
