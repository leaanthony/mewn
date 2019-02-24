package mewn

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/leaanthony/mewn/lib"
)

// AssetDirectory is a collection of file groups
type AssetDirectory struct {
	fileGroups map[string]*FileGroup
}

func NewAssetDirectory() *AssetDirectory {
	return &AssetDirectory{
		fileGroups: make(map[string]*FileGroup),
	}
}

// NewFileGroup creates a new file group
func (a *AssetDirectory) NewFileGroup(baseDirectory string) (*FileGroup, error) {
	_, exists := a.fileGroups[baseDirectory]
	if exists {
		return nil, fmt.Errorf("fileGroup '%s' already registered", baseDirectory)
	}
	result := &FileGroup{
		baseDirectory:  baseDirectory,
		assetDirectory: make(map[string]string),
	}
	a.fileGroups[baseDirectory] = result

	return result, nil
}

// FileGroup holds a collection of files
type FileGroup struct {
	baseDirectory  string
	assetDirectory map[string]string
}

// AddAsset to the filegroup
func (f *FileGroup) AddAsset(name, data string) error {
	_, exists := f.assetDirectory[name]
	if exists {
		return fmt.Errorf("asset '%s' already registered in FileGroup '%s'", name, f.baseDirectory)
	}
	f.assetDirectory[name] = data
	return nil
}

// String returns the asset as a string
// Failure is indicated by a blank string.
// If you need hard failures, use MustString.
func (f *FileGroup) String(filename string) string {
	contents, _ := f.loadAsset(filename)
	return string(contents)
}

// Bytes returns the asset as a Byte slice.
// Failure is indicated by a blank slice.
// If you need hard failures, use MustBytes.
func (f *FileGroup) Bytes(filename string) []byte {
	contents, _ := f.loadAsset(filename)
	return contents
}

// MustString returns the asset as a string.
// If the asset doesn't exist, it hard fails
func (f *FileGroup) MustString(filename string) string {
	contents, err := f.loadAsset(filename)
	if err != nil {
		log.Fatalf("The asset '%s' was not found! Aborting!", filename)
	}
	return string(contents)
}

// MustBytes returns the asset as a string.
// If the asset doesn't exist, it hard fails
func (f *FileGroup) MustBytes(filename string) []byte {
	contents, err := f.loadAsset(filename)
	if err != nil {
		log.Fatalf("The asset '%s' was not found! Aborting!", filename)
	}
	return contents
}

// Entries returns a slice of filenames in the filegroup
func (f *FileGroup) Entries() []string {
	keys := reflect.ValueOf(f.assetDirectory).MapKeys()
	result := []string{}
	for _, key := range keys {
		result = append(result, key.String())
	}
	return result
}

// Reset the filegroup
func (f *FileGroup) Reset() {
	f.assetDirectory = make(map[string]string)
}

// loadAsset loads the asset for the given filename
func (f *FileGroup) loadAsset(filename string) ([]byte, error) {
	// Check internal
	storedAsset := f.assetDirectory[filename]
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
