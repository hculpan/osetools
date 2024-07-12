package utils

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var p = message.NewPrinter(language.English)

func FormatInt(num int) string {
	return p.Sprintf("%d", num)
}
