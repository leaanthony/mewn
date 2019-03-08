package test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"runtime"
	"testing"

	"github.com/leaanthony/mewn/lib"
)

func TestPacking(t *testing.T) {
	ignoreErrors := false
	mewnFiles := lib.GetMewnFiles([]string{"example.go"}, ignoreErrors)
	if len(mewnFiles) != 1 {
		t.Fail()
	}

	referencedAssets, err := lib.GetReferencedAssets(mewnFiles)
	if err != nil {
		t.Fail()
	}

	if len(referencedAssets) != 1 {
		t.Fail()
	}

	theAsset := referencedAssets[0]
	packedFileString, err := lib.GeneratePackFileString(theAsset, ignoreErrors)
	if err != nil {
		t.Fatal(err)
	}
	fixture, err := ioutil.ReadFile("./fixtures/example.go.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Fix line endings for fixture
	if runtime.GOOS == "windows" {
		fixture = bytes.Replace(fixture, []byte{13, 10}, []byte{10}, -1)
	}

	if string(fixture) != packedFileString {
		fmt.Println(string(fixture))
		fmt.Println(packedFileString)
		t.Fail()
	}
}

func TestEmptyPack(t *testing.T) {
	ignoreErrors := false
	mewnFiles := lib.GetMewnFiles([]string{"empty.go"}, ignoreErrors)
	if len(mewnFiles) != 0 {
		t.Fail()
		return
	}

	referencedAssets, err := lib.GetReferencedAssets(mewnFiles)
	if err != nil {
		t.Fail()
		return
	}

	if len(referencedAssets) != 0 {
		t.Fail()
		return
	}

	// No referenced assets so makedummy empty structure
	referencedAsset := &lib.ReferencedAssets{PackageName: "test"}

	packedFileString, err := lib.GeneratePackFileString(referencedAsset, ignoreErrors)
	if err != nil {
		t.Fatal(err)
	}
	fixture, err := ioutil.ReadFile("./fixtures/empty.go.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Fix line endings for fixture
	if runtime.GOOS == "windows" {
		fixture = bytes.Replace(fixture, []byte{13, 10}, []byte{10}, -1)
	}

	if string(fixture) != packedFileString {
		fmt.Println(string(fixture))
		fmt.Println(packedFileString)
		t.Fail()
		return
	}
}

func TestIgnoreErrorPack(t *testing.T) {
	ignoreErrors := true
	referencedAsset := &lib.ReferencedAssets{PackageName: "test"}
	dummyAsset := &lib.ReferencedAsset{AssetPath: "fake/path", Name: "badasset"}
	referencedAsset.Assets = append(referencedAsset.Assets, dummyAsset)
	packedFileString, err := lib.GeneratePackFileString(referencedAsset, ignoreErrors)
	if err != nil {
		t.Fatal(err)
	}

	fixture, err := ioutil.ReadFile("./fixtures/bad.asset.txt")
	if err != nil {
		t.Fatal(err)
	}

	// Fix line endings for fixture
	if runtime.GOOS == "windows" {
		fixture = bytes.Replace(fixture, []byte{13, 10}, []byte{10}, -1)
	}

	if string(fixture) != packedFileString {
		fmt.Println(string(fixture))
		fmt.Println(packedFileString)
		t.Fail()
		return
	}
}
