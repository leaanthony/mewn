package mewn

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"

	"github.com/leaanthony/mewn/lib"
)

// assetDirectory stores all the assets
var assetDirectory = make(map[string]map[string]string)

// AddAsset adds the given asset to the asset directory.
// Duplicate asset registration will result in hard failure.
func AddAsset(fileGroupName, fileID, data string) error {
	if assetDirectory[fileGroupName] == nil {
		assetDirectory[fileGroupName] = make(map[string]string)
	}
	fileGroup := assetDirectory[fileGroupName]
	if fileGroup[fileID] != "" {
		_, file, _, _ := runtime.Caller(1)
		// Get caller
		return fmt.Errorf("duplicate fileID '%s' registered by file '%s'. This is most likely to happen when you have old *-mewn.go files in your project", fileID, file)
	}
	fileGroup[fileID] = data
	return nil
}

// loadAsset loads the asset for the given filename
func loadAsset(fileGroupName, filename string) ([]byte, error) {
	fileGroup := assetDirectory[fileGroupName]
	if fileGroup == nil {
		_, file, _, _ := runtime.Caller(1)
		return nil, fmt.Errorf("Invalid file group name '%s' accessed by file '%s'", fileGroupName, file)
	}
	// Check internal
	storedAsset := fileGroup[filename]
	if storedAsset != "" {
		return lib.DecompressHexString(storedAsset)
	}
	// Get caller directory
	_, file, _, _ := runtime.Caller(1)
	callerDir := filepath.Dir(file)

	// Calculate full path
	fullFilePath := filepath.Join(callerDir, filename)
	return ioutil.ReadFile(fullFilePath)
}

// String returns the asset as a string
// Failure is indicated by a blank string.
// If you need hard failures, use MustString.
func String(filename string) string {
	contents, _ := loadAsset("/", filename)
	return string(contents)
}

// Bytes returns the asset as a Byte slice.
// Failure is indicated by a blank slice.
// If you need hard failures, use MustBytes.
func Bytes(filename string) []byte {
	contents, _ := loadAsset("/", filename)
	return contents
}

// MustString returns the asset as a string.
// If the asset doesn't exist, it hard fails
func MustString(filename string) string {
	contents, err := loadAsset("/", filename)
	if err != nil {
		log.Fatalf("The asset '%s' was not found! Aborting!", filename)
	}
	return string(contents)
}

// MustBytes returns the asset as a string.
// If the asset doesn't exist, it hard fails
func MustBytes(filename string) []byte {
	contents, err := loadAsset("/", filename)
	if err != nil {
		log.Fatalf("The asset '%s' was not found! Aborting!", filename)
	}
	return contents
}

type FileGroup struct {
}

func (f *FileGroup) String(key string) string {
	return "cool"
}

func Group(groupName string) *FileGroup {
	return &FileGroup{}
}
