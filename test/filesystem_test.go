package test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/leaanthony/mewn/lib"
)

var cwd string

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
}

// Source: https://stackoverflow.com/a/15312097
func byteSlicesMatch(a, b []byte) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestCompressDecompress(t *testing.T) {
	files := []string{"hello.txt", "wails_small.png"}
	for _, file := range files {
		fullFilename := filepath.Join(cwd, "assets", file)
		expectedBytes, err := ioutil.ReadFile(fullFilename)
		if err != nil {
			t.Fatal(err)
		}
		packed, err := lib.CompressFile(fullFilename)
		if err != nil {
			t.Fatal(err)
		}

		decompressedBytes, err := lib.DecompressHexString(packed)
		if err != nil {
			t.Fatal(err)
		}
		if !byteSlicesMatch(decompressedBytes, expectedBytes) {
			t.Fail()
		}
	}
}
