package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/leaanthony/mewn/lib"
)

var version = "0.0.0"
var commit = ""
var date = "Now"

type arglist struct {
	args []string
}

func newArglist() *arglist {
	return &arglist{}
}

func (a *arglist) addArgs(args []string) {
	a.args = append(a.args, args...)
}

func (a *arglist) hasMoreArgs() bool {
	return len(a.args) > 0
}

func (a *arglist) pop() string {
	if !a.hasMoreArgs() {
		return ""
	}
	result := a.args[0]
	a.args = a.args[1:]
	return result
}

func (a *arglist) peek() string {
	if !a.hasMoreArgs() {
		return ""
	}
	return a.args[0]
}

func (a *arglist) popAll() []string {
	result := a.args
	a.args = []string{}
	return result
}

func main() {

	message := ""
	if len(commit) > 5 {
		message = fmt.Sprintf("v%s (%s - %s)", version, commit[0:5], date)
	}

	fmt.Printf("Mewn %s\n", message)

	args := newArglist()
	if len(os.Args) > 0 {
		args.addArgs(os.Args[1:])
	}
	// Print help
	if args.peek() == "--help" {
		fmt.Println()
		fmt.Println("Mewn is a tool for packing assets into your Go binary. This cli tool generates asset bundles and can also act as a replacement for go build, where it will pack, compile and clean up in one step.")
		fmt.Println()
		fmt.Println("Usage: mewn [-i] [build|pack]")
		fmt.Println()
		fmt.Println("Flags:")
		fmt.Println("  -i   Ignore files that don't exist")
		fmt.Println()
		fmt.Println("Subcommands:")
		fmt.Println("  build      Generates asset packs, calls 'go build' then cleans up")
		fmt.Println("  pack       Same as build, except will compile with '-w -s' flags")
		fmt.Println()
		fmt.Println("More information at https://github.com/leaanthony/mewn")
		fmt.Println()
		os.Exit(0)
	}

	var ignoreErrors = false
	if args.peek() == "-i" {
		ignoreErrors = true
		args.pop()
	}

	buildMode := ""
	buildArgs := []string{}
	if args.peek() == "build" || args.peek() == "pack" {
		buildMode = args.pop()
		buildArgs = args.popAll()
	}

	mewnFiles := lib.GetMewnFiles(args.popAll(), ignoreErrors)

	if len(mewnFiles) == 0 {
		fmt.Println("No files found to process.")
		os.Exit(1)
	}
	referencedAssets, err := lib.GetReferencedAssets(mewnFiles)
	if err != nil {
		log.Fatal(err)
	}

	targetFiles := []string{}

	for _, referencedAsset := range referencedAssets {
		packfileData, err := lib.GeneratePackFileString(referencedAsset, ignoreErrors)
		if err != nil {
			log.Fatal(err)
		}
		targetFile := filepath.Join(referencedAsset.BaseDir, referencedAsset.PackageName+"-mewn.go")
		targetFiles = append(targetFiles, targetFile)
		ioutil.WriteFile(targetFile, []byte(packfileData), 0644)
	}

	if buildMode == "build" || buildMode == "pack" {

		var cmdargs []string

		cmdargs = append(cmdargs, "build")
		cmdargs = append(cmdargs, buildArgs...)
		if buildMode == "pack" {
			cmdargs = append(cmdargs, "-ldflags")
			cmdargs = append(cmdargs, "-w -s")
		}
		cmd := exec.Command("go", cmdargs...)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error running command! %s\n", err.Error)
			fmt.Printf("From program: %s\n", stdoutStderr)
		}

		// Remove target Files
		for _, filename := range targetFiles {
			err := os.Remove(filename)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
