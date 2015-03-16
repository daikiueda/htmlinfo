package main

import (
	"code.google.com/p/mahonia"
)

func decodeText(text string, charset string) string {
	if charset == "" {
		return text
	}
	return mahonia.NewDecoder(charset).ConvertString(text)
}
