package mewn

import (
	"log"
)

// assetDirectory stores all the assets
var assetDirectory = NewAssetDirectory()
var rootFileGroup *FileGroup
var err error

func init() {
	rootFileGroup, err = assetDirectory.NewFileGroup(".")
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
func AddAsset(name, value string) {
	rootFileGroup.AddAsset(name, value)
}

// Entries returns the file entries as a slice of filenames
func Entries() []string {
	return rootFileGroup.Entries()
}

// Reset clears the file entries
func Reset() {
	rootFileGroup.Reset()
}
