package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/leaanthony/mewn/lib"
)

var version = "0.0.0"
var commit = ""
var date = "Now"

func main() {

	message := "[dev build]"
	if len(commit) > 5 {
		message = fmt.Sprintf("v%s (%s - %s)", version, commit[0:5], date)
	}

	fmt.Printf("Mewn %s\n", message)
	buildMode := ""

	var mewnFiles []string
	inputs := []string{}
	if len(os.Args) > 1 {
		if os.Args[1] == "build" || os.Args[1] == "pack" {
			buildMode = os.Args[1]
		} else {
			inputs = os.Args[1:]
		}
	}
	mewnFiles = lib.GetMewnFiles(inputs...)

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
		packfileData := lib.GeneratePackFileString(referencedAsset)
		targetFile := referencedAsset.PackageName + "-mewn.go"
		targetFiles = append(targetFiles, targetFile)
		ioutil.WriteFile(targetFile, []byte(packfileData), 0644)
	}

	if buildMode == "build" || buildMode == "pack" {

		var args []string

		if buildMode == "pack" {

			args = append(args, "build")
			args = append(args, os.Args[2:]...)
			args = append(args, "-ldflags")
			args = append(args, "-w -s")
		} else {
			args = append(args, os.Args[1:]...)
		}
		cmd := exec.Command("go", args...)
		stdoutStderr, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error running command: %s\n", stdoutStderr)
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
