package fs

import (
	"epic/ctx"
	"log"
	"path/filepath"
	"slices"
	"strings"
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

type Module struct {
	Name string
	Path string
}

func getModules(searchPath string) []Module {
	entries := GetDirectories(searchPath)

	var modules []Module

	for _, entry := range entries {
		if !entry.IsDir {
			continue
		}

		// Accept only user defined modules
		if !slices.Contains(ctx.Modules, strings.ToLower(entry.Name)) {
			continue
		}

		modules = append(modules, Module{
			Name: entry.Name,
			Path: entry.FullPath,
		})
	}

	// Check if there's any missing module
	if len(modules) != len(ctx.Modules) {
		var names []string
		for _, module := range modules {
			names = append(names, strings.ToLower(module.Name))
		}

		for _, module := range ctx.Modules {
			if !slices.Contains(names, module) {
				log.Fatalf("Module '%s' not found (%s)", module, filepath.Join(searchPath, module))
			}
		}
	}

	return modules
}

func GetProjectModules() []Module {
	/*
		Each module is a separate directory in "project/modules/<module>/'
	*/
	searchPath := ProjectPath("modules")

	return getModules(searchPath)
}

func GetOutputModules() []Module {
	/*
		Each module is a separate directory in "output/objects/modules/<module>/'
	*/
	searchPath := OutputPath("objects", "modules")

	return getModules(searchPath)
}
