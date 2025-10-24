package fs

import "strings"

func getExtension(name string) string {
	parts := strings.Split(name, ".")
	last := parts[len(parts)-1]

	return last
}

func IsSourceFile(name string) bool {
	ext := getExtension(name)

	return strings.ToLower(ext) == "c"
}

func IsHeaderFile(name string) bool {
	ext := getExtension(name)

	return strings.ToLower(ext) == "h"
}

func ReplaceExtension(name string, extension string) string {
	parts := strings.Split(name, ".")
	noExtension := parts[:len(parts)-1]

	noExtension = append(noExtension, "o")

	return strings.Join(noExtension, ".")
}
