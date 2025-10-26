package fs

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type FsEntry struct {
	Name     string
	Dir      string
	FullPath string
	IsDir    bool
}

func rawToFsEntry(entry os.DirEntry, entryPath string) FsEntry {
	dir := filepath.Dir(entryPath)
	name := entry.Name()
	absDir := MustGetAbsPath(dir)
	fullPath := filepath.Join(absDir, name)

	return FsEntry{
		Name:     name,
		IsDir:    entry.IsDir(),
		Dir:      filepath.Dir(fullPath),
		FullPath: fullPath,
	}
}

func GetDirectories(path string) []FsEntry {
	/*
		Get all directories (no subdirectories) from path.
	*/
	rawEntries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	var entries []FsEntry

	for _, entry := range rawEntries {
		if !entry.IsDir() {
			continue
		}

		entryPath := filepath.Join(path, entry.Name())
		entries = append(entries, rawToFsEntry(entry, entryPath))
	}

	return entries
}

func GetFilesByExtension(path string, ext string) []FsEntry {
	/*
		Get all files from path and all subdirectories by extension.
	*/
	var files []FsEntry

	err := filepath.WalkDir(path, func(entryPath string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if entry.IsDir() || !HasExtension(entry.Name(), ext) {
			return nil
		}

		files = append(files, rawToFsEntry(entry, entryPath))

		return nil
	})

	if err != nil {
		log.Fatalln(err)
	}

	return files
}

func MustCreateDirTree(path string) {
	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func MustCopyFile(src string, dst string) {
	data, err := os.ReadFile(src)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func MustWriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func EntryExists(path string) bool {
	_, err := os.Stat(path)

	return !errors.Is(err, os.ErrNotExist)
}
