package utils

import (
	"embed"
	"epic/cli"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func HasExtension(name string, ext string) bool {
	return strings.EqualFold(filepath.Ext(name), strings.ToLower(ext))
}

func ReplaceExtension(name, ext string) string {
	return strings.TrimSuffix(name, filepath.Ext(name)) + ext
}

func PathExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func MustIsDir(path string) bool {
	elo, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	return elo.IsDir()
}

func MustGetAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}

	return absPath
}

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

func GetChildDirs(path string) []string {
	/*
		Simple return names of child directories of the provided path.
	*/
	rawEntries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalln(err)
	}

	var dirs []string

	for _, entry := range rawEntries {
		if !entry.IsDir() {
			continue
		}

		dirs = append(dirs, entry.Name())
	}

	return dirs

}

func GetFilesByExtensions(path string, exts []string) []FsEntry {
	/*
		Get all files from path and all subdirectories by extension.
	*/
	var files []FsEntry

	err := filepath.WalkDir(path, func(entryPath string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() {
			return nil
		}
		if !slices.Contains(exts, filepath.Ext(entryPath)) {
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
	err := os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func MustWriteFile(path string, content string) {
	err := os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func ValidateProjectStructure(path string) {
	corePath := filepath.Join(path, "core")
	if !PathExists(corePath) {
		cli.LogErrf("Invalid project structure. Path doesn't exist: %s", corePath)
		os.Exit(1)
	}

	modulesPath := filepath.Join(path, "modules")
	if !PathExists(modulesPath) {
		cli.LogErrf("Invalid project structure. Path doesn't exist: %s", modulesPath)
		os.Exit(1)
	}
}

func ExtractEmbeddedDir(embeddedFS embed.FS, sourceDir, targetDir string) error {
	return fs.WalkDir(embeddedFS, sourceDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Calculate target path
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(targetDir, relPath)

		if d.IsDir() {
			// Create directory
			return os.MkdirAll(targetPath, 0755)
		}

		// Read embedded file
		data, err := embeddedFS.ReadFile(path)
		if err != nil {
			return err
		}

		// Write to disk
		return os.WriteFile(targetPath, data, 0644)
	})
}
