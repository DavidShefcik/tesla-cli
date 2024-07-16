package utils

import "strings"

func RemoveStringWhitespace(input *string) {
	*input = strings.ReplaceAll(*input, " ", "")
}