package common

import "strings"

func InitializeName(name string) string {
	filename := strings.ReplaceAll(name, "http_", "http://")
	filename = strings.ReplaceAll(filename, "https_", "https://")
	filename = strings.ReplaceAll(filename, "%", ":")
	filename = strings.ReplaceAll(filename, "_", "/")
	filename = strings.ReplaceAll(filename, ".png", "")

	return filename
}
