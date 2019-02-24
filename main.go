package mewn

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"runtime"

	"github.com/leaanthony/mewn/lib"
)

// assetDirectory stores all the assets
var assetDirectory = make(map[string]string)

// loadAsset loads the asset for the given filename
func loadAsset(filename string) ([]byte, error) {
	// Check internal
	storedAsset := assetDirectory[filename]
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
	contents, _ := loadAsset(filename)
	return string(contents)
}

// Bytes returns the asset as a Byte slice.
// Failure is indicated by a blank slice.
// If you need hard failures, use MustBytes.
func Bytes(filename string) []byte {
	contents, _ := loadAsset(filename)
	return contents
}

// MustString returns the asset as a string.
// If the asset doesn't exist, it hard fails
func MustString(filename string) string {
	contents, err := loadAsset(filename)
	if err != nil {
		log.Fatalf("The asset '%s' was not found! Aborting!", filename)
	}
	return string(contents)
}

// MustBytes returns the asset as a string.
// If the asset doesn't exist, it hard fails
func MustBytes(filename string) []byte {
	contents, err := loadAsset(filename)
	if err != nil {
		log.Fatalf("The asset '%s' was not found! Aborting!", filename)
	}
	return contents
}

func AddFile(key string, value string) {
	_, exists := assetDirectory[key]
	if exists {
		log.Fatalf("Key '%s' already registered", key)
	}
	assetDirectory[key] = value
}

func Entries() []string {
	keys := reflect.ValueOf(assetDirectory).MapKeys()
	result := []string{}
	for _, key := range keys {
		result = append(result, key.String())
	}
	return result
}

func Reset() {
	assetDirectory = make(map[string]string)
}
