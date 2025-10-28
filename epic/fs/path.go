package fs

import (
	"log"
	"os"
	"path/filepath"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)

	return !os.IsNotExist(err)
}

func MustGetAbsPath(path string) string {
	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalln(err)
	}

	return absPath
}

func OutputPath(parts ...string) string {
	// allParts := append([]string{ctx.OutputPath}, parts...)

	return filepath.Join(parts...)
}

func ProjectPath(parts ...string) string {
	// allParts := append([]string{ctx.ProjectPath}, parts...)

	return filepath.Join(parts...)
}

/*
func xGetModules(searchPath string) []Module {
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

func xGetProjectModules() []Module {
	searchPath := ProjectPath("modules")

	return getModules(searchPath)
}

func xGetOutputModules() []Module {
	searchPath := OutputPath("objects", "modules")

	if !PathExists(searchPath) {
		return []Module{}
	}

	return getModules(searchPath)
}
*/
