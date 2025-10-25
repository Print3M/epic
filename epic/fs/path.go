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
	return filepath.Join(append([]string{ctx.OutputPath}, parts...)...)
}

func ProjectPath(parts ...string) string {
	return filepath.Join(append([]string{ctx.ProjectPath}, parts...)...)
}
