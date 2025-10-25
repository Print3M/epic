package fs

import (
	"epic/ctx"
	"log"
	"path/filepath"
)

func MustGetAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}

	return absPath
}

func OutputPath(parts ...string) string {
	allParts := append([]string{ctx.OutputPath}, parts...)

	return filepath.Join(allParts...)
}

func ProjectPath(parts ...string) string {
	allParts := append([]string{ctx.ProjectPath}, parts...)

	return filepath.Join(allParts...)
}
