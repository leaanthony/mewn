package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/leaanthony/mewn/lib"
)

func main() {

	buildMode := ""

	mewnFiles := []string{}
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
		targetFile := referencedAsset.Caller[:len(referencedAsset.Caller)-3] + "-mewn.go"
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
