package utils

import (
	"fmt"
	"strings"
)

const (
	error   = "❌"
	info    = "ℹ️"
	success = "✅"
	warning = "⚠️"
)

func Text2Emoij(text string) string {
	switch text {
	case "error":
		return error
	case "info":
		return info
	case "success":
		return success
	case "warning":
		return warning
	}
	return info
}
func GenerateHeader(text string) string {
	return strings.ReplaceAll(fmt.Sprintf("\t**%s**\n\n", text), "-", " ")
}
