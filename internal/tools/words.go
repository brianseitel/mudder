package tools

import "regexp"

var slugRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)

func Slug(str string) string {
	return slugRegex.ReplaceAllString(str, "-")
}
