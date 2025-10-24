package fs

import (
	"log"
	"os"
	"path/filepath"
)

type DirEntry struct {
	Name  string
	Path  string
	IsDir bool
}

func ListDir(path string) []DirEntry {
	rawEntries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	var entries []DirEntry

	for _, entry := range rawEntries {
		name := entry.Name()

		entries = append(entries, DirEntry{
			Name:  name,
			IsDir: entry.IsDir(),
			Path:  MustGetAbsPath(filepath.Join(path, name)),
		})
	}

	return entries
}

func GetAllFileEntriesFlat(path string) []DirEntry {
	/*
		Convert all files from all subdirectories (recursively)
		into a flat list of file entries.
	*/

	var entries []DirEntry

	err := filepath.WalkDir(path, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() {
			return nil
		}

		name := entry.Name()
		entries = append(entries, DirEntry{
			Name:  name,
			Path:  MustGetAbsPath(path),
			IsDir: false,
		})

		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	return entries
}
