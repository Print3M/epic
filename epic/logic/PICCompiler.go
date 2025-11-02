package logic

import (
	_ "embed"
	"epic/cli"
	"epic/utils"
	"fmt"
	"os"
	"path/filepath"
)

type PICCompiler struct {
	ProjectPath string
	OutputPath  string
	Strict      bool
}

func (pc *PICCompiler) ValidateProjectPath() error {
	if !utils.PathExists(pc.ProjectPath) {
		return fmt.Errorf("project path does not exist: %s", pc.ProjectPath)
	}

	if !utils.MustIsDir(pc.ProjectPath) {
		return fmt.Errorf("project path must be a directory: %s", pc.ProjectPath)
	}

	if err := utils.ValidateProjectStructure(pc.ProjectPath); err != nil {
		return err
	}

	return nil
}

func (pc *PICCompiler) ValidateOutputPath() error {
	if !utils.PathExists(pc.OutputPath) {
		return fmt.Errorf("output path doesn't exist: %s", pc.OutputPath)
	}

	if !utils.MustIsDir(pc.OutputPath) {
		return fmt.Errorf("output path must be a directory: %s", pc.OutputPath)
	}

	return nil
}

func (pc *PICCompiler) Run() {
	pc.compileCore()
	fmt.Println()

	pc.compileModules()
	fmt.Println()
}

func (pc *PICCompiler) compileCore() {
	/*
		Compile "project/core/**" directory.
	*/
	cli.LogInfo("Compiling core...")

	pc.compileProjectDirectory("core")

	cli.LogOk("Core compiled!")
}

func (pc *PICCompiler) compileModules() {
	/*
		Compile "project/modules/<name>/**" directories.
	*/
	modules := utils.GetChildDirs(filepath.Join(pc.ProjectPath, "modules"))

	for _, module := range modules {
		cli.LogInfof("Compiling '%s' module...", module)

		moduleDir := filepath.Join("modules", module)
		pc.compileProjectDirectory(moduleDir)

		cli.LogOkf("Module '%s' compiled!", module)
	}
}

func (pc *PICCompiler) compileProjectDirectory(targetDir string) {
	/*
		Compile all source files from :inputPath and put it into output/objects/**
		directories. The output structure of directory is mimicking :inputPath.
	*/

	absProjectDir := utils.MustGetAbsPath(filepath.Join(pc.ProjectPath, targetDir))

	for _, file := range utils.GetFilesByExtensions(absProjectDir, []string{".c", ".cpp"}) {
		relPath, err := filepath.Rel(absProjectDir, file.FullPath)
		if err != nil {
			cli.LogErrf("%v", err)
			os.Exit(1)
		}

		inputFile := filepath.Join(pc.ProjectPath, targetDir, file.Name)
		outputDir := filepath.Join(
			pc.OutputPath,
			targetDir,
			filepath.Dir(relPath),
		)
		utils.MustCreateDirTree(outputDir)
		outputFile := filepath.Join(outputDir, utils.ReplaceExtension(file.Name, ".o"))

		cli.LogInfof(" ‣ %s -> %s", inputFile, outputFile)

		// TODO: Check which parameters are actually necessary (some of them are linker params)
		params := []string{
			"-c", file.FullPath,
			"-o", outputFile,
			"-Os",
			"-fPIC",
			"-nostdlib",
			"-nostdinc",
			"-nostartfiles",
			"-ffreestanding",
			"-fno-builtin",
			"-nodefaultlibs",
			"-ffunction-sections",
			"-fdata-sections",
			"-fno-ident",
			"-fno-jump-tables",
			"-falign-jumps=1",
			"-mgeneral-regs-only",
			"-fdiagnostics-color=always",
			"-fcf-protection=none",
			"-mno-sse",
			"-mno-mmx",
			"-mno-red-zone",
			"-mno-stack-arg-probe",
			"-fno-delete-null-pointer-checks",
			"-fno-asynchronous-unwind-tables",
			"-DPIC",
			"-ffixed-rbx",
			"-I", filepath.Join(pc.ProjectPath, "include"),
			"-I", pc.ProjectPath,
		}

		switch filepath.Ext(file.FullPath) {
		case ".cpp":
			params = append(params, "-std=c++20", "-fno-rtti")
		default:
			params = append(params, "-std=c17")
		}

		if pc.Strict {
			params = append(params, "-Wall", "-Wextra", "-pedantic")
		}

		output := utils.MingwGcc(params...)
		if len(output) > 0 {
			fmt.Println(output)
		}
	}
}
