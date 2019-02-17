package test

import (
	"io/ioutil"
	"testing"

	"github.com/leaanthony/mewn/lib"
)

func TestPacking(t *testing.T) {
	mewnFiles := lib.GetMewnFiles("example.go")
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
	packedFileString := lib.GeneratePackFileString(theAsset)
	fixture, err := ioutil.ReadFile("./fixtures/example.go.txt")
	if err != nil {
		t.Fatal(err)
	}
	if string(fixture) != packedFileString {
		t.Fail()
	}
}
