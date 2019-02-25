package mewn

import (
	"fmt"
	"log"

	"github.com/leaanthony/mewn/lib"
)

// mainAssetDirectory stores all the assets
var mainAssetDirectory = lib.NewAssetDirectory()
var rootFileGroup *lib.FileGroup
var err error

func init() {
	rootFileGroup, err = mainAssetDirectory.NewFileGroup(".")
	if err != nil {
		log.Fatal(err)
	}
}

// String gets the asset value by name
func String(name string) string {
	return rootFileGroup.String(name)
}

// MustString gets the asset value by name
func MustString(name string) string {
	return rootFileGroup.MustString(name)
}

// Bytes gets the asset value by name
func Bytes(name string) []byte {
	return rootFileGroup.Bytes(name)
}

// MustBytes gets the asset value by name
func MustBytes(name string) []byte {
	return rootFileGroup.MustBytes(name)
}

// AddAsset adds the given asset to the root context
func AddAsset(groupName, name, value string) {
	fileGroup := mainAssetDirectory.GetGroup(groupName)
	if fileGroup == nil {
		fileGroup, err = mainAssetDirectory.NewFileGroup(groupName)
		if err != nil {
			log.Fatal(err)
		}
	}
	fileGroup.AddAsset(name, value)
}

// Entries returns the file entries as a slice of filenames
func Entries() []string {
	return rootFileGroup.Entries()
}

// Reset clears the file entries
func Reset() {
	rootFileGroup.Reset()
}

// Group holds a group of assets
func Group(name string) *lib.FileGroup {
	result := mainAssetDirectory.GetGroup(name)
	if result == nil {
		result, err = mainAssetDirectory.NewFileGroup(name)
		if err != nil {
			fmt.Println("64")
			log.Fatal(err.Error())
		}
	}
	return result
}
