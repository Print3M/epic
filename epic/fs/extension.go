package fs

import (
	"path/filepath"
	"strings"
)

func HasExtension(name string, ext string) bool {
	return strings.EqualFold(filepath.Ext(name), strings.ToLower(ext))
}

func ReplaceExtension(name, ext string) string {
	return strings.TrimSuffix(name, filepath.Ext(name)) + ext
}
